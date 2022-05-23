package art

type ArtRespon struct {
	ID            int    `json:"id"`
	Name_Art      string `json:"name_art"`
	Name_Artist   string `json:"name_artist"`
	Location      string `json:"location"`
	Type          string `json:"type"`
	Creation_Date string `json:"creation_date"`
}
