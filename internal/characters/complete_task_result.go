package characters

type CompleteTaskResult struct {
	cooldown  *Cooldown
	reward    *Reward
	character *Character
}

func NewCompleteTaskResult(cooldown *Cooldown, reward *Reward, character *Character) *CompleteTaskResult {
	return &CompleteTaskResult{
		cooldown:  cooldown,
		reward:    reward,
		character: character,
	}
}
