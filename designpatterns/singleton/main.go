package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			s := GetInstance4()
			fmt.Println(&s.Name)
		}()
	}
	wg.Wait()
}

//懒汉模式
type Singleton struct {
	Name string
}

var Instance *Singleton

func GetInstance() *Singleton {
	if Instance == nil {
		Instance = &Singleton{}
	}
	return Instance
}

//饿汉模式
var Instance2 *Singleton = &Singleton{}

func GetInstance2() *Singleton {
	return Instance2
}

//带锁的懒汉模式 反复上锁导致效率下降
var Instance3 *Singleton
var mut sync.Mutex

func GetInstance3() *Singleton {
	mut.Lock()
	defer mut.Unlock()
	if Instance3 == nil {
		Instance3 = &Singleton{}
	}
	return Instance3
}

//双重检验锁
var Instance4 *Singleton
var mut2 sync.Mutex

func GetInstance4() *Singleton {
	if Instance4 == nil {
		mut2.Lock()
		defer mut2.Unlock()
		if Instance4 == nil {
			Instance4 = &Singleton{Name: "mamama"}
		}
	}
	return Instance4
}

//once 使用懒汉模式
var Instance5 *Singleton
var once sync.Once

func GetInstance5() *Singleton {
	once.Do(func() {
		Instance5 = &Singleton{}
	})
	return Instance5
}
