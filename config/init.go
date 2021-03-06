package config

import (
	"database/sql"
	"fmt"
	"github.com/Firmansyah845/go_hackaton/utils/logger"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"

	"go.elastic.co/apm/module/apmsql"

	_ "go.elastic.co/apm/module/apmsql/pq"
)

var (
	App *Application
)

type (
	Application struct {
		Name   string  `json:"name"`
		Port   string  `json:"port"`
		Config Config  `json:"app_config"`
		DB     *sql.DB `json:"db"`
	}
	Config struct {
		Port                 string `envconfig:"PORT"`
		DB_Host              string `envconfig:"DB_HOST"`
		DB_Username          string `envconfig:"DB_USERNAME"`
		DB_Port              int    `envconfig:"DB_PORT"`
		DB_Password          string `envconfig:"DB_PASSWORD"`
		DB_Name              string `envconfig:"DB_NAME"`
		URL_FORECAST_SERVICE string `envconfig:"URL_FORECAST_SERVICE"`
	}
)

// Initiate news instances
func init() {
	var err error
	App = &Application{}
	App.Name = "go_hackaton"

	if err = App.LoadConfigs(); err != nil {
		log.Printf("Load config error : %v", err)
		os.Exit(1)
	}

	InitLogger() // initialize logger

	if err = App.DBinit(); err != nil {
		log.Printf("DB init error : %v", err)
		os.Exit(1)
	}

	App.Port = App.Config.Port
}

func (x *Application) Close() (err error) {
	if err = x.DB.Close(); err != nil {
		return err
	}

	return nil
}

// LoadConfigs load configuration from yaml
func (x *Application) LoadConfigs() error {

	err := envconfig.Process("", &x.Config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// DBinit initialize database connection
func (x *Application) DBinit() error {
	dbconf := x.Config

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		dbconf.DB_Host, dbconf.DB_Port, dbconf.DB_Username, dbconf.DB_Password, dbconf.DB_Name)

	db, err := apmsql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err // proper error handling instead of panic
	}
	x.DB = db
	return nil
}

// InitLogger initialize logger instance
func InitLogger() {
	logConfig := logger.Configuration{
		EnableConsole:     true,    // next, get from configuration
		ConsoleJSONFormat: true,    // next, get from configuration
		ConsoleLevel:      "debug", // next, get from configuration
	}

	if err := logger.NewLogger(logConfig, logger.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}
}
