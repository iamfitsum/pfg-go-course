package main

import "fmt"

func CreateMap(key1 string, key2 string, val1 string, val2 string) map[string]string {
	theMap := map[string]string{key1: val1, key2: val2}
	
	return theMap
}

func AddToMap(theMap map[string]int, newKey string, newVal int) map[string]int {
	theMap[newKey] = newVal
	
	return theMap
}

func DeleteFromMap(theMap map[string]string, delKey string) map[string]string {
	delete(theMap, delKey)
	return theMap
}

func main() {
	fmt.Println(CreateMap("key1", "key2", "val1", "val2"))
	fmt.Println(AddToMap(map[string]int{"key1": 1, "key2": 2}, "key3", 3))
	fmt.Println(DeleteFromMap(map[string]string{"key1": "val1", "key2": "val2"}, "key1"))
}
