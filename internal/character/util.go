package character

func GetTestCharacter() *Character {
	c := &Character{
		Name:             "Bernerd",
		Race:             "Elf",
		Class:            "Monk",
		Level:            7,
		ProficiencyBonus: 3,
	}

	abilityModifiers := map[Ability]int{
		Strength:     1,
		Dexterity:    4,
		Constitution: 1,
		Intelligence: 0,
		Wisdom:       2,
		Charisma:     -1,
	}
	c.AbilityModifiers = abilityModifiers

	abilityScores := []*AbilityScore{
		NewAbilityScore(Strength, 11, c.AbilityModifiers[Strength]),
		NewAbilityScore(Dexterity, 14, c.AbilityModifiers[Dexterity]),
		NewAbilityScore(Constitution, 12, c.AbilityModifiers[Constitution]),
		NewAbilityScore(Intelligence, 11, c.AbilityModifiers[Intelligence]),
		NewAbilityScore(Wisdom, 13, c.AbilityModifiers[Wisdom]),
		NewAbilityScore(Charisma, 9, c.AbilityModifiers[Charisma]),
	}
	c.AbilityScores = abilityScores

	return c
}
