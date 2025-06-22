package utils

import (
	"runtime"

	"github.com/ErebusAJ/gocrafter/internal"
)

// setup a Dockerfile based on user's project
// for containerization
func DockerInit(projectName, tmplType string) error {

	// setting up data to embed in template file according to template type
	var data interface{}
	if tmplType == "api" {
		data = struct {
			ProjectName string
			ExposePort	string
		}{
			ProjectName: projectName,
			ExposePort: `# Expose server port
EXPOSE 8080`,
		}
	} else {
		data = struct {
			ProjectName string
		} {
			ProjectName: projectName,
		}
	}

	//template os check
	var template string
	if runtime.GOOS == "windows" {
		template = "docker/ Dockerfile.windows.tmpl"
	} else {
		template = "docker/ Dockerfile.linux.tmpl"
	}

	// generate docker file
	if err := internal.GenerateFiles(projectName, "Dockerfile", template, data); err != nil {
		return  err
	}
	
	return  nil
}