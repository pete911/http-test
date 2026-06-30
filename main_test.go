package main

import (
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.String(a.Key, a.Value.Time().Format("15:04:05.000000000"))
			}
			return a
		},
	})))
}

func TestRunFullTrace(t *testing.T) {
	client := NewClient(DefaultTransport(), Trace)
	res, err := client.Do(NewTestGetRequest(t, "https://google.com"))
	require.NoError(t, err)
	require.NoError(t, DiscardResponse(res))
}

func TestRunLoad(t *testing.T) {
	client := NewClient(DefaultTransport(), ConnTrace)
	for i := 0; i < 100; i++ {
		res, err := client.Do(NewTestGetRequest(t, "https://google.com"))
		require.NoError(t, err)
		require.NoError(t, DiscardResponse(res))
	}
}

func NewTestGetRequest(t *testing.T, url string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	return req
}
