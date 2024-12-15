package main

func VariableDeclaration() string {
	var variableName string = "Hello"
	return variableName
}

func StringConcatenation(string1 string, string2 string) string {
	return string1 + string2
}

func IncrementInt(num int) int {
	return num + 1
}

func main() {
	VariableDeclaration()
	StringConcatenation("Hello", "World")
	IncrementInt(1)
}
