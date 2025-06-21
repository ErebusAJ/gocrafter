package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// setup goose for user's project
// goose - handles db migrations
func GooseInit(projectName string) error {
	fmt.Println("Setting up goose...")

	// goose install
	gooseInstlCmd := exec.Command("go", "install", "github.com/pressly/goose/v3/cmd/goose@latest")
	gooseInstlCmd.Dir = filepath.Join(".", projectName)
	gooseInstlCmd.Stdout = os.Stdout
	gooseInstlCmd.Stderr = os.Stderr
	if err := gooseInstlCmd.Run(); err != nil {
		return err
	}

	fmt.Println("goose setup compelete!!")

	return nil
}