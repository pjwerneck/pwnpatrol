package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pjwerneck/pwnpatrol/pwnpatrolmain"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwnpatrol",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	cobra.OnInitialize(pwnpatrolmain.ConnectDB)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will
	// be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwnpatrol.json)")

	// Cobra also supports local flags, which will only run when this
	// action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pwnpatrol" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pwnpatrol")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.SetDefault("host", "0.0.0.0")
	viper.SetDefault("port", 8007)
	viper.SetDefault("logLevel", "INFO")
	viper.SetDefault("readTimeout", 30)
	viper.SetDefault("writeTimeout", 30)

	viper.BindEnv("dumpFile", "PWNPATROL_DUMPFILE")
	viper.BindEnv("dbFile", "PWNPATROL_DBFILE")
	viper.BindEnv("host", "PWNPATROL_HOST")
	viper.BindEnv("port", "PWNPATROL_PORT")
	viper.BindEnv("logLevel", "PWNPATROL_LOGLEVEL")

}
