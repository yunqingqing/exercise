package main

import (
	"sync"
	"fmt"
)

/*
sync 提供基本的同步原语，比如互斥锁。

func (o *Once) Do(f func())  Do 保证只调用f一次
*/

type singleton map[string]string

var (
    once sync.Once

    instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

func main() {
	var s singleton = New()

	s["this"] = "that"

	var s2 singleton = New()

	fmt.Println("This is ", s2["this"])
}


