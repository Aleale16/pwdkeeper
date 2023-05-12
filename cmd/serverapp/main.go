package main

import (
	"pwdkeeper/internal/app/grpcserver"
	"pwdkeeper/internal/app/initconfig"

	"github.com/rs/zerolog/log"
)

func main() {
	initconfig.SetinitVars()
	
	log.Info().Msg("Starting gRPC server...")
	grpcserver.Grpcserverstart()
}
