package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		// 第一次创建 创建实例
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created")
		}

	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}
