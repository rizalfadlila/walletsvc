package cmd

import (
	"github.com/julo/walletsvc/cmd/rest"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Short: "Mini Wallet Service",
	}
)

func Execute() {
	registerRestCommand()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("failed to execute : \n", err.Error())
		os.Exit(-1)
	}
}

func registerRestCommand() {
	rootCmd.AddCommand(rest.ServeRESTCmd())
}
