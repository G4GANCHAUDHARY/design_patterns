package main

import (
	"fmt"
	"sync"
)

var m = sync.Mutex{}
var instance *singleton

type singleton struct{}

func newSingleton() *singleton {
	if instance == nil {
		m.Lock()
		defer m.Unlock()
		if instance == nil {
			instance = &singleton{}
			fmt.Print("First !\n")
			return instance
		}
		fmt.Print("Returning in 1st else\n")
		return instance
	}
	fmt.Print("Returning in 2nd else\n")
	return instance
}

func main() {
	for i := 0; i < 100; i++ {
		go newSingleton()
	}
	fmt.Scanln()
}
