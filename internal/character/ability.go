package character

import "fmt"

type Ability int

const (
	Strength Ability = iota
	Dexterity
	Constitution
	Intelligence
	Wisdom
	Charisma
)

func (a Ability) String() string {
	return [...]string{
		"STR",
		"DEX",
		"CON",
		"INT",
		"WIS",
		"CHA",
	}[a]
}

type AbilityScore struct {
	Ability
	Base     int
	Modifier int
}

func (as AbilityScore) String() string {
	return fmt.Sprintf("%v : Base: %d\t Modifier: %d\t Total: %d", as.Ability, as.Base, as.Modifier, as.Base+as.Modifier)
}

func NewAbilityScore(a Ability, b, m int) *AbilityScore {
	return &AbilityScore{
		Ability:  a,
		Base:     b,
		Modifier: m,
	}
}
