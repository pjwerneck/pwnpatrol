package cmd

import (
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
		pwnpatrolmain.InitDB(dumpFile)
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)

	initdbCmd.Flags().StringVarP(&dumpFile, "dumpfile", "d", "", "passwords list file")
	initdbCmd.MarkFlagRequired("dumpfile")

}
