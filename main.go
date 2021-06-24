package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rivo/tview"
	"github.com/zmb3/spotify"

	g "github.com/leonardoong/gospotify/globals"
	p "github.com/leonardoong/gospotify/pages"
)

const redirectURI = "http://localhost:8080/callback"

func main() {
	http.HandleFunc("/callback", completeAuth)
	go http.ListenAndServe(":8080", nil)

	g.Auth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)

	g.Auth.SetAuthInfo(g.CLIENT_ID, g.CLIENT_SECRET)

	g.URL = g.Auth.AuthURL(g.State)

	pages := tview.NewPages()

	p.SetUniversalInputCaptures(pages)

	p.ShowLoginPage(pages)

	if err := g.App.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := g.Auth.Token(g.State, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != g.State {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, g.State)
	}

	g.Client = g.Auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	g.CH <- &g.Client
}
