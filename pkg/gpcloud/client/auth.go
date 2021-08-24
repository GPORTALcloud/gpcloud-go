package client

import (
	"context"
	"fmt"
)

type GPCloudAuth struct {
	token string
}

func (a GPCloudAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", a.token),
	}, nil
}

func (a GPCloudAuth) RequireTransportSecurity() bool {
	return true
}
