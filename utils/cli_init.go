package utils

import (
	"os/exec"
	"path/filepath"
)

// executes cli template specific commands and installations
func CliInit(projectName string) error {

	// install cobra 
	cmd := exec.Command("go", "get", "github.com/spf13/cobra@latest")
	cmd.Dir = filepath.Join(projectName)
	cmd.Stderr = nil
	cmd.Stdout = nil
	if err := cmd.Run(); err != nil {
		return err
	}

	// initialize cobra 
	initCmd := exec.Command("cobra", "init")
	initCmd.Dir = filepath.Join(projectName)
	cmd.Stderr = nil
	cmd.Stdout = nil
	if err := initCmd.Run(); err != nil {
		return err
	}

	return nil
}