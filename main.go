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

	dur, err := getCommuteTime("45+Spruce+St+Aurora+ON+L4G+1R9", "43.8316105,-79.3561524", apiKey)
	if err != nil {
		fmt.Printf("Error getting commute time: %v\n", err)
		return
	}

	fmt.Println(dur)
}
