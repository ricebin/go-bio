package ipni

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	hc *http.Client
}

func NewClient(hc *http.Client) *Client {
	return &Client{hc: hc}
}

func (c *Client) GetAuthor(ctx context.Context, id string) (*AuthorRecord, error) {
	authorUrl := fmt.Sprintf("https://www.ipni.org/api/1/a/%s", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, authorUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	var out AuthorRecord
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCitation(ctx context.Context, ipni string) (*CitationRecord, error) {
	authorUrl := fmt.Sprintf("https://www.ipni.org/api/1/n/%s", ipni)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, authorUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	var out CitationRecord
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
