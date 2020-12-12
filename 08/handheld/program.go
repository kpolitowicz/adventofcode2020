package handheld

type Program struct {
	Commands    []Command
	Cursor      int
	Accumulator int
}

type Command struct {
	Instruction string
	Arg         int
}

func (p *Program) executeNextCommand() {
	switch nextCommand := p.Commands[p.Cursor]; nextCommand.Instruction {
	case "acc":
		p.Accumulator += nextCommand.Arg
		p.Cursor++
	case "jmp":
		p.Cursor += nextCommand.Arg
	case "nop":
		p.Cursor++
	}
}
