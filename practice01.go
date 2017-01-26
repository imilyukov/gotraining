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

	if len(os.Args) < 3 {
		fmt.Println("Entered too few arguments. Minimum count of arguments is 2")
		os.Exit(1)
	}

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

	var baseNumb baseutil.BasedNumb = baseutil.Numb10(number10).BaseOf(base)

	fmt.Printf("%s(%d)\n", baseNumb.Numb, baseNumb.Base)
}