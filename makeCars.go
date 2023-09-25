package makeCars

import "fmt"

type Car interface {
	DescribeCar() string
	GetMPG() float64
}

type sedan struct {
	mpg          float64
	color        string
	transmission string
}

func MakeSedan(mpg float64, color string, transmission string) Car {
	return &sedan{
		mpg:          mpg,
		color:        color,
		transmission: transmission,
	}
}

func (s sedan) DescribeCar() string {
	return fmt.Sprintf("The sedan is %v, gets %v miles per gallon, and its transmission is %v.", s.color, s.mpg, s.transmission)
}

func (s sedan) GetMPG() float64 {
	return s.mpg
}

type suv struct {
	mpg          float64
	color        string
	transmission string
}

func MakeSUV(mpg float64, color string, transmission string) Car {
	return &suv{
		mpg:          mpg,
		color:        color,
		transmission: transmission,
	}
}

func (u suv) DescribeCar() string {
	return fmt.Sprintf("The %v, %v SUV gets %v miles per gallon.", u.color, u.transmission, u.mpg)
}

func (u suv) GetMPG() float64 {
	return c.mpg
}

type coupe struct {
	mpg          float64
	color        string
	transmission string
}

func MakeCoupe(mpg float64, color string, transmission string) Car {
	return &coupe{
		mpg:          mpg,
		color:        color,
		transmission: transmission,
	}
}

func (c coupe) DescribeCar() string {
	return fmt.Sprintf("The %v coupe is %v and gets %v miles per gallon.", c.transmission, c.color, c.mpg)
}

func (c coupe) GetMPG() float64 {
	return c.mpg
}
