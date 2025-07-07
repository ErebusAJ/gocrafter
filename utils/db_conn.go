package utils

import (
	"fmt"
	"path/filepath"

	"github.com/ErebusAJ/gocrafter/internal"
)

func DBConnInit(projectName, dbType string) error {
	type templateData struct {
		DBDriver 		string
		DriverPackage	string
	}

	var data templateData

	switch dbType {
	case "postgresql":
		data = templateData{
			DBDriver: "postgres",
			DriverPackage: "github.com/lib/pq",
		}
	
	case "mysql":
		data = templateData{
			DBDriver: "mysql",
			DriverPackage: "github.com/go-sql-driver/mysql",
		}

	case "sqlite":
		data = templateData{
			DBDriver: "sqlite3",
			DriverPackage: "github.com/mattn/go-sqlite3",
		}
	}

	// // install go dot env
	// cmd := exec.Command("go", "get", "github.com/joho/godotenv@latest")
	// cmd.Stderr = nil
	// cmd.Stdout = nil
	// err := cmd.Run(); if err != nil {
	// 	return fmt.Errorf("error installing godotenv dependency")
	// }

	// generate util db connection file
	err := internal.GenerateFiles(filepath.Join(projectName, "internal", "utils"), "db_conn.go", "api/db_conn.go.tmpl", data); if err != nil {
		return fmt.Errorf("error generating file %v", err)
	}

	return nil
}