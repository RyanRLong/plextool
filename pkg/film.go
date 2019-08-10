package plextool

import (
	"fmt"
)

// Film represents a film on the Plex Server
type Film struct {
	Title      string
	ViewOffset float64
	ViewCount  float64
}

// GetElapsedTimePretty returns a string of hours and mintutes
// in the format of "HH:MM"
func (data Film) GetElapsedTimePretty() string {
	var hours, minutes int
	if hours = int(data.ViewOffset / 60); hours < 0 {
		hours = 0
	}

	if minutes = int(data.ViewOffset) % 60; minutes < 0 {
		minutes = 0
	}
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
