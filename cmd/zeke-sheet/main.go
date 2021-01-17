package main

import (
	"fmt"

	"github.com/ecshreve/zeke-sheet/pkg/creator"
	"github.com/kr/pretty"
)

func main() {
	bernerd := creator.CreateCharacter()
	pretty.Print(bernerd)
	for _, s := range bernerd.Skills {
		fmt.Println(s.PrettyPrint())
	}
	//client.RunClient(bernerd)
}
