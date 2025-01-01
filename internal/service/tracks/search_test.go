package tracks

import (
	"context"
	spotifyModel "github.com/jetaimejeteveux/music-catalog/internal/models/spotify"
	spotifyRepo "github.com/jetaimejeteveux/music-catalog/internal/repository/spotify"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func Test_service_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSpotifyOutbond := NewMockSpotifyOutbond(mockCtrl)
	next := "https://api.spotify.com/v1/search?query=bohemian+rhapsody&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.9&offset=10&limit=10"
	type args struct {
		query     string
		pageSize  int
		pageIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    *spotifyModel.SearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:     "bohemian rhapsody",
				pageSize:  10,
				pageIndex: 1,
			},
			want: &spotifyModel.SearchResponse{
				Limit:  10,
				Offset: 0,
				Items: []spotifyModel.SpotifyTrackObject{
					{
						AlbumAlbumType:   "album",
						AlbumTotalTracks: 22,
						AlbumImagesUrl:   []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						Id:               "3z8h0TU7ReDPLIbEnYhWZb",
						Name:             "Bohemian Rhapsody",
					},
					{
						AlbumAlbumType:   "album",
						AlbumTotalTracks: 12,
						AlbumImagesUrl:   []string{"https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0"},
						AlbumName:        "A Night At The Opera (2011 Remaster)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						Id:               "4u7EnebtmKWzUH433cf5Qv",
						Name:             "Bohemian Rhapsody - Remastered 2011",
					},
				},
				Total: 905,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyOutbond.EXPECT().Search(gomock.Any(), args.query, 0, 10).Return(&spotifyRepo.SpotifySearchResponse{
					Tracks: spotifyRepo.SpotifyTracks{
						Href:   "https://api.spotify.com/v1/search?query=bohemian+rhapsody&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.9&offset=0&limit=10",
						Limit:  10,
						Next:   &next,
						Offset: 0,
						Total:  905,
						Items: []spotifyRepo.SpotifyTrackObject{
							{
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:   "album",
									TotalTracks: 22,
									Images: []spotifyRepo.SpotifyAlbumImage{
										{
											URL: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
										},
									},
									Name: "Bohemian Rhapsody (The Original Soundtrack)",
								},
								Artists: []spotifyRepo.SpotifyArtistObject{
									{
										Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
										Name: "Queen",
									},
								},
								Explicit: false,
								Href:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
								Id:       "3z8h0TU7ReDPLIbEnYhWZb",
								Name:     "Bohemian Rhapsody",
							},
							{
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:   "album",
									TotalTracks: 12,
									Images: []spotifyRepo.SpotifyAlbumImage{
										{
											URL: "https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0",
										},
									},
									Name: "A Night At The Opera (2011 Remaster)",
								},
								Artists: []spotifyRepo.SpotifyArtistObject{
									{
										Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
										Name: "Queen",
									},
								},
								Explicit: false,
								Href:     "https://api.spotify.com/v1/tracks/4u7EnebtmKWzUH433cf5Qv",
								Id:       "4u7EnebtmKWzUH433cf5Qv",
								Name:     "Bohemian Rhapsody - Remastered 2011",
							},
						},
					},
				}, nil)
			},
		},
		{
			name: "ERROR: spotify outbond search",
			args: args{
				query:     "bohemian rhapsody",
				pageSize:  10,
				pageIndex: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbond.EXPECT().Search(gomock.Any(), args.query, 0, 10).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				spotifyOutbond: mockSpotifyOutbond,
			}
			got, err := s.Search(context.Background(), tt.args.query, tt.args.pageSize, tt.args.pageIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
