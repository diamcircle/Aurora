package auroraclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	hProtocol "github.com/diamnet/go/protocols/aurora"
	"github.com/diamnet/go/support/errors"
)

// BuildURL creates the endpoint to be queried based on the data in the LedgerRequest struct.
// If no data is set, it defaults to the build the URL for all ledgers
func (lr LedgerRequest) BuildURL() (endpoint string, err error) {
	endpoint = "ledgers"

	if lr.forSequence != 0 {
		endpoint = fmt.Sprintf(
			"%s/%d",
			endpoint,
			lr.forSequence,
		)
	} else {
		queryParams := addQueryParams(cursor(lr.Cursor), limit(lr.Limit), lr.Order)
		if queryParams != "" {
			endpoint = fmt.Sprintf(
				"%s?%s",
				endpoint,
				queryParams,
			)
		}
	}

	_, err = url.Parse(endpoint)
	if err != nil {
		err = errors.Wrap(err, "failed to parse endpoint")
	}

	return endpoint, err
}

// HTTPRequest returns the http request for the ledger endpoint
func (lr LedgerRequest) HTTPRequest(auroraURL string) (*http.Request, error) {
	endpoint, err := lr.BuildURL()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("GET", auroraURL+endpoint, nil)
}

// LedgerHandler is a function that is called when a new ledger is received
type LedgerHandler func(hProtocol.Ledger)

// StreamLedgers streams diamnet ledgers. It can be used to stream all ledgers. Use context.WithCancel
// to stop streaming or context.Background() if you want to stream indefinitely.
// LedgerHandler is a user-supplied function that is executed for each streamed ledger received.
func (lr LedgerRequest) StreamLedgers(ctx context.Context, client *Client,
	handler LedgerHandler) (err error) {
	endpoint, err := lr.BuildURL()
	if err != nil {
		return errors.Wrap(err, "unable to build endpoint for ledger request")
	}

	url := fmt.Sprintf("%s%s", client.fixAuroraURL(), endpoint)
	return client.stream(ctx, url, func(data []byte) error {
		var ledger hProtocol.Ledger
		err = json.Unmarshal(data, &ledger)
		if err != nil {
			return errors.Wrap(err, "error unmarshaling data for ledger request")
		}
		handler(ledger)
		return nil
	})
}
