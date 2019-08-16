package ch06

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gist struct {
	Rawurl string `json:"html_url"`
}

type Doer interface {
	doGistsRequest2(user string) (io.Reader, error)
}

type Client struct {
	Gister Doer
}

type Gister struct {}

var doGistsRequest = func(user string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", user))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}
	return &buf, nil
}

func ListGists(user string) ([]string, error) {
	r, err := doGistsRequest(user)
	if err != nil {
		return nil, err
	}
	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}
	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}
	return urls, nil
}

func (g *Gister) doGistsRequest2(user string) (io.Reader, error) {
	return doGistsRequest(user)
}

func (c *Client) ListGists2(user string) ([]string, error) {
	_, err := c.Gister.doGistsRequest2(user)
	if err != nil {
		return nil, err
	}
	return ListGists(user)
}