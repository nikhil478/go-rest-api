package metrics

import (
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

var (
	Meter            metric.Meter
	HTTPRequestCount metric.Int64Counter
	HTTPRequestDuration metric.Float64Histogram
	
)

func InitOTelMetrics(serviceName string) (*prometheus.Exporter, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create Prometheus exporter: %w", err)
	}
	res, err := resource.Merge(resource.Default(), resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(serviceName)))
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}
	provider := metricsdk.NewMeterProvider(metricsdk.WithReader(exporter), metricsdk.WithResource(res))
	otel.SetMeterProvider(provider)
	Meter = otel.Meter(serviceName)
	HTTPRequestCount, err = Meter.Int64Counter("http_requests_total")
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTPRequestCount counter: %w", err)
	}
	HTTPRequestDuration, err = Meter.Float64Histogram("http_request_duration_seconds")
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTPRequestDuration : %w", err)
	}
	return exporter, nil
}
