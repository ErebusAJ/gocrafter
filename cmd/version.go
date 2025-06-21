package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gorelease var
var (
	Version = "dev"
	Date = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version of gocrafter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %v \nBuilt at: %v \n", Version, Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
