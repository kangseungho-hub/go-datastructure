package game

var (
	whitespaceX float64 = 5
	whitespaceY float64 = 7
	wareSizeX   float64 = 87
	wareSizeY   float64 = 84

	BoxSizeX float64 = 116
	BoxSizeY float64 = 116
)

type Board struct {
	boardPosition [3][4][2]float64
}

func (b *Board) Initialize() {

}

func NewBoard() *Board {
	newBoard := &Board{}
	newBoard.InitalizeBoard()

	return newBoard

}

func (b *Board) InitalizeBoard() {
	var Xstart float64 = 19.5
	var Ystart float64 = 23

	for i := 0; i < len(b.boardPosition); i++ {
		for j := 0; j < len(b.boardPosition[i]); j++ {
			b.boardPosition[i][j][0] = Xstart + float64(j)*BoxSizeX
			b.boardPosition[i][j][1] = Ystart + float64(i)*BoxSizeY
		}
	}
}
