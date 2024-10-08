package main

import (
	"log"

	"github.com/chmenegatti/gorillas/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1024
	screenHeight = 768
)

func main() {

	games := game.NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gorilla Go")

	if err := ebiten.RunGame(games); err != nil {
		log.Fatal(err)
	}
}
