package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type RoundTrip struct {
	to   int
	from int
}

func main() {
	godotenv.Load()

	apiKey := os.Getenv("MAPS_API_KEY")
	if apiKey == "" {
		fmt.Println("error: no maps api key defined")
		return
	}

	target := os.Getenv("TARGET")
	if apiKey == "" {
		fmt.Println("error: no target defined")
		return
	}

	locations, err := getLocations()
	if err != nil {
		fmt.Printf("Error getting locations: %v\n", err)
		return
	}

	for location := range locations {

		//get each way and store to print
		to, err := getCommuteTime(location, target, apiKey)
		if err != nil {
			fmt.Printf("Error getting commute time: %v\n", err)
			return
		}

		//get each way and store to print
		from, err := getCommuteTime(target, location, apiKey)
		if err != nil {
			fmt.Printf("Error getting commute time: %v\n", err)
			return
		}

		locations[location] = RoundTrip{
			to:   to,
			from: from,
		}

		fmt.Println(fmt.Printf("Location: %s, To Target: %v, From Target: %v", location, to, from))
	}

	err = outputTimes(locations)
	if err != nil {
		fmt.Printf("Error getting commute time: %v\n", err)
		return
	}
}

func getLocations() (locations map[string]RoundTrip, err error) {

	locByte, err := os.ReadFile("locations.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	locArray := strings.Split(string(locByte), "\n")

	locations = make(map[string]RoundTrip)

	for i := 0; i < len(locArray); i++ {
		locations[locArray[i]] = RoundTrip{}
	}

	fmt.Println(locations)

	return
}

func outputTimes(locations map[string]RoundTrip) error {
	toRow := []string{time.Now().String()}
	fromRow := []string{time.Now().String()}

	for location := range locations {
		toRow = append(toRow, string(locations[location].to))
		fromRow = append(fromRow, string(locations[location].from))
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
