package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gorelease var
var (
	version = "dev"
	date = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version of gocrafter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %v \nBuilt at: %v \n", version, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
