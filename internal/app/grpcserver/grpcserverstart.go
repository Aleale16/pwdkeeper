package grpcserver

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	pb "pwdkeeper/internal/app/proto"
	"pwdkeeper/internal/app/storage"
	"syscall"

	"google.golang.org/grpc"
)

// Grpcserverstart starts gRPC server
<<<<<<< HEAD
func Grpcserverstart() error{
=======
func Grpcserverstart() (error) {
>>>>>>> 982e7d38dd6e3f0e8ace8216f21751b8f85d91cb
	
	storage.Initdb()

	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}

	// создаём gRPC-сервер без зарегистрированной службы
<<<<<<< HEAD
	s := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterActionsServer(s, &ActionsServer{})
=======
	S := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterActionsServer(S, &ActionsServer{})
>>>>>>> 982e7d38dd6e3f0e8ace8216f21751b8f85d91cb

		fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	//	if err := s.Serve(listen); err != nil {
	//		log.Fatal(err)
	//	}
	errChan := make(chan error)
	stopChan := make(chan os.Signal, 1)

	// Ожидаем события от ОС
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// сообщаем об ошибках в канал
	go func() {
<<<<<<< HEAD
		if err := s.Serve(listen); err != nil {
=======
		if err := S.Serve(listen); err != nil {
>>>>>>> 982e7d38dd6e3f0e8ace8216f21751b8f85d91cb
			errChan <- err
		}
	}()
	defer func() {
<<<<<<< HEAD
		s.GracefulStop()
=======
		S.GracefulStop()
>>>>>>> 982e7d38dd6e3f0e8ace8216f21751b8f85d91cb
	}()

	select {
	case err := <-errChan:
		log.Printf("Fatal error: %v\n", err)
	case <-stopChan:
	}
	return nil
}