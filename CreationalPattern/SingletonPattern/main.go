/**
@author: Ly Tran Vinh
@contact: lytranvinh.work@gmail.com
@content: Abstract Singleton Pattern
*/

package main

import (
	"fmt"
	"sync"
)

// Handle concurrency 
var lock =  &sync.Mutex{}

type single struct {
	// TODO - Add some properties
}

// Variable to store single instance
var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	
	wg := &sync.WaitGroup{}

	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getInstance()
		}()
    }

	wg.Wait()
}
