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

func (sk Skill) String() string {
	return fmt.Sprintf("%v : Proficient: %t\t Modifier: %d\tAbility: %v", sk.SkillName, sk.Proficient, sk.Modifier, sk.Ability)
}

func NewSkill(s SkillName, proficient bool, proficiencyBonus int, abilityScores map[Ability]*AbilityScore) *Skill {
	ability := SkillToAbility[s]
	modifier := abilityScores[ability].Modifier
	if proficient {
		modifier += proficiencyBonus
	}

	return &Skill{
		SkillName:  s,
		Ability:    ability,
		Proficient: proficient,
		Modifier:   modifier,
	}
}

func (s *Skill) PrettyPrint() string {
	profCheckbox := "☐"
	if s.Proficient {
		profCheckbox = "◆"
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
