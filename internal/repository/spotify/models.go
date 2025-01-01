package spotify

type SpotifySearchResponse struct {
	Tracks SpotifyTracks `json:"tracks"`
}

type SpotifyTracks struct {
	Href     string               `json:"href"`
	Limit    int                  `json:"limit"`
	Offset   int                  `json:"offset"`
	Next     *string              `json:"next"`
	Previous *string              `json:"previous"`
	Total    int                  `json:"total"`
	Items    []SpotifyTrackObject `json:"items"`
}

type SpotifyTrackObject struct {
	Album    SpotifyAlbumObject    `json:"album"`
	Artists  []SpotifyArtistObject `json:"artists"`
	Explicit bool                  `json:"explicit"`
	Href     string                `json:"href"`
	Id       string                `json:"id"`
	Name     string                `json:"name"`
}

type SpotifyAlbumObject struct {
	AlbumType   string              `json:"album_type"`
	TotalTracks int                 `json:"total_tracks"`
	Images      []SpotifyAlbumImage `json:"images"`
	Name        string              `json:"name"`
}

type SpotifyAlbumImage struct {
	URL string `json:"url"`
}

type SpotifyArtistObject struct {
	Href string `json:"href"`
	Name string `json:"name"`
}
