package app

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"beef/internal/handler/api"
	pb "beef/internal/pd"
)

func RunClient() {
	flag.Parse()

	cc, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to server ", err)
	}

	clientBeef := pb.NewBeefServiceClient(cc)

	myhandler := api.NewHandler(clientBeef)

	app := myhandler.Init()

	go func() {
		if err := app.Listen(":5000"); err != nil {
			log.Fatal("server error ", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Fatal("Shutting down the server...")

}
