package metrics

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nicklanng/carpark/logging"
	"github.com/peterbourgon/g2s"
)

const (
	requestMetric  = "request"
	responseMetric = "response"
)

var client *statsdClient
var defaultTags []StatsdTag

// statsdClient wraps the g2s client, keeping a pointer to
// the underlying session.
type statsdClient struct {
	session *g2s.Statsd
}

// StatsdTag is a key/value pair used to send additional "tags"
// to StatsD.
type StatsdTag struct {
	Key   string
	Value string
}

// Initialize configures the singleton statsdClient
func Initialize(addr, prefix string) {
	if addr == "" {
		logging.Warn("No statsd endpoint provided")
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		logging.Fatal("Failed to get hostname")
	}

	session, err := g2s.DialWithPrefix("udp", addr, prefix)
	if err != nil {
		logging.Error(fmt.Sprintf("Failed to setup statsd metrics because: %s", err.Error()))
		return
	}

	client = new(statsdClient)
	client.session = session

	defaultTags = []StatsdTag{
		StatsdTag{Key: "host", Value: hostname},
	}
}

func Counter(bucket string, value int, tags ...StatsdTag) {
	if client == nil {
		return
	}
	client.session.Counter(1, buildMetricName(bucket, tags), value)
}

func Timing(bucket string, duration time.Duration, tags ...StatsdTag) {
	if client == nil {
		return
	}
	client.session.Timing(1, buildMetricName(bucket, tags), duration)
}

func Gauge(bucket string, value float64, tags ...StatsdTag) {
	if client == nil {
		return
	}
	client.session.Gauge(1, buildMetricName(bucket, tags), fmt.Sprintf("%f", value))
}

func buildMetricName(bucket string, tags []StatsdTag) string {
	tagStrings := []string{
		bucket,
	}

	for _, tag := range append(defaultTags, tags...) {
		tagStrings = append(tagStrings, fmt.Sprintf("%s=%s", tag.Key, tag.Value))
	}

	return strings.Join(tagStrings, ",")
}
