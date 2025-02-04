// Code generated by codegen; DO NOT EDIT.

package logging

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"

	"cloud.google.com/go/logging/apiv2"

	pb "google.golang.org/genproto/googleapis/logging/v2"

	"google.golang.org/api/option"
)

func createSinks() (*client.Services, error) {
	fakeServer := &fakeSinksServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterConfigServiceV2Server(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := logging.NewConfigClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		LoggingConfigClient: svc,
	}, nil
}

type fakeSinksServer struct {
	pb.UnimplementedConfigServiceV2Server
}

func (f *fakeSinksServer) ListSinks(context.Context, *pb.ListSinksRequest) (*pb.ListSinksResponse, error) {
	resp := pb.ListSinksResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestSinks(t *testing.T) {
	client.MockTestHelper(t, Sinks(), createSinks, client.TestOptions{})
}
