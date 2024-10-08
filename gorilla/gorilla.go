package gorilla

import (
	"log"
	"math/rand"
	"time"

	_ "image/png" // Import para carregar imagens PNG

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Gorila representa o sprite e a posição do gorila no jogo.
type Gorilla struct {
	sprite    *ebiten.Image
	positionX int
	positionY int
}

// GorillaManager gerencia a criação e o posicionamento dos gorilas.
type GorillaManager struct {
	gorillas []*Gorilla
}

// NewGorillaManager cria uma nova instância de GorillaManager.
func NewGorillaManager() *GorillaManager {
	return &GorillaManager{}
}

// LoadGorillaSprite carrega o sprite do gorila a partir do arquivo "gorilla.png".
func (gm *GorillaManager) LoadGorillaSprite() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("gorilla/gorilla.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	return img
}

// PositionGorillas posiciona os gorilas no topo dos prédios, garantindo que fiquem a pelo menos 4 prédios de distância.
func (gm *GorillaManager) PositionGorillas(buildingHeights []int, windowWidths, windowHeight int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	numBuildings := len(buildingHeights)

	// Garantir que pelo menos dois gorilas estejam posicionados
	firstGorillaPos := rnd.Intn(numBuildings-8) + 4
	secondGorillaPos := rnd.Intn(numBuildings-8) + firstGorillaPos + 4

	// Garantir que não haja sobreposição dos gorilas
	if secondGorillaPos >= numBuildings {
		secondGorillaPos = numBuildings - 1
	}

	// Posicionar os gorilas no topo dos prédios
	gm.gorillas = []*Gorilla{
		{
			sprite:    gm.LoadGorillaSprite(),
			positionX: firstGorillaPos * (windowHeight / numBuildings),
			positionY: windowHeight - buildingHeights[firstGorillaPos] - 35,
		},
		{
			sprite:    gm.LoadGorillaSprite(),
			positionX: secondGorillaPos * (windowHeight / numBuildings),
			positionY: windowHeight - buildingHeights[secondGorillaPos] - 35,
		},
	}
}

// DrawGorillas desenha os gorilas no topo dos prédios.
func (gm *GorillaManager) DrawGorillas(screen *ebiten.Image) {
	for _, gorilla := range gm.gorillas {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(gorilla.positionX), float64(gorilla.positionY))
		screen.DrawImage(gorilla.sprite, op)
	}
}
