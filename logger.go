package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type MoonbaseLogger struct {
	token  string
	target *url.URL
	client *http.Client
}

func NewMoonbaseLogger(projectId, token string, opts *slog.HandlerOptions) (*slog.JSONHandler, error) {
	target, err := url.Parse(fmt.Sprintf("https://api.moonbasehq.com/v1/projects/%s/logs", projectId))
	if err != nil {
		return nil, err
	}
	return slog.NewJSONHandler(&MoonbaseLogger{
		token:  token,
		target: target,
		client: &http.Client{},
	}, opts), nil
}

type LogResponse struct {
	Success bool `json:"success"`
}

func (l *MoonbaseLogger) Write(buf []byte) (int, error) {
	req := http.Request{
		Method: http.MethodPost,
		URL:    l.target,
		Body:   io.NopCloser(bytes.NewBuffer(buf)),
		Header: http.Header{
			"Content-Type":     []string{"application/json"},
			"X-Moonbase-Token": []string{l.token},
		},
	}
	resp, err := l.client.Do(&req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var logResp LogResponse
	if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&logResp); err != nil {
		return 0, err
	}
	if !logResp.Success {
		return 0, fmt.Errorf("failed to log to moonbase: %s", body)
	}
	return len(buf), nil
}
