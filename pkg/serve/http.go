package serve

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

// HTTP starts HTTP server on addr using handler logged as service.
// It runs until failed or ctx.Done.
func HTTP(log *zap.SugaredLogger, host string, port int, handler http.Handler) func(context.Context) error {
	return func(ctx context.Context) error {
		srv := &http.Server{
			Addr:    net.JoinHostPort(host, strconv.Itoa(port)),
			Handler: handler,
		}

		errc := make(chan error, 1)
		go func() { errc <- srv.ListenAndServe() }()
		log.Infof("http started: %s:%d", host, port)
		defer log.Info("shutdown")

		var err error
		select {
		case err = <-errc:
		case <-ctx.Done():
			// nolint:contextcheck // close right away
			err = srv.Shutdown(context.Background())
		}
		if err != nil {
			return fmt.Errorf("srv.ListenAndServe: %w", err)
		}

		return nil
	}
}
