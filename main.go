package main

import (
	"fmt"
	"log"
)

func main() {
	var input string
	var after int
	fmt.Print("Queue: ")
	_, e := fmt.Scanf("%s", &input)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Print("After: ")
	_, e = fmt.Scanf("%d", &after)
	if e != nil {
		log.Fatal(e)
	}
	queue, err := NewQueue(input)
	if err != nil {
		log.Println(err)
	}

	queue.ArrangeUntil(after)
	fmt.Println(queue)
}
