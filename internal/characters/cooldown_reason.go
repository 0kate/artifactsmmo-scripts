package characters

type CooldownReason int

const (
	CooldownReasonMovement CooldownReason = iota
	CooldownReasonUnknown
)

func NewCooldownReasonFromString(reason string) CooldownReason {
	switch reason {
	case "movement":
		return CooldownReasonMovement
	default:
		return CooldownReasonUnknown
	}
}
