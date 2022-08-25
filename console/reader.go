package console

import (
	"bufio"
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"os"
	"strings"
	"time"
)

func StartConsole() {
	go func() {
		time.Sleep(time.Millisecond * 500)
		source := &Console{}
		fmt.Println("Type help for commands.")
		// I don't use fmt.Scan() because the fmt package intentionally filters out whitespaces, this is how it is implemented.
		scanner := bufio.NewScanner(os.Stdin)
		reader := bufio.NewReader(os.Stdin)

		for {
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			fmt.Println(text)
			if strings.Compare("hi", text) == 0 {
				fmt.Println("Hello World!")
			}

			if scanner.Scan() {
				commandString := scanner.Text()
				if len(commandString) == 0 {
					continue
				}
				commandName := strings.Split(commandString, " ")[0]
				command, ok := cmd.ByAlias(commandName)

				if !ok {
					output := &cmd.Output{}
					output.Errorf("Unknown command '%v'", commandName)
					for _, e := range output.Errors() {
						fmt.Println(e)
					}
					continue
				}

				command.Execute(strings.TrimPrefix(strings.TrimPrefix(commandString, commandName), " "), source)
			}
		}
	}()
}
