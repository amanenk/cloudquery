// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

//go:generate mockgen -package=mocks -destination=../mocks/directconnect.go -source=directconnect.go DirectconnectClient
type DirectconnectClient interface {
	DescribeConnectionLoa(context.Context, *directconnect.DescribeConnectionLoaInput, ...func(*directconnect.Options)) (*directconnect.DescribeConnectionLoaOutput, error)
	DescribeConnections(context.Context, *directconnect.DescribeConnectionsInput, ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOutput, error)
	DescribeConnectionsOnInterconnect(context.Context, *directconnect.DescribeConnectionsOnInterconnectInput, ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOnInterconnectOutput, error)
	DescribeCustomerMetadata(context.Context, *directconnect.DescribeCustomerMetadataInput, ...func(*directconnect.Options)) (*directconnect.DescribeCustomerMetadataOutput, error)
	DescribeDirectConnectGatewayAssociationProposals(context.Context, *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput, ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error)
	DescribeDirectConnectGatewayAssociations(context.Context, *directconnect.DescribeDirectConnectGatewayAssociationsInput, ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error)
	DescribeDirectConnectGatewayAttachments(context.Context, *directconnect.DescribeDirectConnectGatewayAttachmentsInput, ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error)
	DescribeDirectConnectGateways(context.Context, *directconnect.DescribeDirectConnectGatewaysInput, ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error)
	DescribeHostedConnections(context.Context, *directconnect.DescribeHostedConnectionsInput, ...func(*directconnect.Options)) (*directconnect.DescribeHostedConnectionsOutput, error)
	DescribeInterconnectLoa(context.Context, *directconnect.DescribeInterconnectLoaInput, ...func(*directconnect.Options)) (*directconnect.DescribeInterconnectLoaOutput, error)
	DescribeInterconnects(context.Context, *directconnect.DescribeInterconnectsInput, ...func(*directconnect.Options)) (*directconnect.DescribeInterconnectsOutput, error)
	DescribeLags(context.Context, *directconnect.DescribeLagsInput, ...func(*directconnect.Options)) (*directconnect.DescribeLagsOutput, error)
	DescribeLoa(context.Context, *directconnect.DescribeLoaInput, ...func(*directconnect.Options)) (*directconnect.DescribeLoaOutput, error)
	DescribeLocations(context.Context, *directconnect.DescribeLocationsInput, ...func(*directconnect.Options)) (*directconnect.DescribeLocationsOutput, error)
	DescribeRouterConfiguration(context.Context, *directconnect.DescribeRouterConfigurationInput, ...func(*directconnect.Options)) (*directconnect.DescribeRouterConfigurationOutput, error)
	DescribeTags(context.Context, *directconnect.DescribeTagsInput, ...func(*directconnect.Options)) (*directconnect.DescribeTagsOutput, error)
	DescribeVirtualGateways(context.Context, *directconnect.DescribeVirtualGatewaysInput, ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error)
	DescribeVirtualInterfaces(context.Context, *directconnect.DescribeVirtualInterfacesInput, ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error)
	ListVirtualInterfaceTestHistory(context.Context, *directconnect.ListVirtualInterfaceTestHistoryInput, ...func(*directconnect.Options)) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error)
}
