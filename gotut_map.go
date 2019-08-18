package main

import "fmt"

func main(){
	// var grades map[string]float32
	grades :=make(map[string]float32)//設立新map，關鍵字用make
	grades["Timmy"]=42
	grades["Jess"]=92
	grades["Sam"]=67
	fmt.Println(grades)
	TimsGrade :=grades["Timmy"]
	fmt.Println(TimsGrade)
	delete(grades,"Timmy")
	fmt.Println(grades)
	for k,v := range grades{
		fmt.Println(k,":",v)
	}


}