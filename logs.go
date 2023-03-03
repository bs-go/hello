package main

import (
	"log"
	"os"
)

func Logs() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World")
	}
	log.Panic("Panic: Hello World!")
}
