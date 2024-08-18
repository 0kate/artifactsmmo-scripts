package characters

type GatheringResult struct {
	cooldown  *Cooldown
	details   *Details
	character *Character
}

func NewGatheringResult(details *Details, character *Character) *GatheringResult {
	return &GatheringResult{
		details:   details,
		character: character,
	}
}
