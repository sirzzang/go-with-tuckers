package main

import "fmt"

// named type
type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func ColorToString(color ColorType) string {
	switch color {
	case Red:
		return "red"
	case Blue:
		return "blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	}
	return "Unknown" // default
}

func GetFavoriteColor() ColorType {
	return Yellow
}

func main() {
	fmt.Println(ColorToString((GetFavoriteColor())))
}
