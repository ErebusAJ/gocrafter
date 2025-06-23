package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ErebusAJ/gocrafter/internal"
	"github.com/ErebusAJ/gocrafter/utils"
	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
)

// falgs variables
var name   		string
var module 		string
var template	string
var db			string
var useSqlc		bool
var useGoose	bool
var useDocker	bool
var magefile 	bool


// config struct and file path var
var configFilePath string

type InitConfig struct {
	Name   		string	`yaml:"name"`
 	Module 		string	`yaml:"module"`
 	Template	string	`yaml:"template"`
	Db			string	`yaml:"db"`
	UseSqlc		bool	`yaml:"sqlc"`
	UseGoose	bool	`yaml:"goose"`
	UseDocker	bool	`yaml:"docker"`
	Magefile 	bool	`yaml:"magefile"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initiate a new Go project",
	Long: `Initializes a new Go project with added configuration using flags`,
	RunE: cmdRun,
}


func init() {
	rootCmd.AddCommand(initCmd)

	// config file path
	initCmd.Flags().StringVar(&configFilePath, "config", "", "path to YAML config file")

	// project name
	initCmd.Flags().StringVarP(&name, "name", "n", "", "project name")

	// github module
	initCmd.Flags().StringVarP(&module, "module", "m", "", "go module path")

	// template type
	initCmd.Flags().StringVarP(&template, "template", "t", "base", "type of go project api, cli defaults to base")

	// db type
	initCmd.Flags().StringVarP(&db, "database", "d", "postgresql", "name of database to be used")

	// sqlc flag
	initCmd.Flags().BoolVar(&useSqlc, "sqlc", false, "setup sqlc for sql query generation")

	// goose flag
	initCmd.Flags().BoolVar(&useGoose, "goose", false, "setup goose for database migrations")

	// docker flag
	initCmd.Flags().BoolVar(&useDocker, "docker", false, "setup docker for containerization")

	// magefile flag
	initCmd.Flags().BoolVar(&magefile, "magefile", false, "setup magefile for cli commands automation")
}


// run function for initCmd
func cmdRun(cmd *cobra.Command, args []string) error {

	// config data
	var config InitConfig

	// check for config file path
	if configFilePath != "" {
		if err := configParse(&config); err != nil {
			return err
		}
		if err := initRun(config); err != nil {
			return err
		}
	} else {
		if name == "" {
			return  fmt.Errorf("name: error required flag name not provided")
		}
		if module == "" {
			return  fmt.Errorf("module: error required flag module not provided") 
		}

		config = InitConfig{
			Name: name,
			Module: module,
			Template: template,
			Db: db,
			UseSqlc: useSqlc,
			UseGoose: useGoose,
			UseDocker: useDocker,
			Magefile: magefile,
		}
		if err := initRun(config); err != nil {
			return  err
		}
	}

	fmt.Printf("\n🎉 Your Go project: %v has been initialized. \n", config.Name)
	fmt.Println("Thank You! For using gocrafter")
	return nil
}


// core function of initCmd run
func initRun(config InitConfig) error {
	// check malformed module path
	if !strings.HasPrefix(config.Module, "github.com/") {
		return fmt.Errorf("init: flag --module should be in format github.com/user/myapp")
	}

	// create all the initial directories 
	if err := utils.CreateDirectories(config.Name, config.Template); err != nil {
		return err
	}

	// go mod initialize cmd
	modInitCmd := exec.Command("go", "mod", "init", config.Module)
	modInitCmd.Dir = filepath.Join(".", config.Name)
	modInitCmd.Stdout = nil
	modInitCmd.Stderr = nil
	if err := modInitCmd.Run(); err != nil {
		return err
	}

	fmt.Println()

	// create necessary files
	if err := utils.ProgressTask("Setting up directories", func() error {
		return utils.CreateFiles(config.Name)
	}); err != nil {
		return err
	}

	// case for different types of initializations template specific
	switch config.Template {
	case "base":
		fmt.Println("init: flag -t not provided defaulting to base")
		if err := internal.GenerateFiles(config.Name, "main.go", "base/main.go.tmpl", struct{ProjectName string}{ProjectName: config.Name}); err != nil {
			return err
		}

	case "api" :
		if err := utils.ProgressTask("Setting up API template", func() error {
			return utils.ApiInit(config.Name, config.Module)
		}); err != nil  {
			return err
		}
	
	case "cli" :
		if err := utils.ProgressTask("Setting up CLI template", func() error {
			return utils.CliInit(config.Name)
		}); err != nil {
			return err
		}
	
	default:
		fmt.Println("init: flag -t malformed accpeted (api/cli)")
	}

	// optinal setups
	// sqlc
	if useSqlc{
		if err := utils.ProgressTask("Initializing sqlc", func() error {
			return utils.SqlcInit(config.Name, config.Db)
		}); err != nil {
			return err
		}
	}

	// goose
	if useGoose{
		if err := utils.ProgressTask("Initializing goose", func() error {
			return utils.GooseInit(config.Name, config.Db)
		}); err != nil {
			return err
		}
	}
	
	// docker
	if useDocker{
		if err := utils.ProgressTask("Setting up Dockerfile", func() error {
			return utils.DockerInit(config.Name, config.Template); 
		}); err != nil {
			return err
		}
	}
	
	// magefile
	if magefile{
		if err := utils.ProgressTask("Configuring magefile", func () error {
			return utils.MageInit(config.Name, config.Db, config.UseDocker, config.UseGoose)
		}); err != nil {
			return err
		}
	}

	// go mod tidy cmd
	goTidy := exec.Command("go", "mod", "tidy")
	goTidy.Dir = filepath.Join(".", config.Name)
	goTidy.Stdout = nil
	goTidy.Stderr = nil
	if err := goTidy.Run(); err != nil {
		return err
	}

	// go mod vendor cmd
	goVendor := exec.Command("go", "mod", "vendor")
	goVendor.Dir = filepath.Join(".", config.Name)
	goVendor.Stdout = nil
	goVendor.Stderr = nil
	if err := goVendor.Run(); err != nil {
		return err
	}

	return  nil
}


// config yaml parse 
func configParse(config *InitConfig) error {

	// read yaml file
	if !strings.HasSuffix(configFilePath, "gocrafter.yaml") {
		return fmt.Errorf("config: error need gocrafter.yaml found: %v", configFilePath)
	}
	readb, err := os.ReadFile(configFilePath)
	if err != nil {
		return  fmt.Errorf("config: could not read config file - %v", err)
	}

	// unmarshal
	if err := yaml.Unmarshal(readb, &config); err != nil {
		return  err
	}

	return nil
}
