package main

import (
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	replacer := strings.NewReplacer(",", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetDefault("application.db.max_idle", 5)
	viper.SetDefault("application.db.max_conn", 10)

	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()

	if err != nil {
		log.Error(err)
	}
}

func main() {
	app := NewApplication()

	app.ListenAndServe()
}
