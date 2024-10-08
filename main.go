package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2" // Pacote atualizado
)

type Game struct {
	selectedIndex int
	options       []string
	windowWidth   int
	windowHeight  int
}

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
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyEnter):
		switch g.selectedIndex {
		case 0:
			log.Println("Player vs Player")
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
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Centraliza o título
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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Define o tamanho da janela e mantém as proporções
	g.windowWidth = outsideWidth
	g.windowHeight = outsideHeight
	return outsideWidth, outsideHeight
}

func main() {

	game := &Game{
		options: []string{
			"Player vs Player",
			"Player vs Computer",
			"Exit",
		},
	}

	// Define a resolução inicial
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Menu Example")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
