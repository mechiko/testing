package utility

import (
	"bufio"
	"os"
	"strconv"
)

func PressAnyKey() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
