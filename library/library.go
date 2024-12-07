package library

import "fmt"

func SayHello(name string) { // Exported/public dengan cara nama func huruf kapital
	fmt.Println("Hello")
	introduce(name) // solusi memanggil introduce pada SayHello
}

func introduce(name string)  { // Unexported/public
	fmt.Println("My name is",name)
}