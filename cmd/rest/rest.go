package rest

import (
	"github.com/joho/godotenv"
	"github.com/julo/walletsvc/bootstrap"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	serveRESTCmd = &cobra.Command{
		Use:              "rest",
		Short:            "REST API",
		Long:             "run rest api",
		PersistentPreRun: rootPreRun,
		RunE:             runREST,
	}
)

func rootPreRun(cmd *cobra.Command, args []string) {
	err := godotenv.Load()
	if err != nil {
		log.Err(err).Msg("error load env")
	}
}

func runREST(cmd *cobra.Command, args []string) error {
	bootstrap.NewRest().
		RegisterHandler().
		Serve()

	return nil
}

func ServeRESTCmd() *cobra.Command {
	return serveRESTCmd
}
