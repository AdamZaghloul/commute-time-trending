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

	apiKey := os.Getenv("MAPS_API_KEY")
	if apiKey == "" {
		logMessage("error: no maps api key defined")
		return
	}

	target := os.Getenv("TARGET")
	if apiKey == "" {
		logMessage("error: no target defined")
		return
	}

	locations, err := getLocations()
	if err != nil {
		logMessage(fmt.Sprintf("Error getting locations: %v\n", err))
		return
	}

	for i, trip := range locations {

		to, err := getCommuteTime(trip.location, target, apiKey)
		if err != nil {
			logMessage(fmt.Sprintf("Error getting commute time from %v to %v: %v\n", trip.location, target, err))
		}

		from, err := getCommuteTime(target, trip.location, apiKey)
		if err != nil {
			logMessage(fmt.Sprintf("Error getting commute time from %v to %v: %v\n", target, trip.location, err))
		}

		locations[i] = RoundTrip{
			location: trip.location,
			to:       to,
			from:     from,
		}
	}

	err = outputTimes(locations)
	if err != nil {
		fmt.Printf("Error outputting commute times: %v\n", err)
		return
	}

	logMessage("Output successful.")
}

func getLocations() (locations []RoundTrip, err error) {

	locByte, err := os.ReadFile("locations.txt")
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

func outputTimes(locations []RoundTrip) error {
	toRow := []string{time.Now().String()}
	fromRow := []string{time.Now().String()}

	for _, trip := range locations {
		toRow = append(toRow, strconv.Itoa(trip.to))
		fromRow = append(fromRow, strconv.Itoa(trip.from))
	}

	err := writeFile(toRow, "to-target-times.csv")
	if err != nil {
		return err
	}

	err = writeFile(fromRow, "from-target-times.csv")
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
