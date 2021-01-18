package client

import (
	"fmt"

	"github.com/rivo/tview"

	ch "github.com/ecshreve/zeke-sheet/pkg/character"
)

func BuildInfoWidget(c *ch.Character) *tview.TextView {
	infoStr := fmt.Sprintf("Name: %s\nLevel: %d\nRace: %s\nClass: %s", c.Name, c.Level, c.Race, c.Class)
	infoTextView := tview.NewTextView().SetText(infoStr).SetTextAlign(0).SetWordWrap(false)
	infoTextView.SetBorder(true).SetBorderPadding(1, 0, 2, 1)

	return infoTextView
}
