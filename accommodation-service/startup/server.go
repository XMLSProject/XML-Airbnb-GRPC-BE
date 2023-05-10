package startup

import (
	handler "accomm_module/handlers"
	"accomm_module/infrastructure/persistence"
	"accomm_module/proto/accommodation"
	repo "accomm_module/repository"
	"accomm_module/service"
	"accomm_module/startup/config"
	"fmt"
	"log"
	"net"

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

	accoRepo := &repo.AccommodationRepository{DatabaseConnection: mongoClient}
	accoService := &service.AccommodationService{AccommodationRepo: accoRepo}

	accoHandler := handler.NewAccommodationHandler(accoService)

	server.startGrpcServer(accoHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.OrderingDBHost, server.config.OrderingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) startGrpcServer(orderHandler *handler.AccommodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation.RegisterAccommodationServiceServer(grpcServer, orderHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
