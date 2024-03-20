package geometry

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
