package main

import "time"

func main() {
	var ch chan Connection
	ch = make(chan Connection, 5) //5 users can access
	for i := 0; i < 15; i++ {
		go CheckConnection(ch, i) //all of it want to use program same time

	}
	time.Sleep(12000 * time.Millisecond)

}
