package main

import (
	"fmt"
	"os"
)

func main() {

	t, err := os.Open("./test.txt")

	if err != nil {
		fmt.Println("This is the error code:", err)
	}
	t.Read()
	fmt.Println("This is the error code:", err)
}
