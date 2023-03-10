package jaeger

import (
	"io"
	"log"

	"github.com/mamochiro/beef/internals/config"
	"github.com/opentracing/opentracing-go"
	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
)

func NewJaeger(appConfig config.Configuration) io.Closer {
	log.Println()
	cfg, err := jaegerConf.FromEnv()
	log.Println("ERROR", err)
	panicIfErr(err)
	log.Println("IN JAGER")
	cfg.ServiceName = appConfig.AppName + "-" + appConfig.Env
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter = &jaegerConf.ReporterConfig{
		LogSpans:           true,
		LocalAgentHostPort: appConfig.JaegerAgentHost + ":" + appConfig.JaegerAgentPort,
	}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory
	jMetricsFactory = jMetricsFactory.Namespace(metrics.NSOptions{Name: appConfig.AppName + "-" + appConfig.Env, Tags: nil})

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
		jaegerConf.Observer(rpcmetrics.NewObserver(jMetricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	panicIfErr(err)
	opentracing.SetGlobalTracer(tracer)

	return closer
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
