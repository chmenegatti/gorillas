package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	windowWidth  int
	windowHeight int
	state        int
	menu         *Menu
	buildings    []*ebiten.Image
	generator    *BuildingGenerator
}

const (
	screenWidth  = 1024
	screenHeight = 768
	StateMenu    = iota
	StatePlayerVsPlayer
)

func main() {
	game := &Game{
		state:     StateMenu,
		menu:      NewMenu([]string{"Player vs Player", "Player vs Computer", "Exit"}),
		generator: NewBuildingGenerator(screenWidth, screenHeight),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gorilla Go")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
