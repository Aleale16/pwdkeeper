package grpcserver

import (
	// импортируем пакет со сгенерированными protobuf-файлами
	"context"
	pb "pwdkeeper/internal/app/proto"
	"pwdkeeper/internal/app/storage"
)

// ActionsServer поддерживает все необходимые методы сервера.
type ActionsServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedActionsServer
}

func (s *ActionsServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var response pb.GetUserResponse

	response.Status, response.Publickey = storage.GetUser(in.Login)

	return &response, nil
}

func (s *ActionsServer) StoreUser(ctx context.Context, in *pb.StoreUserRequest) (*pb.StoreUserResponse, error) {
	var response pb.StoreUserResponse

	response.Status, response.Publickey = storage.StoreUser(in.Login, in.Password, in.Publickey)

	return &response, nil
}

func (s *ActionsServer) GetUserAuth(ctx context.Context, in *pb.GetUserAuthRequest) (*pb.GetUserAuthResponse, error) {
	var response pb.GetUserAuthResponse

	response.Status, response.Publickey = storage.AuthenticateUser(in.Login, in.Password)

	return &response, nil
}

func (s *ActionsServer) GetUserRecords(ctx context.Context, in *pb.GetUserRecordsRequest) (*pb.GetUserRecordsResponse, error) {
	var response pb.GetUserRecordsResponse

	response.Status, response.UserRecordsJSON = storage.GetUserRecords(in.Login)

	return &response, nil
}

func (s *ActionsServer) StoreSingleRecord(ctx context.Context, in *pb.StoreSingleRecordRequest) (*pb.StoreSingleRecordResponse, error) {
	var response pb.StoreSingleRecordResponse

	response.Status, response.RecordID = storage.StoreRecord(in.DataName, in.SomeData, in.DataType, in.Login)

	return &response, nil
}