package character

import (
	"fmt"
)

// Ability is a short description of one of the 6 physical or mental characteristics.
type Ability int

const (
	// Strength: physical power.
	Strength Ability = iota

	// Dexterity: agility.
	Dexterity

	// Constitution: endurance.
	Constitution

	// Intelligence: reasoning and memory.
	Intelligence

	// Wisdom: perception and insight.
	Wisdom

	// Charisma: force of personality.
	Charisma
)

// String implements the Stringer interface for the Ability type.
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

// AbilityScore represents a physical or mental characteristic for a character.
type AbilityScore struct {
	Ability
	Base     int
	Modifier int
}

// String implements the Stringer interface for the AbilityScore type.
func (as AbilityScore) String() string {
	return fmt.Sprintf("%v : Base: %d\t Modifier: %d\t Total: %d", as.Ability, as.Base, as.Modifier, as.Base+as.Modifier)
}

// NewAbilityScore returns a pointer to an AbilityScore with the given values.
func NewAbilityScore(a Ability, base, modifier int) *AbilityScore {
	return &AbilityScore{
		Ability:  a,
		Base:     base,
		Modifier: modifier,
	}
}

// PrettyPrint returns a friendly string representation of an AbilityScore to
// use in the TUI.
//
// 	┌────────┐
// 	│   +2   │
// 	└─|────|─┘
// 	  | 18 |
// 	  └────┘
func (as *AbilityScore) PrettyPrint() string {
	modSign := "+"
	modVal := as.Modifier
	if as.Modifier < 0 {
		modSign = "-"
		modVal *= -1
	}

	// Add an extra space if the total is less than two characters (less than 10)
	// so the borders all line up with eachother.
	totVal := as.Base + as.Modifier
	buf := ""
	if totVal < 10 {
		buf = " "
	}

	r1 := "┌────────┐\n"
	r2 := fmt.Sprintf("│   %s%d   │\n", modSign, modVal)
	r3 := "└─┐────┌─┘\n"
	r4 := fmt.Sprintf("  │ %s%d │  \n", buf, totVal)
	r5 := "  └────┘  \n"

	return r1 + r2 + r3 + r4 + r5
}
