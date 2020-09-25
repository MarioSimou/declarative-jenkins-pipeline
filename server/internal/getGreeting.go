package internal

import (
	"fmt"
)

func GetGreeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}