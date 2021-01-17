package main

import (
	"fmt"

	ch "github.com/ecshreve/zeke-sheet/internal/character"
)

func main() {
	abilities := []*ch.AbilityScore{
		ch.NewAbilityScore(ch.Strength, 11, 1),
		ch.NewAbilityScore(ch.Dexterity, 14, 4),
		ch.NewAbilityScore(ch.Constitution, 12, 1),
		ch.NewAbilityScore(ch.Intelligence, 11, 0),
		ch.NewAbilityScore(ch.Wisdom, 13, 2),
		ch.NewAbilityScore(ch.Charisma, 9, -1),
	}

	for _, a := range abilities {
		fmt.Println(a)
	}
}
