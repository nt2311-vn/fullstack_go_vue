package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sc "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type ShutdownTracing func(ctx context.Context) error

func InitTracing(service string) (ShutdownTracing, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		return func(ctx context.Context) error { return nil }, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			sc.SchemaURL,
			sc.ServiceNameKey.String(service),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp.Shutdown, nil
}
