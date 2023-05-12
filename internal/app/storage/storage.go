package storage

import (
	"encoding/hex"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var PGdb *pgxpool.Pool

type authUsers struct{ login string; password string; publickey string }
type dataRecords struct{ idrecord int32; namerecord, datarecord, datatype string; login string }
var (	authUser authUsers;
		record dataRecords;
	)
type storagerUser interface {
	storeuser() (status string, authToken string)
	getuser() (status string, publickey string)
	authenticateuser() (status string, publickey string)
	getuserrecords() (status string, rowsDataRecordJSON string)
}
type storagerData interface {
	storerecord() (status string, recordID string)
	getrecord() (status string, rowDataRecordJSON string)
	
}
var (	SUser storagerUser;
		SData storagerData;
	)

//структура выводимого JSON	 
type rowDataRecord struct {
		IDrecord		int32 	`json:"id,omitempty"`
		Namerecord 		string 	`json:"namerecord"`    
		Datarecord 		string 	`json:"datarecord"`       
		Datatype 		string 	`json:"datatype"`       
	}

func StoreUser(login string, password string, publickey string) (status string, authToken string) {
	log.Debug().Msg("func StoreUser")
	authUser.login = login
	authUser.password = password
	authUser.publickey = publickey
	SUser = authUser
	return SUser.storeuser()
}

func GetUser(login string) (status string, publickey string){
	log.Debug().Msg("func GetUser")
	authUser.login = login
	SUser = authUser
	return SUser.getuser()
}

func AuthenticateUser(login, password string) (status string, publickey string){
	log.Debug().Msg("func GetUser")
	authUser.login = login
	authUser.password = password
	SUser = authUser
	return SUser.authenticateuser()
}

func GetUserRecords(login string) (status string, rowsDataRecordJSON string){
	log.Debug().Msg("func GetUserRecords")
	authUser.login = login
	SUser = authUser
	return SUser.getuserrecords()
}

func StoreRecord(namerecord, datarecord, datatype, login string) (status string, recordID string){
	log.Debug().Msg("func StoreRecord")
	record.namerecord = namerecord
	record.datarecord = hex.EncodeToString([]byte(datarecord))
	record.datatype = datatype
	record.login = login
	SData = record
	return SData.storerecord()
}

func GetRecord(idrecord int32) (status string, rowDataRecordJSON string){
	record.idrecord = idrecord
	SData = record
	return SData.getrecord()
}

func UpdateRecord() {

}
