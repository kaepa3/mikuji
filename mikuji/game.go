package mikuji

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
)

type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

