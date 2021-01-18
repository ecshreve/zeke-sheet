package character

func GetTestCharacter() *Character {
	c := &Character{
		Name:             "Bernerd",
		Race:             "Elf",
		Class:            "Monk",
		Level:            7,
		ProficiencyBonus: 3,
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
