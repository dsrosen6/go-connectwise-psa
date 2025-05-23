package connectwise

import (
	"context"
	"fmt"
)

func boardIdEndpoint(boardId int) string {
	return fmt.Sprintf("service/boards/%d", boardId)
}

func (c *Client) ListBoards(ctx context.Context, params *QueryParams) ([]Board, error) {
	return apiRequestPaginated[Board](ctx, c, "GET", "service/boards", params, nil)
}

func (c *Client) PostBoard(ctx context.Context, board *Board) (*Board, error) {
	return apiRequestNonPaginated[Board](ctx, c, "POST", "service/boards", nil, board)
}

func (c *Client) GetBoard(ctx context.Context, boardId int, params *QueryParams) (*Board, error) {
	return apiRequestNonPaginated[Board](ctx, c, "GET", boardIdEndpoint(boardId), params, nil)
}

func (c *Client) DeleteBoard(ctx context.Context, boardId int) (*Board, error) {
	return apiRequestNonPaginated[Board](ctx, c, "DELETE", boardIdEndpoint(boardId), nil, nil)
}

func (c *Client) PutBoard(ctx context.Context, board *Board) (*Board, error) {
	return apiRequestNonPaginated[Board](ctx, c, "PUT", "service/boards", nil, board)
}

func (c *Client) PatchBoard(ctx context.Context, boardId int, patchOps []PatchOp) (*Board, error) {
	return apiRequestNonPaginated[Board](ctx, c, "PATCH", boardIdEndpoint(boardId), nil, patchOps)
}

// TODO: PATCH Board
// TODO: List Board Usage Count
// TODO: List Board Usage
// TODO: Copy Board
// TODO: Count Boards

// TODO: Board Autoassign Resources Endpoints
// TODO: Board Auto Templates Endpoints
// TODO: Board Excluded Members Endpoints
// TODO: Board Infos Endpoints
// TODO: Board Item Association Endpoints
// TODO: Board Items Endpoints
// TODO: Board Notifications Endpoints
// TODO: Board Skill Mapping Endpoints
// TODO: Board Statuses Endpoints
// TODO: Board Status Info Endpoints
// TODO: Board Status Notifications Endpoints
// TODO: Board Subtype Infos Endpoints
// TODO: Board Subtypes Endpoints
// TODO: Board Team Infos Endpoints
// TODO: Board Teams Endpoints
// TODO: Board Type Infos Endpoints
// TODO: Board Types Endpoints
// TODO: Board Type Subtype Item Association Endpoints
