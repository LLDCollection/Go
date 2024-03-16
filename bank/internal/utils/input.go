package utils

import (
	"os"
	"bufio"
	"strings"
)

func TakeInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}