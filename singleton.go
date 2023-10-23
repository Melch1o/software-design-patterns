package main

import (
	"fmt"
)

type rating struct {
	num int
}

var singleInstance *rating

func getInstance() *rating {
	if singleInstance == nil {
		fmt.Println("Creating single instance")
		singleInstance = &rating{0}
	} else {
		fmt.Println("Instance alredy exists - returning it")
	}
	return singleInstance
}

func (r *rating) upvote() {
	r.num += 1
}

func (r *rating) downvote() {
	r.num -= 1
}

func (r *singleton) showRating() {
	fmt.Println(r.num)
}

func main() {
	r1 := getInstance()
	r1.showRating()
	r1.upvote()
	r1.showRating()

	r2 := getInstance()
	r2.showRating()
	r2.upvote()
	r2.showRating()

	fmt.Println(r1 == r2)
}
