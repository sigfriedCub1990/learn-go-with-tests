package main

import "math"

func Perimeter(rectangle Rectangle) float64  {
    return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
    return rectangle.Width * rectangle.Height
}

type Rectangle struct {
    Width float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Height * r.Width
}

type Circle struct {
    Radius float64
}

func (r Circle) Area() float64 {
    return math.Pi * math.Pow(r.Radius, 2)
}