package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ecshreve/zeke-sheet/internal/character"
)

func main() {
	ss := spew.ConfigState{
		Indent:                  "\t",
		DisableCapacities:       true,
		DisablePointerAddresses: true,
	}

	bernerd := character.GetTestCharacter()

	ss.Dump(bernerd)

	// for _, s := range bernerd.Skills {
	// 	fmt.Println(s.PrettyPrint())
	// }
	//client.RunClient(bernerd)
}
