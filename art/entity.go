package art

import "time"

type Art struct {
	ID            int
	Name_Art      string
	Name_Artist   string
	Type          string
	Location      string
	Creation_Date string
	CreatedAt     time.Time
	UpdateAt      time.Time
}
