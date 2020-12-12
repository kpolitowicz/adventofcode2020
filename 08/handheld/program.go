package handheld

// import "fmt"

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

func (p *Program) Execute() (err *ProgramError) {
	for p.Cursor < len(p.Commands) {
		err = p.executeNextCommand()
		if err != nil {
			break
		}
	}

	return
}

func (p *Program) FixCorruptCommand() {
	for fixCommandCursor := len(p.Commands) - 1; fixCommandCursor >= 0; fixCommandCursor-- {
		if p.Commands[fixCommandCursor].Instruction == "nop" {
			p.Commands[fixCommandCursor].Instruction = "jmp"
			if err := p.Execute(); err == nil {
				return
			}
			p.Commands[fixCommandCursor].Instruction = "nop"
			p.reinitialize()
		}
		if p.Commands[fixCommandCursor].Instruction == "jmp" {
			p.Commands[fixCommandCursor].Instruction = "nop"
			if err := p.Execute(); err == nil {
				return
			}
			p.Commands[fixCommandCursor].Instruction = "jmp"
			p.reinitialize()
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

func (p *Program) reinitialize() {
	p.Accumulator = 0
	p.Cursor = 0

	for i := 0; i < len(p.Commands); i++ {
		p.Commands[i].Executed = false
	}
}
