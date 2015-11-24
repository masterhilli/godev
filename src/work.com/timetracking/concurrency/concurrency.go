package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 1

func main() {
	/*CallingSQFor10Times()
	wg.Wait()*/
	playPingerPonger()
}

func playPingerPonger() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
		time.Sleep(time.Second * 1)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		time.Sleep(time.Second * 2)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func CallingSQFor10Times() {
	myChannel := make(chan int)
	for i := 0; i < 10; i++ {
		go genSQ(myChannel, i)
	}

	//time.Sleep(30 * time.Millisecond)
	for k := 0; k < 10; k++ {
		fmt.Println(<-myChannel)
	}

}

func genSQ(out <-chan int, nums int) {
	fmt.Println("Going to sleep")
	wg.Add(nums)
	defer wg.Done()
	time.Sleep(20 * time.Millisecond)
	out = sq(gen(nums))
}

/*
The first stage, gen, is a function that converts a list of integers to a channel that emits the
integers in the list. The gen function starts a goroutine that sends the integers on the channel
and closes the channel when all the values have been sent:
*/
func gen(nums ...int) <-chan int {
	out := make(chan int)
	fmt.Printf("Call to function gen\n")
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

/*
The second stage, sq, receives integers from a channel and returns a channel that emits the square of each
received integer. After the inbound channel is closed and this stage has sent all the values downstream, it
closes the outbound channel:
*/
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	fmt.Printf("Call to function sq\n")
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
