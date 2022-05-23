package artist

import "time"

type Artist struct {
	ID          int
	Name        string
	Nationality string
	Description string
	Birth_year  string
	Death_year  string
	CreatedAt   time.Time
	UpdateAt    time.Time
}
