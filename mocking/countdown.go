package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
