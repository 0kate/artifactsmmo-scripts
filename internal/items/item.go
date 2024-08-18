package items

type Name = string

type Item struct {
	code Code
	name Name
}

func NewItem(name Name, code Code) *Item {
	return &Item{code, name}
}
