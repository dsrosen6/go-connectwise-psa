package connectwise

import (
	"context"
	"fmt"
)

func (c *Client) GetBoards(ctx context.Context, params map[string]string) ([]Board, error) {
	return apiRequestPaginatedIntoSlice[Board](ctx, c, "GET", "service/boards", params, nil)
}

func (c *Client) GetBoard(ctx context.Context, boardId int, params map[string]string) (*Board, error) {
	b := &Board{}
	if err := c.apiRequest(ctx, "GET", fmt.Sprintf("service/boards/%d", boardId), params, nil, b); err != nil {
		return nil, fmt.Errorf("error getting board %d: %w", boardId, err)
	}

	return b, nil
}
