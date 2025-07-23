package main

import (
	"log"
	"os"
)

func logMessage(message string) {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()

	// Create a logger that writes to the file
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Log a message
	logger.Println(message)
}
