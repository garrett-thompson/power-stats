package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/manifoldco/promptui"
)

func main() {
	// collect data
	ftpValue := promptFtp()
	weightUnit := promptWeightUnit()
	promptWeightValue(weightUnit)

	// display power zones
	powerZones := calculatePowerZones(ftpValue)
	printPowerZones(powerZones)

	// TODO: display watts/kg
}

func promptFtp() float64 {
	ftpPrompt := promptui.Prompt{
		Label: "What is your FTP?",
	}

	// validate ftp is a number
	ftpStringValue, promptErr := ftpPrompt.Run()

	ftpValue, parseIntErr := strconv.ParseFloat(ftpStringValue, 64)

	if parseIntErr != nil {
		fmt.Printf("Failed to parse FTP value %v\n", parseIntErr)
		// TODO: throw error?
	}

	if promptErr != nil {
		fmt.Printf("Failed to get FTP value %v\n", promptErr)
		// TODO: throw error?
	}

	return ftpValue
}

func promptWeightUnit() string {
	weightUnitPrompt := promptui.Select{
		Label: "Next we'll need to get your weight. Do you use the metric system (kgs), or the imperial system (lbs)?",
		Items: []string{"Metric", "Imperial"},
	}

	weightUnitMap := map[string]string{"Metric": "kgs", "Imperial": "lbs"}

	_, weightUnit, err := weightUnitPrompt.Run()

	if err != nil {
		fmt.Printf("Failed to get weight unit %v\n", err)
		// TODO: throw error?
	}

	return weightUnitMap[weightUnit]
}

// TODO: Would be nice to verify the unit conforms to our union type "Metric" | "Imperial"
func promptWeightValue(unit string) string {
	weightValuePrompt := promptui.Prompt{
		Label: fmt.Sprintf("Thanks. What is your weight in %v?", unit),
	}

	// validate weight is a number
	weightValue, err := weightValuePrompt.Run()

	if err != nil {
		fmt.Printf("Failed to get weight value %v\n", err)
		// TODO: throw error?
	}

	return weightValue
}

// TODO: extract this type so we can DRY this up
func calculatePowerZones(ftp float64) [][]int {
	zone1Max := int(math.Round(ftp * float64(0.55)))
	zone2Max := int(math.Round(ftp * float64(0.75)))
	zone3Max := int(math.Round(ftp * float64(0.90)))
	zone4Max := int(math.Round(ftp * float64(1.05)))
	zone5Max := int(math.Round(ftp * float64(1.20)))
	zone6Max := int(math.Round(ftp * float64(1.50)))

	// TODO: extract this type so we can DRY this up
	return [][]int{
		{0, zone1Max},
		{zone1Max + 1, zone2Max},
		{zone2Max + 1, zone3Max},
		{zone3Max + 1, zone4Max},
		{zone4Max + 1, zone5Max},
		{zone5Max + 1, zone6Max},
		{zone6Max + 1, 0},
	}
}

func printPowerZones(powerZones [][]int) {
	for zone := range powerZones {
		// can you destructure these in a 1-liner similar to JS?
		min := powerZones[zone][0]
		max := powerZones[zone][1]

		var zoneRange string

		if min == 0 {
			zoneRange = "< " + strconv.Itoa(max)
		} else if max == 0 {
			zoneRange = "> " + strconv.Itoa(min)
		} else {
			zoneRange = strconv.Itoa(min) + " - " + strconv.Itoa(max)
		}

		fmt.Println("Zone " + strconv.Itoa(zone+1) + ": " + zoneRange)
	}
}
