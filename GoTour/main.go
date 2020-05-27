package main

import (
	"fmt"
	"net"
)

type Bootcamp struct {
	Lat float64
	Lon float64
}

func main() {
	// x := new(Bootcamp)
	// y := &Bootcamp{}
	// fmt.Println(&x)
	// fmt.Println(&y)
	// fmt.Println(&*x)
	// fmt.Println(*y)

	name := "eth0"
	inf, err := net.InterfaceByName(name)
	if err != nil {
		fmt.Println("msg", "error getting interface", "inf", name, "err", err)
	}

	addrs, err := inf.Addrs()
	if err != nil {
		fmt.Println("msg", "error getting addresses for interface", "inf", name, "err", err)
	}

	if len(addrs) <= 0 {
		fmt.Println("msg", "no addresses found for interface", "inf", name, "err", err)
	}
}
