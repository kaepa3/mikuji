package mikuji

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/labstack/gommon/log"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
	fontSize     = 10
	dpi          = 72
)

var (
	arcadeFont font.Face
)

type Game struct {
	boardImage *ebiten.Image
}

func init() {
	rand.NewSource(time.Now().UnixNano())
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Error(err)
	}
}

func NewGame() (*Game, error) {
	return &Game{}, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("HelloWorld:%d", 22), arcadeFont, 10, 10, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
