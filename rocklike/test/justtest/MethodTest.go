package main

import "fmt"

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() {
	v.X++
	v.Y++
}
func (v *Vertex) Abs2_pointerRef() {
	v.X++
	(*v).Y++
}

func main() {
	v := Vertex{1, 1}
	v.Abs()
	fmt.Println(v)
	v = Vertex{1, 1}
	(&v).Abs()
	fmt.Println(v)
	v = Vertex{1, 1}
	(&v).Abs2_pointerRef()
	fmt.Println(v)

}
