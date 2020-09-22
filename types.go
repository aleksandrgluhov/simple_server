package main

import "time"

// Data types definition

// ApplicationConfig
type ApplicationConfig struct {
	Debug    bool   `default:"false" envconfig:"debug"`
	IDLength int    `default:3`
	ID       string `default:"None"`
	BindHost string `default:"0.0.0.0" envconfig:"bind_host"`
	BindPort string `default:"8080" envconfig:"bind_port"`
}

// The Response template
type Response struct {
	ID        string    `json:"id"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}
