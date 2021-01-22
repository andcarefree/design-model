package singleton

import (
	"fmt"
	"sync"
)

type singleton1 int

var once sync.Once
var instance1 *singleton1

//GetInstance1 use sync.Once return a singleton1
func GetInstance1() *singleton1 {
	once.Do(func() {
		fmt.Println("new instance1 !!!")
		instance1 = new(singleton1)
	})
	return instance1
}
