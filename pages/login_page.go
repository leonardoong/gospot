package pages

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	g "github.com/leonardoong/gospotify/globals"
)


func ShowLoginPage(pages *tview.Pages) {
	form := tview.NewForm()

	form.SetButtonsAlign(tview.AlignCenter).
		SetLabelColor(tcell.ColorWhite).
		SetTitle("Terminal Spotify").
		SetTitleColor(tcell.ColorOrange).
		SetBorder(true).
		SetBorderColor(tcell.ColorGreen)

	form.AddButton("Login", func() {
		if attemptLogin(pages) {
			pages.RemovePage(g.LoginPageID)
			ShowMainPage(pages)
		}
	})

	grid := tview.NewGrid().SetColumns(0).SetRows(0)
	grid.AddItem(form, 0, 0, 3, 3, 0, 0, true).
		AddItem(form, 1, 1, 1, 1, 32, 70, true)

	pages.AddPage(g.LoginPageID, grid, true, false)
	g.App.SetFocus(grid)
	pages.SwitchToPage(g.LoginPageID)
}

func attemptLogin(pages *tview.Pages) bool {

	openBrowser(g.URL)

	client := <- g.CH

	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
		return false
	}
	g.User = user.User
	return true
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}


