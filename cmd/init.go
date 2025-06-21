package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ErebusAJ/gocrafter/internal"
	"github.com/ErebusAJ/gocrafter/utils"
	"github.com/spf13/cobra"
	"github.com/go-yaml/yaml"
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

	fmt.Printf("Your Go project: %v has been initialized. \n", name)
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
	if err := utils.CreateDirectories(config.Name); err != nil {
		return err
	}

	// go mod initialize cmd
	modInitCmd := exec.Command("go", "mod", "init", config.Module)
	modInitCmd.Dir = filepath.Join(".", config.Name)
	modInitCmd.Stdout = os.Stdout
	modInitCmd.Stderr = os.Stderr
	if err := modInitCmd.Run(); err != nil {
		return err
	}

	// create necessary files
	if err := utils.CreateFiles(config.Name); err != nil {
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
		if err := utils.ApiInit(config.Name, config.Module); err != nil  {
			return err
		}
	
	case "cli" :
		if err := utils.CliInit(); err != nil {
			return err
		}
	
	default:
		fmt.Println("init: flag -t malformed accpeted (api/cli)")
	}

	// optinal setups
	if config.UseSqlc {
		if config.Db == "" || (config.Db != "postgresql" && config.Db != "mysql" && config.Db != "sqlite") {
			return fmt.Errorf("sqlc: error database type should be specified supported (postgresql, mysql, sqlite)")
		}
		if err := utils.SqlcInit(config.Name, config.Db); err != nil {
			return err
		}
	}

	if config.UseGoose {
		if config.Db == "" || (config.Db != "postgresql" && config.Db != "mysql" && config.Db != "sqlite") {
			return fmt.Errorf("sqlc: error database type should be specified supported (postgresql, mysql, sqlite)")
		}
		if err := utils.GooseInit(config.Name); err != nil {
			return err
		}
	}
	
	if config.UseDocker {
		if err := utils.DockerInit(config.Name, config.Template); err != nil {
			return err
		}
	}
	
	
	// magefile setup
	if config.Magefile {
		if err := utils.MageInit(config.Name, config.Db, config.UseDocker, config.UseGoose); err != nil {
			return err
		}
	}


	// go mod tidy cmd
	goTidy := exec.Command("go", "mod", "tidy")
	goTidy.Dir = filepath.Join(".", config.Name)
	goTidy.Stdout = os.Stdout
	goTidy.Stderr = os.Stderr
	if err := goTidy.Run(); err != nil {
		return err
	}

	// go mod vendor cmd
	goVendor := exec.Command("go", "mod", "vendor")
	goVendor.Dir = filepath.Join(".", config.Name)
	goVendor.Stdout = os.Stdout
	goVendor.Stderr = os.Stderr
	if err := goVendor.Run(); err != nil {
		return err
	}

	return  nil
}


// config yaml parse 
func configParse(config *InitConfig) error {

	// read yaml file
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
