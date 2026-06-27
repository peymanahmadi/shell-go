package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type HandlerFunc func(args string)

var builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
}

var handlers map[string]HandlerFunc

func init() {
	handlers = map[string]HandlerFunc{
		"echo": func(args string) { fmt.Println(args) },
		"exit": func(_ string) { os.Exit(0) },
		"type": handleType,
	}
}

func handleType(args string) {
	if builtins[args] {
		fmt.Printf("%s is a shell builtin\n", args)
	} else {
		fmt.Printf("%s: not found\n", args)
	}
}

func dispatch(input string) {
	name, args, _ := strings.Cut(input, " ")
	if handler, ok := handlers[name]; ok {
		handler(args)
	} else {
		fmt.Printf("%s: command not found\n", name)
	}
}

func readLoop(r *bufio.Reader) {
	for {
		fmt.Print("$ ")
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		if line = strings.TrimSpace(line); line != "" {
			dispatch(line)
		}
	}
}

func main() {
	readLoop(bufio.NewReader(os.Stdin))
}
