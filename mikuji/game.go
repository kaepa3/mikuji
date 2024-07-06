package mikuji

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/kaepa3/mikuji/mikuji/seiza"
	"github.com/kaepa3/mikuji/mikuji/uranai"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
	fontSize     = 24
	dpi          = 72
)

var (
	fontFace *text.GoTextFace
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
	result     *uranai.UranaiResult
}

//go:embed ADTNumeric.ttc
var arabicTTF []byte

func init() {
	rand.NewSource(time.Now().UnixNano())
	//font
	src, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}
	fontFace = &text.GoTextFace{Source: src, Size: fontSize}
}

func NewGame() (*Game, error) {
	s := seiza.NewSeizaInfo()
	return &Game{seiza: s}, nil
}

func (g *Game) Update() error {
	switch g.mode {
	case Select:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			go func() {
				r, err := uranai.Today()
				if err != nil {
					log.Println(err)
				} else {
					g.result = r.GetRecord(g.seiza.GetCurrentValue())
				}
			}()
			g.mode = Result
		} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.seiza.Next()
		} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.seiza.Before()
		}
	case Result:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = Select
			g.result = nil
		}
	}
	return nil
}

var once sync.Once

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.mode {
	case Select:
		seizaName := g.seiza.GetCurrent()
		text.Draw(screen, fmt.Sprintf("Seiza:%s", seizaName), fontFace, nil)
	case Result:
		if g.result == nil {
			text.Draw(screen, "Wait....", fontFace, nil)
		} else {
			for i, v := range g.result.GetArrayResult() {
				op := &text.DrawOptions{}
				op.GeoM.Translate(0, fontSize*float64(i))
				op.LineSpacing = 48 * 1.5
				switch val := v.(type) {
				case string:
					text.Draw(screen, val, fontFace, op)
				case int:
					txt := strconv.Itoa(val)
					text.Draw(screen, txt, fontFace, op)
				default:
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
