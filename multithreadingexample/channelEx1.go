package multithreadingexample

// import "fmt"

// type Car struct {
// 	val string
// }

// func main() {
// 	chan1 := make(chan string)

// 	chan1 <- "kangseungho"

// 	Chan1 := make(chan Car)
// 	Chan2 := make(chan Car)
// 	Chan3 := make(chan Car)

// 	Chan1 <- Car{val: "Car 1 : "}
// 	go MakeTire(Chan1, Chan2)
// 	go MakeEngine(Chan2, Chan3)

// 	fmt.Println(<-Chan3)

// }

// func MakeTire(CarChan chan Car, OutChan chan Car) {
// 	car := <-CarChan
// 	car.val += " + Tire "

// 	OutChan <- car
// }

// func MakeEngine(CarChan chan Car, OutChan chan Car) {
// 	car := <-CarChan
// 	car.val += " +Engine"

// 	OutChan <- car

// }
