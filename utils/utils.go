package utils

import "fmt"

func CreatePayload(errorString string, message string) []byte {
	if errorString != "" {
		return []byte(fmt.Sprintf(`{"message": %s, "error": %s}`, message, errorString))
	}

	return []byte(fmt.Sprintf(`{"message": %s}`, message))
}
