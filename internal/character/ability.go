package character

import (
	"fmt"
)

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
	buf := "\t"
	if as.Base+as.Modifier < 10 {
		buf = "\t\t"
	}
	return fmt.Sprintf("%v : Base: %d%s Modifier: %d\t Total: %d", as.Ability, as.Base, buf, as.Modifier, as.Base+as.Modifier)
}

func NewAbilityScore(a Ability, b, m int) *AbilityScore {
	return &AbilityScore{
		Ability:  a,
		Base:     b,
		Modifier: m,
	}
}

// PrettyPrint returns a string representation of an AbilityScore to in the TUI.
//
// ┌────────┐
// │   +2   │
// └─|────|─┘
//   | 18 |
//   └────┘
func (as *AbilityScore) PrettyPrint() string {
	modSign := "+"
	modVal := as.Modifier
	if as.Modifier < 0 {
		modSign = "-"
		modVal *= -1
	}

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
