package characters

import (
	"time"
)

type CooldownTotalSeconds = int64
type CooldownRemainingSeconds = int64
type CooldownExpiration = time.Time

type Cooldown struct {
	totalSeconds     CooldownTotalSeconds
	remainingSeconds CooldownRemainingSeconds
	startedAt        time.Time
	expiration       CooldownExpiration
	reason           CooldownReason
}

func NewCooldown(
	totalSeconds CooldownTotalSeconds,
	remainingSeconds CooldownRemainingSeconds,
	startedAt time.Time,
	expiration time.Time,
	reason CooldownReason,
) *Cooldown {
	return &Cooldown{
		totalSeconds:     totalSeconds,
		remainingSeconds: remainingSeconds,
		startedAt:        startedAt,
		expiration:       expiration,
		reason:           reason,
	}
}

func (c *Cooldown) RemainingSeconds() CooldownRemainingSeconds {
	return c.remainingSeconds
}

func (c *Cooldown) RemainingSecondsInDuration() time.Duration {
	return time.Duration(c.remainingSeconds)
}
