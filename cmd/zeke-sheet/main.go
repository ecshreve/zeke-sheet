package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ecshreve/zeke-sheet/pkg/character"
)

func main() {
	ss := spew.ConfigState{
		Indent:                  "\t",
		DisableCapacities:       true,
		DisablePointerAddresses: true,
	}

	bernerd := character.GetTestCharacter()

	ss.Dump(bernerd)
	//client.RunClient(bernerd)
}
