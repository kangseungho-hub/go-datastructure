package game



type Board struct {
	RedArea  [6]*RedSpace
	BLueArea [6]*BlueSpace
	Spaces   *[12]int
}

func NewBoard() *Board {
	var Spaces [12]int

	for i, _ := range Spaces {
		Spaces[i] = i + 1
	}

	board := &Board{Spaces: &Spaces}
	return board
}

type RedSpace struct {
	Space *[6]int
}

type BlueSpace struct {
	Index *[6]int
}
