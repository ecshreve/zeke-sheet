package creator

import (
	ch "github.com/ecshreve/zeke-sheet/internal/character"
)

func CreateCharacter() *ch.Character {
	ps := map[ch.SkillName]bool{
		ch.Acrobatics:    true,
		ch.Insight:       true,
		ch.Investigation: true,
		ch.Perception:    true,
		ch.Religion:      true,
	}
	as := getInitialAbilityScores()
	am := getAbilityModifiers(as)

	c := &ch.Character{
		Name:             "Bernerd",
		Race:             "Elf",
		Class:            "Monk",
		Level:            7,
		ProficiencyBonus: 3,
		AbilityScores:    as,
		AbilityModifiers: am,
		ProficientSkills: ps,
	}

	c.PopulateSkills()

	return c
}

func getInitialAbilityScores() []*ch.AbilityScore {
	abilities := []*ch.AbilityScore{
		ch.NewAbilityScore(ch.Strength, 11, 1),
		ch.NewAbilityScore(ch.Dexterity, 14, 4),
		ch.NewAbilityScore(ch.Constitution, 12, 1),
		ch.NewAbilityScore(ch.Intelligence, 11, 0),
		ch.NewAbilityScore(ch.Wisdom, 13, 2),
		ch.NewAbilityScore(ch.Charisma, 9, -1),
	}

	return abilities
}

func getAbilityModifiers(as []*ch.AbilityScore) map[ch.Ability]int {
	abilityModifiers := make(map[ch.Ability]int)
	for _, a := range as {
		abilityModifiers[a.Ability] = a.Modifier
	}
	return abilityModifiers
}
