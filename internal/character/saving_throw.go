package character

import "fmt"

type SavingThrow struct {
	Ability
	Proficient bool
	Modifier   int
}

func (c *Character) PopulateSavingThrows() {
	var saves []*SavingThrow
	for _, a := range c.AbilityScores {
		p := c.ProficientSaves[a.Ability]
		s := c.GetSavingThrow(a.Ability, p)
		saves = append(saves, s)
	}
	c.SavingThrows = saves
}

func (c *Character) GetSavingThrow(a Ability, p bool) *SavingThrow {
	mod := c.AbilityScores[a].Modifier
	if p {
		mod += c.ProficiencyBonus
	}

	return &SavingThrow{
		Ability:    a,
		Proficient: p,
		Modifier:   mod,
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
