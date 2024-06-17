package grpc

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockery --name grpc_client
type grpcClient struct {
	conn *grpc.ClientConn
}

// Close implements GrpcClient.
func (g *grpcClient) Close() error {
	return g.conn.Close()
}

// GetGrpcConnection implements GrpcClient.
func (g *grpcClient) GetGrpcConnection() *grpc.ClientConn {
	return g.conn
}

type GrpcClient interface {
	GetGrpcConnection() *grpc.ClientConn
	Close() error
}

func NewGrpcClient(config *GrpcConfig) (GrpcClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s%s", config.Host, config.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClient{conn: conn}, nil
}
