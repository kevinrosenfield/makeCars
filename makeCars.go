package makeCars

import "fmt"

type Car interface {
	DescribeCar() string
}

type sedan struct {
	mpg          int
	color        string
	transmission string
}

func MakeSedan() Car {
	return &sedan{}
}

func (s sedan) DescribeCar() string {
	return fmt.Sprintf("The sedan is %v, gets %v miles per gallon, and its transmission is %v.", s.color, s.mpg, s.transmission)
}

type suv struct {
	mpg          int
	color        string
	transmission string
}

func MakeSUV() Car {
	return &suv{}
}

func (u suv) DescribeCar() string {
	return fmt.Sprintf("The %v, %v SUV gets %v miles per gallon.", u.color, u.transmission, u.mpg)
}

type coupe struct {
	mpg          int
	color        string
	transmission string
}

func MakeCoupe() Car {
	return &coupe{}
}

func (c coupe) DescribeCar() string {
	return fmt.Sprintf("The %v coupe is %v and gets %v miles per gallon.", c.transmission, c.color, c.mpg)
}
