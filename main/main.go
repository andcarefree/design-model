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

	//singleton:
	switch os.Args[1] {
	case "singleton1": //use sync.Once get a singleton instance
		testSingleton("singleton1")
		return
	case "singleton2": //use DCL get a sinleton instance
		testSingleton("singleton2")
		return
	}

	fmt.Println("no this design model")
}

func testSingleton(model string) {
	var wg sync.WaitGroup
	count := 20
	wg.Add(count)
	for i := 0; i < count; i++ {
		switch model {
		case "singleton1":
			go func() {
				fmt.Println(singleton.GetInstance1())
				wg.Done()
			}()
		case "singleton2":
			go func() {
				fmt.Println(singleton.GetInstance2())
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
