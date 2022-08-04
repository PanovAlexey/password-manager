package service

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type UserMetadataFromContextGetter struct {
}

func GetUserMetadataFromContextGetterService() UserMetadataFromContextGetter {
	return UserMetadataFromContextGetter{}
}

func (s UserMetadataFromContextGetter) GetUserIdFromContext(ctx context.Context) string {
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

func (s UserMetadataFromContextGetter) GetTraceIdFromContext(ctx context.Context) string {
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
