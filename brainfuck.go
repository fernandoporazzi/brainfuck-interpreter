package brainfuck

import (
	"bufio"
	"fmt"
	"os"
)

// this struct should never be exposed to the outside
type interpreter struct {
	size         int
	instructions []byte
	pointer      int
}

// NewInterpreter returns a pointer to the type Interpreter with the given parameters
func NewInterpreter(size int, instructions []byte, pointer int) *interpreter {
	return &interpreter{
		size:         size,
		instructions: instructions,
		pointer:      pointer,
	}
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func incrementMemory(i *interpreter) {
	i.instructions[i.pointer] += 1
}

func decrementMemory(i *interpreter) {
	i.instructions[i.pointer] -= 1
}

func incrementPointer(i *interpreter) {
	if i.pointer == i.size-1 {
		i.pointer = 0
	}

	i.pointer += 1
}

func decrementPointer(i *interpreter) {
	if i.pointer == 0 {
		i.pointer = i.size - 1
	}

	i.pointer -= 1
}

// Run interprets and execute brainfuck code.
func (i *interpreter) Run(code string) {
	i.runWith(code)

	fmt.Println("\n\nEND")
}

func (i *interpreter) runWith(code string) {
	loopStart := -1
	loopEnd := -1
	ignore := 0
	skipClosingLoop := 0

	for index, char := range code {
		if ignore == 1 {
			if char == '[' {
				skipClosingLoop += 1
			}

			if char == ']' {
				if skipClosingLoop != 0 {
					skipClosingLoop -= 1
					continue
				}

				loopEnd = index
				ignore = 0

				if loopStart == loopEnd {
					loopStart = -1
					loopEnd = -1
					continue
				}

				loop := code[loopStart:loopEnd]

				for i.instructions[i.pointer] > 0 {
					i.runWith(loop)
				}
			}
			continue
		}

		if char == '+' {
			incrementMemory(i)
		}

		if char == '-' {
			decrementMemory(i)
		}

		if char == '>' {
			incrementPointer(i)
		}

		if char == '<' {
			decrementPointer(i)
		}

		if char == '.' {
			fmt.Printf("%c", rune(i.instructions[i.pointer]))
		}

		if char == ',' {
			fmt.Print("input: ")

			reader := bufio.NewReader(os.Stdin)
			input, _, err := reader.ReadRune()

			check(err)

			i.instructions[i.pointer] = byte(input)
		}

		if char == '[' {
			loopStart = index + 1
			ignore = 1
		}
	}
}
