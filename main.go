package main

import (
	"fmt"
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

	screen.DrawImage(BgImg, nil)

	fmt.Println(ebiten.CursorPosition())
	for _, name := range BlueWares.WareNames {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(BlueWares.Wares[name].PositionX, BlueWares.Wares[name].PositionY)
		screen.DrawImage(BlueWares.Wares[name].Img, opt)
	}

	for _, name := range RedWares.WareNames {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(RedWares.Wares[name].PositionX, RedWares.Wares[name].PositionY)
		screen.DrawImage(RedWares.Wares[name].Img, opt)
	}

	return nil
}

func main() {
	b := game.NewBoard()
	fmt.Println(b)
	BgImg = game.NewImage("game/images/bgimg.png", ebiten.FilterDefault)
	err := ebiten.Run(RunGame, WindowWidth, WindowHeight, windowScale, "twelveJanggi")

	if err != nil {
		log.Fatalf("Run Failed width error %v", err)
	}
}
