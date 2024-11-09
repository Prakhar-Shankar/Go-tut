package main 

import "fmt"

func main(){
	var a [5]int 
	fmt.Println("emp: ", a)


	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a)) // this prints out the length of the array

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl: ", b)

	b = [...]int{1,2,3,4,5} // ... - here the compiler counts the number of elements
	fmt.Println("dcl: ", b)

	b= [...]int{100, 3:400, 500} // third index will be 400, before that will be zero that is first and second index will be zero
	fmt.Println("idx: ", b)
	fmt.Println("len of b: ", len(b))

	//lets create a two D array using for loop

	var twoD [2][3] int
	for i := 0; i < 2; i++{
		for j:=0; j< 3; j++{
			twoD[i][j] = i + j;
		}
	}
	fmt.Println("2D: ", twoD)
// Adding prefixed elements here - 

	twoD = [2][3] int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D: ", twoD)
}