package multithreadingexample

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go Push1(c1)
// 	go Push2(c2)

// 	timeChan := time.After(3 * time.Second)

// 	for {
// 		select {
// 		case val1 := <-c1:
// 			fmt.Println(val1)

// 		case val2 := <-c2:
// 			fmt.Println(val2)
// 		case <-timeChan:
// 			fmt.Println("time out!")
// 			return

// 		default:
// 			fmt.Println("empty..")
// 			time.Sleep(1000)
// 		}
// 	}
// }

// func Push1(c1 chan string) {
// 	for {
// 		c1 <- "c1"
// 	}

// }

// func Push2(c2 chan string) {
// 	for {
// 		c2 <- "c2"
// 	}
// }
