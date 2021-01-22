package singleton

import (
	"fmt"
	"sync"
)

type singleton2 int

var instance2 *singleton2
var mu sync.Mutex

//GetInstance2 use DCL ensure return a singleton
func GetInstance2() *singleton2 {
	if instance2 == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance2 == nil {
			fmt.Println("new singleton2 !!!")
			instance2 = new(singleton2)
		}
	}
	return instance2
}
