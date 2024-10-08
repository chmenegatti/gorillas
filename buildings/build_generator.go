package buildings

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func (bg *BuildingGenerator) GenerateBuildings() ([]*ebiten.Image, []int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	buildings := make([]*ebiten.Image, bg.buildingCount)
	buildingWidth := bg.windowWidth / bg.buildingCount
	buildingHeights := make([]int, bg.buildingCount)

	// Define alturas fixas
	minHeight := int(0.3*float64(bg.windowHeight)/10) * 10 // Mínimo de 30%
	maxHeight := int(0.7*float64(bg.windowHeight)/10) * 10 // Máximo de 70%

	buildingColors := []color.RGBA{
		{128, 128, 128, 255}, // Cinza médio
		{64, 64, 64, 255},    // Cinza escuro
		{222, 184, 135, 255}, // Bege médio
		{160, 82, 45, 255},   // Bege escuro
		{102, 51, 0, 255},    // Marrom escuro
		{139, 69, 19, 255},   // Marrom médio
		{165, 42, 42, 255},   // Marrom avermelhado
		{0, 0, 255, 255},     // Azul clássico
		{135, 206, 235, 255}, // Azul claro
		{30, 144, 255, 255},  // Azul royal
		{0, 128, 0, 255},     // Verde escuro
		{153, 204, 51, 255},  // Verde oliva
	}

	for i := 0; i < bg.buildingCount; i++ {
		// Gera a altura do prédio de forma aleatória, múltiplo de 10
		buildingHeight := minHeight + rnd.Intn((maxHeight-minHeight)/10+1)*10
		building := ebiten.NewImage(buildingWidth, buildingHeight)

		// Escolhe uma cor diferente da anterior
		var buildingColor color.RGBA
		for {
			buildingColor = buildingColors[rnd.Intn(len(buildingColors))]
			if i == 0 || buildingColor != buildings[i-1].At(0, 0).(color.RGBA) {
				break
			}
		}

		building.Fill(buildingColor)

		// Desenhar janelas
		bg.drawWindows(building, buildingWidth, buildingHeight)
		buildings[i] = building
	}

	return buildings, buildingHeights
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
	startY := int(0.88 * float64(buildingHeight))

	// Calcula o offset de 5% do topo do prédio para desenhar a última linha de janelas

	// Desenha as janelas, começando de baixo para cima, e adiciona uma linha extra no topo
	for y := startY; y > int(float64(windowHeight)*1.5)-10; y -= int(float64(windowHeight) * 1.5) {
		for col := 0; col < 3; col++ {
			x := offset + col*columnWidth + (columnWidth-windowWidth)/2
			windowRect := image.Rect(x, y, x+windowWidth, y+windowHeight)
			building.SubImage(windowRect).(*ebiten.Image).Fill(windowColor)
		}
	}

}
