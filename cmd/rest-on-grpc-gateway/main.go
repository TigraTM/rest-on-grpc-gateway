package main

import (
	"context"
	logStd "log"
	"os/signal"
	"rest-on-grpc-gateway/modules/user"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	userModule := user.Service{
		Log: log,
	}

	err = userModule.RunServe(ctx)
	if err != nil {
		// nolint:gocritic // If the service fails to start, you need to call fatal
		log.Fatalf("userModule.RunServe: %v", err)
	}
}
