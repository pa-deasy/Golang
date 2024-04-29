package main

import "fmt"

func main() {
	printBoolFrom("true")
	printBoolFrom("yes")
	printBoolFrom("1")
	printBoolFrom("false")
	printBoolFrom("no")
	printBoolFrom("0")
	printBoolFrom("yep")
}

func printBoolFrom(input string) {
	switch input {
	case "true", "yes", "1":
		fmt.Println(input, "equats to", true)
	case "false", "no", "0":
		fmt.Println(input, "equats to", false)
	default:
		fmt.Println(input, "value is not a valid bool")

	}
}
