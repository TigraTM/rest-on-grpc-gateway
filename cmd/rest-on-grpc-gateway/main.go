package main

import (
	"context"
	"fmt"
	logStd "log"
	"os/signal"
	"rest-on-grpc-gateway/modules/payment"
	"rest-on-grpc-gateway/modules/user"
	"rest-on-grpc-gateway/pkg/serve"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"

	"go.uber.org/zap"
)

type embeddedService interface {
	Name() string
	Init(ctx context.Context, log *zap.SugaredLogger) (err error)
	RunServe(ctx context.Context) error
}

var embeddedServices = []embeddedService{
	&user.Service{},
	&payment.Service{},
}

func main() {
	zap, err := zap.NewDevelopment()
	if err != nil {
		logStd.Fatalf("couldn't init logger: %+v \n", err)
	}
	defer func() {
		err := zap.Sync()
		if err != nil {
			logStd.Fatalf("err: %v", err)
		}
	}()
	log := zap.Sugar()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ctxWithLog := ctxzap.ToContext(ctx, log.Desugar())
	go forceShutdown(ctxWithLog)

	if err = runServices(ctxWithLog, log); err != nil {
		// nolint:gocritic // ...
		log.Fatalf("failed run service: %s", err)
	}
}

func runServices(ctx context.Context, log *zap.SugaredLogger) (err error) {
	services := make([]func(context.Context) error, len(embeddedServices))
	for i := range embeddedServices {
		log := log.Named(embeddedServices[i].Name())

		err = embeddedServices[i].Init(ctx, log)
		if err != nil {
			return fmt.Errorf("failed service - %s init: %w", embeddedServices[i].Name(), err)
		}

		runServe := embeddedServices[i].RunServe
		services[i] = func(ctxShutdown context.Context) error {
			return runServe(ctxShutdown)
		}
	}

	return serve.Start(ctx, services...)
}

func forceShutdown(ctx context.Context) {
	log := ctxzap.Extract(ctx)
	const shutdownDelay = 15 * time.Second

	<-ctx.Done()
	time.Sleep(shutdownDelay)

	log.Fatal("failed to graceful shutdown") //nolint:revive // By design.
}
