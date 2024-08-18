package monsters

type Code int

const (
	YellowSlime Code = iota
	Unknown
)

func NewCodeFromString(code string) Code {
	switch code {
	case "yellow_slime":
		return YellowSlime
	default:
		return Unknown
	}
}

func (c Code) String() string {
	switch c {
	case YellowSlime:
		return "yellow_slime"
	default:
		return "unknown"
	}
}
