package main

import "sync"

type singleton struct{}

var instance *singleton

var once sync.Once

//GetInstance func
func GetInstance() *singleton {

	once.Do(func() {
		if instance == nil {
			instance = &singleton{}
		}
	})
	return instance
}

func main() {

	instan := GetInstance()

}
