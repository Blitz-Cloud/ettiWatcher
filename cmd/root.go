package cmd

import (
	"fmt"
	"os"
)

func RootCmd() error {
	initCommands()
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Invalid command")
		Help(args)
		return nil
	} else {
		for _, cmd := range CommandRegistry {
			if cmd.name == args[0] {
				err := cmd.Run(os.Args[2:])
				if err != nil {
					fmt.Println(err)
				}
				return nil
			}
		}
		fmt.Println("\tThis is an invalid command\n\tCheck the help menu")
		return nil
	}

}
