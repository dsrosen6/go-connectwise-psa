package connectwise

import "context"

func (c *Client) GetBoards(ctx context.Context, params map[string]string) ([]Board, error) {
	return apiRequestPaginatedIntoSlice[Board](ctx, c, "GET", "service/boards", params, nil)
}
