package grpc_helper

import (
	"context"
)

func alwaysLoggingDeciderServer(_ context.Context, _ string, _ interface{}) bool { return true }

func alwaysLoggingDeciderClient(_ context.Context, _ string) bool { return true }
