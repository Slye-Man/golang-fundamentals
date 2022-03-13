package main

import (
	"fmt"
	"errors"
)

type Stuff struct {
	values []int
}

func (stu *Stuff) Get(index int) (int, error) { //return value when found
												//return error when not found
	if index > len(stu.values) {
		return 0, errors.New(fmt.Sprintf("No element at index %v", index))
	} else {
		return stu.values[index], nil //no error occurs return nil as error
	}

}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value is", value)
	}
}
