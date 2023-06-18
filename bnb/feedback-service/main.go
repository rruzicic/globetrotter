package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/controllers"
	grpc_server "github.com/rruzicic/globetrotter/bnb/feedback-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/repos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func main() {
	repos.Connect()
	go ginSetup()
	grpc_server.InitServer()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	//TODO implement controllers
	Feedback := r.Group("/feedback")

	HostFeedback := Feedback.Group("/HostFeedback")
	HostFeedback.POST("/", controllers.CreateHostReview)
	HostFeedback.GET("/id/:id", controllers.GetHostReviewById)
	HostFeedback.GET("/user/:user_id", controllers.GetHostReviewsByUserId)
	HostFeedback.GET("/host/:host_id", controllers.GetHostReviewsByHostId)
	HostFeedback.DELETE("/:id", controllers.DeleteHostReview)
	HostFeedback.PUT("/", controllers.UpdateHostReview)

	AccommodationFeedback := Feedback.Group("AccommodationFeedback")
	AccommodationFeedback.POST("/", controllers.CreateAccommodationReview)
	AccommodationFeedback.GET("/id/:id", controllers.GetAccommodationReviewById)
	AccommodationFeedback.GET("/user/:user_id", controllers.GetAccommodationReviewsByUserId)
	AccommodationFeedback.GET("/accommodation/:accommodation_id", controllers.GetAccommodationReviewsByAccommodationId)
	AccommodationFeedback.DELETE("/:id", controllers.DeleteAccommodationReview)
	AccommodationFeedback.PUT("/", controllers.UpdateAccommodationReview)

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
		resource.WithAttributes(semconv.ServiceNameKey.String("feedback-service"),
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
