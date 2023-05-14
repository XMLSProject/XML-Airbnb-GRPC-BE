package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example/gateway/config"
	"example/gateway/proto/greeter"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.GetConfig()

	conn, err := grpc.DialContext(
		context.Background(),
		cfg.GreeterServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	client := greeter.NewGreeterServiceClient(conn)
	err = greeter.RegisterGreeterServiceHandlerClient(
		context.Background(),
		gwmux,
		client,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	///////////////////////////////////////////////////////

	connn, errr := grpc.DialContext(
		context.Background(),
		cfg.LoginServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if errr != nil {
		log.Fatalln("Failed to dial serverr:", errr)
	}
	// Register Greeter
	clientt := greeter.NewLoginServiceClient(connn)
	err = greeter.RegisterLoginServiceHandlerClient(
		context.Background(),
		gwmux,
		clientt,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	///////////////////////////////////////////////////////

	connn7, errr7 := grpc.DialContext(
		context.Background(),
		cfg.ReservationServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if errr7 != nil {
		log.Fatalln("Failed to dial serverr:", errr7)
	}
	// Register Greeter
	clientt7 := greeter.NewReservationServiceClient(connn7)
	errr7 = greeter.RegisterReservationServiceHandlerClient(
		context.Background(),
		gwmux,
		clientt7,
	)
	if errr7 != nil {
		log.Fatalln("Failed to register gateway:", errr7)
	}

	///////////////////////////////////////////////////////

	con4, err4 := grpc.DialContext(
		context.Background(),
		cfg.AccommodationServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err4 != nil {
		log.Fatalln("Failed to dial serverr:", err4)
	}
	//Register Accommodation
	accoClient := greeter.NewAccommodationServiceClient(con4)
	accoErr := greeter.RegisterAccommodationServiceHandlerClient(
		context.Background(),
		gwmux,
		accoClient,
	)
	if accoErr != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	headersOk := gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization",
		"accept", "origin", "Cache-Control", "X-Requested-With"})
	originsOk := gorillaHandlers.AllowedOrigins([]string{"*"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	cors := gorillaHandlers.CORS(headersOk, originsOk, methodsOk)

	gwServer := &http.Server{
		Addr:    cfg.Address,
		Handler: cors(gwmux),
	}

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
