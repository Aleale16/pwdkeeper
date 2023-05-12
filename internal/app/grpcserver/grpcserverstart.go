package grpcserver

import (
	"fmt"
	"log"
	"net"
	pb "pwdkeeper/internal/app/proto"
	"pwdkeeper/internal/app/storage"

	"google.golang.org/grpc"
)

// Grpcserverstart starts gRPC server
func Grpcserverstart() {
	
	storage.Initdb()

	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}

	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterActionsServer(s, &ActionsServer{})

		fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
		if err := s.Serve(listen); err != nil {
			log.Fatal(err)
		}
}