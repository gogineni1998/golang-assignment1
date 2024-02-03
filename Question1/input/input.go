package stdinput

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetInput() int {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	value, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		panic(err)
	}

	return value
}
