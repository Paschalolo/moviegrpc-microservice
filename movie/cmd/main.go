package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
	"movieexample.com/gen"
	"movieexample.com/movie/internal/controller/movie"
	metaGateway "movieexample.com/movie/internal/gateway/metadata/grpc"
	ratinggateway "movieexample.com/movie/internal/gateway/rating/grpc"
	grpchandler "movieexample.com/movie/internal/handler/grpc"
	"movieexample.com/pkg/discovery"
	"movieexample.com/pkg/discovery/consul"
	"movieexample.com/pkg/tracing"
)

const serviceName = "movie"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	var port int
	// change this to base.yaml before building docekr file
	f, err := os.Open("configs/base.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var cfg serviceConfig
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	flag.IntVar(&port, "port", cfg.APIConfig.Port, "Api handler port")
	flag.Parse()
	log.Printf("starting the movie service on port: %d ", port)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tp, err := tracing.NewJaegarProvider(ctx, cfg.Jaeger.URL, serviceName)
	if err != nil {
		logger.Fatal("Failed to initialize Jaeger provider", zap.Error(err))
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			logger.Fatal("Failed to shut down Jaeger prodiver", zap.Error(err))
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("failed to report healthy state :" + err.Error())
			}
			time.Sleep(time.Second * 1)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	metadataGateway := metaGateway.New(registry)
	ratingDatagateway := ratinggateway.New(registry)
	svc := movie.New(ratingDatagateway, metadataGateway)
	h := grpchandler.New(svc)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen on port :%d", port)
	}

	srv := grpc.NewServer()
	gen.RegisterMovieServiceServer(srv, h)
	reflection.Register(srv)
	srv.Serve(lis)
	// h := httphandler.New(svc)
	// http.HandleFunc("/movie", h.GetMovieDetails)
	// if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
	// 	panic(err)
	// }
}
