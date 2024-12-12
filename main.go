package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// RollDice simulates rolling a dice with the given number of sides for a specified number of times.
func rollDice(sides int, numTimes int) {
	// Create a new random number generator with a unique seed based on the current Unix timestamp
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < numTimes; i++ {
		// Generate a random number between 1 and 'sides'
		roll := rng.Intn(sides) + 1
		fmt.Println(roll)
	}
}

func main() {
	// Command-line flags
	help := flag.Bool("h", false, "Shows this help menu")
	helpQuestionMark := flag.Bool("?", false, "Shows this help menu (alternative flag)")
	numRolls := flag.Int("n", 1, "Number of rolls")
	rollType := flag.String("d", "", "Specify dice type (e.g., -d 2-?)")
	// Custom flag usage to print flags in specific order
	flag.Usage = func() {
		// Print custom usage information in desired order
		fmt.Printf("\nUsage: xfdice [flags]\n\n")
		fmt.Println("\t-h\tShows this help menu")
		fmt.Println("\t-?\tShows this help menu (alternative flag)")
		fmt.Println("\t-d\tSpecify dice type (e.g., -d 2-?)")
		fmt.Println("\t-n\tNumber of rolls (default is 1)")
	}

	// Parse command-line flags
	flag.Parse()

	// Show help if requested or if no valid dice type is provided
	if *help || *helpQuestionMark || *rollType == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Parse the dice type (e.g., d6, d20)
	// 	if !strings.HasPrefix(*rollType, "d") {
	// 		fmt.Println("Invalid dice type. The dice type should start with 'd' (e.g., d6, d20).")
	// 		os.Exit(1)
	// 	}

	// Extract the number after 'd'
	// 	sidesStr := (*rollType)[1:]
	// 	sides, err := strconv.Atoi(sidesStr)
	sides, err := strconv.Atoi(*rollType)
	if err != nil {
		fmt.Println("Invalid dice number. Please provide a valid number after 'd'.")
		os.Exit(1)
	}

	// Validate the dice sides (should be between 2 and 100)
	if sides < 2 {
		fmt.Println("The number of sides should be 2 or larger.")
		os.Exit(1)
	}

	// Perform the dice roll
	fmt.Printf("Rolling d%d(s):\n", sides)
	rollDice(sides, *numRolls)
}
