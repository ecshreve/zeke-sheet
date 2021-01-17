package client

import "github.com/rivo/tview"

var (
	app         *tview.Application
	infoFlex    *tview.Flex
	abilityFlex *tview.Flex
)

func RunClient() {
	app = tview.NewApplication()

	infoFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	abilityFlex = tview.NewFlex().SetDirection(tview.FlexColumn)

	infoFlex.SetBorder(true).SetTitle(" info ").SetBorderPadding(1, 1, 1, 1)
	abilityFlex.SetBorder(true).SetTitle(" abilities ").SetBorderPadding(1, 1, 1, 1)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(infoFlex, 0, 1, false).
			AddItem(abilityFlex, 0, 1, false), 0, 5, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
