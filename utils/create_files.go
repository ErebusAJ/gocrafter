package utils

import "github.com/ErebusAJ/gocrafter/internal"


func CreateFiles(projectName string) error {
	// create .env
	if err := internal.GenerateFiles(projectName, ".env", "base/env.tmpl", nil); err != nil {
		return err
	}

	// create .gitignore
	if err := internal.GenerateFiles(projectName, ".gitignore", "base/.gitignore.tmpl", struct{ProjectName string}{ProjectName: projectName}); err != nil {
		return err
	}

	// create README.md
	if err := internal.GenerateFiles(projectName, "README.md", "base/readme.md.tmpl", struct{ProjectName string}{ProjectName: projectName}); err != nil {
		return err
	}

	return nil
}