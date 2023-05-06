package client

import "fmt"

type ReversePrinter struct{}

func (r ReversePrinter) PrintHello() {
	fmt.Println("tneilc olleh")
}
