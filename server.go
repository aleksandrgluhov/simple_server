package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Log handler
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// Root handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	responseStructure := &Response{
		ID:        cfg.ID,
		Payload:   "Greeting from the Autoscale Example Application!",
		Timestamp: time.Now(),
	}

	responseJSON, _ := json.Marshal(responseStructure)
	fmt.Fprintf(w, string(responseJSON))
}

// Web server
func startServer(cfg *ApplicationConfig) {
	http.HandleFunc("/", rootHandler)
	log.Printf("Starting server at: http://%s:%s/ \n", cfg.BindHost, cfg.BindPort)
	log.Printf("Press Ctrl+C to stop \n\n")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.BindHost, cfg.BindPort), logRequest(http.DefaultServeMux)); err != nil {
		processExceptionHard(err)
	}
}
