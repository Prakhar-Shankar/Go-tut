package main

import (
	"fmt"
	"time"
)

func main(){
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3: 
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekend")
	default:
		fmt.Println("It is a weekday")
	}

	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("I am a bool")
		case int: 
			fmt.Println("I am an int")
		default:
			fmt.Println(" Do not know the type ", t)
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("hey")
}