package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/realtime"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type TransactionsAPI struct {
	http        *http.Client
	realtime    *realtime.Client
	workspaceID string
}

func NewTransactionsAPI(http *http.Client, workspaceID string, rt *realtime.Client) *TransactionsAPI {
	return &TransactionsAPI{
		http:        http,
		realtime:    rt,
		workspaceID: workspaceID,
	}
}

func (api *TransactionsAPI) GetTransactions(ctx context.Context, query ...*types.TransactionsQuery) (*types.Transactions, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions", api.workspaceID)
	var result types.Transactions
	var queryParam *types.TransactionsQuery
	if len(query) > 0 && query[0] != nil {
		queryParam = query[0]
	}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  queryParam,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CreateTransaction(ctx context.Context, body types.CreateTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions", api.workspaceID)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CreateMultipleTransactions(ctx context.Context, body types.CreateTransactions) (*types.Transactions, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/bulk-create", api.workspaceID)
	var result types.Transactions
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) DryRunTransaction(ctx context.Context, body types.CreateTransaction) (*types.DryRunTransaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/dry-run", api.workspaceID)
	var result types.DryRunTransaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) GetTransactionByID(ctx context.Context, transactionId string) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) AcceptDepositOffer(ctx context.Context, transactionId string, body types.OfferActions) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/accept-deposit-offer", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) ApproveTransaction(ctx context.Context, transactionId string, body types.ApproveTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/approve", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CancelTransaction(ctx context.Context, transactionId string, body types.CancelTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/cancel", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CreateSigningRequest(ctx context.Context, transactionId string) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/create-signing-request", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) DeclineTransaction(ctx context.Context, transactionId string, body types.CancelTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/decline", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) GetTransactionEvents(ctx context.Context, transactionId string) (*types.TransactionEvents, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/events", api.workspaceID, transactionId)
	var result types.TransactionEvents
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

// Subscribe opens a WebSocket subscription that mirrors GetTransactions(query).
//
// The first frame on the returned stream is the historical match (same shape
// as GetTransactions); subsequent frames are live updates, each typically a
// single-element transactions list. Same query DTO, same response DTO — WS
// is "GET extended". The same filters (transactionStatuses, transactionTypes,
// accountId, createdAtFrom/To, etc.) apply to both the initial replay and to
// each live update.
//
// To skip the initial replay (only live updates), pass a query with Limit
// set to "0" — the server returns an empty transactions list and then begins
// streaming.
//
// Always call stream.Close() when done — it sends UNSUBSCRIBE and tears down
// the connection. Channel closes when ctx is cancelled, Close is called, or
// the connection drops; check stream.Err() afterwards to inspect the cause.
func (api *TransactionsAPI) Subscribe(ctx context.Context, query *types.TransactionsQuery) (*TransactionsStream, error) {
	return api.SubscribeWithFilter(ctx, query)
}

// SubscribeWithFilter is the generic-filter variant of Subscribe — accepts
// any value (typed query, map, etc.) to send as the SUBSCRIBE envelope body.
// Useful when the typed *TransactionsQuery shape doesn't quite fit (the
// generator emits *string for some numeric fields because they're URL query
// params at REST level; over WS they need to be JSON numbers).
func (api *TransactionsAPI) SubscribeWithFilter(ctx context.Context, filter interface{}) (*TransactionsStream, error) {
	if api.realtime == nil {
		return nil, fmt.Errorf("transactions: realtime client is not configured")
	}
	uri := fmt.Sprintf("/workspaces/%s/transactions", api.workspaceID)
	raw, err := api.realtime.Subscribe(ctx, realtime.Subscription{URI: uri, Filter: filter})
	if err != nil {
		return nil, err
	}
	return newTransactionsStream(raw), nil
}

// TransactionsStream is a typed view over a realtime.Stream; Updates yields
// decoded *types.Transactions frames (initial match first, then live updates
// — same shape both times).
type TransactionsStream struct {
	updates chan *types.Transactions
	inner   *realtime.Stream
	final   func() error
}

func newTransactionsStream(inner *realtime.Stream) *TransactionsStream {
	out := make(chan *types.Transactions, 16)
	var decodeErr error
	go func() {
		defer close(out)
		for f := range inner.Frames() {
			if f.Status >= 400 {
				decodeErr = fmt.Errorf("subscribe failed (status=%d): %s", f.Status, string(f.Body))
				return
			}
			if len(f.Body) == 0 {
				continue
			}
			var v types.Transactions
			if err := json.Unmarshal(f.Body, &v); err != nil {
				decodeErr = fmt.Errorf("decode transactions frame: %w", err)
				return
			}
			out <- &v
		}
	}()
	return &TransactionsStream{
		updates: out,
		inner:   inner,
		final: func() error {
			if decodeErr != nil {
				return decodeErr
			}
			return inner.Err()
		},
	}
}

// Updates returns the decoded frame channel. Closed when the stream ends.
func (s *TransactionsStream) Updates() <-chan *types.Transactions { return s.updates }

// Err returns the connection-ending error after Updates() is closed, or nil
// for a clean shutdown.
func (s *TransactionsStream) Err() error { return s.final() }

// Close sends UNSUBSCRIBE and tears down the connection. Safe to call
// multiple times.
func (s *TransactionsStream) Close() error { return s.inner.Close() }

func (api *TransactionsAPI) RejectOutgoingOffer(ctx context.Context, transactionId string, body types.OfferActions) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/reject-outgoing-offer", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

