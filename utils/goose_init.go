package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// setup goose for user's project
// goose - handles db migrations
func GooseInit(projectName, dbType string) error {
	if dbType == "" || (dbType != "postgresql" && dbType != "mysql" && dbType != "sqlite") {
			return fmt.Errorf("sqlc: error database type should be specified supported (postgresql, mysql, sqlite)")
		}

	// goose install
	gooseInstlCmd := exec.Command("go", "install", "github.com/pressly/goose/v3/cmd/goose@latest")
	gooseInstlCmd.Dir = filepath.Join(".", projectName)
	gooseInstlCmd.Stdout = nil
	gooseInstlCmd.Stderr = os.Stderr
	if err := gooseInstlCmd.Run(); err != nil {
		return err
	}


	return nil
}