package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ErebusAJ/gocrafter/internal"
)

// setup sqlc package for user
// sqlc - is used to generate type safe go queries
func SqlcInit(projectName, dbType string) error {
	if dbType == "" || (dbType != "postgresql" && dbType != "mysql" && dbType != "sqlite") {
		return fmt.Errorf("sqlc: error database type should be specified supported (postgresql, mysql, sqlite)")
	}
	// install sqlc
	sqlcInstlCmd := exec.Command("go", "install", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
	sqlcInstlCmd.Dir = filepath.Join(".", projectName)
	sqlcInstlCmd.Stderr = os.Stderr
	sqlcInstlCmd.Stdout = nil
	if err := sqlcInstlCmd.Run(); err != nil {
		return err
	}

	// create schema directory for database schema files 
	if err := CreateSubDir(projectName, filepath.Join("migrations", "schema"), 0755); err != nil {
		return err
	}
	
	// create queries directory for sql queries files 
	if err := CreateSubDir(projectName, filepath.Join("migrations", "sql"), 0755); err != nil {
		return err
	}

	// create sqlc.yaml file
	if err := internal.GenerateFiles(projectName, "sqlc.yaml", "sqlc/sqlc.yaml.tmpl", struct{DBType string}{DBType: dbType}); err != nil {
		return err
	}

	// create sample table schema
	outputPath := filepath.Join(projectName, "migrations", "schema")
	if err := internal.GenerateFiles(outputPath, "001_users.sql", "sqlc/table_users.sql.tmpl", nil); err != nil {
		return err
	}

	// create sample sql query 
	outputPath = filepath.Join(projectName, "migrations", "sql")
	if err := internal.GenerateFiles(outputPath, "queries_users.sql", "sqlc/sql_user.sql.tmpl", struct{ProjectName string}{ProjectName: projectName}); err != nil {
		return err
	}

	// sqlc generate
	sqlcCmd := exec.Command("sqlc", "generate")
	sqlcCmd.Dir = filepath.Join(".", projectName)
	sqlcCmd.Stderr = os.Stderr
	sqlcCmd.Stdout = nil
	if err := sqlcCmd.Run(); err != nil {
		return err
	}

	return nil
}