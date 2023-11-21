package main

import "log"

/*
func main() {
	// Declare a map of strings to strings, initialized using a map literal
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }
	// prints 2023/11/16 14:51:20 map[green:#00ff00 red:#ff0000]

	// var colors map[string]string
	// prints 2023/11/16 14:51:20 map[]

	// What does the make function do?
	// 1. Creates a new map
	// 2. Initializes the map
	// 3. Returns a reference to the map
	// ==> Also used for slices, channels. Not for arrays. Arrays are fixed length. Slices are dynamic.
	// ==> make() allocates memory on the heap and initializes and puts zero or empty values.
	// Prototype: func make(t Type, size ...IntegerType) Type
	// Parameters:
	// ==> t: Type of the map
	// ==> size: Length of the map. Optional. Default is 0.
	// ==> capacity: Capacity of the map. Optional. Default is length.
	// Return value:
	// ==> A reference to the map, channel, slice that is allocated on the memory.

	colors := make(map[string]string)
	// Empty map
	colors["white"] = "#ffffff"
	// Map with one key-value pair (white: #ffffff)
	delete(colors, "white") // What does the delete func do? Deletes the key-value pair from the map
	// Empty map again
	colors["red"] = "#ff0000"
	// Map with one key-value pair (red: #ff0000)

	log.Println(colors)
}

*/

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
		"black": "#000000",
		"blue":  "#0000ff",
	}

	// Print out each key-value pair in the colors map
	func() {
		for color, hex := range colors {
			log.Println("Hex code for", color, "is", hex)
		}
	}()
}
