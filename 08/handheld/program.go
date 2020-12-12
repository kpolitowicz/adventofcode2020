package handheld

type Program struct {
	Commands    []Command
	Cursor      int
	Accumulator int
}

type Command struct {
	Instruction string
	Arg         int
	Executed    bool
}

func (p *Program) Execute() {
	for p.Cursor < len(p.Commands) {
		err := p.executeNextCommand()
		if err != "" {
			break
		}
	}
}

func (p *Program) executeNextCommand() (err string) {
	nextCommand := &p.Commands[p.Cursor]
	if nextCommand.Executed {
		return "Error: infinite loop"
	}

	switch nextCommand.Instruction {
	case "acc":
		p.Accumulator += nextCommand.Arg
		p.Cursor++
	case "jmp":
		p.Cursor += nextCommand.Arg
	case "nop":
		p.Cursor++
	}

	nextCommand.Executed = true

	return
}
