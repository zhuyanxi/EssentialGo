package main

import "fmt"

type Bootcamp struct {
	Lat float64
	Lon float64
}

func main() {
	x := new(Bootcamp)
	y := &Bootcamp{}
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(&*x)
	fmt.Println(*y)
}
