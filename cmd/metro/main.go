package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/razorpay/metro/internal/boot"
	"github.com/razorpay/metro/pkg/logger"
	"github.com/razorpay/metro/service/producer"
	"google.golang.org/grpc"
)

func main() {
	// Initialize context
	ctx, cancel := context.WithCancel(boot.NewContext(context.Background()))
	defer cancel()

	// Init app dependencies
	env := boot.GetEnv()
	err := boot.InitProducer(ctx, env)
	if err != nil {
		log.Fatalf("failed to init metro: %v", err)
	}

	// Shutdown tracer
	defer boot.Closer.Close()

	// initialize producer service and start it
	producerService := producer.NewService(ctx)
	producerService.Start()

	// Handle SIGINT & SIGTERM - Shutdown gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	// Block until signal is received.
	<-c

	logger.Ctx(ctx).Infow("stopping metro")
	// stop producer service
	err = producerService.Stop()
	if err != nil {
		panic(err)
	}

	logger.Ctx(ctx).Infow("stopped metro")
}

func getInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{}
}
