package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func newRectangle(width, height float64) shape {
	return &rectangle{
		width:  width,
		height: height,
	}
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}

func (r *rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func newCircle(radius float64) shape {
	return &circle{radius: radius}
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *circle) volume() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	c := newCircle(21)
	r := newRectangle(10, 20)

	var v interface{}
	v = "10"
	if temp, ok := v.(int); ok {
		v = temp * 10
	}
	fmt.Println(v)

	value, ok := r.(*circle)
	if ok {
		fmt.Println(ok, value.area())
	}

	print("Rectangle", r)
	print("Circle", c)

	mahasiswa := map[string]interface{}{
		"nama":    "Hacktiv8",
		"umur":    20,
		"kelas":   "Golang Basic",
		"isBasic": true,
	}
	temp := mahasiswa["umur"]
	fmt.Println(temp.(int) * 10)
}

func print(t string, s shape) {
	fmt.Println("======================")
	fmt.Printf("%s area : %.1f \n", t, s.area())
	fmt.Printf("%s perimeter : %.1f \n", t, s.perimeter())
}
