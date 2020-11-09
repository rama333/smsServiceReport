package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
)

var Config appConfig

type appConfig struct {
	DB               *sqlx.DB
	DBerr            error
	RESTAPIPort      int     `mapstructure:"rest_api_port"`
	DBURL            string  `mapstructure:"db_url"`
	DIAGPORT         int     `mapstructure:"diag_port"`
	RABBITMQWAITTIME float64 `mapstructure:"rabbit_wait_time"`
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("restful")

	v.AutomaticEnv()

	//Config.RESTAPIPort = v.Get("8080").(string)
	//Config. = v.Get("API_KEY").(string)
	v.SetDefault("rest_api_port", 8080)
	v.SetDefault("diag_port", 8081)
	v.SetDefault("db_url", "tcp://192.168.114.145:9000?debug=true")
	v.SetDefault("rabbit_wait_time", 60)

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	err := v.ReadInConfig()
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}

	log.Println("veper", v.AllKeys())

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
