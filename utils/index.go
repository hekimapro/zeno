package utils

import "fmt"

// CreateError generates a formatted error message.
func CreateError(message string) error {
	return fmt.Errorf("zenopay error: %s", message)
}
