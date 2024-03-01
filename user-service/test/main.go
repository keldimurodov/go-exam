package main

import (
	"fmt"
	"os"
)

func main() {
	value, exist := os.LookupEnv("TEST1")
	if !exist {
		fmt.Println("NOT FOUND OS VARIABLE")
	}
	fmt.Println(value)

}
