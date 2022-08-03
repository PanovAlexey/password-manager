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

func (s UserIdFromContextGetter) getUserIdFromContext(ctx context.Context) string {
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
