package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/imilyukov/gotraining/baseutil"
)

func main()  {
	var (
		base int8
		number10 int64
		cpl int
	)

	if len(os.Args) < 4 {
		fmt.Println("Entered too few arguments. Minimum count of arguments is 3")
		os.Exit(1)
	}

	_, err1 := fmt.Sscanf(os.Args[1], "%d", &base)
	_, err2 := fmt.Sscanf(os.Args[2], "%d", &number10)
	_, err3 := fmt.Sscanf(os.Args[3], "%d", &cpl)

	isError := (err1 != nil)
	if isError{
		fmt.Printf("Unable to read <result_base>, reason: %s\n", err1)
	}

	isError = isError && (err2 != nil)
	if err2 != nil {
		fmt.Printf("Unable to read <base10_number>, reason: %s\n", err2)
	}

	isError = (err3 != nil)
	if isError{
		fmt.Printf("Unable to read <count_per_line>, reason: %s\n", err3)
	}

	if isError {
		os.Exit(1)
	}

	ch := make(chan baseutil.BasedNumb, base - 1)
	wg := sync.WaitGroup{}

	for i := base; i > 1; i-- {
		wg.Add(1)

		go func (base int8, number10 int64) {
			wg.Done()

			ch <- baseutil.Numb10(number10).BaseOf(base)
		} (i, number10)
	}

	wg.Wait()

	for i := 1; len(ch) > 0; i++ {
		var bn baseutil.BasedNumb = <-ch
		if len(ch) == 0 || (i % cpl) == 0 {
			fmt.Printf("%s(%d) \n", bn.Numb, bn.Base)
			continue
		}

		fmt.Printf("%s(%d), ", bn.Numb, bn.Base)
	}
}