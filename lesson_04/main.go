package main

import (
"fmt"
"math"
)

type Shape interface {
	Area() float64
	Perimetr() float64
}
func (s Circle) String() string{
	return fmt.Sprintf("Circle: radius %.2f", s.radius)
}
func (s Rectangle) String() string{
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f", s.height,s.width)
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	height float64
	width float64
}
func (s Circle) Area() float64{
	return math.Pi*math.Pow(s.radius,2)
}
func (s Rectangle) Area() float64{
	return s.height*s.width
}
func (s Circle) Perimetr() float64{
	return 2*math.Pi*s.radius
}
func (s Rectangle) Perimetr() float64{
	return 2*s.height+2*s.width
}
func DescribeShape(s Shape){
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimetr: %.2f\n", s.Perimetr())
}
func main(){
	c:=Circle{radius: 8}
	r:=Rectangle{
		height: 9,
		width: 3,
	}
	DescribeShape(c)
	DescribeShape(r)
}