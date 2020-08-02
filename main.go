package main

import (
	"game"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	RedWares  *game.RedWares  = game.NewRedWares()
	BlueWares *game.BlueWares = game.NewBlueWares()

	BgImg *ebiten.Image = game.NewImage("game/images/bgimg.png", 1)
)

const (
	WindowWidth  int     = 500
	WindowHeight int     = 362
	windowScale  float64 = 1.0
)

func RunGame(screen *ebiten.Image) error {
	opt := &ebiten.DrawImageOptions{}

	screen.DrawImage(BgImg, nil)

	for _, ware := range BlueWares.Wares {
		opt.GeoM.Translate(ware.PositionX, ware.PositionY)
		screen.DrawImage(ware.Img, opt)
	}

	return nil
}

func main() {
	BgImg = game.NewImage("game/images/bgimg.png", ebiten.FilterDefault)
	err := ebiten.Run(RunGame, WindowWidth, WindowHeight, windowScale, "twelveJanggi")

	if err != nil {
		log.Fatalf("Run Failed width error %v", err)
	}
}
