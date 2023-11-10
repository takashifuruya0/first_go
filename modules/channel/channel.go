package channel

import (
	"fmt"
)

func Run() {
	// channel
	ch := make(chan string)
	go server(ch)
	for ss := range ch {
		fmt.Println(ss)
	}
	ch2 := make(chan string)
	go server(ch2)
	counter := 0
final:
	for {
		select {
		case ss := <-ch2:
			if ss == "" {
				fmt.Println("Completed")
				break final
			} else {
				fmt.Println(ss)
			}
		default:
			counter++
			if counter > 1000 {
				break final
			} else {
				fmt.Printf("waiting %d\n", counter)
			}
		}
	}
}

// channel
func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}
