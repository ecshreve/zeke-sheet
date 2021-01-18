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
	ProficientSaves  map[Ability]bool
	Skills           []*Skill
	SavingThrows     []*SavingThrow

	// Refactoring this type... refactored fields are below this comment.

	// Speed determines	how	far	you	can	move when traveling	and	fighting.
	Speed int

	// Inspiration is a rule the game master can use to reward you for playing
	// your character in a way that’s true to his or her personality traits,
	// ideal, bond, and flaw.
	Inspiration bool

	// Initiative determines the order of turns during combat. When combat
	// starts, every participant makes a Dexterity check to determine their
	// place in the initiative order.
	InitiativeBonus int

	// Armor Class (AC) represents how well your character avoids being wounded
	// in battle. Things that contribute to your AC include the armor you wear,
	// the shield you carry, and your Dexterity modifier.
	ArmorClass int
}
