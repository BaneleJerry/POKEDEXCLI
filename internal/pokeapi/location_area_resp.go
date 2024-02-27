package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationAreasResp(pageurl *string) (LocationAreasResp, error) {

	endpoint := "location"
	fullURL := baseURl + endpoint

	if pageurl != nil{
		fullURL = *pageurl
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpsClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return LocationAreasResp{}, fmt.Errorf("%v", resp.Status)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}
