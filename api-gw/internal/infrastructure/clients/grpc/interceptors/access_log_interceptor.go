package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AccessLogInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	traceId := ctx.Value("trace-id") //@ToDo move key trace-id to service
	userId := "11"                   // @ToDo set userId value from context

	mdOut := metadata.Pairs(
		"trace-id", traceId.(string),
		"user-id", userId,
	)
	ctx = metadata.NewOutgoingContext(ctx, mdOut)

	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}
