package main

import (
	"fmt"
)

func main() {
	// map in go, object in js, dict in python
	// image.png

	//METHOD 1:
	// colorHexMap := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"blue":  "#0000ff",
	// }

	//METHOD 2:
	// var colorHexMap map[string]string

	//METHOD 3:
	colorHexMap := make(map[string]string)
	colorHexMap["red"] = "#ff0000"
	colorHexMap["green"] = "#4bf745"
	colorHexMap["blue"] = "#00000f"
	colorHexMap["white"] = "#ffffff"

	fmt.Println(colorHexMap)

	colorHexMap["blue"] = "#0000ff"
	delete(colorHexMap, "white")
	fmt.Println(colorHexMap)

	printMap(colorHexMap)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Color: " + color + ", Hex: " + hex)
	}
}
