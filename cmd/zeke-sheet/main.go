package main

import (
	"github.com/ecshreve/zeke-sheet/internal/client"
	"github.com/ecshreve/zeke-sheet/pkg/creator"
)

func main() {
	bernerd := creator.CreateCharacter()
	client.RunClient(bernerd)
}
