package pages

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"

	g "github.com/leonardoong/gospotify/globals"
)

func ShowTracksPage(pages *tview.Pages, playlist *spotify.SimplePlaylist) {
	var ga []int
	for i := 0; i < 15; i++ {
		ga = append(ga, -1)
	}
	grid := tview.NewGrid().SetColumns(ga...).SetRows(ga...)

	grid.SetTitleColor(tcell.ColorOrange).
		SetBorderColor(tcell.ColorLightGrey).
		SetTitle("Playlist Tracks").
		SetBorder(true)

	SetTrackPageHandlers(pages, grid)

	table := tview.NewTable()

	songHeader := tview.NewTableCell("Song Name").
		SetAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorLightYellow).
		SetSelectable(true)
	artistHeader := tview.NewTableCell("Artist Name").
		SetAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorLightSkyBlue).
		SetSelectable(false)
	albumHeader := tview.NewTableCell("Album").
		SetAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorSpringGreen).
		SetSelectable(false)

	table.SetCell(0, 0, songHeader).
		SetCell(0, 1, artistHeader).
		SetCell(0, 2, albumHeader).
		SetFixed(1, 0)

	go func() {
		setTracksTable(pages, table, playlist)
	}()

	table.SetSelectable(true, false).
		SetSeparator('|').
		SetBordersColor(tcell.ColorGrey).
		SetTitle(playlist.Name).
		SetTitleColor(tcell.ColorLightSkyBlue).
		SetBorder(true)

	grid.AddItem(table, 0, 0, 15, 15, 0, 0, true)

	pages.AddPage(g.TrackPageID, grid, true, false)
	g.App.SetFocus(grid)
	pages.SwitchToPage(g.TrackPageID)
}

func setTracksTable(pages *tview.Pages, table *tview.Table, playlist *spotify.SimplePlaylist) {

	plTrack, err := g.Client.GetPlaylistTracks(playlist.ID)
	if err != nil {
		g.App.QueueUpdateDraw(func() {
			ShowModal(pages, g.GenericAPIErrorModalID, "Error getting tracks", []string{"OK"},
				func(i int, label string) {
					pages.RemovePage(g.GenericAPIErrorModalID)
				})
		})
		return
	}

	for i, track := range plTrack.Tracks {
		g.App.QueueUpdateDraw(func() {
			songName := tview.NewTableCell(fmt.Sprintf("%-40s", track.Track.Name)).SetMaxWidth(40).
				SetTextColor(tcell.ColorLightYellow)

			var artistList strings.Builder
			for i, artist := range track.Track.Artists {
				artistList.WriteString(artist.Name)
				if i != len(track.Track.Artists) - 1 {
					artistList.WriteString(" ,")
				}
			}

			artistName := tview.NewTableCell(fmt.Sprintf("%-40s", artistList.String())).SetMaxWidth(40).
				SetTextColor(tcell.ColorLightSkyBlue)

			albumName := tview.NewTableCell(fmt.Sprintf("%-20s", track.Track.Album.Name)).SetMaxWidth(20).
				SetTextColor(tcell.ColorSpringGreen)

			table.SetCell(i+1, 0, songName)
			table.SetCell(i+1, 1, artistName)
			table.SetCell(i+1, 2, albumName)
		})
	}
}
