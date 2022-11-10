package main

import "fmt"

type TPoint struct {
	X, Y float64
}

func (p *TPoint) add(z float64) {
	p.Y += z
	p.X += z
}

func main() {
	x := &TPoint{
		1.0, 2.0,
	}

	//fmt.Println(*x)
	x.add(1)
	(*x).add(2)
	fmt.Println(*x)
}
