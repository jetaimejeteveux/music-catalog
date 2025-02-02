package tracks

import (
	"context"
	spotifyModel "github.com/jetaimejeteveux/music-catalog/internal/models/spotify"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivities"
	spotifyRepo "github.com/jetaimejeteveux/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userId uint) (*spotifyModel.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * limit

	trackDetails, err := s.spotifyOutbond.Search(ctx, query, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("error Spotify Outbond Search")
		return nil, err
	}
	trackIds := make([]string, len(trackDetails.Tracks.Items))
	for idx, item := range trackDetails.Tracks.Items {
		trackIds[idx] = item.Id
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userId, trackIds)
	if err != nil {
		log.Error().Err(err).Msg("error Get trackActivities")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil
}

func modelToResponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotifyModel.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotifyModel.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {
		artistsName := make([]string, len(item.Artists))
		for idx, artists := range item.Artists {
			artistsName[idx] = artists.Name
		}

		imageUrl := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imageUrl[idx] = image.URL
		}

		items = append(items, spotifyModel.SpotifyTrackObject{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesUrl:   imageUrl,
			AlbumName:        item.Album.Name,

			// artist related field
			ArtistsName: artistsName,

			// track related fileds
			Explicit: item.Explicit,
			Id:       item.Id,
			Name:     item.Name,
			IsLiked:  mapTrackActivities[item.Id].IsLiked,
		})
	}

	return &spotifyModel.SearchResponse{
		Items:  items,
		Limit:  data.Tracks.Limit,
		Total:  data.Tracks.Total,
		Offset: data.Tracks.Offset,
	}

}
