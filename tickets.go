package connectwise

import (
	"context"
	"fmt"
)

func ticketIdEndpoint(ticketId int) string {
	return fmt.Sprintf("service/tickets/%d", ticketId)
}

func (c *Client) ListTickets(ctx context.Context, params *QueryParams) ([]Ticket, error) {
	return ApiRequestPaginated[Ticket](ctx, c, "GET", "service/tickets", params, nil)
}

func (c *Client) PostTicket(ctx context.Context, ticket *Ticket) (*Ticket, error) {
	return ApiRequestNonPaginated[Ticket](ctx, c, "POST", "service/tickets", nil, ticket)
}

func (c *Client) GetTicket(ctx context.Context, ticketId int, params *QueryParams) (*Ticket, error) {
	return ApiRequestNonPaginated[Ticket](ctx, c, "GET", ticketIdEndpoint(ticketId), params, nil)
}

func (c *Client) DeleteTicket(ctx context.Context, ticketId int) error {
	if _, err := ApiRequestNonPaginated[struct{}](ctx, c, "DELETE", ticketIdEndpoint(ticketId), nil, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PutTicket(ctx context.Context, ticketId int, ticket *Ticket) (*Ticket, error) {
	return ApiRequestNonPaginated[Ticket](ctx, c, "PUT", ticketIdEndpoint(ticketId), nil, ticket)
}

func (c *Client) PatchTicket(ctx context.Context, ticketId int, patchOps []PatchOp) (*Ticket, error) {
	return ApiRequestNonPaginated[Ticket](ctx, c, "PATCH", ticketIdEndpoint(ticketId), nil, patchOps)
}

// ListServiceTicketNotes gets all ticket notes, regardless of if they have a time entry.
//
// This is most likely the one you want to use unless you consistently uncheck the time entry box.
func (c *Client) ListServiceTicketNotes(ctx context.Context, ticketId int, params *QueryParams) ([]ServiceTicketNoteAll, error) {
	return ApiRequestPaginated[ServiceTicketNoteAll](ctx, c, "GET", ticketIdEndpoint(ticketId)+"/allNotes", params, nil)
}

// ListServiceNotes gets all notes that are not time entry.
//
// Not recommended since you will probably get what you need through ListServiceTicketNotes
func (c *Client) ListServiceNotes(ctx context.Context, ticketId int, params *QueryParams) ([]ServiceTicketNote, error) {
	return ApiRequestPaginated[ServiceTicketNote](ctx, c, "GET", ticketIdEndpoint(ticketId)+"notes", params, nil)
}

func (c *Client) PostServiceTicketNote(ctx context.Context, ticketId int, note *ServiceTicketNote) (*ServiceTicketNote, error) {
	return ApiRequestNonPaginated[ServiceTicketNote](ctx, c, "POST", ticketIdEndpoint(ticketId)+"notes", nil, note)
}
