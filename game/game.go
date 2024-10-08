package game

import (
	"log"

	"github.com/chmenegatti/gorillas/buildings"
	"github.com/chmenegatti/gorillas/gorilla"
	"github.com/chmenegatti/gorillas/menu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	windowWidth  int
	windowHeight int
	state        int
	menu         *menu.Menu
	buildings    []*ebiten.Image
	generator    *buildings.BuildingGenerator
	gorillas     *gorilla.GorillaManager
}

func NewGame() *Game {
	return &Game{
		state:     StateMenu,
		menu:      menu.NewMenu([]string{"Player vs Player", "Player vs Computer", "Exit"}),
		generator: buildings.NewBuildingGenerator(1024, 768),
		gorillas:  gorilla.NewGorillaManager(),
	}
}

const (
	StateMenu = iota
	StatePlayerVsPlayer
)

func (g *Game) Update() error {
	var buildingHeights []int
	switch g.state {
	case StateMenu:
		selectedOption := g.menu.Update()
		if selectedOption >= 0 {
			switch selectedOption {
			case 0:
				g.state = StatePlayerVsPlayer
				g.buildings, buildingHeights = g.generator.GenerateBuildings()
				g.gorillas.PositionGorillas(buildingHeights, g.windowWidth, g.windowHeight)

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
	default:
		panic("unhandled default case")
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
			g.gorillas.DrawGorillas(screen)
		}
	default:
		panic("unhandled default case")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.windowWidth = outsideWidth
	g.windowHeight = outsideHeight
	return outsideWidth, outsideHeight
}
