package main

import "fmt"

type S1 struct {
}

type S2 struct {
}

type F interface {
	f()
}

func (s S1) f() {}

func (s *S2) f() {}

func main() {
	var i F
	s1Val := S1{}
	//s1Ptr := &S1{}
	//s2Val := S2{}
	s2Ptr := &S2{}

	i = s1Val

	//i = s1Ptr
	//i = s2Val
	i = s2Ptr

	fmt.Println(i)
}
