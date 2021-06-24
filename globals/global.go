package globals

import (
	"github.com/rivo/tview"
	"github.com/zmb3/spotify"
)

var (
	App = tview.NewApplication()
)

const (
	LoginPageID  = "login_page"
	MainPageID   = "main_page"
	TrackPageID = "track_page"

	LoginLogoutFailureModalID   = "login_failure_modal"
	LoginLogoutCfmModalID       = "logout_modal"
	GenericAPIErrorModalID		= "api_error_modal"
)

const (
	UsrDir       = "usr"
	CredFileName = "cred"
	DownloadDir  = "downloads"
	CLIENT_ID = "36e39e23def54b94a8cacbd5e0a2fe23"
	CLIENT_SECRET = "e1ca9890b0b844dc99974a60939505b7"
)

var (
	URL   string
	CH    = make(chan *spotify.Client)
	State = "spotifyterminal"
	Auth  spotify.Authenticator
	User spotify.User
	Client spotify.Client
)
