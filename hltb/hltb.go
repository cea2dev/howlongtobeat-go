package hltb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const baseUrl = "https://howlongtobeat.com"
const detailUrl = baseUrl + "/game"
const searchUrl = baseUrl + "/api/search"
const imageUrl = baseUrl + "/games"

type Client struct {
	httpClient *http.Client
	rp         requestProvider
}

func New(httpClient *http.Client) (Client, error) {
	if httpClient == nil {
		return Client{}, errors.New("httpClient is nil")
	}

	client := Client{
		httpClient: httpClient,
		rp:         &defaultRequestProvider{},
	}

	return client, nil
}

func (c *Client) Search(args SearchArgs) ([]GameEntry, error) {
	req, err := c.newSearchRequest(args)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, errors.New(string(b))
	}

	var searchResp searchResponse

	err = json.NewDecoder(resp.Body).Decode(&searchResp)

	if err != nil {
		return nil, err
	}

	return searchResp.getGameEntries(imageUrl, args.Term), nil
}

func (c *Client) newSearchRequest(args SearchArgs) (*http.Request, error) {
	if len(strings.TrimSpace(args.Term)) == 0 {
		return nil, errors.New("search request must have at least one search term")
	}

	if args.Page < 1 {
		args.Page = 1
	}

	if args.PageSize < 1 {
		args.PageSize = 20
	}

	payload := getSearchRequestPayload(args)

	data, err := json.Marshal(&payload)

	if err != nil {
		return nil, err
	}

	r, err := c.rp.NewRequest("POST", searchUrl, bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	// TODO find a better way to set this
	r.Header.Set("User-Agent", "PostmanRuntime/7.29.2")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Origin", baseUrl)
	r.Header.Set("Referer", baseUrl)

	return r, nil
}
