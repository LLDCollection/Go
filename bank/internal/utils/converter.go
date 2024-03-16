package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStringAmountToInt(input string) (int, error) {
	inputArr := strings.Fields(input)
	if len(inputArr) == 0 || len(inputArr) > 2 {
		return 0, fmt.Errorf("invalid input")
	}

	var amount int
	if inputArr[0][len(inputArr[0])-1:] == "D" {
		val, err := strconv.Atoi(inputArr[0][:len(inputArr[0])-1])
		if err != nil {
			return 0, fmt.Errorf("invalid input")
		}

		amount += val * 100
	} else if inputArr[0][len(inputArr[0])-1:] == "C" {
		val, err := strconv.Atoi(inputArr[0][:len(inputArr[0])-1])
		if err != nil {
			return 0, fmt.Errorf("invalid input")
		}

		amount += val
	}

	if len(inputArr) == 2 && inputArr[1][len(inputArr[1])-1:] == "C" {
		val, err := strconv.Atoi(inputArr[1][:len(inputArr[1])-1])
		if err != nil {
			return 0, fmt.Errorf("invalid input")
		}

		amount += val
	}

	return amount, nil
}

func ConvertAmountToString(amount int) string {
	dollars := 0
	cents := 0

	cents = amount % 100
	dollars = amount / 100

	return fmt.Sprintf("%dD %dC", dollars, cents)
}
