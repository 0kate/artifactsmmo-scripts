package characters

import (
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type MoveResult struct {
	cooldown    *Cooldown
	destination *Destination
	character   *Character
}

func NewMoveResult(cooldown *Cooldown, destination *Destination, character *Character) *MoveResult {
	return &MoveResult{
		cooldown:    cooldown,
		destination: destination,
		character:   character,
	}
}

func (r *MoveResult) Cooldown() *Cooldown {
	return r.cooldown
}

type Destination struct {
	name string
	skin Skin

	position *shared.Position

	content struct {
		type_ string
		code  string
	}
}

func NewDestination(
	name string,
	skin Skin,
	position *shared.Position,
	contentType string,
	contentCode string,
) *Destination {
	return &Destination{
		name:     name,
		skin:     skin,
		position: position,
		content: struct {
			type_ string
			code  string
		}{
			type_: contentType,
			code:  contentCode,
		},
	}
}
