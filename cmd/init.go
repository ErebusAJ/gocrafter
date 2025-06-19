package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ErebusAJ/gocrafter/internal"
	"github.com/ErebusAJ/gocrafter/utils"
	"github.com/spf13/cobra"
)

var name   string
var module string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initiate a new Go project",
	Long: `Initializes a new Go project with added configuration using flags`,
	RunE: cmdRun,
}


func init() {
	rootCmd.AddCommand(initCmd)


	// project name
	initCmd.Flags().StringVarP(&name, "name", "n", "", "project name")
	initCmd.MarkFlagRequired("name")

	// github module
	initCmd.Flags().StringVarP(&module, "module", "m", "", "go module path")
	initCmd.MarkFlagRequired("module")
}

func cmdRun(cmd *cobra.Command, args []string) error {
	
	// check malformed module path
	if !strings.HasPrefix(module, "github.com/") {
		return fmt.Errorf("init: flag --module should be in format github.com/user/myapp")
	}

	// create all the initial directories 
	if err := utils.CreateDirectories(name); err != nil {
		return err
	}

	modInitCmd := exec.Command("go", "mod", "init", module)
	modInitCmd.Stdout = os.Stdout
	modInitCmd.Stderr = os.Stderr
	if err := modInitCmd.Run(); err != nil {
		return err
	}

	// create main.go entry point
	if err := internal.GenerateFiles(name, "main.go.tmpl", struct{ProjectName string}{ProjectName: name}); err != nil {
		return err
	}

	

	fmt.Printf("Your Go project: %v has been initialized. \n", name)
	fmt.Println("Thank You! For using gocrafter")
	return nil
}


