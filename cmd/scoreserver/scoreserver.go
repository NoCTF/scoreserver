package main

import (
	"log"
	"os"
	"regexp"

	flags "github.com/jessevdk/go-flags"
	"github.com/noctf/scoreserver"
	"github.com/noctf/scoreserver/db"
	"github.com/spf13/viper"
)

type options struct {
	ConfigFile string `short:"c" long:"config" description:"path to config file without extension"`
	Port       string `short:"p" long:"port" default:":8080" description:"Listen Port"`
}

func main() { os.Exit(_main()) }

func _main() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Printf("%s", err)
		return 1
	}
	cf := opts.ConfigFile
	if cf == "" {
		if err := db.Init(""); err != nil {
			log.Printf("%s", err)
			return 1
		}
	} else {
		viper.SetConfigName(opts.ConfigFile)
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.scoreserver")
		err := viper.ReadInConfig()
		if err != nil {
			log.Printf("%s", err)
			return 1
		}
		if err := db.Init(viper.GetString("database.DSN")); err != nil {
			log.Printf("%s", err)
			return 1
		}
		if p := viper.GetString("port"); p != "" {
			opts.Port = p
		}
	}
	portPrefix, err := regexp.Compile(`:\d+$`)
	if err != nil {
		log.Printf("%s", err)
		return 1
	}
	if !portPrefix.MatchString(opts.Port) {
		opts.Port = ":" + opts.Port
	}
	log.Printf("Server listening on %s", opts.Port)
	if err := scoreserver.Run(opts.Port); err != nil {
		log.Printf("%s", err)
		return 1
	}
	return 0
}
