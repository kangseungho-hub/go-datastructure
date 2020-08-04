package game

type RedWares struct {
	Wares     map[string]Ware
	WareNames []string
}

var (
	board *Board = NewBoard()
)

//Ware 87 , 84
//board : 480, 362

func (NewRedWares *RedWares) InitalizeWares() {
	NewRedWares.Wares["king"] = Ware{
		Name:      "king",
		PositionX: board.boardPosition[1][3][0],
		PositionY: board.boardPosition[1][3][1],
		Img:       NewImage("game/images/red_jang.png", 1)}

	NewRedWares.Wares["jang"] = Ware{
		Name:      "jang",
		PositionX: board.boardPosition[2][3][0],
		PositionY: board.boardPosition[2][3][1],
		Img:       NewImage("game/images/red_sang.png", 1)}

	NewRedWares.Wares["sang"] = Ware{
		Name:      "sang",
		PositionX: board.boardPosition[0][3][0],
		PositionY: board.boardPosition[0][3][1],
		Img:       NewImage("game/images/red_ja.png", 1)}

	NewRedWares.Wares["son"] = Ware{
		Name:      "son",
		PositionX: board.boardPosition[1][2][0],
		PositionY: board.boardPosition[1][2][1],
		Img:       NewImage("game/images/red_wang.png", 1)}
}
func NewRedWares() *RedWares {
	NewRedWares := &RedWares{Wares: make(map[string]Ware)}
	NewRedWares.InitalizeWares()

	NewRedWares.WareNames = append(NewRedWares.WareNames, "king", "jang", "sang", "son")
	return NewRedWares
}

type BlueWares struct {
	Wares     map[string]Ware
	WareNames []string
}

func (NewBlueWares *BlueWares) InitalizeWares() {

	NewBlueWares.Wares["king"] = Ware{
		Name:      "King",
		PositionX: board.boardPosition[1][0][0],
		PositionY: board.boardPosition[1][0][1],
		Img:       NewImage("game/images/green_wang.png", 1)}

	NewBlueWares.Wares["sang"] = Ware{
		Name:      "Sang",
		PositionX: board.boardPosition[0][0][0],
		PositionY: board.boardPosition[0][0][1],
		Img:       NewImage("game/images/green_sang.png", 1)}

	NewBlueWares.Wares["jang"] = Ware{
		Name:      "Jang",
		PositionX: board.boardPosition[2][0][0],
		PositionY: board.boardPosition[2][0][1],
		Img:       NewImage("game/images/green_jang.png", 1)}

	NewBlueWares.Wares["son"] = Ware{
		Name:      "Son",
		PositionX: board.boardPosition[1][1][0],
		PositionY: board.boardPosition[1][1][1],
		Img:       NewImage("game/images/green_ja.png", 1)}

}

//BlueWares Creator Function
func NewBlueWares() *BlueWares {
	NewBlueWares := &BlueWares{Wares: make(map[string]Ware)}
	NewBlueWares.InitalizeWares()

	NewBlueWares.WareNames = append(NewBlueWares.WareNames, "king", "sang", "jang", "son")
	return NewBlueWares
}

type Wares interface {
	InitializeWares()
}
