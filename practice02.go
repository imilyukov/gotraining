package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/imilyukov/gotraining/baseutil"
)

func main()  {
	var (
		base, baseTo byte
		number10 int64
	)

	_, err1 := fmt.Sscanf(os.Args[1], "%d", &base)
	_, err2 := fmt.Sscanf(os.Args[2], "%d", &number10)
	_, err3 := fmt.Sscanf(os.Args[3], "%d", &baseTo)

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
		fmt.Printf("Unable to read <up_to_base>, reason: %s\n", err3)
	}

	if isError {
		os.Exit(1)
	}

	ch := make(chan *string, int(baseTo) + 1)
	wg := sync.WaitGroup{}

	for i := base; i >= base - baseTo; i-- {
		wg.Add(1)

		go func (base byte, number10 int64) {
			wg.Done()
			nb := baseutil.Convert(base, number10)
			ch <- &nb
		} (i, number10)
	}

	wg.Wait()
	close(ch)

	for len(ch) > 0 {
		fmt.Printf("%s ", *<-ch)
	}
}