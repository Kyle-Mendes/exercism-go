package main

import "fmt"

const TestVersion = 1

func main() {
	var (
		input string
	)

	fmt.Print("Enter a name! \n> ")
	fmt.Scanln(&input)
	fmt.Println(HelloWorld(input))
}

func HelloWorld(input string) string {

	if input == "" {
		input = "World"
	}

	return "Hello, " + input + "!"

}
