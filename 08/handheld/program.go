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

type ProgramError struct {
	Message string
}

func (e *ProgramError) Error() string {
	return e.Message
}

func NewError(msg string) (err *ProgramError) {
	return &ProgramError{msg}
}

func (p *Program) Execute() {
	for p.Cursor < len(p.Commands) {
		err := p.executeNextCommand()
		if err != nil {
			break
		}
	}
}

func (p *Program) executeNextCommand() (err *ProgramError) {
	nextCommand := &p.Commands[p.Cursor]
	if nextCommand.Executed {
		return NewError("Error: infinite loop")
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
