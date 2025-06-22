package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/ErebusAJ/gocrafter/internal"
)

// setup magefile for cli commands automations
func MageInit(projectName, db string, useDocker, useGoose bool) error {
	fmt.Println("setting up magefile...")

	// install go mage
	cmd := exec.Command("go", "install", "github.com/magefile/mage@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return  err
	}

	
	// set correct goose driver
	var dbType string
	if useGoose {
		switch db {
		case "postgresql":
			dbType = "postgres"
		
		case "mysql":
			dbType = "mysql"
		
		case "sqlite":
			dbType = "sqlite3"
		}
	}
	
	// construct template embedding struct
	data := struct {
		ProjectName string
		UseDocker	bool
		UseGoose 	bool
		DbType 		string
		OsWin		string
	}{
		ProjectName: projectName,
		UseDocker: useDocker,
		UseGoose: useGoose,
		DbType: dbType,
		OsWin: runtime.GOOS,
	}

	// creat magefile
	if err := internal.GenerateFiles(projectName, "magefile.go", "build/magefile.go.tmpl", data); err != nil {
		return  err
	}

	fmt.Println("magefile setup complete!!")

	return nil
} 
