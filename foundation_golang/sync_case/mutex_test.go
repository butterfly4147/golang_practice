package sync_case

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type data struct {
	sync.Mutex
	name string
	age  int
}

func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
	d.name = "a"
	d.age = 2
}

func TestMutex(t *testing.T) {
	var d data
	var wg sync.WaitGroup
	wg.Add(2)

	go func(m *sync.WaitGroup) {
		defer wg.Done()
		//time.Sleep(time.Second)
		d.test("write")
	}(&wg)

	go func(m *sync.WaitGroup) {
		defer wg.Done()
		//time.Sleep(time.Second)
		d.test("read")
	}(&wg)

	wg.Wait()
	fmt.Println(d)
}
