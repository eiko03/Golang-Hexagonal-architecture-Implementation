package main

import (
	"fmt"
	"hexArch/internal/adapters/core/arithmetic"
	"hexArch/internal/ports"
	"os"
)

func main() {

	var core ports.ArithmeticPort = arithmetic.NewAdapter()

	result, err := core.Addition(1, 3)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

}
