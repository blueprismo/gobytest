package main

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

type Shape interface{
	Area() float64
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

func (c Circle) Area() float64{
	return math.Pow(c.Radius,2) * math.Pi
}

func (t Triangle) Area() float64{
	return t.Base * t.Height / 2 +1
}


func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
