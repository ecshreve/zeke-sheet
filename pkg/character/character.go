package character

type Character struct {
	Name  string
	Level int

	// Race is the humanoid species of this character.
	Race string

	// Class broadly describes the character’s vocation, special talents, and
	// the tactics the character is most likely to employ.
	Class string

	// Proficiency Bonus represents your experience from a roleplaying POV.
	ProficiencyBonus int

	// Speed determines how far you can move when traveling and fighting.
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

	// AbilityScores contains the quick descriptions of the character's physical
	// and mental characteristics.
	AbilityScores map[Ability]*AbilityScore

	// SavingThrows (sometimes just called a "save") represents an attempt to
	// resist a spell, a trap, a poison, a disease, or similar threat.
	SavingThrows map[Ability]*SavingThrow

	// Skills represents a specific aspect of an ability score, and a character's
	// proficiency in a skill demonstrates a focus on that aspect.
	Skills map[SkillName]*Skill
}
