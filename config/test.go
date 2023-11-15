package config

// TODO write tests
import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func returnNumber(input int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(input)
}

// write into channel: chan <- [type]
// read from channel: <- chan [type
// send and read chan: [type]

func returnNumberWithChannel(input int, ch chan<- int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(input)
	result := input * 2
	ch <- result
}

func main() {
	ch := make(chan int)

	go returnNumber(1)

	go returnNumberWithChannel(2, ch)
	result := <-ch
	fmt.Printf("result: %v", result)
}
