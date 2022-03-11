package main

import "github.com/fernandoporazzi/brainfuck-interpreter"

func main() {
	size := 30000
	instructions := make([]byte, size, size)
	pointer := 0

	interpreter := brainfuck.NewInterpreter(size, instructions, pointer)

	// prints 666
	// interpreter.Run(">+++++++++[<++++++>-]<...>++++++++++.")

	// prints Hello World
	interpreter.Run("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
}
