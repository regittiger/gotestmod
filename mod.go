package gotestmod

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("hello: %s", name)
}
