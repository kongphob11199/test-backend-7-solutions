package app

import (
	"beef/internal/handler/gapi"
	"beef/internal/repository"
	"beef/internal/service"
	"flag"
	"fmt"
	"log"
	"net"

	pb "beef/internal/pd"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func RunServer() {
	fmt.Println("Grpc Server running ...")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository()

	services := service.NewService(service.Deps{Repository: repository})

	beefGapi := gapi.NewBeefHandlerGrpcHandler(services.Beef)

	s := grpc.NewServer()

	pb.RegisterBeefServiceServer(s, beefGapi)

	fmt.Println("Server running on port : 50051")

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
