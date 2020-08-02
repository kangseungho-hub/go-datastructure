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

type RedWares struct {
	Wares     map[string]Ware
	WareNames []string
}

//Ware 87 , 84
//board : 480, 362

func NewRedWares() *RedWares {
	NewRedWares := &RedWares{Wares: make(map[string]Ware)}

	NewRedWares.Wares["king"] = Ware{PositionX: 0, PositionY: 0, Img: NewImage("game/images/red_jang.png", 1)}
	NewRedWares.Wares["jang"] = Ware{PositionX: 0, PositionY: 0, Img: NewImage("game/images/red_sang.png", 1)}
	NewRedWares.Wares["sang"] = Ware{PositionX: 0, PositionY: 0, Img: NewImage("game/images/red_ja.png", 1)}
	NewRedWares.Wares["son"] = Ware{PositionX: 0, PositionY: 0, Img: NewImage("game/images/red_wang.png", 1)}

	NewRedWares.WareNames = append(NewRedWares.WareNames, "king", "jang", "sang", "son")
	return NewRedWares
}

type BlueWares struct {
	Wares     map[string]Ware
	WareNames []string
}

//BlueWares Creator Function
func NewBlueWares() *BlueWares {
	NewBlueWares := &BlueWares{Wares: make(map[string]Ware)}
	NewBlueWares.Wares["jang"] = Ware{Name: "Jang", PositionX: 5, PositionY: 23, Img: NewImage("game/images/green_jang.png", 1)}
	NewBlueWares.Wares["sang"] = Ware{Name: "Sang", PositionX: 5, PositionY: 53, Img: NewImage("game/images/green_sang.png", 1)}
	NewBlueWares.Wares["son"] = Ware{Name: "Son", PositionX: 5, PositionY: 41, Img: NewImage("game/images/green_ja.png", 1)}
	NewBlueWares.Wares["king"] = Ware{Name: "King", PositionX: 5, PositionY: 123, Img: NewImage("game/images/green_wang.png", 1)}

	NewBlueWares.WareNames = append(NewBlueWares.WareNames, "king", "jang", "sang", "son")
	return NewBlueWares
}
