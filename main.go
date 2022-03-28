package main

import (
	"fmt"
	"os"
	"os/user"

	_ "github.com/keisuke713/monkey/lexer"
	_ "github.com/keisuke713/monkey/parser"
	"github.com/keisuke713/monkey/repl"
)

func main() {
	// for debug
	// input := "a + add(b * c) + d"
	// l := lexer.New(input)
	// p := parser.New(l)
	// if program := p.ParseProgram(); program != nil {
	// 	fmt.Print(program.Statements)
	// }
	// ex) b ./main.go:16

	// original

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! this is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
