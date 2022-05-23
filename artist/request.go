package artist

type Creator struct {
	Name        string `json:"name" binding:"required"`
	Nationality string `json:"nationality" binding:"required"`
	Description string `json:"description" binding:"required"`
	Birth_year  string `json:"birth_year" binding:"required"`
	Death_year  string `json:"death_year" binding:"required"`
}
