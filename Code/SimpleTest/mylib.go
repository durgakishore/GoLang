package simpletest

import (
	"fmt"
)

func messageWriter(greeting, name string) string {

	message := fmt.Sprintf("%v %v", greeting, name)
	fmt.Println(message)
	return message
}

 