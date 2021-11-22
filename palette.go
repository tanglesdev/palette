package main

import (
	"fmt"
	"image"
	"os"

	"github.com/RobCherry/vibrant"
	"golang.org/x/image/draw"

	_ "image/jpeg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./palette PATH_TO_IMAGE")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	decodedImage, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Use a custom resize image area and scaler.
	palette := vibrant.NewPaletteBuilder(decodedImage).ResizeImageArea(320 * 320).Scaler(draw.CatmullRom).Generate()
	swatch := palette.DarkVibrantSwatch()
	if swatch == nil {
		fmt.Println("no dark vibrant swatch found")
		return
	}
	hsl := swatch.HSL()
	fmt.Printf("link: hsla(%f, %f%%, %f%%, %d)\n", hsl.H, hsl.S*100, hsl.L*100, hsl.A)

	swatch = palette.LightMutedSwatch()
	if swatch == nil {
		fmt.Println("no light muted swatch found")
		return
	}
	hsl = swatch.HSL()
	fmt.Printf("background: hsla(%f, %f%%, %f%%, %d)\n", hsl.H, hsl.S*100, hsl.L*100, hsl.A)
}
