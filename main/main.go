package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/myrepo/designModel/singleton"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("the name of model can not be nil")
		return
	}

	{
	singleton:
		switch os.Args[1] {
		case "singleton2": //use DCL get a sinleton instance
			testSingleton2()
		default:
			break singleton
		}
	}

}

func testSingleton2() {
	var wg sync.WaitGroup
	count := 20
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			fmt.Println(singleton.GetInstance2())
			wg.Done()
		}()
	}
	wg.Wait()
}
