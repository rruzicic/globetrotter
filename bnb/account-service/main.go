package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/controllers"
	grpcserver "github.com/rruzicic/globetrotter/bnb/account-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/account-service/jwt"
	"github.com/rruzicic/globetrotter/bnb/account-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/account-service/repos"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const serviceName = "account-service"

func main() {
	ctx := context.Background()
	shutdown, err := InitProviderWithJaegerExporter(ctx)
	if err != nil {
		log.Fatalf("%s: %v", "Failed to initialize opentelemetry provider", err)
	}
	defer shutdown(ctx)

	repos.Connect()
	go ginSetup()
	grpcserver.InitServer()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(otelgin.Middleware(serviceName))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	public := r.Group("/user")

	protected := r.Group("/user")
	protected.Use(jwt.AnyUserAuthMiddleware())

	public.GET("/health", controllers.HealthCheck)
	public.GET("/ping", controllers.Ping)
	public.GET("/all", controllers.GetAll)
	public.GET("/id/:id", controllers.GetById)
	public.GET("/email/:email", controllers.GetByEmail)
	public.GET("/api-key", controllers.AddAPIKeyToUser)

	public.POST("/register/host", controllers.RegisterHost)
	public.POST("/register/guest", controllers.RegisterGuest)
	protected.POST("/update", controllers.UpdateUser)
	public.POST("/login", controllers.Login)

	protected.DELETE("/delete/:id", controllers.DeleteUser)
	r.Run(":8080")
	log.Println("HTTP server running on port 8080")
}

func InitProviderWithJaegerExporter(ctx context.Context) (func(context.Context) error, error) {
	exp, err := exporterToJaeger()
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	tp := trace.NewTracerProvider(
		trace.WithSampler(getSampler()),
		trace.WithBatcher(exp),
		trace.WithResource(newResource(ctx)),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown, nil
}

func exporterToJaeger() (*jaeger.Exporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://otel_collector:14278/api/traces")))
}

func newResource(ctx context.Context) *resource.Resource {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName),
			attribute.String("environment", "dev"),
		),
	)
	if err != nil {
		log.Fatalf("%s: %v", "Failed to create resource", err)
	}
	return res
}

func getSampler() trace.Sampler {
	return trace.AlwaysSample()
}
