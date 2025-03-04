package cmd

var CommandRegistry []Command

func initCommands() {
	incrementCmd := &Command{
		name:    "incr",
		short:   "This command is used to increase the values of the current university year and semester",
		long:    "",
		example: "semHelper incr year/semester To increase the values of year or semester",
		exec:    Increment,
	}
	helpCmd := &Command{
		name:    "help",
		short:   "This command displays this message",
		long:    "",
		example: "",
		exec:    Help,
	}
	setCmd := &Command{
		name:    "set",
		short:   "This command is used for setting the text editor and location for the new projects",
		long:    "",
		example: "semHelper set location /home/user/uni/labs\t semHelper set editor code",
		exec:    Set,
	}
	newCmd := &Command{
		name:    "new",
		short:   "This command it is used to create new projects or docs",
		long:    "",
		example: "semHelper new lab name or semHelper new blog name",
		exec:    New,
	}
	generateMarkdownFromLabs := &Command{
		name:    "generateMd",
		short:   "",
		long:    "",
		example: "",
		exec:    generateMd,
	}
	generateAiCompletionForLabs := &Command{
		name:    "generateAi",
		short:   "",
		long:    "",
		example: "",
		exec:    func(s []string) error { return nil },
	}
	RegisterCmd(incrementCmd, helpCmd, setCmd, newCmd, generateMarkdownFromLabs, generateAiCompletionForLabs)
}
