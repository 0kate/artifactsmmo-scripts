package dto

import (
	"time"

	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type Cooldown struct {
	TotalSeconds     characters.CooldownTotalSeconds     `json:"total_seconds"`
	RemainingSeconds characters.CooldownRemainingSeconds `json:"remaining_seconds"`
	StartedAt        string                              `json:"started_at"`
	Expiration       string                              `json:"expiration"`
	Reason           string                              `json:"reason"`
}

func (r *Cooldown) ToActionResult() *characters.Cooldown {
	startedAtTime, _ := time.Parse(time.RFC3339, r.StartedAt)
	expirationTime, _ := time.Parse(time.RFC3339, r.Expiration)

	return characters.NewCooldown(
		r.TotalSeconds,
		r.RemainingSeconds,
		startedAtTime,
		expirationTime,
		characters.NewCooldownReasonFromString(r.Reason),
	)
}
