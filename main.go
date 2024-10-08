package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

//go:embed font.ttf
var pixelFontRegular []byte

var pixelFontRegularSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pixelFontRegular))
	if err != nil {
		log.Fatal(err)
	}
	pixelFontRegularSource = s
}

// Menu struct para encapsular lógica de menu
type Menu struct {
	selectedIndex int
	options       []string
}

func NewMenu(options []string) *Menu {
	return &Menu{
		options: options,
	}
}

func (m *Menu) Update() int {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return m.selectedIndex
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.selectedIndex++
		if m.selectedIndex >= len(m.options) {
			m.selectedIndex = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.selectedIndex--
		if m.selectedIndex < 0 {
			m.selectedIndex = len(m.options) - 1
		}
	}
	return -1 // Nenhuma seleção foi feita
}

func (m *Menu) Draw(screen *ebiten.Image, width int) {
	menuTitle := "Gorilla Go"
	menuWidth := 324.0
	menuX := float32((width - int(menuWidth)) / 2)

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(menuX), 100)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, menuTitle, &text.GoTextFace{
		Size:   48,
		Source: pixelFontRegularSource,
	}, op)

	textStartX := int(menuX)
	for i, option := range m.options {
		y := 180 + i*60
		optionText := "  " + option
		if i == m.selectedIndex {
			optionText = "> " + option
		}
		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(textStartX), float64(y))
		text.Draw(screen, optionText, &text.GoTextFace{
			Size:   32,
			Source: pixelFontRegularSource,
		}, op)
	}
}

// BuildingGenerator struct para encapsular a lógica de geração de prédios
type BuildingGenerator struct {
	windowWidth   int
	windowHeight  int
	buildingCount int
}

func NewBuildingGenerator(windowWidth, windowHeight int) *BuildingGenerator {
	return &BuildingGenerator{
		windowWidth:   windowWidth,
		windowHeight:  windowHeight,
		buildingCount: 10,
	}
}

func (bg *BuildingGenerator) GenerateBuildings() []*ebiten.Image {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	buildings := make([]*ebiten.Image, bg.buildingCount)
	buildingWidth := bg.windowWidth / bg.buildingCount
	minHeight := int(0.3 * float64(bg.windowHeight))
	maxHeight := int(0.8 * float64(bg.windowHeight))

	buildingColors := []color.RGBA{
		{128, 128, 128, 255}, // Cinza médio
		{64, 64, 64, 255},    // Cinza escuro

		// Tons de bege
		{222, 184, 135, 255}, // Bege médio
		{160, 82, 45, 255},   // Bege escuro

		// Tons de marrom
		{102, 51, 0, 255},  // Marrom escuro
		{139, 69, 19, 255}, // Marrom médio
		{165, 42, 42, 255}, // Marrom avermelhado

		// Tons de azul
		{0, 0, 255, 255},     // Azul clássico
		{135, 206, 235, 255}, // Azul claro
		{30, 144, 255, 255},  // Azul royal

		// Tons de verde
		{0, 128, 0, 255},    // Verde escuro
		{153, 204, 51, 255}, // Verde oliva

	}

	for i := 0; i < bg.buildingCount; i++ {
		buildingHeight := minHeight + rnd.Intn(maxHeight-minHeight+1)
		building := ebiten.NewImage(buildingWidth, buildingHeight)
		buildingColor := buildingColors[rnd.Intn(len(buildingColors))]
		building.Fill(buildingColor)

		// Desenhar janelas
		bg.drawWindows(building, buildingWidth, buildingHeight)
		buildings[i] = building
	}

	return buildings
}

func (bg *BuildingGenerator) drawWindows(building *ebiten.Image, buildingWidth, buildingHeight int) {

	windowColor := color.RGBA{200, 255, 255, 100}
	windowWidth := buildingWidth / 5
	windowHeight := 20 // Tamanho fixo para as janelas

	// Margem lateral de 5%
	offset := int(0.05 * float64(buildingWidth))

	// Largura da coluna, considerando 3 colunas de janelas
	columnWidth := (buildingWidth - 2*offset) / 3

	// Offset inicial de 13% a partir da parte inferior do prédio
	startY := int(0.87 * float64(buildingHeight))

	// Desenha as janelas, começando de baixo para cima
	for y := startY; y > windowHeight*2; y -= windowHeight * 2 {
		for col := 0; col < 3; col++ {
			x := offset + col*columnWidth + (columnWidth-windowWidth)/2
			windowRect := image.Rect(x, y, x+windowWidth, y+windowHeight)
			building.SubImage(windowRect).(*ebiten.Image).Fill(windowColor)
		}
	}
}

// Implementação da interface do jogo
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
