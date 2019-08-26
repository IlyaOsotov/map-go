package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "RED",
		"green": "GREEN",
	}

	colors["white"] = "WHITE"
	delete(colors, "red")
	delete(colors, "blue")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, upperColor := range c {
		fmt.Println(upperColor, "is a upper of", color)
	}
}
