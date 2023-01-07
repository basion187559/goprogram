package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 7
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

}
