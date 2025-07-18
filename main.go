package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("MAPS_API_KEY")
	if apiKey == "" {
		fmt.Println("error: no maps api key")
		return
	}

	possibilities, target, err := getLocations()
	if err != nil {
		fmt.Printf("Error getting locations: %v\n", err)
		return
	}

	for i := 0; i < len(possibilities); i++ {

		//get each way and store to print
		dur, err := getCommuteTime(possibilities[i], target, apiKey)
		if err != nil {
			fmt.Printf("Error getting commute time: %v\n", err)
			return
		}

		fmt.Println(dur)

		//get each way and store to print
		dur, err = getCommuteTime(target, possibilities[i], apiKey)
		if err != nil {
			fmt.Printf("Error getting commute time: %v\n", err)
			return
		}

		fmt.Println(dur)
	}
}

func getLocations() (possibilities []string, target string, err error) {

}
