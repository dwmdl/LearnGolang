package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		var userWeight, userHeight float64 = getUserInput()
		IMT, err := calculateIMT(userWeight, userHeight)

		if err != nil {
			fmt.Println(err)
			continue
		}

		outputResultIMT(IMT)

		var needToRepeatCalculate = checkRepeatCalculate()
		if !needToRepeatCalculate {
			break
		}
	}
}

func getUserInput() (userWeight, userHeight float64) {
	fmt.Print("Enter your height in santimeters: ")
	fmt.Scan(&userHeight)

	fmt.Print("Enter your weight: ")
	fmt.Scan(&userWeight)

	return
}

func calculateIMT(userWeight, userHeight float64) (float64, error) {
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("INCORRECT_INPUT_DATA")
	}

	return userWeight / math.Pow(userHeight/100, IMTPower), nil
}

func outputResultIMT(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println("You have a severe body weight deficiency")
	case IMT <= 18.5:
		fmt.Println("You have a body weight deficiency")
	case IMT <= 25:
		fmt.Println("You have normal body weight")
	case IMT <= 30:
		fmt.Println("You have overweight")
	default:
		fmt.Println("You have a degree of obesity")
	}

	fmt.Printf("Your IMT equals: %.0f\n", IMT)
}

func checkRepeatCalculate() bool {
	var oneMoreCalculate string
	fmt.Println("Do you want to calculate another IMT? y/n")
	fmt.Scan(&oneMoreCalculate)

	return oneMoreCalculate == "y"
}
