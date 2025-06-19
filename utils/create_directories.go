package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// creates all the necessary directories for initial project
func CreateDirectories(projectName string) error {
	// create project directory
	err := os.Mkdir(fmt.Sprintf("%v", projectName), 0755)
	if err != nil {
		return err
	}

	// create sub directories
	err = CreateSubDir(projectName, "cmd", 0755)
	if err != nil {
		return err
	}

	err = CreateSubDir(projectName, "pkg", 0755)
	if err != nil {
		return err
	}

	err = CreateSubDir(projectName, "api", 0755)
	if err != nil {
		return err
	}
	
	err = CreateSubDir(projectName, "config", 0755)
	if err != nil {
		return err
	}

	err = CreateSubDir(projectName, "db", 0755)
	if err != nil {
		return err
	}

	err = CreateSubDir(projectName, "migrations", 0755)
	if err != nil {
		return err
	}

	// internal directory and it's sub directories

	err = CreateSubDir(projectName, "internal", 0755)
	if err != nil {
		return err
	}

	err = CreateSubDir(projectName, filepath.Join("internal", "handlers"), 0755)
	if err != nil {
		return err
	}

	err = CreateSubDir(projectName, filepath.Join("internal", "middlewares"), 0755)
	if err != nil {
		return err
	}

	err = CreateSubDir(projectName, filepath.Join("internal", "utils"), 0755)
	if err != nil {
		return err
	}
	

	return nil
}


// helper to create sub directories
func CreateSubDir(projectName, dirName string, perm os.FileMode) error {
	err := os.Mkdir(filepath.Join(".", projectName, dirName), perm)
	if err != nil {
		return err
	}

	return nil
}