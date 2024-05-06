// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("go.opentelemetry.io/collector/internal/receiver/samplereceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("go.opentelemetry.io/collector/internal/receiver/samplereceiver")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	BatchSizeTriggerSend metric.Int64Counter
	RequestDuration      metric.Float64Histogram
}

// telemetryBuilderOption applies changes to default builder.
type telemetryBuilderOption func(*TelemetryBuilder)

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...telemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	var err, errs error
	meter := Meter(settings)
	builder.BatchSizeTriggerSend, err = meter.Int64Counter(
		"batch_size_trigger_send",
		metric.WithDescription("Number of times the batch was sent due to a size trigger"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.RequestDuration, err = meter.Float64Histogram(
		"request_duration",
		metric.WithDescription("Duration of request"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
