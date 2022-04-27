package config

// Transport contains data for start and connected backend.
type Transport struct {
	Host       string `envconfig:"PAYMENT_HOST" default:"0.0.0.0"`
	GRPCPort   int    `envconfig:"PAYMENT_GRPC_PORT" default:"8090"`
	GRPCGWPort int    `envconfig:"PAYMENT_GRPC_GW_PORT" default:"8091"`
}
