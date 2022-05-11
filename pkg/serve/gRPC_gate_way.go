package serve

import (
	"context"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GateWayConfig is config for building gRPC-Gateway proxy for WEB clients.
type GateWayConfig struct {
	FS             fs.FS
	GRPCServerPort int
	Namespace      string
	GRPCGWPattern  string                                                           // Pattern for http.ServeMux to serve grpc-gateway.
	OpenAPIPattern string                                                           // Pattern for http.ServeMux to serve swagger.json.
	Register       func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error // Register gRPC server.
}

// GRPCGateWay starts HTTP-proxy server for gRPC serer, for using gRPC endpoints from WEB.
func GRPCGateWay(log *zap.Logger, host string, port int, cfg GateWayConfig) func(context.Context) error {
	return func(ctx context.Context) error {
		conn, err := grpc.DialContext(ctx, net.JoinHostPort(host, strconv.Itoa(cfg.GRPCServerPort)), grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("grpc_helper.Dial: %w", err)
		}

		gw := runtime.NewServeMux()
		err = cfg.Register(ctx, gw, conn)
		if err != nil {
			return fmt.Errorf("cfg.Register: %w", err)
		}

		mux := http.NewServeMux()
		mux.Handle(cfg.GRPCGWPattern, corsAllowAll(gw))
		mux.Handle(cfg.OpenAPIPattern, http.StripPrefix(cfg.OpenAPIPattern, http.FileServer(http.FS(cfg.FS))))

		return HTTP(log, host, port, mux)(ctx)
	}
}

func corsAllowAll(next http.Handler) http.Handler {
	return cors.AllowAll().Handler(next)
}
