package gotestmod

import (
	"fmt"

	_ "xorm.io/xorm"
)

func Hello(name, lang string) string {
	return fmt.Sprintf("hello: %s lang: %s", name, lang)
}
