package main

import (
	"fmt"
	"image"
)

func images() {
	fmt.Println("\n========== Images demo ==========")

	// The image package defines the Image interface
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
