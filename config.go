package puppetquery

import (
	"log"
	"os"
	"path/filepath"

	"github.com/nightlyone/ini"
)

// Locations of site/global config and per user config.
var (
	GlobalConfig = "/etc/puppetquery/config.ini"
	UserConfig   = ".config/puppetquery/config.ini"
)

// Endpoint, if we didn't configure any
var DefaultEndpoint = "http://localhost:8080"

/// This part does the puppet endpoint autoconfiguration
var config *ini.File
var endpoint = DefaultEndpoint

func init() {
	config = loadConfig()
	if puppetdb, ok := config.Global["url"]; ok {
		endpoint = puppetdb
	}
}

func loadConfig() *ini.File {
	if home := os.Getenv("HOME"); home != "" {
		c, err := tryLoadConfig(filepath.Join(home, UserConfig))
		if err != nil {
			log.Println("ERROR: reading per user config: ", err)
			return new(ini.File)
		}
		if c != nil {
			return c
		}
	}
	c, err := tryLoadConfig(GlobalConfig)
	if err != nil {
		log.Println("ERROR: reading global config: ", err)
		return new(ini.File)
	}
	if c != nil {
		return c
	}
	return new(ini.File)
}

func tryLoadConfig(filename string) (*ini.File, error) {
	c, err := ini.ReadFile(filename)
	switch {
	case err == nil:
		return c, nil
	case os.IsNotExist(err):
		return nil, nil
	default:
		return nil, err
	}
}
