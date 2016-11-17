package main

import (
	"fmt"
)

type node struct {
	next  *node
	prev *node
	value float64
}

func main() {
	startElement := new(node)
	scndElement := new(node)
	thrdElement := new(node)
	frthElement := new(node)

	startElement.value = 1
	startElement.next = scndElement

	scndElement.value = 2
	scndElement.next = thrdElement
	scndElement.prev = startElement

	thrdElement.value = 3
	thrdElement.next = nil
	thrdElement.prev = scndElement

	for i := startElement; ; i = i.next {
		fmt.Println("Value is: ", i.value)
		if i.next == nil {
			break
		}
	}

	fmt.Println("--------------------------------------")

	frthElement.value = 2.5
	frthElement.next = thrdElement
	frthElement.prev = scndElement

	thrdElement.prev = frthElement
	scndElement.next = frthElement

	for i := startElement; ; i = i.next {
		fmt.Println("Value is: ", i.value)
		if i.next == nil {
			break
		}
	}
}
