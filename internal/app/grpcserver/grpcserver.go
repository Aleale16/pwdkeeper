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

	response.Status, response.Fek = storage.GetUser(in.Login)

	return &response, nil
}

func (s *ActionsServer) StoreUser(ctx context.Context, in *pb.StoreUserRequest) (*pb.StoreUserResponse, error) {
	var response pb.StoreUserResponse

	response.Status, response.Fek = storage.StoreUser(in.Login, in.Password, in.Fek)

	return &response, nil
}

func (s *ActionsServer) GetUserAuth(ctx context.Context, in *pb.GetUserAuthRequest) (*pb.GetUserAuthResponse, error) {
	var response pb.GetUserAuthResponse

	response.Status, response.Fek = storage.AuthenticateUser(in.Login, in.Password)

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

func (s *ActionsServer) UpdateRecord(ctx context.Context, in *pb.UpdateRecordRequest) (*pb.UpdateRecordResponse, error) {
	var response pb.UpdateRecordResponse

	response.Status= storage.UpdateRecord(in.RecordID, in.EncryptedData)

	return &response, nil
}

func (s *ActionsServer) DeleteRecord(ctx context.Context, in *pb.DeleteRecordRequest) (*pb.DeleteRecordResponse, error) {
	var response pb.DeleteRecordResponse

	response.Status= storage.DeleteRecord(in.RecordID)

	return &response, nil
}

func (s *ActionsServer) GetSingleRecord(ctx context.Context, in *pb.GetSingleRecordRequest) (*pb.GetSingleRecordResponse, error) {
	var response pb.GetSingleRecordResponse

	//response.EncryptedData, response.DataType = storage.GetRecord(in.RecordID)
	response.EncryptedData, response.DataType = storage.GetRecord(in.RecordID)

	return &response, nil
}

func (s *ActionsServer) GetSingleNameRecord(ctx context.Context, in *pb.GetSingleNameRecordRequest) (*pb.GetSingleNameRecordResponse, error) {
	var response pb.GetSingleNameRecordResponse

	//response.EncryptedData, response.DataType = storage.GetRecord(in.RecordID)
	response.DataName = storage.GetNameRecord(in.RecordID)

	return &response, nil
}