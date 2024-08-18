package maps

import (
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type MapsRepository interface {
	// Find all maps
	FindAll(content *shared.Content, pages *shared.Pages[*Map]) (*shared.Pages[*Map], error)

	// Find map by position
	Find(position *shared.Position) (*Map, error)
}
