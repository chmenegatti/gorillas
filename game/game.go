package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	switch g.state {
	case StateMenu:
		selectedOption := g.menu.Update()
		if selectedOption >= 0 {
			switch selectedOption {
			case 0:
				g.state = StatePlayerVsPlayer
				g.buildings = g.generator.GenerateBuildings()
			case 1:
				log.Println("Player vs Computer")
			case 2:
				log.Println("Exit")
				return ebiten.Termination
			}
		}
	case StatePlayerVsPlayer:
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.state = StateMenu
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateMenu:
		g.menu.Draw(screen, g.windowWidth)
	case StatePlayerVsPlayer:
		screen.Clear()
		for i, building := range g.buildings {
			x := i * (g.windowWidth / 10)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(g.windowHeight-building.Bounds().Dy()))
			screen.DrawImage(building, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.windowWidth = outsideWidth
	g.windowHeight = outsideHeight
	return outsideWidth, outsideHeight
}
