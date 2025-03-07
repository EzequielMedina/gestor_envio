package api

import (
	"io"
	"net/http"

	"github.com/sony/gobreaker/v2"
)

var cb *gobreaker.CircuitBreaker[[]byte]

func init() {
	var st gobreaker.Settings
	st.Name = "api get"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests > 3 && failureRatio > 0.6
	}

	cb = gobreaker.NewCircuitBreaker[[]byte](st)
}

func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() ([]byte, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body, nil
}
