package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/controllers"
	grpcserver "github.com/rruzicic/globetrotter/bnb/reservation-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const serviceName = "reservation-service"

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

	res := r.Group("/reservation")
	res.POST("/", controllers.CreateReservation)
	res.GET("/:id", controllers.GetReservationById)
	res.GET("/accommodation/:id", controllers.GetReservationsByAccommodationId)
	res.GET("/user/:id", controllers.GetReservationsByUserId)
	res.DELETE("/:id", controllers.DeleteReservation)
	res.POST("/approve/:id", controllers.ApproveReservation)
	res.POST("/reject/:id", controllers.RejectReservation)
	res.POST("/accommodation/:acc_id/reservation/:res_id", controllers.AddReservationToAccommodation)
	res.GET("/test/:msg", controllers.TestConnection)

	r.Run(":8080")
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
