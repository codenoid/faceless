package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

import "github.com/Kagami/go-face"

func main() {
	fmt.Println("starting...")
	// Init the recognizer.
	// ./ must be contain *.dat files
	rec, err := face.NewRecognizer(".")
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	// Free the resources when you're finished.
	defer rec.Close()

	// Recognize faces on that image.
	faces, err := rec.RecognizeFile("./family.jpg")
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	image1, err := os.Open("family.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	base, err := jpeg.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	b := base.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, base, image.ZP, draw.Src)

	rect := image.NewRGBA(image.Rect(0, 0, 220, 220)) // x1,y1,  x2,y2
	mygreen := color.RGBA{0, 100, 0, 255}             //  R, G, B, Alpha

	// backfill entire surface with green
	draw.Draw(rect, rect.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)

	for _, face := range faces {
		fmt.Println(face.Rectangle)
		face_pos := image.Pt(face.Rectangle.Min.X, face.Rectangle.Min.Y)
		draw.Draw(image3, rect.Bounds().Add(face_pos), rect, image.ZP, draw.Over)
	}

	third, err := os.Create("result.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}
