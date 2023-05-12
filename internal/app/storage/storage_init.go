package storage

import (
	"context"
	"fmt"
	"os"
	"pwdkeeper/internal/app/initconfig"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func Initdb() {
	//----------------------------//
	//Подключаемся к СУБД postgres
	//----------------------------//
	urlExample := "postgres://postgres:1@localhost:5432/pwdkeeper"
	os.Setenv("DATABASE_DSN", urlExample)
	initconfig.PostgresDBURL = urlExample
	if initconfig.PostgresDBURL != "" {
		poolConfig, err := pgxpool.ParseConfig(initconfig.PostgresDBURL)
		if err != nil {
			log.Error().Err(err)
			log.Fatal().Msg("Unable to parse DATABASE_DSN")
		}
		log.Debug().Msgf("poolConfig: %v", poolConfig)

		PGdb, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err != nil {
			fmt.Println("ERROR! PGdb NOT OPENED")
			panic(err)
		}
	}
}