package config

// Transport contains data for start and connected backend.
type Transport struct {
	Host             string `envconfig:"USER_HOST" default:"0.0.0.0"`
	GRPCExternalPort int    `envconfig:"USER_GRPC_EXTERNAL_PORT" default:"8080"`
	GRPCInternalPort int    `envconfig:"USER_GRPC_INTERNAL_PORT" default:"8081"`
	GRPCGWPort       int    `envconfig:"USER_GRPC_GW_PORT" default:"8082"`
}
