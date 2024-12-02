package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type Event struct {
	Version string   `json:"version"`
	Session struct{} `json:"session"`
	Request struct {
		Command           string
		Nlu               struct{}
		OriginalUtterance string `json:"original_utterance"`
	} `json:"request"`
}

type ResponceResult struct {
	Text       string `json:"text"`
	Tts        string `json:"tts,omitempty"`
	EndSession bool   `json:"end_session"`
}

type Response struct {
	Version string         `json:"version"`
	Session struct{}       `json:"session"`
	Result  ResponceResult `json:"response"`
}

func Handler(ctx context.Context, event []byte) (*Response, error) {
	var input Event
	err := json.Unmarshal(event, &input)
	if err != nil {
		return nil, fmt.Errorf("an error has occurred when parsing event: %v", err)
	}

	text := "Привет! Я помогу тебе составить  письмо."

	if input.Request.OriginalUtterance != "" {
		text = input.Request.OriginalUtterance
	}

	return &Response{
		Version: input.Version,
		Session: input.Session,
		Result: ResponceResult{
			Text:       text,
			EndSession: false,
		},
	}, nil
}
