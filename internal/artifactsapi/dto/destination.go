package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type Destination struct {
	Name    string  `json:"name"`
	Skin    string  `json:"skin"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	Content Content `json:"content"`
}

func (r *Destination) ToActionResult() *characters.Destination {
	position := shared.NewPosition(r.X, r.Y)

	return characters.NewDestination(r.Name, characters.NewSkin(r.Skin), position, r.Content.Type, r.Content.Code)

}
