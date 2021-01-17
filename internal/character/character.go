package character

type Character struct {
	Name          string
	Race          string
	Class         string
	Level         int
	AbilityScores []*AbilityScore
}
