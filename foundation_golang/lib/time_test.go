package lib

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	go func() {
		for {
			select {
			case <-time.After(time.Second * 5):
				fmt.Println("timeout...")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick := time.Tick(time.Second)

		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	//<-(chan struct{})(nil)
	<-(chan struct{})(nil)
}
