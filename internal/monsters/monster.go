package monsters

type MonsterName = string
type MonsterLevel = int
type MonsterHP = int

type Monster struct {
	name  MonsterName
	code  Code
	level MonsterLevel
	hp    MonsterHP
}

func NewMonster(name MonsterName, code Code, level MonsterLevel, hp MonsterHP) *Monster {
	return &Monster{name, code, level, hp}
}

func (m *Monster) Name() MonsterName {
	return m.name
}

func (m *Monster) Code() Code {
	return m.code
}

func (m *Monster) Level() MonsterLevel {
	return m.level
}

func (m *Monster) HP() MonsterHP {
	return m.hp
}
