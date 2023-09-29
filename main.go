package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("sava")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}

	viper.SetDefault("port", 3000)
	viper.SetDefault("data-dir", "data")

	app := newApp(&appConfig{
		dataDir: viper.GetString("data-dir"),
	})
	log.Fatal(app.Listen(":" + viper.GetString("port")))
}
