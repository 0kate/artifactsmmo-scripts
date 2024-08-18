package maps

import (
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type Map struct {
	name     string
	skin     string
	position *shared.Position
	content  *shared.Content
}

func NewMap(name string, skin string, position *shared.Position, content *shared.Content) *Map {
	return &Map{
		name:     name,
		skin:     skin,
		position: position,
		content:  content,
	}
}

func (m *Map) Name() string {
	return m.name
}

func (m *Map) Skin() string {
	return m.skin
}

func (m *Map) Position() *shared.Position {
	return m.position
}

func (m *Map) Content() *shared.Content {
	return m.content
}
