package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type Events []struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ParseEvents(r io.Reader) (string, error) {
	// buffering standard output
	reader := bufio.NewReader(r)
	// peek top 1 byte
	b, _ := reader.Peek(1)
	if string(b) == "[" {
		// top 1 byte is "[" so this is maybe JSON data
		// decode as JSON
		var events Events
		dec := json.NewDecoder(reader)
		if err := dec.Decode(&events); err != nil {
			return "", err
		}
		return fmt.Sprintf("%v", events), nil
	} else {
		// read one line
		line, err := reader.ReadString('\n')
		return line, err
	}
}
