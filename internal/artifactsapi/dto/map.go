package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/maps"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type Map struct {
	Name    string  `json:"name"`
	Skin    string  `json:"skin"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	Content Content `json:"content"`
}

func (m *Map) ToMap() *maps.Map {
	var content *shared.Content
	if m.Content.Type != "" && m.Content.Code != "" {
		content = shared.NewContent(
			shared.ContentType(m.Content.Type),
			shared.ContentCode(m.Content.Code),
		)
	}

	return maps.NewMap(
		m.Name,
		m.Skin,
		shared.NewPosition(m.X, m.Y),
		content,
	)
}
