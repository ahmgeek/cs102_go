package main

import "fmt"

type node struct {
	next  *node
	value int
}

func main() {
	startElement := new(node)
	scndElement := new(node)
	thrdElement := new(node)

	startElement.value = 1
	startElement.next = scndElement

	scndElement.value = 3
	scndElement.next = thrdElement

	thrdElement.value = 5
	thrdElement.next = nil

	for i := startElement; ; i = i.next {
		fmt.Println("Value is: ", i.value)
		if i.next == nil {
			break
		}
	}
}
