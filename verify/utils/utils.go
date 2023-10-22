package utils

import "fmt"

func FromFormat(from, email string) string {
	return fmt.Sprintf("%s <%s>", from, email)
}
