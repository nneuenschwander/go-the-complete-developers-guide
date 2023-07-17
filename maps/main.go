package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#008000",
	}

	colors["white"] = "#FFFFFF"
	delete(colors, "red")

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("hex code for", color, "is", hex)
	}
}
