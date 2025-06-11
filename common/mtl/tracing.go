package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// provider.WithExportEndpoint(":4317"),
		provider.WithExportEndpoint("jaeger-all-in-one:4317"),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)

	// defer p.Shutdown(context.Background())
	return p
}
