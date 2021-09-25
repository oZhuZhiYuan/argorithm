package main

import "fmt"

type Test struct {
	x int
}

func (this *Test) add() {
	this.x++
}

func main() {
	t := Test{100}
	t.add()
	fmt.Println(t.x)
}
