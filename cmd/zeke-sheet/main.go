package main

import (
	"github.com/ecshreve/zeke-sheet/internal/character"
	"github.com/ecshreve/zeke-sheet/internal/client"
)

func main() {
	bernerd := character.GetTestCharacter()
	//pretty.Print(bernerd)
	// for _, s := range bernerd.Skills {
	// 	fmt.Println(s.PrettyPrint())
	// }
	client.RunClient(bernerd)
}
