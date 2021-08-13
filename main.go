package main

import (
	pro "argo/problems"
	"fmt"
)

func main() {
	nums := []int{1, 0, -1, 1}
	ret := pro.ThreeSum(nums)
	fmt.Println(ret)
}
