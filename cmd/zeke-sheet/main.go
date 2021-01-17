package main

import (
	"github.com/ecshreve/zeke-sheet/internal/client"
	"github.com/ecshreve/zeke-sheet/pkg/creator"
)

func main() {
	bernerd := creator.CreateCharacter()
	// pretty.Print(bernerd)
	// for _, s := range bernerd.Skills {
	// 	fmt.Println(s.PrettyPrint())
	// }
	client.RunClient(bernerd)
}
