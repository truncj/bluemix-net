package cmd

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

func getPath() *string {
	var dir, configPath string
	if dir, err = homedir.Dir(); err != nil {
		log.Fatal(err)
	}
	if len(config) != 0 {
		configPath = config
	} else if _, err = os.Stat(dir); err == nil {
		configPath = dir + "/config"
	} else {
		configPath = "./config"
	}
	return &configPath
}
