package client

import (
	ch "github.com/ecshreve/zeke-sheet/internal/character"
	"github.com/rivo/tview"
)

var (
	app         *tview.Application
	infoFlex    *tview.Flex
	abilityFlex *tview.Flex
	skillFlex   *tview.Flex
)

func RunClient(c *ch.Character) {
	app = tview.NewApplication()

	infoFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	abilityFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	skillFlex = tview.NewFlex().SetDirection(tview.FlexColumn)

	infoFlex.SetBorder(true).SetTitle(" info ").SetBorderPadding(1, 1, 1, 1)
	abilityFlex.SetBorder(true).SetTitle(" abilities ").SetBorderPadding(1, 1, 1, 1)
	skillFlex.SetBorder(true).SetTitle(" skills ").SetBorderPadding(1, 1, 1, 1)

	for _, as := range c.AbilityScores {
		abilityFlex.AddItem(BuildAbilityScoreWidget(as), 0, 1, false)
	}

	skillFlex.AddItem(BuildSkillsWidget(c.Skills), 0, 1, false)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(infoFlex, 0, 1, false).
			AddItem(abilityFlex, 0, 1, false).
			AddItem(skillFlex, 0, 1, false), 0, 5, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
