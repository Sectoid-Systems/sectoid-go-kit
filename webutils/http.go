package webutils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpDoFunc func(req *http.Request) (*http.Response, error)

func DoAndParseJson[T any](do HttpDoFunc, req *http.Request, parsed T) error {
	res, err := do(req)
	if err != nil {
		return fmt.Errorf("error doing request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", res.Status)
	}

	bts, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(bts, parsed)
	if err != nil {
		return fmt.Errorf("error unmarshalling json body: %w", err)
	}

	return nil
}
