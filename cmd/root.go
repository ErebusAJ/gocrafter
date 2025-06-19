package cmd

import (
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocrafter",
	Short: "scaffolds go project quickly",
	Long: `GoCrafter is a tool to give your Go project a headstart.
Making it easy to setup a Go project with sqlc, goose, database connections, and JWT auth setup.

Save yourself the repetitive task of setting up same things. 
`,
	Run: func(cmd *cobra.Command, args []string) { 
		if len(os.Args) == 1 {
			figure.NewColorFigure("Go-Crafter", "isometric1", "cyan", true).Print()
		}
	},
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


