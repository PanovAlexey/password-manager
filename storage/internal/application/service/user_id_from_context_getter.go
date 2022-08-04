package service

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type UserIdFromContextGetter struct {
}

func GetUserIdFromContextGetterService() UserIdFromContextGetter {
	return UserIdFromContextGetter{}
}

func (s UserIdFromContextGetter) GetUserIdFromContext(ctx context.Context) string {
	userId := ""

	md, ok := metadata.FromIncomingContext(ctx)

	if ok {
		values := md.Get("user-id")

		if len(values) > 0 {
			userId = values[0]
		}
	}

	return userId
}

func (s UserIdFromContextGetter) GetTraceIdFromContext(ctx context.Context) string {
	traceId := ""

	md, ok := metadata.FromIncomingContext(ctx)

	if ok {
		values := md.Get("trace-id")

		if len(values) > 0 {
			traceId = values[0]
		}
	}

	return traceId
}
