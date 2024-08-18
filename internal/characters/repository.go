package characters

type CharactersRepository interface {
	// Delete a character
	Delete(character *Character) error
}
