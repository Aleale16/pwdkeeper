package msgsender

import (
	"context"
	pb "pwdkeeper/internal/app/proto"

	"github.com/rs/zerolog/log"
)

func SendUserGetmsg(c pb.ActionsClient, login string) (status string, key1enc string) {
	resp, err := c.GetUser(context.Background(), &pb.GetUserRequest{
		Login: login,
	})
	if err != nil {
		log.Fatal().Err(err)
	}
	//if resp.Error != "" {
	//	fmt.Println(resp.Error)
	//}
	return resp.Status, resp.Publickey
}

func SendUserAuthmsg(c pb.ActionsClient, login, password string) (status string){
	resp, err := c.GetUserAuth(context.Background(), &pb.GetUserAuthRequest{
		Login: login,
		Password: password,
	})
	if err != nil {
		log.Fatal().Err(err)
	}
	//if resp.Error != "" {
	//	fmt.Println(resp.Error)
	//}
	return resp.Status
}

func SendUserStoremsg(c pb.ActionsClient, login, password, publickey string) (status string){
	resp, err := c.StoreUser(context.Background(), &pb.StoreUserRequest{
		Login: login,
		Password: password,
		Publickey: publickey,
	})
	if err != nil {
		log.Fatal().Err(err)
	}
	//if resp.Error != "" {
	//	fmt.Println(resp.Error)
	//}
	return resp.Status
}

func SendUserGetRecordsmsg(c pb.ActionsClient, login string) (status string, UserRecordsJSON string){
	resp, err := c.GetUserRecords(context.Background(), &pb.GetUserRecordsRequest{
		Login: login,
	})
	if err != nil {
		log.Fatal().Err(err)
	}
	//if resp.Error != "" {
	//	fmt.Println(resp.Error)
	//}
	return resp.Status, resp.UserRecordsJSON
}

func SendUserStoreRecordmsg(c pb.ActionsClient, namerecord, datarecord, datatype, login string) (status string, RecordID string){
	resp, err := c.StoreSingleRecord(context.Background(), &pb.StoreSingleRecordRequest{
		DataName: namerecord,
		SomeData: datarecord,
		DataType: datatype,
		Login: login,
	})
	if err != nil {
		log.Fatal().Err(err)
	}
	//if resp.Error != "" {
	//	fmt.Println(resp.Error)
	//}
	return resp.Status, resp.RecordID
}