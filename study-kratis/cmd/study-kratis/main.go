package main

import (
	"flag"
	"os"

	"study-kratis/internal/conf"
	"study-kratis/internal/pkg/zap"

	"github.com/go-kratos/kratos/contrib/config/apollo/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/internal/endpoint"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/genproto/googleapis/cloud/aiplatform/logging"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func setLoggerInit(){
	logger := log.With(zap.Logger(),
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

func setConfig(){
	c := config.New(
		config.WithSource(
			appllo.NewSource(
				apollo.WithAppID(appId),
				apollo.WithCluster(cluser),
				apollo.WithEndpoint(endpoint),
				apollo.WithNamespace(namespace),
				apollo.WithEnableBackup(),
				apollo.WithSecret(secret),
			)
		)
	)
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}


//引入链路追踪
func setTracerProvider(url string)error{
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	//配置链路追踪
	tp := trace.NewTracerProvider(
		// 采样率设置(100%)
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(1.0))),
		// 采样地址
		trace.WithBatcher(exp),
		// 程序信息设置
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.DBSystemMySQL,
			semconv.ServiceNameKey.String(bc.Dbs.ProjectName),

			attribute.String("ID", id),
			attribute.String("Env", bc.Dbs.Env),
			attribute.String("Version", bc.Dbs.Version),
		)),
	)
	//注册全局链路追踪
	otel.SetTracerProvider(tp)
	return nil
}
