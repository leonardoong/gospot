package pages

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	g "github.com/leonardoong/gospotify/globals"
)

func ShowMainPage(pages *tview.Pages) {
	var ga []int
	for i := 0; i < 15; i++ {
		ga = append(ga, -1)
	}
	grid := tview.NewGrid().SetColumns(ga...).SetRows(ga...)

	grid.SetTitle(fmt.Sprintf("Welcome to GoSpotify, [lightgreen]%s!", g.User.DisplayName))

	grid.SetTitleColor(tcell.ColorOrange).
		SetBorderColor(tcell.ColorLightGrey).
		SetBorder(true)

	table := tview.NewTable()

	table.SetSelectable(true, false).
		SetSeparator('|').
		SetBordersColor(tcell.ColorGrey).
		SetTitleColor(tcell.ColorLightSkyBlue).
		SetBorder(true)

	grid.AddItem(table, 0, 0, 15, 15, 0, 0, true)

	setupPlaylist(pages, table)

	pages.AddPage(g.MainPageID, grid, true, false)
	g.App.SetFocus(grid)
	pages.SwitchToPage(g.MainPageID)
}

func setupPlaylist(pages *tview.Pages, table *tview.Table) {
	titleColor := tcell.ColorOrange
	ownerColor := tcell.ColorLightGrey

	playlistTitleHeader := tview.NewTableCell("Playlist Name").
		SetAlign(tview.AlignCenter).
		SetTextColor(titleColor).
		SetSelectable(false)

	ownerTitleHeader := tview.NewTableCell("Owner").
		SetAlign(tview.AlignCenter).
		SetTextColor(ownerColor).
		SetSelectable(false)

	table.SetCell(0, 0, playlistTitleHeader).
		SetCell(0, 1, ownerTitleHeader).
		SetFixed(1, 0)

	go func() {
		// config := &clientcredentials.Config{
		// 	ClientID:     g.CLIENT_ID,
		// 	ClientSecret: g.CLIENT_SECRET,
		// 	TokenURL:     spotify.TokenURL,
		// }

		// token, err := config.Token(context.Background())
		// if err != nil {
		// 	log.Fatalf("couldn't get token: %v", err)
		// }

		// client = g.Client

		msg, playlist, err := g.Client.FeaturedPlaylists()
		if err != nil {
			log.Fatal(err)
		}

		table.SetTitle(msg)

		table.SetSelectedFunc(func(row, column int) {
			ShowTracksPage(pages, &(playlist.Playlists[row-1]))
		})

		for i, pl := range playlist.Playlists {

			plNameCell := tview.NewTableCell(fmt.Sprintf("%-50s", pl.Name)).
				SetMaxWidth(50)
			plNameCell.Color = titleColor

			plOwnerCell := tview.NewTableCell(fmt.Sprintf("%-50s", pl.Owner.DisplayName)).
				SetMaxWidth(50)
			plOwnerCell.Color = ownerColor

			g.App.QueueUpdateDraw(func() {
				table.SetCell(i+1, 0, plNameCell).
					SetCell(i+1, 1, plOwnerCell)
			})
		}

	}()
}
