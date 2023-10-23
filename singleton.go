package main

import (
	"fmt"
)

type singleton struct {
	num int
}

var singleInstance *singleton

func getInstance() *singleton {
	if singleInstance == nil {
		fmt.Println("Creating single instance")
		singleInstance = &singleton{0}
	} else {
		fmt.Println("Instance alredy exists - returning it")
	}
	return singleInstance
}

func (s *singleton) addOne() {
	s.num += 1
}

func (s *singleton) showNum() {
	fmt.Println(s.num)
}

func main() {
	s1 := getInstance()
	s1.showNum()
	s1.addOne()
	s1.showNum()

	s2 := getInstance()
	s2.showNum()
	s2.addOne()
	s2.showNum()

	fmt.Println(s1 == s2)
}
