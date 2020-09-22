package main

import (
	"log"
	"os"

	"github.com/dchest/uniuri"
	"github.com/kelseyhightower/envconfig"
)

func loadConfigFromEnv(cfg *ApplicationConfig) {
	err := envconfig.Process(ConstApplicationPrefix, cfg)
	if err != nil {
		processExceptionSoft(err)
	}

	if !(cfg.IDLength > 0) {
		cfg.IDLength = 3
	}
	cfg.ID = uniuri.NewLen(cfg.IDLength)
}

func showConfig(c *ApplicationConfig) {
	cCopy := *c
	log.Printf("Application: %s\n", ConstApplicationPrefix)
	log.Printf("Version: %s\n", ConstVersion)
	log.Printf("-----------------------------------\n\n")

	if cCopy.Debug {
		cwd, err := os.Getwd()
		if err != nil {
			processExceptionHard(err)
		}
		log.Printf("Current working directory: %s\n", cwd)
		log.Printf("Current application configuration: ")
		log.Printf("%+v\n\n", cCopy)
	}
}
