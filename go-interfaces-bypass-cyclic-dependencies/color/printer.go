package color

import "fmt"

type GreenPrinter struct {
}

func (g GreenPrinter) PrintHello() {
	fmt.Println(green, "hello client", reset)
}

type YellowPrinter struct {
}

func (y YellowPrinter) PrintHello() {
	fmt.Println(yellow, "hello client", reset)
}
