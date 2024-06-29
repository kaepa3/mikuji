package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaepa3/mikuji"
)

func main() {
	game, err := mikuji.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(twenty48.ScreenWidth, twenty48.ScreenHeight)
	ebiten.SetWindowTitle("2048 (Ebitengine Demo)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
