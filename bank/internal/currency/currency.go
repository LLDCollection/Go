package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type currency struct {
	Code             string
	SubUnitConverter int
	MainUnitSymbol   string
	SubUnitSymbol    string
}

var currencies = map[string]currency{
	"USD": {Code: "USD", SubUnitConverter: 100, MainUnitSymbol: "D", SubUnitSymbol: "C"},
}

func ConvertStringToAmount(input, currencyCode string) (int, error) {
	cur, ok := currencies[currencyCode]
	if !ok {
		return 0, fmt.Errorf("invalid currency code: %v", currencyCode)
	}

	inputArr := strings.Fields(input)
	if len(inputArr) > 2 {
		return 0, fmt.Errorf("too many items")
	}

	var amount int
	for _, part := range inputArr {
		sym := part[len(part)-1:]
		valStr := part[:len(part)-1]

		val, err := strconv.Atoi(valStr)
		if err != nil {
			return 0, fmt.Errorf("incorrect value: %v", valStr)
		}

		if sym == cur.MainUnitSymbol {
			amount += val * cur.SubUnitConverter
		} else if sym == cur.SubUnitSymbol {
			amount += val
		} else {
			return 0, fmt.Errorf("incorrect symbol: %v", sym)
		}
	}

	return amount, nil
}

func ConvertAmountToString(amount int, currencyCode string) (string, error) {
	cur, ok := currencies[currencyCode]
	if !ok {
		return "", fmt.Errorf("unsupported currency: %s", currencyCode)
	}

	dollars := amount / cur.SubUnitConverter
	cents := amount % cur.SubUnitConverter

	return fmt.Sprintf("%d%s %d%s", dollars, cur.MainUnitSymbol, cents, cur.SubUnitSymbol), nil
}
