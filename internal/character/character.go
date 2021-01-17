package character

type Character struct {
	Name             string
	Race             string
	Class            string
	Level            int
	ProficiencyBonus int
	AbilityScores    []*AbilityScore
	AbilityModifiers map[Ability]int
	ProficientSkills map[SkillName]bool
	Skills           []*Skill
}
