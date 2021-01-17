package character

import "fmt"

type SkillName int

const (
	Acrobatics SkillName = iota
	AnimalHandling
	Arcana
	Athletics
	Deception
	History
	Insight
	Intimidation
	Investigation
	Medicine
	Nature
	Perception
	Performance
	Persuasion
	Religion
	SleightOfHand
	Stealth
	Survival
)

var SkillNames = []SkillName{
	Acrobatics,
	AnimalHandling,
	Arcana,
	Athletics,
	Deception,
	History,
	Insight,
	Intimidation,
	Investigation,
	Medicine,
	Nature,
	Perception,
	Performance,
	Persuasion,
	Religion,
	SleightOfHand,
	Stealth,
	Survival,
}

func (sn SkillName) String() string {
	return [...]string{
		"Acrobatics",
		"Animal Handling",
		"Arcana",
		"Athletics",
		"Deception",
		"History",
		"Insight",
		"Intimidation",
		"Investigation",
		"Medicine",
		"Nature",
		"Perception",
		"Performance",
		"Persuasion",
		"Religion",
		"Sleight of Hand",
		"Stealth",
		"Survival",
	}[sn]
}

var SkillToAbility = map[SkillName]Ability{
	Acrobatics:     Dexterity,
	AnimalHandling: Wisdom,
	Arcana:         Intelligence,
	Athletics:      Strength,
	Deception:      Charisma,
	History:        Intelligence,
	Insight:        Wisdom,
	Intimidation:   Charisma,
	Investigation:  Intelligence,
	Medicine:       Wisdom,
	Nature:         Intelligence,
	Perception:     Wisdom,
	Performance:    Charisma,
	Persuasion:     Charisma,
	Religion:       Intelligence,
	SleightOfHand:  Dexterity,
	Stealth:        Dexterity,
	Survival:       Wisdom,
}

type Skill struct {
	SkillName
	Ability
	Proficient bool
	Modifier   int
}

func (c *Character) PopulateSkills() {
	var skills []*Skill
	for _, sn := range SkillNames {
		skills = append(skills, c.GetSkill(sn))
	}
	c.Skills = skills
}

func (c *Character) GetSkill(s SkillName) *Skill {
	ability := SkillToAbility[s]
	abilityMod := c.AbilityModifiers[ability]
	proficient := c.ProficientSkills[s]

	if proficient {
		abilityMod += c.ProficiencyBonus
	}

	return &Skill{
		SkillName:  s,
		Ability:    ability,
		Proficient: proficient,
		Modifier:   abilityMod,
	}
}

func (s *Skill) PrettyPrint() string {
	profCheckbox := "☐"
	if s.Proficient {
		profCheckbox = "◆" //"☑︎"
	}

	modSign := "+"
	modVal := s.Modifier
	if s.Modifier < 0 {
		modSign = "-"
		modVal *= -1
	}

	spacer := "\t"
	if len(s.SkillName.String()) < 8 {
		spacer = "\t\t"
	}

	return fmt.Sprintf("%s  %s%d\t%s%s%s", profCheckbox, modSign, modVal, s.SkillName, spacer, s.Ability)
}
