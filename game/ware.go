package game

import "github.com/hajimehoshi/ebiten"

type Ware struct {
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

type RedWares struct {
	Wares map[string]Ware
}

//Ware 87 , 84
//board : 480, 362
func NewRedWares() *RedWares {
	NewRedWares := &RedWares{Wares: make(map[string]Ware)}

	NewRedWares.Wares["king"] = Ware{PositionX: 0, PositionY: 0, Img: NewImage("game/images/red_jang.png", 1)}
	NewRedWares.Wares["jang"] = Ware{PositionX: 0, PositionY: 0, Img: NewImage("game/images/red_sang.png", 1)}
	NewRedWares.Wares["sang"] = Ware{PositionX: 0, PositionY: 0, Img: NewImage("game/images/red_ja.png", 1)}
	NewRedWares.Wares["son"] = Ware{PositionX: 0, PositionY: 0, Img: NewImage("game/images/red_wang.png", 1)}

	return NewRedWares
}

type BlueWares struct {
	Wares map[string]Ware
}

//BlueWares Creator Function
func NewBlueWares() *BlueWares {
	NewBlueWares := &BlueWares{Wares: make(map[string]Ware)}
	NewBlueWares.Wares["jang"] = Ware{PositionX: 5, PositionY: 123, Img: NewImage("game/images/green_jang.png", 1)}
	NewBlueWares.Wares["sang"] = Ware{PositionX: 5, PositionY: 23, Img: NewImage("game/images/green_sang.png", 1)}
	NewBlueWares.Wares["son"] = Ware{PositionX: 5, PositionY: 223, Img: NewImage("game/images/green_ja.png", 1)}
	NewBlueWares.Wares["king"] = Ware{PositionX: 5, PositionY: 323, Img: NewImage("game/images/green_wang.png", 1)}

	return NewBlueWares
}
