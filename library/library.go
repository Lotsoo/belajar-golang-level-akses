package library

import "fmt"

func SayHello() { // Exported/public dengan cara nama func huruf kapital
	fmt.Println("Hello")
}

func introduce(name string)  { // Unexported/public
	fmt.Println("My name is",name)
}