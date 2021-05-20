package main

import "fmt"

func main() {

	person := map[string]int{
		"karn": 29,
		"set":  30,
	}
	fmt.Println(person)

	colors := make(map[string]string)
	colors["red"] = "#FF0000"
	colors["black"] = "#000000"
	updateMap(colors, "black", "#ffffff")
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hexcode := range colors {
		fmt.Println(color, hexcode)
	}
}

func updateMap(colors map[string]string, color string, code string) {
	colors[color] = code
}
