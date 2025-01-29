package utils

import (
	"fmt"
	"strings"
)

// CreateError generates a formatted error message.
func CreateError(message string) error {
	return fmt.Errorf("[ZENOPAY ERROR]: %s", message)
}

func FormatPhoneNumber(input string) string {
	if strings.HasPrefix(input, "0") {
		return "255" + input[1:]
	}
	return input
}
