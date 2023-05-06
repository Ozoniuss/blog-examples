package main

import (
	"printer/client"
	"printer/color"
)

func main() {
	client.PrintHelloRed()
	client.PrintHelloYello()

	g := color.GreenPrinter{}
	y := color.YellowPrinter{}
	r := client.ReversePrinter{}

	client.Print(g)
	client.Print(y)
	client.Print(r)
}
