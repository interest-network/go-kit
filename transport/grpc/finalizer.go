package grpc

import "context"

// ServerFinalizerFunc can be used to perform work at the end of an gRPC
// request, after the response has been written to the client.
type ServerFinalizerFunc func(ctx context.Context, err error)

// ClientFinalizerFunc can be used to perform work at the end of a client gRPC
// request, after the response is returned. The principal
// intended use is for error logging. Additional response parameters are
// provided in the context under keys with the ContextKeyResponse prefix.
// Note: err may be nil. There maybe also no additional response parameters depending on
// when an error occurs.
type ClientFinalizerFunc func(ctx context.Context, err error)
