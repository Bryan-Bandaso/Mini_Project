package art

type Arts struct {
	Name_Art      string `json:"name_art" binding:"required"`
	Name_Artist   string `json:"name_artist" binding:"required"`
	Type          string `json:"type" binding:"required"`
	Location      string `json:"location" binding:"required"`
	Creation_Date string `json:"creation_date" binding:"required"`
}
