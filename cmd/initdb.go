package cmd

import (
	"fmt"
	"github.com/pjwerneck/pwnpatrol/pwnpatrolmain"
	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Initializes the password database",
	Long: `Initializes the password database using the dumpfile from haveibeenpwned.com

    `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initdb called")
		pwnpatrolmain.InitDB()
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)

}
