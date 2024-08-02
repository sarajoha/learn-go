// Extract/Inline variable. Taking magic values and giving them a name lets you simplify your code quickly.
// command .

// Rename. You should be able to rename symbols across files confidently.
package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
