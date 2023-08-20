package main

import "fmt"

func main() {
	e := NewEnigma()
	output := e.Encrypt("naik kereta api tut tut tut")
	fmt.Println(output)
}
