package main

import (
	"fmt"
	"math"
)

type geometry interface{ area() float64 }
type retangle struct{ width, height float64 }
type circle struct{ radius float64 }

func (r retangle) area() float64 { return r.width * r.height }
func (c circle) area() float64   { return math.Pi * c.radius * c.radius }

func main() {
	objArea := func(g geometry) { fmt.Println(g.area()) }

	r := retangle{width: 4, height: 6}
	c := circle{radius: 5}

	objArea(r)
	objArea(c)
}
