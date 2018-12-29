package cmd

import (
	"github.com/pjwerneck/pwnpatrol/pwnpatrolmain"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP API server",

	Run: func(cmd *cobra.Command, args []string) {
		pwnpatrolmain.Main()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
