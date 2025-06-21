package utils

import (
	"fmt"

	"github.com/ErebusAJ/gocrafter/internal"
)

// setup magefile for cli commands automations
func MageInit(projectName, db string, useDocker, useGoose bool) error {
	fmt.Println("setting up magefile...")

	
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
	}{
		ProjectName: projectName,
		UseDocker: useDocker,
		UseGoose: useGoose,
		DbType: dbType,
	}

	// creat magefile
	if err := internal.GenerateFiles(projectName, "magefile.go", "build/magefile.go.tmpl", data); err != nil {
		return  err
	}

	fmt.Println("magefile setup complete!!")

	return nil
} 
