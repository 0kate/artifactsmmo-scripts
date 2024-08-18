package characters

type Slot int

const (
	SlotWeapon Slot = iota
	SlotUnknown
)

func NewSlot(s string) Slot {
	switch s {
	case "weapon":
		return SlotWeapon
	default:
		return SlotUnknown
	}
}

func (s Slot) String() string {
	switch s {
	case SlotWeapon:
		return "weapon"
	default:
		return "unknown"
	}
}
