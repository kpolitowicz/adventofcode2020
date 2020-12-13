package rainrisk

type Position struct {
	X, Y int
}

type Ferry struct {
	Pos    Position
	Facing int
}

type CommandExecutor func(*Ferry, NavigationCmd)

func NewFerry() *Ferry {
	return &Ferry{
		Pos:    Position{0, 0},
		Facing: 90,
	}
}

func (f *Ferry) ExecuteNavigation(data NavigationData, executor CommandExecutor) {
	for _, cmd := range data {
		executor(f, cmd)
		// f.executeCommand(cmd)
	}
}

func (f *Ferry) ExecuteCommand(cmd NavigationCmd) {
	switch cmd.Cmd {
	case 'N':
		f.Pos = Position{f.Pos.X, f.Pos.Y + cmd.Arg}
	case 'S':
		f.Pos = Position{f.Pos.X, f.Pos.Y - cmd.Arg}
	case 'E':
		f.Pos = Position{f.Pos.X + cmd.Arg, f.Pos.Y}
	case 'W':
		f.Pos = Position{f.Pos.X - cmd.Arg, f.Pos.Y}
	case 'R':
		f.Facing = (f.Facing + cmd.Arg) % 360
	case 'L':
		f.Facing = (f.Facing - cmd.Arg + 360) % 360
	case 'F':
		switch f.Facing {
		case 0:
			f.Pos = Position{f.Pos.X, f.Pos.Y + cmd.Arg}
		case 180:
			f.Pos = Position{f.Pos.X, f.Pos.Y - cmd.Arg}
		case 90:
			f.Pos = Position{f.Pos.X + cmd.Arg, f.Pos.Y}
		case 270:
			f.Pos = Position{f.Pos.X - cmd.Arg, f.Pos.Y}
		}
	}
}
