/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocrafter",
	Short: "GoCrafter",
	Long: `GoCrafter is a tool to give your Go project a headstart.
Making it easy to setup a Go project and save yourself the repetitive task
of setting up same things. 
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		figure.NewColorFigure("Go-Crafter", "isometric1", "cyan", true).Print()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gocrafter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// var help bool


	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().BoolVarP(&help, "help", "h", false, "Help manual for gocrafter")
}


