package main

import "fmt"

type Vertex struct {
	x, y int
}

type Vertex3D struct {
	Vertex
	z int
}

func main() {
	v := New(3, 4, 5)

	fmt.Println(v.Area3D())

	v.Scale3D(10)
	fmt.Println(v.Area3D())
}

/*
func Area(v Vertex) int {
	return v.X * v.Y
}
*/

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}
