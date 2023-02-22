package rest

import (
	"context"
	"fmt"
	"strings"
)

func (c *Client) BeefSummary(ctx context.Context) (*[]string, error) {
	url := fmt.Sprintf("/?type=meat-and-filler&paras=99&format=text")
	resp, err := c.httpClient.Client.R().
		SetContext(ctx).
		Get(url)
	if err != nil {
		return nil, err
	}
	words := strings.Fields(string(resp.Body()))

	return &words, nil
}
