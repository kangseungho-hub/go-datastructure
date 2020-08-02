package game

//img size
//Ware 87 , 84

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func NewImage(path string, option ebiten.Filter) *ebiten.Image {
	bgImg, _, err := ebitenutil.NewImageFromFile(path, option)

	if err != nil {
		log.Fatalf("load img file error with : %v", err)
	}

	return bgImg
}
