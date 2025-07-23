package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type RoundTrip struct {
	location string
	to       int
	from     int
}

func main() {
	godotenv.Load()

	logFilePath := os.Getenv("LOG_FILE_PATH")
	if logFilePath == "" {
		fmt.Println("error: no log file path defined")
		return
	}

	apiKey := os.Getenv("MAPS_API_KEY")
	if apiKey == "" {
		logMessage("error: no maps api key defined", logFilePath)
		return
	}

	target := os.Getenv("TARGET")
	if apiKey == "" {
		logMessage("error: no target defined", logFilePath)
		return
	}

	locationsFilePath := os.Getenv("LOCATIONS_FILE_PATH")
	if locationsFilePath == "" {
		logMessage("error: no locations file path defined", logFilePath)
		return
	}

	toFilePath := os.Getenv("TO_FILE_PATH")
	if toFilePath == "" {
		logMessage("error: no to file path defined", logFilePath)
		return
	}

	fromFilePath := os.Getenv("FROM_FILE_PATH")
	if fromFilePath == "" {
		logMessage("error: no from file path defined", logFilePath)
		return
	}

	locations, err := getLocations(locationsFilePath)
	if err != nil {
		logMessage(fmt.Sprintf("Error getting locations: %v\n", err), logFilePath)
		return
	}

	for i, trip := range locations {

		to, err := getCommuteTime(trip.location, target, apiKey)
		if err != nil {
			logMessage(fmt.Sprintf("Error getting commute time from %v to %v: %v\n", trip.location, target, err), logFilePath)
		}

		from, err := getCommuteTime(target, trip.location, apiKey)
		if err != nil {
			logMessage(fmt.Sprintf("Error getting commute time from %v to %v: %v\n", target, trip.location, err), logFilePath)
		}

		locations[i] = RoundTrip{
			location: trip.location,
			to:       to,
			from:     from,
		}
	}

	err = outputTimes(locations, toFilePath, fromFilePath)
	if err != nil {
		fmt.Printf("Error outputting commute times: %v\n", err)
		return
	}

	logMessage("Output successful.", logFilePath)
}

func getLocations(locationsFilePath string) (locations []RoundTrip, err error) {

	locByte, err := os.ReadFile(locationsFilePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	locArray := strings.Split(string(locByte), "\n")

	locations = []RoundTrip{}

	for i := 0; i < len(locArray); i++ {
		locations = append(locations, RoundTrip{
			location: locArray[i],
		})
	}

	return
}

func outputTimes(locations []RoundTrip, toFilePath, fromFilePath string) error {
	toRow := []string{time.Now().String()}
	fromRow := []string{time.Now().String()}

	for _, trip := range locations {
		toRow = append(toRow, strconv.Itoa(trip.to))
		fromRow = append(fromRow, strconv.Itoa(trip.from))
	}

	err := writeFile(toRow, toFilePath)
	if err != nil {
		return err
	}

	err = writeFile(fromRow, fromFilePath)
	if err != nil {
		return err
	}

	return nil

}

func writeFile(row []string, target string) error {

	file, err := os.OpenFile(target, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.UseCRLF = true

	err = writer.Write(row)
	if err != nil {
		return err
	}

	writer.Flush()

	err = writer.Error()
	if err != nil {
		return err
	}

	return nil
}
