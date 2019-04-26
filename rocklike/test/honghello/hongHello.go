package main

import (
	"fmt"
	"golang-example/rocklike/test/helloutil"
	"math"
)

func testMath() {
	fmt.Println(math.Pi)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	s := "012345"
	s = helloutil.Reverse(s)
	testMath()
	fmt.Print(add(10, 20))
}
