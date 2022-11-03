// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

//go:generate mockgen -package=mocks -destination=../mocks/configservice.go -source=configservice.go ConfigserviceClient
type ConfigserviceClient interface {
	BatchGetAggregateResourceConfig(context.Context, *configservice.BatchGetAggregateResourceConfigInput, ...func(*configservice.Options)) (*configservice.BatchGetAggregateResourceConfigOutput, error)
	BatchGetResourceConfig(context.Context, *configservice.BatchGetResourceConfigInput, ...func(*configservice.Options)) (*configservice.BatchGetResourceConfigOutput, error)
	DescribeAggregateComplianceByConfigRules(context.Context, *configservice.DescribeAggregateComplianceByConfigRulesInput, ...func(*configservice.Options)) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error)
	DescribeAggregateComplianceByConformancePacks(context.Context, *configservice.DescribeAggregateComplianceByConformancePacksInput, ...func(*configservice.Options)) (*configservice.DescribeAggregateComplianceByConformancePacksOutput, error)
	DescribeAggregationAuthorizations(context.Context, *configservice.DescribeAggregationAuthorizationsInput, ...func(*configservice.Options)) (*configservice.DescribeAggregationAuthorizationsOutput, error)
	DescribeComplianceByConfigRule(context.Context, *configservice.DescribeComplianceByConfigRuleInput, ...func(*configservice.Options)) (*configservice.DescribeComplianceByConfigRuleOutput, error)
	DescribeComplianceByResource(context.Context, *configservice.DescribeComplianceByResourceInput, ...func(*configservice.Options)) (*configservice.DescribeComplianceByResourceOutput, error)
	DescribeConfigRuleEvaluationStatus(context.Context, *configservice.DescribeConfigRuleEvaluationStatusInput, ...func(*configservice.Options)) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
	DescribeConfigRules(context.Context, *configservice.DescribeConfigRulesInput, ...func(*configservice.Options)) (*configservice.DescribeConfigRulesOutput, error)
	DescribeConfigurationAggregatorSourcesStatus(context.Context, *configservice.DescribeConfigurationAggregatorSourcesStatusInput, ...func(*configservice.Options)) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error)
	DescribeConfigurationAggregators(context.Context, *configservice.DescribeConfigurationAggregatorsInput, ...func(*configservice.Options)) (*configservice.DescribeConfigurationAggregatorsOutput, error)
	DescribeConfigurationRecorderStatus(context.Context, *configservice.DescribeConfigurationRecorderStatusInput, ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	DescribeConfigurationRecorders(context.Context, *configservice.DescribeConfigurationRecordersInput, ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConformancePackCompliance(context.Context, *configservice.DescribeConformancePackComplianceInput, ...func(*configservice.Options)) (*configservice.DescribeConformancePackComplianceOutput, error)
	DescribeConformancePackStatus(context.Context, *configservice.DescribeConformancePackStatusInput, ...func(*configservice.Options)) (*configservice.DescribeConformancePackStatusOutput, error)
	DescribeConformancePacks(context.Context, *configservice.DescribeConformancePacksInput, ...func(*configservice.Options)) (*configservice.DescribeConformancePacksOutput, error)
	DescribeDeliveryChannelStatus(context.Context, *configservice.DescribeDeliveryChannelStatusInput, ...func(*configservice.Options)) (*configservice.DescribeDeliveryChannelStatusOutput, error)
	DescribeDeliveryChannels(context.Context, *configservice.DescribeDeliveryChannelsInput, ...func(*configservice.Options)) (*configservice.DescribeDeliveryChannelsOutput, error)
	DescribeOrganizationConfigRuleStatuses(context.Context, *configservice.DescribeOrganizationConfigRuleStatusesInput, ...func(*configservice.Options)) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error)
	DescribeOrganizationConfigRules(context.Context, *configservice.DescribeOrganizationConfigRulesInput, ...func(*configservice.Options)) (*configservice.DescribeOrganizationConfigRulesOutput, error)
	DescribeOrganizationConformancePackStatuses(context.Context, *configservice.DescribeOrganizationConformancePackStatusesInput, ...func(*configservice.Options)) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error)
	DescribeOrganizationConformancePacks(context.Context, *configservice.DescribeOrganizationConformancePacksInput, ...func(*configservice.Options)) (*configservice.DescribeOrganizationConformancePacksOutput, error)
	DescribePendingAggregationRequests(context.Context, *configservice.DescribePendingAggregationRequestsInput, ...func(*configservice.Options)) (*configservice.DescribePendingAggregationRequestsOutput, error)
	DescribeRemediationConfigurations(context.Context, *configservice.DescribeRemediationConfigurationsInput, ...func(*configservice.Options)) (*configservice.DescribeRemediationConfigurationsOutput, error)
	DescribeRemediationExceptions(context.Context, *configservice.DescribeRemediationExceptionsInput, ...func(*configservice.Options)) (*configservice.DescribeRemediationExceptionsOutput, error)
	DescribeRemediationExecutionStatus(context.Context, *configservice.DescribeRemediationExecutionStatusInput, ...func(*configservice.Options)) (*configservice.DescribeRemediationExecutionStatusOutput, error)
	DescribeRetentionConfigurations(context.Context, *configservice.DescribeRetentionConfigurationsInput, ...func(*configservice.Options)) (*configservice.DescribeRetentionConfigurationsOutput, error)
	GetAggregateComplianceDetailsByConfigRule(context.Context, *configservice.GetAggregateComplianceDetailsByConfigRuleInput, ...func(*configservice.Options)) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error)
	GetAggregateConfigRuleComplianceSummary(context.Context, *configservice.GetAggregateConfigRuleComplianceSummaryInput, ...func(*configservice.Options)) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error)
	GetAggregateConformancePackComplianceSummary(context.Context, *configservice.GetAggregateConformancePackComplianceSummaryInput, ...func(*configservice.Options)) (*configservice.GetAggregateConformancePackComplianceSummaryOutput, error)
	GetAggregateDiscoveredResourceCounts(context.Context, *configservice.GetAggregateDiscoveredResourceCountsInput, ...func(*configservice.Options)) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error)
	GetAggregateResourceConfig(context.Context, *configservice.GetAggregateResourceConfigInput, ...func(*configservice.Options)) (*configservice.GetAggregateResourceConfigOutput, error)
	GetComplianceDetailsByConfigRule(context.Context, *configservice.GetComplianceDetailsByConfigRuleInput, ...func(*configservice.Options)) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
	GetComplianceDetailsByResource(context.Context, *configservice.GetComplianceDetailsByResourceInput, ...func(*configservice.Options)) (*configservice.GetComplianceDetailsByResourceOutput, error)
	GetComplianceSummaryByConfigRule(context.Context, *configservice.GetComplianceSummaryByConfigRuleInput, ...func(*configservice.Options)) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
	GetComplianceSummaryByResourceType(context.Context, *configservice.GetComplianceSummaryByResourceTypeInput, ...func(*configservice.Options)) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
	GetConformancePackComplianceDetails(context.Context, *configservice.GetConformancePackComplianceDetailsInput, ...func(*configservice.Options)) (*configservice.GetConformancePackComplianceDetailsOutput, error)
	GetConformancePackComplianceSummary(context.Context, *configservice.GetConformancePackComplianceSummaryInput, ...func(*configservice.Options)) (*configservice.GetConformancePackComplianceSummaryOutput, error)
	GetCustomRulePolicy(context.Context, *configservice.GetCustomRulePolicyInput, ...func(*configservice.Options)) (*configservice.GetCustomRulePolicyOutput, error)
	GetDiscoveredResourceCounts(context.Context, *configservice.GetDiscoveredResourceCountsInput, ...func(*configservice.Options)) (*configservice.GetDiscoveredResourceCountsOutput, error)
	GetOrganizationConfigRuleDetailedStatus(context.Context, *configservice.GetOrganizationConfigRuleDetailedStatusInput, ...func(*configservice.Options)) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error)
	GetOrganizationConformancePackDetailedStatus(context.Context, *configservice.GetOrganizationConformancePackDetailedStatusInput, ...func(*configservice.Options)) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error)
	GetOrganizationCustomRulePolicy(context.Context, *configservice.GetOrganizationCustomRulePolicyInput, ...func(*configservice.Options)) (*configservice.GetOrganizationCustomRulePolicyOutput, error)
	GetResourceConfigHistory(context.Context, *configservice.GetResourceConfigHistoryInput, ...func(*configservice.Options)) (*configservice.GetResourceConfigHistoryOutput, error)
	GetStoredQuery(context.Context, *configservice.GetStoredQueryInput, ...func(*configservice.Options)) (*configservice.GetStoredQueryOutput, error)
	ListAggregateDiscoveredResources(context.Context, *configservice.ListAggregateDiscoveredResourcesInput, ...func(*configservice.Options)) (*configservice.ListAggregateDiscoveredResourcesOutput, error)
	ListConformancePackComplianceScores(context.Context, *configservice.ListConformancePackComplianceScoresInput, ...func(*configservice.Options)) (*configservice.ListConformancePackComplianceScoresOutput, error)
	ListDiscoveredResources(context.Context, *configservice.ListDiscoveredResourcesInput, ...func(*configservice.Options)) (*configservice.ListDiscoveredResourcesOutput, error)
	ListStoredQueries(context.Context, *configservice.ListStoredQueriesInput, ...func(*configservice.Options)) (*configservice.ListStoredQueriesOutput, error)
	ListTagsForResource(context.Context, *configservice.ListTagsForResourceInput, ...func(*configservice.Options)) (*configservice.ListTagsForResourceOutput, error)
}
