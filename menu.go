package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
