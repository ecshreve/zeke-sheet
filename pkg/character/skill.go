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

// SkillNames is a slice of SkillName in alphabetical order, this is the order
// we'll want use for the TUI.
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

// String imiplements the Stringer interface for the SkillName type.
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

// SkillToAbility is a map of each SkillName to it's associated Ability. This
// map is mainly for convenience.
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

// Skill represents a specific aspect of an ability score, if a character is
// proficienct in a skill, that demonstrates a focus on that aspect.
type Skill struct {
	SkillName
	Ability
	Proficient bool
	Modifier   int
}

// String implements the Stringer interface for the Skill type.
func (sk Skill) String() string {
	return fmt.Sprintf("%v : Proficient: %t\t Modifier: %d\tAbility: %v", sk.SkillName, sk.Proficient, sk.Modifier, sk.Ability)
}

// NewSkill returns a pointer to a Skill based on the given values.
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

// PrettyPrint returns a friendly string representation of a Skill to
// use in the TUI. Printing out a few Skills would look like this:
//
// 	◆ +7	Acrobatics	DEX
// 	☐ +2	AnimalHandling	WIS
// 	☐ +0	Arcana		INT
// 	...
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
