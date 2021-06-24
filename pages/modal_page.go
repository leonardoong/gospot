package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ShowModal(pages *tview.Pages, label, text string, buttons []string, f func(buttonIndex int, buttonLabel string)) {
	m := tview.NewModal()
	m.SetText(text).
		AddButtons(buttons).
		SetDoneFunc(f).
		SetFocus(0).
		SetBackgroundColor(tcell.ColorDarkSlateGrey)

	pages.AddPage(label, m, true, false)
	pages.ShowPage(label)
}
