package runestones

import "time"

// Runestone represents runestone item in the system.
type Runestone struct {
	Id          string
	URL         string
	CreatedTime time.Time
}
