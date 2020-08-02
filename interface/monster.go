package interface



import "fmt"

type monster interface {
	Attacking() int
	Information() string
}

type Goblin struct {
	Name string
	Ad   int
	AP   int
	Hp   float32
}

type Slime struct {
	Name string
	Ad   int
	AP   int
	Hp   float32
}

func (g Goblin) Attacking() int {
	fmt.Println("몽둥이로 후려 갈겨버리기~")
	return g.Ad * 3
}

func (s Slime) Attacking() int {
	fmt.Println("몸집으로 깔아 뭉개버리기~")
	return s.AP * 2
}

func (g Goblin) Information() string {
	return string(g.Name) + "Ad"
}
func (s Slime) Information() string {
	return string(s.Name) + "Ad"
}

func AttackingInformation(m monster) {
	m.Attacking()
	fmt.Println(m.Information())
}
