package client

import (
	"fmt"

	"github.com/rivo/tview"

	ch "github.com/ecshreve/zeke-sheet/internal/character"
)

func BuildAbilityScoreWidget(as *ch.AbilityScore) *tview.TextView {
	asStr := as.PrettyPrint()
	asTextView := tview.NewTextView().SetText(asStr).SetTextAlign(1).SetWordWrap(false)
	asTextView.SetBorder(true).SetBorderPadding(1, 0, 2, 1).SetTitle(fmt.Sprintf(" %s ", as.Ability))

	return asTextView
}
