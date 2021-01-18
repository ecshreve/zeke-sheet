package main

import (
	"github.com/ecshreve/zeke-sheet/pkg/character"
	"github.com/ecshreve/zeke-sheet/pkg/client"
)

func main() {
	// ss := spew.ConfigState{
	// 	Indent:                  "\t",
	// 	DisableCapacities:       true,
	// 	DisablePointerAddresses: true,
	// }

	bernerd := character.GetTestCharacter()

	//ss.Dump(bernerd)
	client.RunClient(bernerd)
}
