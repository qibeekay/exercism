
// Package twofer provides a function for sharing things "two for one".
package twofer

import "fmt"


// ShareWith returns a sentence offering one item to name, or to "you" if name is empty.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
