package mobileappreactnative

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func parseQuery(r *http.Request) url.Values {
	return r.URL.Query()
}

func parseJSON(r *http.Request) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	return data, err
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func stringify(obj interface{}) (string, error) {
	return json.Marshal(obj)
}

func extractError(r *http.Response) error {
	if r.StatusCode >= http.StatusBadRequest {
		return r.Status
	}
	return nil
}