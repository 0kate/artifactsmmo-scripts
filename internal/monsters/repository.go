package monsters

import (
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type MonstersRepository interface {
	FindAll() (*shared.Pages[*Monster], error)
}
