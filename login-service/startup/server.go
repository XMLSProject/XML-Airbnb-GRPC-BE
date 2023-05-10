package startup

import (
	handler "first_init/handlers"
	"first_init/infrastructure/persistence"
	"first_init/proto/login"
	repo "first_init/repository"
	"first_init/service"
	"first_init/startup/config"
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

	userRepo := &repo.UserRepository{DatabaseConnection: mongoClient}
	userService := &service.UserService{UserRepo: userRepo}

	loginHandler := handler.NewAuthenticationHandler(userService)

	server.startGrpcServer(loginHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.OrderingDBHost, server.config.OrderingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) startGrpcServer(orderHandler *handler.LoginHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	login.RegisterLoginServiceServer(grpcServer, orderHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
