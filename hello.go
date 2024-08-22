// Extract/Inline variable. Taking magic values and giving them a name lets you simplify your code quickly.
// command .

// Rename. You should be able to rename symbols across files confidently.
package main

import "fmt"

const HelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return HelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
