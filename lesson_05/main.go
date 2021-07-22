package main
import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() (float64, error)
	Perimetr() (float64, error)
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
func (s Circle) Area() (float64, error){
	if s.radius <= 0{
		return 0, errors.New("Invalid value\n")
	}
	return math.Pi*math.Pow(s.radius,2), nil

}
func (s Rectangle) Area() (float64, error){
	if s.height<=0 || s.width<=0{
		return 0, errors.New("Invalid values\n")
	}
	return s.height*s.width, nil
}
func (s Circle) Perimetr() (float64, error){
	if s.radius <= 0{
		return 0, errors.New("Invalid value\n")
	}
	return 2*math.Pi*s.radius, nil
}
func (s Rectangle) Perimetr() (float64, error){
	if s.width <= 0 || s.height <= 0{
		return 0, errors.New("Invalid values\n")
	}
	return 2*s.height+2*s.width, nil
}
func DescribeShape(s Shape){
	fmt.Println(s)
	res,err := s.Area()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("Area: %.2f\n", res)
	}
	res,err = s.Perimetr()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("Perimetr: %.2f\n", res)
	}
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