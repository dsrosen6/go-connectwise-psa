package connectwise

import (
	"context"
	"fmt"
)

func (c *Client) GetBoards(ctx context.Context, params *QueryParams) ([]Board, error) {
	return apiRequestPaginated[Board](ctx, c, "GET", "service/boards", params, nil)
}

func (c *Client) GetBoard(ctx context.Context, boardId int, params *QueryParams) (*Board, error) {
	endpoint := fmt.Sprintf("service/boards/%d", boardId)
	return apiRequestNonPaginated[Board](ctx, c, "GET", endpoint, params, nil)
}
