package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// var DoOnce sync.Once
var numRolls int

// different dice
var d2 = []int{1, 2}
var d3 = []int{1, 2, 3}
var d4 = []int{1, 2, 3, 4}
var d6 = []int{1, 2, 3, 4, 5, 6}
var d8 = []int{1, 2, 3, 4, 5, 6, 7, 8}
var d10 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var d12 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12}
var d20 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
var d50 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
var d100 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}

func rollDice(value []int, numTimes *int) {
	var myRoll int
	for num := *numTimes; num >= 1; num-- {
		//DoOnce.Do(func() {
		//rand.Seed(time.Now().UnixNano())
		//})
		rand.Seed(time.Now().UnixNano())
		myRoll = value[rand.Intn(len(value))]
		fmt.Printf("%d\n", myRoll)
	}

}

func main() {
	flag.Usage = func() {
		flatSet := flag.CommandLine
		fmt.Printf("\nONLY ROLLS ONE DICE TYPE AT A TIME\n\nUsage: rolldice -d[#] [-n #]\n")
		order := []string{"n", "d2", "d3", "d4", "d6", "d8", "d10", "d12", "d20", "d50", "d100"}
		for _, name := range order {
			flag := flatSet.Lookup(name)
			fmt.Printf("\t-%s\t%s\n\n", flag.Name, flag.Usage)
		}
	}

	help := flag.Bool("h", false, "Shows this help menu")
	numRolls := flag.Int("n", 1, "-n # to roll a number of same dice")
	rollD2 := flag.Bool("d2", false, "to flip a coin basically")
	rollD3 := flag.Bool("d3", false, "to roll a d3")
	rollD4 := flag.Bool("d4", false, "to roll a d4")
	rollD6 := flag.Bool("d6", false, "to roll a d6")
	rollD8 := flag.Bool("d8", false, "to roll a d8")
	rollD10 := flag.Bool("d10", false, "to roll a d10")
	rollD12 := flag.Bool("d12", false, "to roll a d12")
	rollD20 := flag.Bool("d20", false, "to roll a d20")
	rollD50 := flag.Bool("d50", false, "to roll a d50")
	rollD100 := flag.Bool("d100", false, "to roll a d100")
	flag.Parse()

	switch {
	case *rollD2 != false:
		fmt.Println("Rolling d2(s): ")
		rollDice(d2, numRolls)
		os.Exit(0)

	case *rollD3 != false:
		fmt.Println("Rolling d3(s): ")
		rollDice(d3, numRolls)
		os.Exit(0)

	case *rollD4 != false:
		fmt.Println("Rolling d4(s): ")
		rollDice(d4, numRolls)
		os.Exit(0)

	case *rollD6 != false:
		fmt.Println("Rolling d6(s): ")
		rollDice(d6, numRolls)
		os.Exit(0)

	case *rollD8 != false:
		fmt.Println("Rolling d8(s): ")
		rollDice(d8, numRolls)
		os.Exit(0)

	case *rollD10 != false:
		fmt.Println("Rolling d10(s): ")
		rollDice(d10, numRolls)
		os.Exit(0)

	case *rollD12 != false:
		fmt.Println("Rolling d12(s): ")
		rollDice(d12, numRolls)
		os.Exit(0)

	case *rollD20 != false:
		fmt.Println("Rolling d20(s): ")
		rollDice(d20, numRolls)
		os.Exit(0)

	case *rollD50 != false:
		fmt.Println("Rolling d50(s): ")
		rollDice(d50, numRolls)
		os.Exit(0)

	case *rollD100 != false:
		fmt.Println("Rolling d100(s): ")
		rollDice(d100, numRolls)
		os.Exit(0)

	case *help:
		flag.Usage()
		os.Exit(0)

	default:
		flag.Usage()
		os.Exit(0)
	}

}
