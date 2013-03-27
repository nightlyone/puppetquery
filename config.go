package puppetquery

import (
	"code.google.com/p/goconf/conf"
	"log"
	"os"
	"path/filepath"
)

// Locations of site/global config and per user config.
var (
	GlobalConfig = "/etc/puppetquery/config.ini"
	UserConfig   = ".config/puppetquery/config.ini"
)

// Endpoint, if we didn't configure any
var DefaultEndpoint = "http://localhost:8080"

/// This part does the puppet endpoint autoconfiguration
var config *conf.ConfigFile
var endpoint string

func init() {
	config = loadConfig()
	if puppetdb, err := config.GetString("default", "url"); err != nil {
		endpoint = DefaultEndpoint
	} else {
		endpoint = puppetdb
	}
}

func loadConfig() *conf.ConfigFile {
	c := conf.NewConfigFile()
	global, err := os.Open(GlobalConfig)
	if err == nil {
		defer global.Close()
		if err = c.Read(global); err != nil {
			log.Println("ERROR: reading global config: ", err)
			return c
		}
	}

	home := os.Getenv("HOME")
	user, err := os.Open(filepath.Join(home, UserConfig))
	if err == nil {
		defer user.Close()
		if err = c.Read(user); err != nil {
			log.Println("ERROR: reading per user config: ", err)
			return c
		}
	}

	return c
}
