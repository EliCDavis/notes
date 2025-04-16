package notes

import "time"

type Image struct {
	Path         string
	Created      time.Time
	OriginalPath string
}
