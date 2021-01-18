package character

import "fmt"

type SavingThrow struct {
	Ability
	Proficient bool
	Modifier   int
}

func (st SavingThrow) String() string {
	return fmt.Sprintf("%v : Proficient: %t\t Modifier: %d", st.Ability, st.Proficient, st.Modifier)
}

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
