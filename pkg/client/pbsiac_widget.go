package client

import (
	ch "github.com/ecshreve/zeke-sheet/pkg/character"
	"github.com/rivo/tview"
)

func BuildPBSIACWidget(pb, s, i, ac int) *tview.Flex {
	pbsiacFlex := tview.NewFlex().SetDirection(tview.FlexColumn)

	pbTextView := tview.NewTextView().SetText(ch.PrettyPrintProficiencyBonus(pb)).SetTextAlign(1).SetWordWrap(false)
	pbTextView.SetBorder(true).SetBorderPadding(1, 0, 2, 1).SetTitle(" prof ")
	pbsiacFlex.AddItem(pbTextView, 0, 1, false)

	sTextView := tview.NewTextView().SetText(ch.PrettyPrintSpeed(s)).SetTextAlign(1).SetWordWrap(false)
	sTextView.SetBorder(true).SetBorderPadding(1, 0, 2, 1).SetTitle(" speed ")
	pbsiacFlex.AddItem(sTextView, 0, 1, false)

	iTextView := tview.NewTextView().SetText(ch.PrettyPrintInitiative(i)).SetTextAlign(1).SetWordWrap(false)
	iTextView.SetBorder(true).SetBorderPadding(1, 0, 2, 1).SetTitle(" init ")
	pbsiacFlex.AddItem(iTextView, 0, 1, false)

	acTextView := tview.NewTextView().SetText(ch.PrettyPrintArmorClass(ac)).SetTextAlign(1).SetWordWrap(false)
	acTextView.SetBorder(true).SetBorderPadding(1, 0, 2, 1).SetTitle(" ac ")
	pbsiacFlex.AddItem(acTextView, 0, 1, false)

	return pbsiacFlex
}
