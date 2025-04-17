package notes

import "time"

type Entry struct {
	Created time.Time `json:"created"`
	Path    string    `json:"path"`
	Tags    []string  `json:"tags"`
}
