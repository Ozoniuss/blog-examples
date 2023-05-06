package client

import "printer/color"

func Print(h helloPrinter) {
	h.PrintHello()
}

func PrintHelloGreen() {
	p := color.GreenPrinter{}
	p.PrintHello()
}

func PrintHelloYello() {
	p := color.YellowPrinter{}
	p.PrintHello()
}
