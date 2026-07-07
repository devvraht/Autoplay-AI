package main

import (
	"fmt"
	"log"

	"github.com/devvraht/Autoplay-AI/internal/capture"
	"github.com/devvraht/Autoplay-AI/internal/vision"
)

func main() {
	fmt.Println("AutoPlay AI Started")

	img, err := capture.CaptureScreen()
	if err != nil {
		log.Fatal(err)
	}

	game := vision.CropGameArea(img)

	err = capture.SaveImage(game, "game.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Game image saved.")
}
