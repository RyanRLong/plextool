package plextool

import (
	"fmt"
	"path/filepath"

	"gopkg.in/toast.v1"
)

// DisplayToast takes a Film and IncomingData structs
// and displays a toast message on Windows 10
func DisplayToast(film Film, data IncomingData) {

	imagePath, _ := filepath.Abs("./img/plex.png") // TODO Verify image exists or remove from Notification
	notification := toast.Notification{
		AppID:    "{1AC14E77-02E7-4E5D-B744-2EB1AE5198B7}\\WindowsPowerShell\\v1.0\\powershell.exe",
		Title:    "Plex Activity",
		Duration: "short",
		Message:  fmt.Sprintf("%s (%s)\nView Count: %.0f\nTime Elapsed: %s", film.Title, data.Event, film.ViewCount, film.GetElapsedTimePretty()),
		Icon:     imagePath, // This file must exist (remove this line if it doesn't)
		Audio:    "silent",
		Actions:  []toast.Action{
			//TODO {"protocol", "I'm a button", ""},
			//TODO {"protocol", "Me too!", ""},
		},
	}
	notification.Push()
}
