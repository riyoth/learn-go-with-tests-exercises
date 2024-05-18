package interactions

import "fmt"

func Greet(name string) string {
	response := "World"
	if name != "" {
		response = name
	}
	return fmt.Sprintf("Hello, %s", response)
}
