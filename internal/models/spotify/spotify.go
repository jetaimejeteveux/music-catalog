package spotify

type SearchResponse struct {
	Limit  int                  `json:"limit"`
	Items  []SpotifyTrackObject `json:"items"`
	Total  int                  `json:"total"`
	Offset int                  `json:"offset"`
}

type SpotifyTrackObject struct {
	// album related fields
	AlbumType        string   `json:"albumType"`
	AlbumTotalTracks int      `json:"totalTracks"`
	AlbumImagesUrl   []string `json:"albumImagesUrl"`
	AlbumName        string   `json:"albumName"`

	// artist related field
	ArtistsName []string `json:"artists"`

	// track related fileds
	Explicit bool   `json:"explicit"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}
