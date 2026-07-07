package main

import (
	"fmt"
	"log"

	"github.com/devvraht/autoplay-ai/internal/capture"
)

func main() {
	fmt.Println("AutoPlay AI Started")

	img, err := capture.CaptureScreen()
	if err != nil {
		log.Fatal(err)
	}

	err = capture.SaveImage(img, "screenshot.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Screenshot saved successfully!")
}
