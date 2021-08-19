package main

import "fmt"

func main() {
	pre := ^int(^uint(0) >> 1)
	fmt.Println(pre)
}
