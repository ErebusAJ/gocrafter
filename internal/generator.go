package internal

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*/*.tmpl
var tmplFS embed.FS


// generates file from given template
func GenerateFiles(outputDir, filename, tmplUse string, data interface{}) error {
	tmpl, err := template.ParseFS(tmplFS, fmt.Sprintf("templates/%v", tmplUse))
	if err != nil {
		return err
	}

	outPath := filepath.Join(outputDir, filename)
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()


	return tmpl.Execute(outFile, data)
}