package character

import "fmt"

// SavingThrow (sometimes just called a "save") represents an attempt to resist
// a spell, a trap, a poison, a disease, or similar threat.
type SavingThrow struct {
	Ability
	Proficient bool
	Modifier   int
}

// String implements the Stringer interface for the SavingThrow type.
func (st SavingThrow) String() string {
	return fmt.Sprintf("%v : Proficient: %t\t Modifier: %d", st.Ability, st.Proficient, st.Modifier)
}

// NewSavingThrow returns a pointer to a SavingThrow with the given values.
func NewSavingThrow(a *AbilityScore, proficient bool, proficiencyBonus int) *SavingThrow {
	modifier := a.Modifier
	if proficient {
		modifier += proficiencyBonus
	}

	return &SavingThrow{
		Ability:    a.Ability,
		Proficient: proficient,
		Modifier:   modifier,
	}
}

// PrettyPrint returns a friendly string representation of a SavingThrow to
// use in the TUI. Printing out all 6 SavingThrows would look like this:
//
// 	◆ +4	STR
// 	◆ +7	DEX
// 	☐ +1	CON
// 	☐ +0	INT
// 	☐ +2	WIS
// 	☐ -1	CHA
func (st *SavingThrow) PrettyPrint() string {
	profCheckbox := "☐"
	if st.Proficient {
		profCheckbox = "◆"
	}

	modSign := "+"
	modVal := st.Modifier
	if st.Modifier < 0 {
		modSign = "-"
		modVal *= -1
	}

	return fmt.Sprintf("%s %s%d\t%s", profCheckbox, modSign, modVal, st.Ability)
}
