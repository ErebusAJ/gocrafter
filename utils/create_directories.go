package utils

import (
	"os"
	"path/filepath"
)

// creates all the necessary directories for initial project
func CreateDirectories(projectName, template string) error {
	// setting up directories to make
	var dirList []string
	if template == "cli" {
		dirList = []string{
			"cmd",
			"pkg",
			"config",
			"internal",
			filepath.Join("internal", "utils"),
			filepath.Join("internal", "templates"),
		}
	} else {
		dirList = []string{
			"cmd",
			"pkg",
			"api",
			"config",
			"db",
			"migrations",
			"internal",
			filepath.Join("internal", "handlers"),
			filepath.Join("internal", "middlewares"),
			filepath.Join("internal", "utils"),
		}
	}

	// create project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return err
	}

	// create sub directories
	
	for _, dir := range dirList {
		err = CreateSubDir(projectName, dir, 0755)
		if err != nil {
			return err
		}
	}
	

	return nil
}


// helper to create sub directories
// ./projectName/newDir
func CreateSubDir(projectName, dirName string, perm os.FileMode) error {
	err := os.Mkdir(filepath.Join(".", projectName, dirName), perm)
	if err != nil {
		return err
	}

	return nil
}