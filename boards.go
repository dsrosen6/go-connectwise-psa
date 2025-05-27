package connectwise

import (
	"context"
	"fmt"
)

func boardIdEndpoint(boardId int) string {
	return fmt.Sprintf("service/boards/%d", boardId)
}

func (c *Client) ListBoards(ctx context.Context, params *QueryParams) ([]Board, error) {
	return ApiRequestPaginated[Board](ctx, c, "GET", "service/boards", params, nil)
}

func (c *Client) PostBoard(ctx context.Context, board *Board) (*Board, error) {
	return ApiRequestNonPaginated[Board](ctx, c, "POST", "service/boards", nil, board)
}

func (c *Client) GetBoard(ctx context.Context, boardId int, params *QueryParams) (*Board, error) {
	return ApiRequestNonPaginated[Board](ctx, c, "GET", boardIdEndpoint(boardId), params, nil)
}

func (c *Client) DeleteBoard(ctx context.Context, boardId int) error {
	if _, err := ApiRequestNonPaginated[struct{}](ctx, c, "DELETE", boardIdEndpoint(boardId), nil, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PutBoard(ctx context.Context, boardId int, board *Board) (*Board, error) {
	return ApiRequestNonPaginated[Board](ctx, c, "PUT", boardIdEndpoint(boardId), nil, board)
}

func (c *Client) PatchBoard(ctx context.Context, boardId int, patchOps []PatchOp) (*Board, error) {
	return ApiRequestNonPaginated[Board](ctx, c, "PATCH", boardIdEndpoint(boardId), nil, patchOps)
}
