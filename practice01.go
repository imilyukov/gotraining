package main

import (
	"fmt"
	"os"

	"github.com/imilyukov/gotraining/baseutil"
)

func main()  {
	var (
		base byte
		number10 int64
	)

	_, err1 := fmt.Sscanf(os.Args[1], "%d", &base)
	_, err2 := fmt.Sscanf(os.Args[2], "%d", &number10)

	isError := (err1 != nil)
	if isError{
		fmt.Printf("Unable to read <result_base>, reason: %s\n", err1)
	}

	isError = isError && (err2 != nil)
	if err2 != nil {
		fmt.Printf("Unable to read <base10_number>, reason: %s\n", err2)
	}

	if isError {
		os.Exit(1)
	}

	fmt.Print(baseutil.Convert(base, number10))
}