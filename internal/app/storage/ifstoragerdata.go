package storage

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func (data dataRecords) storerecord() (status string, recordID string) {
	
err := PGdb.QueryRow(context.Background(), `INSERT into data(namerecord, datarecord, datatype, login_fkey) values($1,decode($2,'hex'),$3,$4) RETURNING (id)`, data.namerecord, data.datarecord, data.datatype, data.login ).Scan(&recordID)
	//result, err := PGdb.Exec(context.Background(), `INSERT into data(namerecord, datarecord, login_fkey) values($1,$2,$3)`, data.namerecord, data.datarecord, data.login)
	if err != nil {
		log.Error().Msg(err.Error())
		return "500", ""
	}
	log.Info().Msgf("NEW data record %v inserted successfully.", recordID)
	status = "200"
	return status, recordID
}

func (data dataRecords) getrecord() (status string, DataRecordJSON string) {
	var namerecord, datarecord, datatype string
	var rowDataRecordJSON []rowDataRecord
	err := PGdb.QueryRow(context.Background(), `SELECT data.namerecord, data.datarecord, data.datatype FROM data WHERE id=$1`, data.idrecord ).Scan(&namerecord, &datarecord, &datatype)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	rowDataRecordJSON = append(rowDataRecordJSON, rowDataRecord{
		Namerecord:			namerecord,
		Datarecord:			datarecord,
		Datatype:			datatype,
	})
	JSONdata, err := json.MarshalIndent(rowDataRecordJSON, "", "  ")
	if err != nil {
		log.Fatal().Str("JSONdata","rowDataRecordJSON").Msg(err.Error())
	}
	log.Info().Msgf("Data record row %v extracted successfully.", string(JSONdata))
	status = "200"
	return status, string(JSONdata)
}

