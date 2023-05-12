package startup

import (
	handler "accomm_module/handlers"
	"accomm_module/infrastructure/persistence"
	"accomm_module/proto/accommodation"
	repo "accomm_module/repository"
	"accomm_module/service"
	"accomm_module/startup/config"
	"context"
	"fmt"
	"log"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(exampleAuthFunc)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(exampleAuthFunc)))
	accommodation.RegisterAccommodationServiceServer(grpcServer, orderHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func parseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func userClaimFromToken(claims jwt.MapClaims) string {

	sub, ok := claims["user_id"].(string)
	if !ok {
		return ""
	}

	return sub
}

var jwtKey = []byte("secret_key")

func exampleAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v for %s", err, token)
	}
	user_id := userClaimFromToken(tokenInfo)

	grpc_ctxtags.Extract(ctx).Set("auth.sub", user_id)

	newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)

	return newCtx, nil
}
