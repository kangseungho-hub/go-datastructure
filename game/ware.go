package game

import "github.com/hajimehoshi/ebiten"

type Ware struct {
	Name      string
	Img       *ebiten.Image
	PositionX float64
	PositionY float64
}

type Jang struct {
}

func (jang Jang) Move() {

}

type Sang struct {
}

func (sang Sang) Move() {
}

type King struct {
}

func (king King) Move() {

}

type Son struct {
}

func (son Son) Move() {

}
