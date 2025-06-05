package main

import "fmt"

func main(){
     
	var name string = "nadun"
	var vilage  = "matara"
	    hobby:= "coding"

		var ages = [3]int{10, 20, 30}

	names := [4]string{"Alice", "Bob", "Charlie", "Diana"}
	names[1] = "luigi"


	fmt.Println("Hello " + name)
	fmt.Println("Im from " + vilage)
	fmt.Println("my hobby is " + hobby)
   
	fmt.Println("Hello, World!")



	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	var scores =[]int{10, 20, 30}
	scores[2] = 25
	scores = append(scores,85)

	fmt.Println(scores, len(scores))
	////slice ranges

	rangeone := scores[0:2]
	rangetwo := names[1:3]
	ramgethee := names[2:]

    fmt.Println(rangeone,len(rangeone))
    fmt.Println(rangetwo,len(rangetwo))
    fmt.Println(ramgethee,len(ramgethee))
}