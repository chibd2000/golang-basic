package main

import (
	"fmt"
)

type Point struct{ X, Y int }

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func main() {
	fmt.Println(Scale(Point{1, 2}, 5)) // "{5 10}"
}
