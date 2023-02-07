package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/GPORTALcloud/gpcloud-go/pkg/gpcloud/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const DefaultEndpoint = "legacy.g-portal.cloud:443"

type Client struct {
	grpcClient *grpc.ClientConn
}

// UserClient Returns the ptypes.UserAPIClient for authenticated requests
func (c *Client) UserClient() ptypes.UserAPIClient {
	client := ptypes.NewUserAPIClient(c.grpcClient)

	return client
}

// PublicClient Returns the ptypes.PublicAPIClient for unauthenticated requests
func (c *Client) PublicClient() ptypes.PublicAPIClient {
	client := ptypes.NewPublicAPIClient(c.grpcClient)

	return client
}

// NewClient Returns a new GRPC client
func NewClient(accessToken string, options ...grpc.DialOption) (*Client, error) {
	cl := &Client{}

	// Certificate pinning
	options = append(options, grpc.WithTransportCredentials(credentials.NewTLS(getTlsOptions())))

	// User Agent
	options = append(options, grpc.WithUserAgent(fmt.Sprintf("GPCloud Golang Client [%s]", Version)))

	// Access Token
	if accessToken != "" {
		options = append(options, grpc.WithPerRPCCredentials(GPCloudAuth{
			token: accessToken,
		}))
	}

	clientConn, err := grpc.Dial(DefaultEndpoint, options...)
	if err != nil {
		return nil, err
	}

	cl.grpcClient = clientConn
	return cl, nil
}

func getTlsOptions() *tls.Config {
	rootCA := x509.NewCertPool()
	rootCA.AppendCertsFromPEM(getCertificate())

	return &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    rootCA,
	}
}
