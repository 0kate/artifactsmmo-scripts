package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/monsters"
)

type Monster struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Level int    `json:"level"`
	HP    int    `json:"hp"`
}

func (m *Monster) ToMonster() *monsters.Monster {
	return monsters.NewMonster(
		m.Name,
		monsters.NewCodeFromString(m.Code),
		monsters.MonsterLevel(m.Level),
		monsters.MonsterHP(m.HP),
	)
}
