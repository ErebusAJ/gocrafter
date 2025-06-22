package utils

import (
	"os/exec"
	"path/filepath"

	"github.com/ErebusAJ/gocrafter/internal"
)

// executes api template specific commands and installations
func ApiInit(projectName, projectModule string) error {
	// generate main.go
	if err := internal.GenerateFiles(projectName, "main.go", "api/api_main.go.tmpl", 
	struct{
		ProjectName string 
		ProjectModule string
	}{
		ProjectName: projectName, 
		ProjectModule: projectModule,
	}); err != nil {
		return err
	}

	// generate routes and test handler file
	outputDir := filepath.Join(projectName, "internal", "handlers")
	if err := internal.GenerateFiles(outputDir, "routes.go", "api/routes.go.tmpl", nil); err != nil {
		return err
	}

	if err := internal.GenerateFiles(outputDir, "greetings.go", "api/greetings.go.tmpl", nil); err != nil {
		return err
	}

	// install required packages
	if err := installPackage(projectName, "github.com/gin-gonic/gin@latest"); err != nil {
		return err
	} 

	if err := installPackage(projectName, "github.com/joho/godotenv@latest"); err != nil {
		return err
	} 

	return nil
}

func installPackage(projectName, packageName string) error {
	cmd := exec.Command("go", "get", packageName)
	cmd.Dir = projectName
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}