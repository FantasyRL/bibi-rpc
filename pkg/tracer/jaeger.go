package tracer

import (
	"bibi/config"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitJaegerTracer(serviceName string) /*(opentracing.Tracer, io.Closer) */ {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: false,
			// 按实际情况替换你的 ip
			CollectorEndpoint: config.Jaeger.Addr,
		},
	}

	tracer, _, err := cfg.NewTracer(
		jaegercfg.Logger(jaeger.StdLogger),
		jaegercfg.ZipkinSharedRPCSpan(true),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	//return tracer, closer
}
