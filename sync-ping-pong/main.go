package main

import (
	"fmt"
	"sync"
	"time"
)

type Ball struct {
	hit int
}

func player(name string, table chan *Ball, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ball := <-table
		ball.hit++
		fmt.Println(name, ball.hit)
		time.Sleep(1500 * time.Millisecond)
		table <- ball
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	table := make(chan *Ball, 1)

	go player("playerA", table, &wg)
	go player("playerB", table, &wg)

	table <- &Ball{}
	wg.Wait()

}
