package gotestmod

import (
	"fmt"

	_ "xorm.io/xorm"
)

func Hello(name string) string {
	return fmt.Sprintf("hello: %s", name)
}
