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
	"github.com/hajimehoshi/ebiten/v2/text/v2" // Pacote atualizado
)

type Game struct {
	selectedIndex int
	options       []string
	windowWidth   int
	windowHeight  int
	state         int // Novo campo para rastrear o estado do jogo
	buildings     []*ebiten.Image
}

const (
	screenWidth      = 1024
	screenHeight     = 768
	StateMenu    int = iota
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

func (g *Game) Update() error {
	// Navegação pelo menu com o teclado
	switch g.state {
	case StateMenu:
		switch {
		case inpututil.IsKeyJustPressed(ebiten.KeyEnter):
			switch g.selectedIndex {
			case 0:
				g.state = StatePlayerVsPlayer
				g.generateBuildings()
			case 1:
				log.Println("Player vs Computer")
			case 2:
				log.Println("Exit")
				return ebiten.Termination
			}
		case inpututil.IsKeyJustPressed(ebiten.KeyArrowDown):
			g.selectedIndex++
			if g.selectedIndex >= len(g.options) {
				g.selectedIndex = 0
			}
		case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp):
			g.selectedIndex--
			if g.selectedIndex < 0 {
				g.selectedIndex = len(g.options) - 1
			}
		}

	case StatePlayerVsPlayer:
		// Adicione a lógica para desenhar os prédios aqui
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.state = StateMenu
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateMenu:
		title := "Gorilla Go"
		menuWidth := 324.0

		menuX := float32((g.windowWidth - int(menuWidth)) / 2)

		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(menuX), 100)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, title, &text.GoTextFace{
			Size:   48,
			Source: pixelFontRegularSource,
		}, op)

		// Define a posição inicial do texto dentro da moldura (alinhado à esquerda da moldura)
		textStartX := int(menuX) // Posição inicial do texto, alinhado à esquerda da moldura com margem de 20px

		// Desenha as opções de menu
		for i, option := range g.options {
			y := 180 + i*60
			// Muda a cor da opção selecionada
			if i == g.selectedIndex {
				op := &text.DrawOptions{}
				op.GeoM.Translate(float64(textStartX), float64(y))
				text.Draw(screen, "> "+option, &text.GoTextFace{
					Size:   32,
					Source: pixelFontRegularSource,
				}, op)
			} else {
				op := &text.DrawOptions{}
				op.GeoM.Translate(float64(textStartX), float64(y))
				text.Draw(screen, "  "+option, &text.GoTextFace{
					Size:   32,
					Source: pixelFontRegularSource,
				}, op)
			}
		}
	case StatePlayerVsPlayer:
		screen.Clear() // Limpa a tela
		for i, building := range g.buildings {
			x := i * (g.windowWidth / 10)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(g.windowHeight-building.Bounds().Dy()))
			screen.DrawImage(building, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Define o tamanho da janela e mantém as proporções
	g.windowWidth = outsideWidth
	g.windowHeight = outsideHeight
	return outsideWidth, outsideHeight
}

func (g *Game) generateBuildings() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	g.buildings = make([]*ebiten.Image, 10)
	buildingWidth := g.windowWidth / 10
	minHeight := int(0.3 * float64(g.windowHeight))
	maxHeight := int(0.8 * float64(g.windowHeight))

	buildingColors := []color.RGBA{
		{169, 169, 169, 255}, // Cinza
		{139, 69, 19, 255},   // Marrom
		{70, 130, 180, 255},  // Azul
	}

	for i := 0; i < 10; i++ {
		buildingHeight := minHeight + rnd.Intn(maxHeight-minHeight+1)
		building := ebiten.NewImage(buildingWidth, buildingHeight)
		buildingColor := buildingColors[rnd.Intn(len(buildingColors))]
		building.Fill(buildingColor)

		// Desenhar janelas
		windowColor := color.RGBA{255, 255, 255, 255} // Branco
		windowWidth := buildingWidth / 5
		windowHeight := buildingHeight / 15

		offset := int(0.05 * float64(buildingWidth)) // Offset de 5%
		columnWidth := (buildingWidth - 2*offset) / 3

		for y := windowHeight; y < buildingHeight; y += windowHeight * 2 {
			for col := 0; col < 3; col++ {
				x := offset + col*columnWidth + (columnWidth-windowWidth)/2
				windowRect := image.Rect(x, y, x+windowWidth, y+windowHeight)
				building.SubImage(windowRect).(*ebiten.Image).Fill(windowColor)
			}
		}

		g.buildings[i] = building
	}
}

func main() {

	game := &Game{
		options: []string{
			"Player vs Player",
			"Player vs Computer",
			"Exit",
		},
		state: StateMenu,
	}

	// Define a resolução inicial
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Menu Example")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
