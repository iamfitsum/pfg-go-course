package main

import (
	"fmt"
	"slices"
)

func CreateSlice() []string {
	slice := []string{"dog", "cat", "monkey"}
	
	return slice
}

func ModifySlice(slice []string) []string {
	if(len(slice) < 3){
		return slice
	} else {
		slice[1] = "something else"
		return slice
	}
}

func PopSliceValue(slice []string) []string {
	slice = slices.Delete(slice, 2,3)
	
	return slice
}

func main() {
	fmt.Println(CreateSlice())
	fmt.Println(ModifySlice(CreateSlice()))
	fmt.Println(PopSliceValue([]string{"1","2","3","4","5"}))
}
