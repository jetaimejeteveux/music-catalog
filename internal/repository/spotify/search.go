package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/url"
	"strconv"
)

func (o *outbond) Search(ctx context.Context, query string, offset, limit int) (*SpotifySearchResponse, error) {
	baseUrl := `https://api.spotify.com/v1/search`

	params := url.Values{}
	params.Set("q", query)
	params.Set("type", "track")
	params.Set("limit", strconv.Itoa(limit))
	params.Set("offset", strconv.Itoa(offset))

	urlPath := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error creating new request to spotify")
		return nil, err
	}

	accessToken, tokenType, err := o.GetTokenDetails()
	if err != nil {
		log.Error().Err(err).Msg("Error getting token details")
		return nil, err
	}
	bearerToken := fmt.Sprintf("%s %s", tokenType, accessToken)
	req.Header.Set("Authorization", bearerToken)

	resp, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Error sending request to spotify")
		return nil, err
	}
	defer resp.Body.Close()

	var response SpotifySearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("Error decoding response from spotify")
		return nil, err
	}

	return &response, nil
}
