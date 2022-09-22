package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("otterdb started.")
	for {
		fmt.Print("# ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input = strings.Replace(input, "\n", "", -1)
		processInput(input)
	}
}

func processInput(input string) {
	fmt.Printf("Input: %s", input)
}
