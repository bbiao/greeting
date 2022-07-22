package greeting

import "fmt"

func Greeting(i any) string {
	return fmt.Sprintf("%s: %v", Message, i)
}
