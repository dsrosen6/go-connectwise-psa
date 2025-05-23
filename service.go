package connectwise

import (
	"context"
	"fmt"
	"github.com/dsrosen6/go-connectwise-psa/types"
)

// BOARD ENDPOINTS
func boardIdEndpoint(boardId int) string {
	return fmt.Sprintf("service/boards/%d", boardId)
}

// TODO: Board Auto Assign Resources
// TODO: Board Auto Templates
// TODO: Board Excluded Members
// TODO: Board Infos
// TODO: Board Item Associations
// TODO: Board Items
// TODO: Board Notifications

func (c *Client) ListBoards(ctx context.Context, params *QueryParams) ([]types.Board, error) {
	return apiRequestPaginated[types.Board](ctx, c, "GET", "service/boards", params, nil)
}

func (c *Client) PostBoard(ctx context.Context, board *types.Board) (*types.Board, error) {
	return apiRequestNonPaginated[types.Board](ctx, c, "POST", "service/boards", nil, board)
}

func (c *Client) GetBoard(ctx context.Context, boardId int, params *QueryParams) (*types.Board, error) {
	return apiRequestNonPaginated[types.Board](ctx, c, "GET", boardIdEndpoint(boardId), params, nil)
}

func (c *Client) DeleteBoard(ctx context.Context, boardId int) error {
	if _, err := apiRequestNonPaginated[types.Board](ctx, c, "DELETE", boardIdEndpoint(boardId), nil, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PutBoard(ctx context.Context, boardId int, board *types.Board) (*types.Board, error) {
	return apiRequestNonPaginated[types.Board](ctx, c, "PUT", boardIdEndpoint(boardId), nil, board)
}

func (c *Client) PatchBoard(ctx context.Context, boardId int, patchOps []PatchOp) (*types.Board, error) {
	return apiRequestNonPaginated[types.Board](ctx, c, "PATCH", boardIdEndpoint(boardId), nil, patchOps)
}

// TODO: Get Board Usage Count, Get List of Usage, Post Copy Board, Get Count of Board

// TODO: Board Skill Mappings
// TODO: Board Statuses
// TODO: Board Status Info
// TODO: Board Status Notifications
// TODO: Board Subtype Infos
// TODO: Board Subtypes
// TODO: Board Team Infos
// TODO: Board Teams
// TODO: Board Type Infos
// TODO: Board Types
// TODO: Board Type Subtype Item Association

// TICKET ENDPOINTS

func ticketIdEndpoint(ticketId int) string {
	return fmt.Sprintf("service/tickets/%d", ticketId)
}

func (c *Client) ListTickets(ctx context.Context, params *QueryParams) ([]types.Ticket, error) {
	return apiRequestPaginated[types.Ticket](ctx, c, "GET", "service/tickets", params, nil)
}

func (c *Client) PostTicket(ctx context.Context, ticket *types.Ticket) (*types.Ticket, error) {
	return apiRequestNonPaginated[types.Ticket](ctx, c, "POST", "service/tickets", nil, ticket)
}

func (c *Client) GetTicket(ctx context.Context, ticketId int, params *QueryParams) (*types.Ticket, error) {
	return apiRequestNonPaginated[types.Ticket](ctx, c, "GET", ticketIdEndpoint(ticketId), params, nil)
}

func (c *Client) DeleteTicket(ctx context.Context, ticketId int) error {
	if _, err := apiRequestNonPaginated[struct{}](ctx, c, "DELETE", ticketIdEndpoint(ticketId), nil, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PutTicket(ctx context.Context, ticketId int, ticket *types.Ticket) (*types.Ticket, error) {
	return apiRequestNonPaginated[types.Ticket](ctx, c, "PUT", ticketIdEndpoint(ticketId), nil, ticket)
}

func (c *Client) PatchTicket(ctx context.Context, ticketId int, patchOps []PatchOp) (*types.Ticket, error) {
	return apiRequestNonPaginated[types.Ticket](ctx, c, "PATCH", ticketIdEndpoint(ticketId), nil, patchOps)
}

// ListServiceTicketNotes gets all ticket notes, regardless of if they have a time entry.
//
// This is most likely the one you want to use unless you consistently uncheck the time entry box.
func (c *Client) ListServiceTicketNotes(ctx context.Context, ticketId int, params *QueryParams) ([]types.ServiceTicketNote, error) {
	return apiRequestPaginated[types.ServiceTicketNote](ctx, c, "GET", ticketIdEndpoint(ticketId)+"/allNotes", params, nil)
}

// ListServiceNotes gets all notes that are not time entry.
//
// Not recommended since you will probably get what you need through ListServiceTicketNotes
func (c *Client) ListServiceNotes(ctx context.Context, ticketId int, params *QueryParams) ([]types.ServiceNote, error) {
	return apiRequestPaginated[types.ServiceNote](ctx, c, "GET", ticketIdEndpoint(ticketId)+"notes", params, nil)
}

func (c *Client) PostServiceTicketNote(ctx context.Context, ticketId int, note *types.ServiceNote) (*types.ServiceNote, error) {
	return apiRequestNonPaginated[types.ServiceNote](ctx, c, "POST", ticketIdEndpoint(ticketId)+"notes", nil, note)
}

// TODO: Rest of Ticket Endpoints
