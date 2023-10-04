package main

import (
	"log"
	"regexp"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	confFile = flag.String("config", "", "config file (default is $PWD/sava.yaml)")
	port     = flag.String("port", "3000", "config file (default is $PWD/sava.yaml)")
	host     = flag.String("host", "localhost", "config file (default is $PWD/sava.yaml)")
)

func main() {
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

	viper.SetDefault("port", 3000)
	viper.SetDefault("host", "localhost")
	viper.SetDefault("data-dir", "data")

	app := newApp(&appConfig{
		dataDir:               viper.GetString("data-dir"),
		corsAllowOriginRegexp: *regexp.MustCompile("localhost"),
	})
	log.Fatal(app.Listen(viper.GetString("host") + ":" + viper.GetString("port")))
}
