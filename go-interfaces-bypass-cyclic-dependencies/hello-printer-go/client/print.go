package client

import "printer/color"

func Print(h helloPrinter) {
	h.PrintHello()
}

func PrintHelloRed() {
	p := color.GreenPrinter{}
	p.PrintHello()
}

func PrintHelloYello() {
	p := color.YellowPrinter{}
	p.PrintHello()
}
