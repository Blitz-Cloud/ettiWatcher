package cmd

import "fmt"

func Help(args []string) error {
	fmt.Println("This is a CLI app that helps me automate some task for EttiHelper\n")
	fmt.Printf("%s", "USAGE:")
	fmt.Printf("\t smeHelper [COMMAND] [VALUE]\n\n")
	fmt.Println("COMMANDS AND EXAMPLES:\n")
	for _, cmd := range CommandRegistry {
		fmt.Printf("\t%s\t\t%s\n\n\t\t\t %s\n\n", cmd.name, cmd.short, cmd.example)

	}
	return nil
}
