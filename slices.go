package main

import (
	"fmt"
	"slices"
	//"slices"
)

func main(){
	//Initialized a slice, where the lenght of the slice is 0, there is no elements in it, hence s is nil.

	var s []string
	fmt.Println("uninitialized: ", s, s==nil, len(s)==0)

	// Initialized an empty slice, but the lenght is three, the capacity of the slice is 3.

	s = make([]string, 3)
	fmt.Println("empty:", s, "len:", len(s), "capacity:", cap(s))


	// we are now adding elements to the slice 
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	//appending the elements of the slice

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended:", s)

	//copying string using copy() function

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied slice:", c)

	//slicing the slice at the end, here it will print the slice of length five that is 0, 1, 2, 3, 4 index

	l := s[:5]
	fmt.Println("sl1:", l)

	// again slicing the slice from the start it will cut off first two indexes of the slice

	l = s[2:]
	fmt.Println("sl3:", l)

	

	t:= []string {"g", "h", "i"}
	fmt.Println("another slice:", t)

	t2 := []string {"g", "h", "i"}
	if slices.Equal(t, t2){
		fmt.Println("t == t2")
	}
}

