package storage

import (
	"context"
	"encoding/hex"

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

func (data dataRecords) updaterecord() (status string) {	
_, err := PGdb.Exec(context.Background(), `UPDATE data SET datarecord = decode($1,'hex') where data.id = $2`, data.datarecord, data.idrecord )
	if err != nil {
		log.Error().Msg(err.Error())
		return "500"
	}
	log.Info().Msgf("Data record %v updated successfully.", data.idrecord)
	status = "200"
	return status
}

func (data dataRecords) deleterecord() (status string) {	
_, err := PGdb.Exec(context.Background(), `DELETE FROM data WHERE data.id = $1`, data.idrecord )
	if err != nil {
		log.Error().Msg(err.Error())
		return "500"
	}
	log.Info().Msgf("Data record %v deleted successfully.", data.idrecord)
	status = "200"
	return status
}

func (data dataRecords) getrecord() (datarecord string, datatype string) {
	var namerecord string
	err := PGdb.QueryRow(context.Background(), `SELECT data.namerecord, encode(data.datarecord,'hex'), data.datatype FROM data WHERE id=$1`, data.idrecord).Scan(&namerecord, &datarecord, &datatype)
	//err := PGdb.QueryRow(context.Background(), `SELECT data.namerecord, data.datarecord, data.datatype FROM data WHERE id=$1`, data.idrecord).Scan(&namerecord, &datarecord, &datatype)

	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	datarecordbyte, _ := hex.DecodeString(datarecord)
	log.Info().Msgf("Data record row %v extracted successfully.", data.idrecord)
	return string(datarecordbyte), datatype
}

func (data dataRecords) getnamerecord() (namerecord string) {
	//var id int32
	//err := PGdb.QueryRow(context.Background(), `INSERT into data(namerecord, login_fkey) values('namerecord','333') RETURNING (id)`).Scan(&recordID)
	//err := PGdb.QueryRow(context.Background(), `SELECT users.fek FROM users WHERE login='333' AND password='password'`).Scan(&namerecord)
	err := PGdb.QueryRow(context.Background(), `SELECT data.namerecord FROM data WHERE id=$1`, data.idrecord).Scan(&namerecord)
	//_, err := PGdb.Query(context.Background(), `SELECT data.id, data.namerecord FROM data`)
	if err != nil {
		log.Error().Msgf(err.Error())
		}

	log.Info().Msgf("Data name record row %v extracted successfully.", data.idrecord)
	return namerecord
}

