package main

import (
	"fmt"

	"github.com/ecshreve/zeke-sheet/pkg/creator"
)

func main() {
	bernerd := creator.CreateCharacter()
	for _, as := range bernerd.AbilityScores {
		fmt.Println(as.PrettyPrint())
	}
	// pretty.Print(bernerd)
	//client.RunClient()
}
