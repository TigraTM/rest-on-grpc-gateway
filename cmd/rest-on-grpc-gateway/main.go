package main

import (
	"context"
	"fmt"
	"os/signal"
	"rest-on-grpc-gateway/modules/payment"
	"rest-on-grpc-gateway/modules/user"
	"rest-on-grpc-gateway/pkg/serve"
	"syscall"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

type embeddedService interface {
	Name() string
	Init(ctx context.Context, log *zap.Logger) (err error)
	RunServe(ctx context.Context) error
}

var embeddedServices = []embeddedService{
	&user.Service{},
	&payment.Service{},
}

func main() {
	logCfg := zap.NewProductionConfig()
	logCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logCfg.EncoderConfig.TimeKey = "timestamp"
	logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	log, err := logCfg.Build(
		zap.WithClock(zapcore.DefaultClock),
		zap.AddCaller(),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := log.Sync()
		if err != nil {
			panic(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go forceShutdown(ctx, log)

	if err = runServices(ctx, log); err != nil {
		// nolint:gocritic // ...
		log.Fatal("failed run service: %s", zap.Error(err))
	}
}

func runServices(ctx context.Context, log *zap.Logger) (err error) {
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

func forceShutdown(ctx context.Context, log *zap.Logger) {
	const shutdownDelay = 15 * time.Second

	<-ctx.Done()
	time.Sleep(shutdownDelay)

	log.Fatal("failed to graceful shutdown") //nolint:revive // By design.
}
