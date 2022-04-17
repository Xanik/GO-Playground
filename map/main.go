package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#fffffff",
		"white": "#00000",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
}
