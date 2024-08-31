package main

import (
	"fmt"
	"rust_compiler/lexer"
	"rust_compiler/tokens"
	"time"
)

func main() {
	input := `	let five = 5;
				let ten = 10;
				let add = fn(x,y){
					x + y
				}`
	l := lexer.New(input)

	for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
		time.Sleep(200 * time.Millisecond)
	}
}
