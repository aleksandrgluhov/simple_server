package main

// imports
import (
	_ "github.com/joho/godotenv/autoload"
)

// constants
const ConstVersion string = "1.0"
const ConstApplicationPrefix string = "SIMPLE_SERVER"
const ConstEnvConcatPattern string = "%s_%s"

var cfg ApplicationConfig

// program entry point
func main() {
	loadConfigFromEnv(&cfg)
	showConfig(&cfg)
	startServer(&cfg)
}
