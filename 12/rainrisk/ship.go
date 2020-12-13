package rainrisk

type Position struct {
	X, Y int
}

type Ferry struct {
	Pos         Position
	Facing      int
	WaypointPos Position
}

type CommandExecutor func(*Ferry, NavigationCmd)

func NewFerry() *Ferry {
	return &Ferry{
		Pos:         Position{0, 0},
		Facing:      90,
		WaypointPos: Position{10, 1},
	}
}

func (f *Ferry) ExecuteNavigation(data NavigationData, executor CommandExecutor) {
	for _, cmd := range data {
		executor(f, cmd)
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

func (f *Ferry) ExecuteWaypointCommand(cmd NavigationCmd) {
	switch cmd.Cmd {
	case 'N':
		f.WaypointPos = Position{f.WaypointPos.X, f.WaypointPos.Y + cmd.Arg}
	case 'S':
		f.WaypointPos = Position{f.WaypointPos.X, f.WaypointPos.Y - cmd.Arg}
	case 'E':
		f.WaypointPos = Position{f.WaypointPos.X + cmd.Arg, f.WaypointPos.Y}
	case 'W':
		f.WaypointPos = Position{f.WaypointPos.X - cmd.Arg, f.WaypointPos.Y}
	case 'R':
		f.rotateWaypoint(cmd.Arg)
	case 'L':
		rightRotation := (-cmd.Arg + 360) % 360
		f.rotateWaypoint(rightRotation)
	case 'F':
		f.Pos = Position{f.Pos.X + cmd.Arg*f.WaypointPos.X, f.Pos.Y + cmd.Arg*f.WaypointPos.Y}
	}
}

func (f *Ferry) rotateWaypoint(degrees int) {
	switch degrees {
	case 90:
		f.WaypointPos = Position{f.WaypointPos.Y, -f.WaypointPos.X}
	case 180:
		f.WaypointPos = Position{-f.WaypointPos.X, -f.WaypointPos.Y}
	case 270:
		f.WaypointPos = Position{-f.WaypointPos.Y, f.WaypointPos.X}
	}
}
