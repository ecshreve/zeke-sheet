package character

import "fmt"

// GetTestCharacter builds a Character and returns a pointer to it. The test
// character is my Elf / Monk named Bernerd from my current campaign.
func GetTestCharacter() *Character {
	c := &Character{
		Name:               "Bernerd",
		Level:              7,
		Race:               "Elf",
		Class:              "Monk",
		ProficiencyBonus:   3,
		Speed:              45,
		Inspiration:        false,
		InitiativeBonus:    4,
		ArmorClass:         16,
		HitPointsCurrent:   45,
		HitPointsMax:       45,
		HitPointsTemporary: 0,
	}

	abilityScores := map[Ability]*AbilityScore{
		Strength:     NewAbilityScore(Strength, 11, 1),
		Dexterity:    NewAbilityScore(Dexterity, 14, 4),
		Constitution: NewAbilityScore(Constitution, 12, 1),
		Intelligence: NewAbilityScore(Intelligence, 11, 0),
		Wisdom:       NewAbilityScore(Wisdom, 13, 2),
		Charisma:     NewAbilityScore(Charisma, 9, -1),
	}
	c.AbilityScores = abilityScores

	savingThrows := map[Ability]*SavingThrow{
		Strength:     NewSavingThrow(abilityScores[Strength], true, 3),
		Dexterity:    NewSavingThrow(abilityScores[Dexterity], true, 3),
		Constitution: NewSavingThrow(abilityScores[Constitution], false, 3),
		Intelligence: NewSavingThrow(abilityScores[Intelligence], false, 3),
		Wisdom:       NewSavingThrow(abilityScores[Wisdom], false, 3),
		Charisma:     NewSavingThrow(abilityScores[Charisma], false, 3),
	}
	c.SavingThrows = savingThrows

	skills := map[SkillName]*Skill{
		Acrobatics:     NewSkill(Acrobatics, true, 3, abilityScores),
		AnimalHandling: NewSkill(AnimalHandling, false, 3, abilityScores),
		Arcana:         NewSkill(Arcana, false, 3, abilityScores),
		Athletics:      NewSkill(Athletics, false, 3, abilityScores),
		Deception:      NewSkill(Deception, false, 3, abilityScores),
		History:        NewSkill(History, false, 3, abilityScores),
		Insight:        NewSkill(Insight, true, 3, abilityScores),
		Intimidation:   NewSkill(Intimidation, false, 3, abilityScores),
		Investigation:  NewSkill(Investigation, true, 3, abilityScores),
		Medicine:       NewSkill(Medicine, false, 3, abilityScores),
		Nature:         NewSkill(Nature, false, 3, abilityScores),
		Perception:     NewSkill(Perception, true, 3, abilityScores),
		Performance:    NewSkill(Performance, false, 3, abilityScores),
		Persuasion:     NewSkill(Persuasion, false, 3, abilityScores),
		Religion:       NewSkill(Religion, true, 3, abilityScores),
		SleightOfHand:  NewSkill(SleightOfHand, false, 3, abilityScores),
		Stealth:        NewSkill(Stealth, false, 3, abilityScores),
		Survival:       NewSkill(Survival, false, 3, abilityScores),
	}
	c.Skills = skills

	return c
}

// PrettyPrintProficiencyBonus returns a friendly string representation of a
// proficiency bonus to use in the TUI.
//
// 	┌────────┐
// 	│   +3   │
// 	└────────┘
func PrettyPrintProficiencyBonus(pb int) string {
	r1 := "┌────────┐\n"
	r2 := fmt.Sprintf("│   +%d   │\n", pb)
	r3 := "└────────┘\n"

	return r1 + r2 + r3
}

// PrettyPrintInitiative returns a friendly string representation of an initiative
// bonus to use in the TUI.
//
// 	┌────────┐
// 	│.. +3 ..│
// 	└────────┘
//
func PrettyPrintInitiative(i int) string {
	r1 := "┌────────┐\n"
	r2 := fmt.Sprintf("│.. +%d ..│\n", i)
	r3 := "└────────┘\n"

	return r1 + r2 + r3
}

// PrettyPrintSpeed returns a friendly string representation of a walking speed
// to use in the TUI.
//
// 	┌────────┐
// 	│** 45 **│
// 	└────────┘
//
func PrettyPrintSpeed(s int) string {
	r1 := "┌────────┐\n"
	r2 := fmt.Sprintf("│** %d **│\n", s)
	r3 := "└────────┘\n"

	return r1 + r2 + r3
}

// PrettyPrintArmorClass returns a friendly string representation of an AC
// to use in the TUI.
//
// 	┌────────┐
// 	│## 16 ##│
// 	└────────┘
//
func PrettyPrintArmorClass(s int) string {
	r1 := "┌────────┐\n"
	r2 := fmt.Sprintf("│## %d ##│\n", s)
	r3 := "└────────┘\n"

	return r1 + r2 + r3
}
