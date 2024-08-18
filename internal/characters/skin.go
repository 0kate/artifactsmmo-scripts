package characters

type Skin int

const (
	SkinMen1 Skin = iota
	SkinMen2
	SkinMen3
	SkinWomen1
	SkinWomen2
	SkinWomen3
	SkinUnknown
)

func NewSkin(s string) Skin {
	switch s {
	case "Men1":
		return SkinMen1
	case "Men2":
		return SkinMen2
	case "Men3":
		return SkinMen3
	case "Women1":
		return SkinWomen1
	case "Women2":
		return SkinWomen2
	case "Women3":
		return SkinWomen3
	default:
		return SkinUnknown
	}
}

func (s Skin) String() string {
	switch s {
	case SkinMen1:
		return "Men1"
	case SkinMen2:
		return "Men2"
	case SkinMen3:
		return "Men3"
	case SkinWomen1:
		return "Women1"
	case SkinWomen2:
		return "Women2"
	case SkinWomen3:
		return "Women3"
	default:
		return "Unknown"
	}
}
