package cmd

import (
	"fmt"

	"github.com/pjwerneck/pwnpatrol/pwnpatrolmain"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP API server",

	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%v:%v", host, port)

		pwnpatrolmain.Main(addr)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&host, "host", "0.0.0.0", "server ip address to bind to")
	serveCmd.Flags().IntVar(&port, "port", 8007, "server port")
}
