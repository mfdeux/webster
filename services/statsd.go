package services

import (
	"github.com/smira/go-statsd"
)

// NewStatsDService makes a new statsd client
func NewStatsDService() *statsd.Client {
	client := statsd.NewClient("localhost:8125",
		statsd.TagStyle(statsd.TagFormatInfluxDB),
		statsd.DefaultTags(statsd.StringTag("app", "billing")),
		statsd.MaxPacketSize(1400),
		statsd.MetricPrefix("web."))
	return client
}
