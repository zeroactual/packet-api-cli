package util

import (
	"bufio"
	"fmt"
	. "os"
	"strconv"
)

func NextInt(max int) (e int) {
	scanner := bufio.NewScanner(Stdin)
	fmt.Print("Option: ")
	scanner.Scan()
	input := scanner.Text()
	e, err := strconv.Atoi(input)

	if (err != nil) && (e >= max) || (0 > e) {
		e = NextInt(max)
	}

	return
}