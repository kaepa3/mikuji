package mikuji

import (
	"fmt"
	"image/color"
	"math/rand"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/labstack/gommon/log"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kaepa3/mikuji/mikuji/seiza"
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

type Mode int

const (
	Select Mode = iota
	Result
)

type Game struct {
	boardImage *ebiten.Image
	mode       Mode
	seiza      *seiza.SeizaInfo
}

func init() {
	rand.NewSource(time.Now().UnixNano())
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
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
	s := seiza.NewSeizaInfo()
	return &Game{seiza: s}, nil
}

func (g *Game) Update() error {
	switch g.mode {
	case Select:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = Result
		} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.seiza.Next()
		} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.seiza.Before()
		}
	case Result:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = Select
		}
	}
	return nil
}

var once sync.Once

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.mode {
	case Select:
		seizaName := g.seiza.GetCurrent()
		text.Draw(screen, fmt.Sprintf("Seiza:%s", seizaName), arcadeFont, 10, 10, color.White)
	case Result:
		text.Draw(screen, fmt.Sprintf("Result:%s", "fuck!!!"), arcadeFont, 10, 10, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
