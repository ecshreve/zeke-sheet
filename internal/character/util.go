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

	return c
}
