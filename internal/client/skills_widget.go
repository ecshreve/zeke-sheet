package client

import (
	"strings"

	ch "github.com/ecshreve/zeke-sheet/internal/character"
	"github.com/rivo/tview"
)

func BuildSkillsWidget(skills []*ch.Skill) *tview.TextView {
	var skillStrs []string
	for _, skill := range skills {
		skillStrs = append(skillStrs, skill.PrettyPrint())
	}

	skillStr := strings.Join(skillStrs, "\n")
	skillTextView := tview.NewTextView().SetText(skillStr).SetTextAlign(0).SetWordWrap(false)
	skillTextView.SetBorder(true).SetBorderPadding(1, 0, 2, 1)

	return skillTextView
}
