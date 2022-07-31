package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func func1() {
	for i := 0; i < 15; i++ {
		fmt.Printf("1.goroutine %d\n", i)
	}
	wg.Done() //this goroutine done.
}

func func2() {
	for i := 0; i < 15; i++ {
		fmt.Printf("2. goroutine %d\n", i)
	}
	wg.Done() //this goroutine done.
}

func func3() {
	for i := 0; i < 15; i++ {
		fmt.Printf("3. goroutine %d\n", i)
	}
	wg.Done() //this goroutine done.
}

func main() {
	fmt.Println("Welcome to Race Condition :)")
	wg.Add(3) //for 3 goroutine wait
	go func1()
	go func2()
	go func3()
	wg.Wait()
	//if all goroutine give done signal main can continue
	fmt.Println("You can see winner in terminal")
}
