package config

// Transport contains data for start and connected backend.
type Transport struct {
	Host       string `envconfig:"PAYMENT_HOST" required:"true" default:"0.0.0.0"`
	GRPCPort   int    `envconfig:"PAYMENT_GRPC_PORT" required:"true" default:"8090"`
	GRPCGWPort int    `envconfig:"PAYMENT_GRPC_GW_PORT" required:"true" default:"8091"`
}
