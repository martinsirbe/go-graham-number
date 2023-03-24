package alphavantage_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/martinsirbe/go-graham-number/internal/alphavantage"
)

func TestGetGlobalQuote(t *testing.T) {
	for name, tc := range map[string]struct {
		overviewResponseStatusCode    int
		overviewResponseBody          string
		globalQuoteResponseStatusCode int
		globalQuoteResponseBody       string
		expectedStockDetails          *alphavantage.StockDetails
		assertErr                     assert.ErrorAssertionFunc
	}{
		"Success": {
			globalQuoteResponseStatusCode: http.StatusOK,
			globalQuoteResponseBody:       `{"Global Quote": {"05. price": "206.0100"}}`,
			overviewResponseStatusCode:    http.StatusOK,
			overviewResponseBody:          `{"BookValue": "48.09","EPS": "8.87"}`,
			expectedStockDetails: &alphavantage.StockDetails{
				EPS: 8.87, BVPS: 48.09, Price: 206.01},
			assertErr: assert.NoError,
		},

		"FailToGetGlobalQuote_HTTP_500": {
			globalQuoteResponseStatusCode: http.StatusInternalServerError,
			globalQuoteResponseBody:       `{"Global Quote": {"05. price": "206.0100"}}`,
			overviewResponseStatusCode:    http.StatusOK,
			overviewResponseBody:          `{"BookValue": "48.09","EPS": "8.87"}`,
			expectedStockDetails:          nil,
			assertErr:                     assert.Error,
		},
		"FailToGetGlobalQuote_MissingPrice": {
			overviewResponseStatusCode:    http.StatusOK,
			overviewResponseBody:          `{"BookValue": "48.09","EPS": "8.87"}`,
			globalQuoteResponseStatusCode: http.StatusOK,
			globalQuoteResponseBody:       `{"Global Quote": {"05. price": ""}}`,
			expectedStockDetails:          nil,
			assertErr:                     assert.Error,
		},
		"FailToGetGlobalQuote_BadResponseBody": {
			globalQuoteResponseStatusCode: http.StatusOK,
			globalQuoteResponseBody:       `}`,
			overviewResponseStatusCode:    http.StatusOK,
			overviewResponseBody:          `{"BookValue": "48.09","EPS": "8.87"}`,
			expectedStockDetails:          nil,
			assertErr:                     assert.Error,
		},

		"FailToGetOverview_HTTP_500": {
			globalQuoteResponseStatusCode: http.StatusOK,
			globalQuoteResponseBody:       `{"Global Quote": {"05. price": "206.0100"}}`,
			overviewResponseStatusCode:    http.StatusInternalServerError,
			overviewResponseBody:          `{"BookValue": "48.09","EPS": "8.87"}`,
			expectedStockDetails:          nil,
			assertErr:                     assert.Error,
		},
		"FailToGetOverview_MissingBVPS": {
			globalQuoteResponseStatusCode: http.StatusOK,
			globalQuoteResponseBody:       `{"Global Quote": {"05. price": "206.0100"}}`,
			overviewResponseStatusCode:    http.StatusOK,
			overviewResponseBody:          `{"BookValue": "","EPS": "8.87"}`,
			expectedStockDetails:          nil,
			assertErr:                     assert.Error,
		},
		"FailToGetOverview_MissingEPS": {
			globalQuoteResponseStatusCode: http.StatusOK,
			globalQuoteResponseBody:       `{"Global Quote": {"05. price": "206.0100"}}`,
			overviewResponseStatusCode:    http.StatusOK,
			overviewResponseBody:          `{"BookValue": "48.09","EPS": ""}`,
			assertErr:                     assert.Error,
		},
		"FailToGetOverview_BadResponseBody": {
			globalQuoteResponseStatusCode: http.StatusOK,
			globalQuoteResponseBody:       `{"Global Quote": {"05. price": "206.0100"}}`,
			overviewResponseStatusCode:    http.StatusOK,
			overviewResponseBody:          `}`,
			expectedStockDetails:          nil,
			assertErr:                     assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					require.Equal(t, r.URL.Query().Get("symbol"), "TEST")
					require.Equal(t, r.URL.Query().Get("apikey"), "TEST")

					switch r.URL.Query().Get("function") {
					case "OVERVIEW":
						w.WriteHeader(tc.overviewResponseStatusCode)
						_, err := w.Write([]byte(tc.overviewResponseBody))
						require.NoError(t, err)
					case "GLOBAL_QUOTE":
						w.WriteHeader(tc.globalQuoteResponseStatusCode)
						_, err := w.Write([]byte(tc.globalQuoteResponseBody))
						require.NoError(t, err)
					}
				}))
			defer server.Close()

			client := alphavantage.NewClient("TEST")
			client.SetURL(server.URL)

			sd, err := client.GetStockDetails("TEST")
			tc.assertErr(t, err)
			if err == nil {
				assert.Equal(t, tc.expectedStockDetails, sd)
			}
		})
	}
}
