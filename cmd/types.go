package cmd

import (
	"fmt"
)

type Command struct {
	name    string
	short   string
	long    string
	example string
	exec    func([]string) error
}

func RegisterCmd(cmds ...*Command) {
	for _, cmd := range cmds {
		CommandRegistry = append(CommandRegistry, *cmd)
	}
}

func (cmd Command) Run(args []string) error {
	if cmd.exec != nil {
		err := cmd.exec(args)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}
	fmt.Println("Command run successfully ")
	return nil
}
