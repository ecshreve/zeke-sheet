package client

import (
	ch "github.com/ecshreve/zeke-sheet/pkg/character"
	"github.com/rivo/tview"
)

var (
	app          *tview.Application
	infoFlex     *tview.Flex
	baseStatFlex *tview.Flex
	abilityFlex  *tview.Flex
	pbsiacFlex   *tview.Flex
	skillFlex    *tview.Flex
)

func RunClient(c *ch.Character) {
	app = tview.NewApplication()

	infoFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	infoFlex.SetBorder(true).SetTitle(" info ").SetBorderPadding(1, 1, 1, 1)
	infoFlex.AddItem(BuildInfoWidget(c), 0, 1, false)

	baseStatFlex = tview.NewFlex().SetDirection(tview.FlexColumn)

	abilityFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	abilityFlex.SetBorder(true).SetTitle(" abilities ").SetBorderPadding(1, 1, 1, 1)
	for _, as := range c.AbilityScores {
		abilityFlex.AddItem(BuildAbilityScoreWidget(as), 0, 1, false)
	}

	// pbsiac: proficiency bonus, speed, initiative, armor class.
	pbsiacFlex = BuildPBSIACWidget(c.ProficiencyBonus, c.Speed, c.InitiativeBonus, c.ArmorClass)

	baseStatFlex.AddItem(abilityFlex, 0, 2, false)
	baseStatFlex.AddItem(pbsiacFlex, 0, 1, false)

	skillFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	skillFlex.SetBorder(true).SetTitle(" skills ").SetBorderPadding(1, 1, 1, 1)
	skillFlex.AddItem(BuildSkillsWidget(c.Skills), 0, 1, false)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(infoFlex, 0, 1, false).
			AddItem(baseStatFlex, 0, 1, false).
			AddItem(skillFlex, 0, 2, false), 0, 5, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
