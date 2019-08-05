// Package plextool .
//
// Copyright 2015 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
