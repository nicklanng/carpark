package metrics

import (
	"net/http"
	"strconv"
)

// CountHTTPRequest increments the request count by 1.
func CountHTTPRequest(request *http.Request) {
	if client == nil {
		return
	}

	statsdTags := []StatsdTag{
		StatsdTag{Key: "method", Value: request.Method},
	}
	Counter(requestMetric, 1, statsdTags...)
}

// CountHTTPResponse increments the response count by 1.
func CountHTTPResponse(response *http.Response) {
	if client == nil {
		return
	}

	statsdTags := []StatsdTag{
		StatsdTag{Key: "statuscode", Value: strconv.Itoa(response.StatusCode)},
	}
	Counter(responseMetric, 1, statsdTags...)
}
