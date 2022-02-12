package main

import "fmt"

//https://golang.gr/2022/02/11/range-channels/

func main() {
	//make a channel
	ch := make(chan int)
	/*start go function
	send into the channel six numbers
	close the channel
	*/
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}()
	//use range to receive the contents of the channel
	for j := range ch {
		fmt.Println(j)
	}
}

//next:buffered-channels
