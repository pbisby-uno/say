// package for printing name
package say

import "fmt"

// Hello returns a string saying hello to the name of
// the person that has been provided
func Hello(name string) string {
	return fmt.Sprintf("Hello %s.", name)
}

func Bye(name string) string {
	return fmt.Sprintf("Good bye %s.", name)
}
