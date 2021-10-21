package main

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
)

// 定义一个常量名，是包名的FQDN，用于trace的名字，相同包的trace名字相同，用于追踪到包
const name = "goexamples/tracing/main"

func Fib(n uint) (uint, error) {
	if n <= 1 {
		return n, nil
	}
	if n > 93 {
		return 0, errors.New("unsupported than 90")
	}
	var n2, n1 uint = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n2 + n1, nil
}

func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r, l}
}

type App struct {
	r io.Reader
	l *log.Logger
}

func (a *App) Write(ctx context.Context, u uint) {
	var span trace.Span
	ctx, span = otel.Tracer(name).Start(ctx, "Write")
	defer span.End()

	n, err := func(ctx context.Context) (uint, error) {
		_, span := otel.Tracer(name).Start(ctx, "Fibonacci")
		defer span.End()
		n, err := Fib(u)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		return n, err
	}(ctx)

	if err != nil {
		a.l.Println("error occur", n, err)
	} else {
		a.l.Println("n is ", n)
	}
}

func (a *App) Poll(ctx context.Context) (uint, error) {
	a.l.Println("What Fibonacci number would you like to know: ")
	// span 就是一个调用了，名称可以是函数名/方法名
	_, span := otel.Tracer(name).Start(ctx, "Poll")
	defer span.End()
	var n uint
	_, err := fmt.Fscanf(a.r, "%d\n", &n)
	nStr := strconv.FormatUint(uint64(n), 10)
	// 给span设置一些其他属于，用于debug更多信息
	span.SetAttributes(attribute.String("request.n", nStr))
	return n, err
}

func (a *App) Run(ctx context.Context) error {
	for {
		// 创建一个父span
		var span trace.Span
		ctx, span = otel.Tracer(name).Start(ctx, "Run")
		// span的ctx继续传给下面的调用，因为是函数调用函数，所以继续将ctx传入
		n, err := a.Poll(ctx)
		if err != nil {
			return err
		}
		a.Write(ctx, n)
		span.End()
	}
}

// 创建一个Exporter
func newExporter(w io.Writer) (tracesdk.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

/*
资源是一种特殊类型的属性，适用于由流程生成的所有跨度。这些应该用于表示关于非临时进程的底层元数据——例如，进程的主机名或它的实例ID。
资源应该在跟踪程序提供程序初始化时分配给它，并且创建的方式类似于属性.可以看成是一个全局的属性
 */
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

func NewJaegerTracerProvider(url string, service, env string, id int64) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	hostname, _ := os.Hostname()
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", env),
			attribute.Int64("ID", id),
			attribute.String("hostname", hostname),
		)),
	)
	return tp, nil
}

func main() {
	l := log.New(os.Stdout, "", 0)
	/*
		// 创建exporter,追踪信息写入文件
		f, _ := os.Create("strace.txt")
		defer f.Close()
		exp, err := newExporter(f)
		if err != nil {
			log.Fatalln(err)
		}
		tp := tracesdk.NewTracerProvider(
			tracesdk.WithBatcher(exp),
			tracesdk.WithResource(newResource()),
		)
	*/
	// 如果jaegerURL不可达，程序会崩溃。要换成UDP才行
	jaegerUrl := "http://128.0.255.10:14268/api/traces"
	tp, err := NewJaegerTracerProvider(jaegerUrl, name, "dev", int64(10000))
	if err != nil {
		log.Fatalln("init jaeger provider failed ", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	app := NewApp(os.Stdin, l)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	errCh := make(chan error)
	go func() {
		parentContext := context.Background()
		errCh <- app.Run(parentContext)
	}()
	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
