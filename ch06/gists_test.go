package ch06

import (
	"io"
	"strings"
	"testing"
)

type dummyDoer struct{}

func TestListGists(t *testing.T) {
	doGistsRequest = func(user string) (io.Reader, error) {
		return strings.NewReader(`
[
	{"html_url": "https://gists.github.com/example1"},
	{"html_url": "https://gists.github.com/example2"}
]
		`), nil
	}
	urls, err := ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}

func (d *dummyDoer) doGistsRequest2(user string) (io.Reader, error) {
	return strings.NewReader(`
[
	{"html_url": "https://gists.github.com/example1"},
	{"html_url": "https://gists.github.com/example2"}
]
		`), nil
}

func TestListGists2(t *testing.T) {
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists2("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
