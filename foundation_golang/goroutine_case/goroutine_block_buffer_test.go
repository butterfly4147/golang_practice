package main

import "testing"

const (
	max     = 500_00_000
	block   = 500
	bufSize = 100
)

func BenchmarkTNM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := make(chan struct{})
		c := make(chan int, bufSize)

		go func() {
			count := 0
			for x := range c {
				count += x
			}
			//println("count = ", count)

			close(done)
		}()

		for j := 0; j < max; j++ {
			c <- j
		}
		close(c)

		<-done
	}
}

func BenchmarkBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := make(chan struct{})
		c := make(chan [block]int, bufSize)

		go func() {
			count := 0
			for x := range c {
				for _, elem := range x {
					count += elem
				}
			}
			//println("count = ", count)

			close(done)
		}()
		for j := 0; j < max; j += block {
			var buf [block]int
			for k := 0; k < block; k++ {
				buf[k] = j + k
			}

			c <- buf
		}
		close(c)

		<-done
	}
}

func TestNormalMode(t *testing.T) {
	done := make(chan struct{})
	c := make(chan int, bufSize)

	go func() {
		count := 0
		for x := range c {
			count += x
		}
		println("count = ", count)

		close(done)
	}()

	for i := 0; i < max; i++ {
		c <- i
	}
	close(c)

	<-done
}

func TestBlock(t *testing.T) {
	done := make(chan struct{})
	c := make(chan [block]int, bufSize)

	go func() {
		count := 0
		for x := range c {
			for _, elem := range x {
				count += elem
			}
		}
		println("count = ", count)

		close(done)
	}()
	for j := 0; j < max; j += block {
		var buf [block]int
		for k := 0; k < block; k++ {
			buf[k] = j + k
		}

		c <- buf
	}
	close(c)

	<-done
}

//250368900
//10395437600
