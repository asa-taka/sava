package main

import (
	"log"
	"regexp"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	confFile := flag.StringP("config", "c", "", "config file (default is $PWD/sava.yaml)")
	flag.StringP("port", "p", "3000", "config file (default is 3000)")
	flag.StringP("host", "H", "localhost", "config file (default is localhost)")
	flag.StringP("data-dir", "d", "data", "config file (default is data)")
	flag.Parse()

	if *confFile != "" {
		viper.SetConfigFile(*confFile)
	} else {
		viper.SetConfigName("sava.yaml")
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}

	viper.BindPFlags(flag.CommandLine)

	app := newApp(&appConfig{
		dataDir:               viper.GetString("data-dir"),
		corsAllowOriginRegexp: *regexp.MustCompile("localhost"),
	})
	log.Fatal(app.Listen(viper.GetString("host") + ":" + viper.GetString("port")))
}
