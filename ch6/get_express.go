package main

import "fmt"

type Pointx struct{ X, Y float64 }

func (p Pointx) Add(q Pointx) Pointx { return Pointx{p.X + q.X, p.Y + q.Y} }
func (p Pointx) Sub(q Pointx) Pointx { return Pointx{p.X - q.X, p.Y - q.Y} }

type Pathx []Pointx

func (path Pathx) TranslateBy(offset Pointx, add bool) {
	var op func(p, q Pointx) Pointx
	if add {
		op = Pointx.Add
	} else {
		op = Pointx.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main() {
	//// 第一种静态方法调用方式
	//scale := (*Pointx).ScaleBy
	//scale(&p, 2)
	//fmt.Println(p)            // "{2 4}"
	//fmt.Printf("%T\n", scale) // "func(*Point, float64)"
	//
	//// 第二种静态方法调用方式
	//(*Pointx).ScaleBy(&p, 3)
	//fmt.Println(p)

	p := Pathx{Pointx{1, 2}}
	p.TranslateBy(Pointx{1, 2}, true)
	fmt.Println(p)
}
