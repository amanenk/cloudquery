// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

//go:generate mockgen -package=mocks -destination=../mocks/kafka.go -source=kafka.go KafkaClient
type KafkaClient interface {
	DescribeCluster(context.Context, *kafka.DescribeClusterInput, ...func(*kafka.Options)) (*kafka.DescribeClusterOutput, error)
	DescribeClusterOperation(context.Context, *kafka.DescribeClusterOperationInput, ...func(*kafka.Options)) (*kafka.DescribeClusterOperationOutput, error)
	DescribeClusterV2(context.Context, *kafka.DescribeClusterV2Input, ...func(*kafka.Options)) (*kafka.DescribeClusterV2Output, error)
	DescribeConfiguration(context.Context, *kafka.DescribeConfigurationInput, ...func(*kafka.Options)) (*kafka.DescribeConfigurationOutput, error)
	DescribeConfigurationRevision(context.Context, *kafka.DescribeConfigurationRevisionInput, ...func(*kafka.Options)) (*kafka.DescribeConfigurationRevisionOutput, error)
	GetBootstrapBrokers(context.Context, *kafka.GetBootstrapBrokersInput, ...func(*kafka.Options)) (*kafka.GetBootstrapBrokersOutput, error)
	GetCompatibleKafkaVersions(context.Context, *kafka.GetCompatibleKafkaVersionsInput, ...func(*kafka.Options)) (*kafka.GetCompatibleKafkaVersionsOutput, error)
	ListClusterOperations(context.Context, *kafka.ListClusterOperationsInput, ...func(*kafka.Options)) (*kafka.ListClusterOperationsOutput, error)
	ListClusters(context.Context, *kafka.ListClustersInput, ...func(*kafka.Options)) (*kafka.ListClustersOutput, error)
	ListClustersV2(context.Context, *kafka.ListClustersV2Input, ...func(*kafka.Options)) (*kafka.ListClustersV2Output, error)
	ListConfigurationRevisions(context.Context, *kafka.ListConfigurationRevisionsInput, ...func(*kafka.Options)) (*kafka.ListConfigurationRevisionsOutput, error)
	ListConfigurations(context.Context, *kafka.ListConfigurationsInput, ...func(*kafka.Options)) (*kafka.ListConfigurationsOutput, error)
	ListKafkaVersions(context.Context, *kafka.ListKafkaVersionsInput, ...func(*kafka.Options)) (*kafka.ListKafkaVersionsOutput, error)
	ListNodes(context.Context, *kafka.ListNodesInput, ...func(*kafka.Options)) (*kafka.ListNodesOutput, error)
	ListScramSecrets(context.Context, *kafka.ListScramSecretsInput, ...func(*kafka.Options)) (*kafka.ListScramSecretsOutput, error)
	ListTagsForResource(context.Context, *kafka.ListTagsForResourceInput, ...func(*kafka.Options)) (*kafka.ListTagsForResourceOutput, error)
}
