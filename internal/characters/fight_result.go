package characters

type FightResult struct {
	cooldown  *Cooldown
	fight     *Fight
	character *Character
}

func NewFightResult(cooldown *Cooldown, fight *Fight, character *Character) *FightResult {
	return &FightResult{
		cooldown:  cooldown,
		fight:     fight,
		character: character,
	}
}

func (r *FightResult) Cooldown() *Cooldown {
	return r.cooldown
}

type FightResultType = string

const (
	FightResultWin  FightResultType = "win"
	FightResultLose FightResultType = "lose"
)

type Fight struct {
	result FightResult
}
