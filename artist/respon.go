package artist

type CreatorRespon struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Description string `json:"description"`
	Birth_year  string `json:"birth_year"`
	Death_year  string `json:"death_year"`
}
