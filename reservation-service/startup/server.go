package startup

import (
	"fmt"
	"log"
	"net"
	handler "res_init/handlers"
	"res_init/infrastructure/persistence"
	"res_init/proto/reservation"
	repo "res_init/repository"
	"res_init/service"
	"res_init/startup/config"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "order_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()

	userRepo := &repo.ResRepository{DatabaseConnection: mongoClient}
	userService := &service.ResService{ResRepo: userRepo}

	loginHandler := handler.NewReservationHandler(userService)

	server.startGrpcServer(loginHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.OrderingDBHost, server.config.OrderingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) startGrpcServer(orderHandler *handler.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reservation.RegisterReservationServiceServer(grpcServer, orderHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
