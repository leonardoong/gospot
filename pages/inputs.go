package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	g "github.com/leonardoong/gospotify/globals"
)

func SetUniversalInputCaptures(pages *tview.Pages) {
	g.App.EnableMouse(true)
}

func SetTrackPageHandlers(pages *tview.Pages, grid *tview.Grid) {
	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc: // User wants to go back.
			pages.RemovePage(g.TrackPageID)
		}
		return event
	})
}