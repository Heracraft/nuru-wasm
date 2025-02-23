package main

import (
	// "embed"
	"fmt"
	// "strconv"
	// "log"
	// "os"
	// "strings"

	// prompt "github.com/AvicennaJr/GoPrompt"
	"github.com/NuruProgramming/Nuru/evaluator"
	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/object"
	"github.com/NuruProgramming/Nuru/parser"

	// "github.com/NuruProgramming/Nuru/styles"
	// "github.com/charmbracelet/bubbles/list"
	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
	// zone "github.com/lrstanley/bubblezone",
	"syscall/js"
)

const PROMPT = ">>> "

func Read(contents string) string {
	env := object.NewEnvironment()

	l := lexer.New(contents)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Kuna Errors Zifuatazo:")

		for _, msg := range p.Errors() {
			fmt.Println("\t" + msg)
		}

	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			return evaluated.Inspect()
		} else {
			return ""
		}
	} else {
		return "Error"
	}

}


func runCode(this js.Value, args []js.Value) interface{} {
	print("Running code")
	code := args[0].String()
	return js.ValueOf(Read(code))
}

func main() {
	fmt.Println("Go WASM initialized")
	js.Global().Set("runCode", js.FuncOf(runCode))
	<-make(chan bool)
}
