package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// ClientRequestFunc may take information from context and use it to construct
// metadata headers to be transported to the server. ClientRequestFuncs are
// executed after creating the request but prior to sending the gRPC request to
// the server.
type ClientRequestFunc func(context.Context, *metadata.MD) context.Context

// ServerRequestFunc may take information from the received metadata header and
// use it to place items in the request scoped context. ServerRequestFuncs are
// executed prior to invoking the endpoint.
type ServerRequestFunc func(context.Context, metadata.MD) context.Context

// ServerResponseFunc may take information from a request context and use it to
// manipulate the gRPC response metadata headers and trailers. ResponseFuncs are
// only executed in servers, after invoking the endpoint but prior to writing a
// response.
type ServerResponseFunc func(ctx context.Context, header *metadata.MD, trailer *metadata.MD) context.Context

// ClientResponseFunc may take information from a gRPC metadata header and/or
// trailer and make the responses available for consumption. ClientResponseFuncs
// are only executed in clients, after a request has been made, but prior to it
// being decoded.
type ClientResponseFunc func(ctx context.Context, header metadata.MD, trailer metadata.MD) context.Context
