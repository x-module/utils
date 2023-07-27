/**
 * Created by Goland.
 * @file   jeager.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/7/20 15:01
 * @desc   jeager.go
 */

package tracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

var Tracer opentracing.Tracer

func NewJaegerTracer(serviceName string, jaegerHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  jaegerHostPort,
		},
	}
	var closer io.Closer
	var err error
	Tracer, closer, err = cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(Tracer)
	return Tracer, closer, err
}
