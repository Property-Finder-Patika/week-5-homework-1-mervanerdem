package main

import (
	"fmt"
	"time"
)

type Licance interface {
	String() string
}

type Connection int

func (c Connection) String() string {
	return fmt.Sprintf("Disconnected:%d\n", c) //leaved.

}

func newAccess(ch chan Connection, i Connection) Licance {
	fmt.Printf("Connectted:%d\n", i)
	time.Sleep(1000 * time.Millisecond) //do some stuff in here
	i = <-ch                            //time to leave
	return Connection(i)
}

func CheckConnection(ch chan Connection, i int) {
	a := Connection(i)
	ch <- a                 //only allow for 5 people and those without permission have to wait for those with leave
	acc := newAccess(ch, a) // give access who has permission
	fmt.Println(acc)

}
