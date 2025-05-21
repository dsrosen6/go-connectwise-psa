package connectwise

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://api-na.myconnectwise.net/v4_6_release/apis/3.0"
)

type PaginationDetails struct {
	HasMorePages bool
	NextLink     string
}

type doRequestResult struct {
	data       []byte
	header     http.Header
	statusCode int
}

func (c *Client) apiRequest(ctx context.Context, method, endpoint string, payload io.Reader, target interface{}) error {
	result, err := c.doRequest(ctx, method, createFullUrl(endpoint), payload)
	if err != nil {
		return err
	}

	if result.statusCode < 200 || result.statusCode >= 300 {
		return fmt.Errorf("bad status: %d", result.statusCode)
	}

	if target != nil {
		if err := json.Unmarshal(result.data, target); err != nil {
			return fmt.Errorf("unmarshaling the response to json: %w", err)
		}
	}

	return nil
}

func (c *Client) apiRequestPaginated(ctx context.Context, method, url string, body io.Reader, handlePage func(data []byte) error) error {
	for {
		result, err := c.doRequest(ctx, method, url, body)
		if err != nil {
			return err
		}

		if result.statusCode < 200 || result.statusCode >= 300 {
			return fmt.Errorf("bad status: %d", result.statusCode)
		}

		if err := handlePage(result.data); err != nil {
			return err
		}

		linkHeader := result.header.Get("Link")
		nextUrl, found := parseLinkHeader(linkHeader, "next")
		if !found {
			break
		}
		url = nextUrl
	}

	return nil
}

func apiRequestPaginatedIntoSlice[T any](ctx context.Context, c *Client, method, endpoint string, body io.Reader) ([]T, error) {
	var allItems []T
	u := createFullUrl(endpoint)
	err := c.apiRequestPaginated(ctx, method, u, body, func(data []byte) error {
		var pageItems []T
		if err := json.Unmarshal(data, &pageItems); err != nil {
			return fmt.Errorf("unmarshaling page: %w", err)
		}
		allItems = append(allItems, pageItems...)
		return nil
	})

	return allItems, err
}

func (c *Client) doRequest(ctx context.Context, method, url string, body io.Reader) (*doRequestResult, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating the request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("clientId", c.clientId)
	req.Header.Set("Authorization", c.encodedCreds)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending the request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}

	return &doRequestResult{
		data:       data,
		header:     res.Header,
		statusCode: res.StatusCode,
	}, nil
}

func parseLinkHeader(linkHeader, rel string) (string, bool) {
	links := strings.Split(linkHeader, ",")
	for _, link := range links {
		parts := strings.Split(strings.TrimSpace(link), ";")
		if len(parts) < 2 {
			continue
		}
		urlPart := strings.Trim(parts[0], "<>")
		relPart := strings.TrimSpace(parts[1])
		if relPart == fmt.Sprintf(`rel="%s"`, rel) {
			return urlPart, true
		}
	}

	return "", false
}

func createFullUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", baseUrl, endpoint)
}
