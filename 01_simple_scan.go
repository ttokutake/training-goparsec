package main

import "fmt"

import "github.com/prataprc/goparsec"

func main() {
	var input = []byte("abc")

	var parser parsec.Parser = parsec.Ident()
	var scanner parsec.Scanner = parsec.NewScanner(input)

	result, nextScanner := parser(scanner)
	fmt.Println("*** parsing result ***")
	fmt.Println("result:", result)
	fmt.Println("nextScanner:", nextScanner)

	identifier, ok := result.(*parsec.Terminal)
	fmt.Println("*** cast result ***")
	fmt.Println("identifier", identifier)
	fmt.Println("ok:", ok)

	fmt.Println("*** fields of identifier ***")
	fmt.Println("Name:", identifier.Name)
	fmt.Println("Value:", identifier.Value)
	fmt.Println("Position:", identifier.Position)
}
