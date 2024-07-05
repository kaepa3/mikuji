package main

import (
	"log"

	"github.com/kaepa3/mikuji/mikuji"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := mikuji.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(mikuji.ScreenWidth, mikuji.ScreenHeight)
	ebiten.SetWindowTitle("mikuji")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
