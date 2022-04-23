package main

import (
	"context"
	logStd "log"
	"os/signal"
	"rest-on-grpc-gateway/modules/user"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"

	"go.uber.org/zap"
)

type embeddedService interface {
	Init(ctx context.Context, log *zap.SugaredLogger) (err error)
	RunServe(ctx context.Context) error
}

var embeddedServices = []embeddedService{
	&user.Service{},
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

	for _, service := range embeddedServices {
		err = service.Init(ctxWithLog, log)
		if err != nil {
			// nolint:gocritic // fatal will exit.
			log.Fatalf("failed to init service: %s", err)
		}

		err = service.RunServe(ctxWithLog)
		if err != nil {
			log.Fatalf("failed to run service: %s", err)
		}
	}
}

func forceShutdown(ctx context.Context) {
	log := ctxzap.Extract(ctx)
	const shutdownDelay = 15 * time.Second

	<-ctx.Done()
	time.Sleep(shutdownDelay)

	log.Fatal("failed to graceful shutdown") //nolint:revive // By design.
}
