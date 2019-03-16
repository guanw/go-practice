package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"strconv"
)

func main() {
	background := flag.String("bg","000000", "background color")
	flag.Parse()

	v, err := strconv.ParseInt(*background, 16, 64)
	if err != nil {
		log.Fatal("Cannot parse the rgb color")
	}
	b := uint8(v % 256)
	g := uint8((v / 256) % 256)
	r := uint8((v / 256 / 256) % 256)
	c := color.RGBA{R: r, G: g, B: b, A: 255}
	fmt.Println("background is %s\n", *background)

	draw(c)
}

func draw(c color.RGBA) {
	r, g, b, a := c.RGBA()
	r, g, b, a = r/256, g/256, b/256, a/256
	fmt.Println("drawing with background rgba(%v, %v, %v, %v)\n", r, g, b, a)
}
