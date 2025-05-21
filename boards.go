package connectwise

import "context"

func (c *Client) GetBoards(ctx context.Context) ([]Board, error) {
	return apiRequestPaginatedIntoSlice[Board](ctx, c, "GET", "service/boards", nil)
}
