package main

import "log"

func processExceptionHard(err error) {
	log.Fatalf("Unrecoverable error occured: %s", err)
}

func processExceptionSoft(err error) {
	log.Println(err)
}
