package main

import (
	"github.com/ecshreve/zeke-sheet/pkg/creator"
	"github.com/kr/pretty"
)

func main() {
	bernerd := creator.CreateCharacter()
	pretty.Print(bernerd)
	//client.RunClient(bernerd)
}
