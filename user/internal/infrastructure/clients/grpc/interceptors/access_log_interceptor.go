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
	userId := ctx.Value("user-id")

	traceIdString := ""
	userIdString := ""

	if traceId != nil {
		traceIdString = traceId.(string)
	}

	if userId != nil {
		userIdString = userId.(string)
	}

	mdOut := metadata.Pairs(
		"trace-id", traceIdString,
		"user-id", userIdString,
	)
	ctx = metadata.NewOutgoingContext(ctx, mdOut)
	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}
