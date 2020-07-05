package main

import (
	"fmt"
)

func main(){
	var names = []string{"Bob", "Ted", "Green"}

	for idx, name := range(names) {
		names[idx] = "Mr " + name
	}

	fmt.Println(names)
}
