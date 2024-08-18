package characters

type CraftingResult struct {
	cooldown  *Cooldown
	details   *Details
	character *Character
}

func NewCraftingResult(cooldown *Cooldown, details *Details, character *Character) *CraftingResult {
	return &CraftingResult{
		cooldown:  cooldown,
		details:   details,
		character: character,
	}
}

type Details struct {
}
