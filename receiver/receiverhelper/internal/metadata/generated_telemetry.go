// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("go.opentelemetry.io/collector/receiver/receiverhelper")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("go.opentelemetry.io/collector/receiver/receiverhelper")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	ReceiverAcceptedLogRecords   metric.Int64Counter
	ReceiverAcceptedMetricPoints metric.Int64Counter
	ReceiverAcceptedSpans        metric.Int64Counter
	ReceiverRefusedLogRecords    metric.Int64Counter
	ReceiverRefusedMetricPoints  metric.Int64Counter
	ReceiverRefusedSpans         metric.Int64Counter
}

// telemetryBuilderOption applies changes to default builder.
type telemetryBuilderOption func(*TelemetryBuilder)

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...telemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	for _, op := range options {
		op(&builder)
	}
	var err, errs error
	meter := Meter(settings)
	builder.ReceiverAcceptedLogRecords, err = meter.Int64Counter(
		"receiver_accepted_log_records",
		metric.WithDescription("Number of log records successfully pushed into the pipeline."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ReceiverAcceptedMetricPoints, err = meter.Int64Counter(
		"receiver_accepted_metric_points",
		metric.WithDescription("Number of metric points successfully pushed into the pipeline."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ReceiverAcceptedSpans, err = meter.Int64Counter(
		"receiver_accepted_spans",
		metric.WithDescription("Number of spans successfully pushed into the pipeline."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ReceiverRefusedLogRecords, err = meter.Int64Counter(
		"receiver_refused_log_records",
		metric.WithDescription("Number of log records that could not be pushed into the pipeline."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ReceiverRefusedMetricPoints, err = meter.Int64Counter(
		"receiver_refused_metric_points",
		metric.WithDescription("Number of metric points that could not be pushed into the pipeline."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ReceiverRefusedSpans, err = meter.Int64Counter(
		"receiver_refused_spans",
		metric.WithDescription("Number of spans that could not be pushed into the pipeline."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
