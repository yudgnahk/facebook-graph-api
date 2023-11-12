package utils

import "fmt"

// SliceToString converts a slice of anonymous type T to a string
func SliceToString[T any](slice []T) string {
	var result string
	for i, v := range slice {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("%v", v)
	}
	return result
}
