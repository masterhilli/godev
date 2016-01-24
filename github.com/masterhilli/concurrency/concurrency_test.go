package main

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

type ConcurrencyTestEngine struct {
}

func TestConcurrencyTestEngine(t *testing.T) {
	Suite(&ConcurrencyTestEngine{})
	TestingT(t)
}

/*
The main function sets up the pipeline and runs the final stage: it receives values from the second stage and
 prints each one, until the channel is closed:
*/
func (cte *ConcurrencyTestEngine) TestGenAndSQ(c *C) {
	fmt.Printf("Call to function TestGenAndSQ\n")
	// Set up the pipeline.
	channel := gen(2, 3)
	out := sq(channel)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

/*
Since sq has the same type for its inbound and outbound channels, we can compose it any number of times. We can
 also rewrite main as a range loop, like the other stages:
*/
func (cte *ConcurrencyTestEngine) TestGenSQSQRange(c *C) {
	fmt.Printf("Call to function TestGenSQSQRange\n")
	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}

/* FAN In - FAN Out
Multiple functions can read from the same channel until that channel is closed; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that's closed when all the inputs are closed. This is called fan-in.

We can change our pipeline to run two instances of sq, each reading from the same input channel. We introduce a new function, merge, to fan in the results:
*/
func (cte *ConcurrencyTestEngine) TestFanInFanOut(c *C) {
	fmt.Printf("Call to function TestFanInFanOut\n")
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
