package items

type Code uint

const (
	WoodenStaff Code = iota
	Unknown
)

func (c Code) String() string {
	switch c {
	case WoodenStaff:
		return "wooden_staff"
	default:
		return "unknown"
	}
}
