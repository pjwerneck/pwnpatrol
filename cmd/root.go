package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pjwerneck/pwnpatrol/pwnpatrolmain"
)

var dumpFile string
var dbFile string
var host string
var port int




// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwnpatrol",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		pwnpatrolmain.ConnectDB(dbFile)
	},
}

// Execute adds all child commands to the root command and sets flags
// appropriately.  This is called by main.main(). It only needs to
// happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will
	// be global for your application.
	rootCmd.PersistentFlags().StringVar(&dbFile, "dbfile", "./pwnpatrol.db", "database file")

	// Cobra also supports local flags, which will only run when this
	// action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetDefault("readTimeout", 30)
	viper.SetDefault("writeTimeout", 30)
}
