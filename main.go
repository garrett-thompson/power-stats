package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	weightUnit := promptWeightUnit()

	promptWeightValue(weightUnit)
}

func promptWeightUnit() string {
	weightUnitPrompt := promptui.Select{
		Label: "We'll need to get your weight. Do you use the metric system (kgs), or the imperial system (lbs)?",
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

	weightValue, err := weightValuePrompt.Run()

	if err != nil {
		fmt.Printf("Failed to get weight value %v\n", err)
		// TODO: throw error?
	}

	return weightValue
}
