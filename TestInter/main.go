package TestInter

import (
	"math"
)

type rectangle struct {
	length, breadth float64
}
type triangle struct {
	s1, s2, s3   float64
	height, base float64
}
type circle struct {
	radius float64
}
type square struct {
	side float64
}

type geometry interface {
	area() float64
	perimeter() float64
}

func (r rectangle) area() float64 {
	if r.length < 0 || r.breadth < 0 {
		return -1
	}
	return r.length * r.breadth
}
func (t triangle) area() float64 {
	if t.height < 0 || t.base < 0 {
		return -1
	}
	return 0.5 * t.base * t.height
}
func (c circle) area() float64 {
	if c.radius < 0 {
		return -1
	}
	x := math.Pi * math.Pow(c.radius, 2)
	return math.Floor(x*100) / 100
}
func (s square) area() float64 {
	if s.side < 0 {
		return -1
	}
	return math.Pow(s.side, 2)
}

func (r rectangle) perimeter() float64 {
	if r.length < 0 || r.breadth < 0 {
		return -1
	}
	return 2 * (r.length + r.breadth)
}
func (t triangle) perimeter() float64 {
	if t.s1 < 0 || t.s3 < 0 || t.s2 < 0 {
		return -1
	}
	return t.s3 + t.s2 + t.s1
}
func (c circle) perimeter() float64 {
	if c.radius < 0 {
		return -1
	}
	x := 2 * math.Pi * c.radius
	return math.Floor(x*100) / 100
}
func (s square) perimeter() float64 {
	if s.side < 0 {
		return -1
	}
	return 4 * s.side
}
