// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

//go:generate mockgen -package=mocks -destination=../mocks/iot.go -source=iot.go IotClient
type IotClient interface {
	DescribeAccountAuditConfiguration(context.Context, *iot.DescribeAccountAuditConfigurationInput, ...func(*iot.Options)) (*iot.DescribeAccountAuditConfigurationOutput, error)
	DescribeAuditFinding(context.Context, *iot.DescribeAuditFindingInput, ...func(*iot.Options)) (*iot.DescribeAuditFindingOutput, error)
	DescribeAuditMitigationActionsTask(context.Context, *iot.DescribeAuditMitigationActionsTaskInput, ...func(*iot.Options)) (*iot.DescribeAuditMitigationActionsTaskOutput, error)
	DescribeAuditSuppression(context.Context, *iot.DescribeAuditSuppressionInput, ...func(*iot.Options)) (*iot.DescribeAuditSuppressionOutput, error)
	DescribeAuditTask(context.Context, *iot.DescribeAuditTaskInput, ...func(*iot.Options)) (*iot.DescribeAuditTaskOutput, error)
	DescribeAuthorizer(context.Context, *iot.DescribeAuthorizerInput, ...func(*iot.Options)) (*iot.DescribeAuthorizerOutput, error)
	DescribeBillingGroup(context.Context, *iot.DescribeBillingGroupInput, ...func(*iot.Options)) (*iot.DescribeBillingGroupOutput, error)
	DescribeCACertificate(context.Context, *iot.DescribeCACertificateInput, ...func(*iot.Options)) (*iot.DescribeCACertificateOutput, error)
	DescribeCertificate(context.Context, *iot.DescribeCertificateInput, ...func(*iot.Options)) (*iot.DescribeCertificateOutput, error)
	DescribeCustomMetric(context.Context, *iot.DescribeCustomMetricInput, ...func(*iot.Options)) (*iot.DescribeCustomMetricOutput, error)
	DescribeDefaultAuthorizer(context.Context, *iot.DescribeDefaultAuthorizerInput, ...func(*iot.Options)) (*iot.DescribeDefaultAuthorizerOutput, error)
	DescribeDetectMitigationActionsTask(context.Context, *iot.DescribeDetectMitigationActionsTaskInput, ...func(*iot.Options)) (*iot.DescribeDetectMitigationActionsTaskOutput, error)
	DescribeDimension(context.Context, *iot.DescribeDimensionInput, ...func(*iot.Options)) (*iot.DescribeDimensionOutput, error)
	DescribeDomainConfiguration(context.Context, *iot.DescribeDomainConfigurationInput, ...func(*iot.Options)) (*iot.DescribeDomainConfigurationOutput, error)
	DescribeEndpoint(context.Context, *iot.DescribeEndpointInput, ...func(*iot.Options)) (*iot.DescribeEndpointOutput, error)
	DescribeEventConfigurations(context.Context, *iot.DescribeEventConfigurationsInput, ...func(*iot.Options)) (*iot.DescribeEventConfigurationsOutput, error)
	DescribeFleetMetric(context.Context, *iot.DescribeFleetMetricInput, ...func(*iot.Options)) (*iot.DescribeFleetMetricOutput, error)
	DescribeIndex(context.Context, *iot.DescribeIndexInput, ...func(*iot.Options)) (*iot.DescribeIndexOutput, error)
	DescribeJob(context.Context, *iot.DescribeJobInput, ...func(*iot.Options)) (*iot.DescribeJobOutput, error)
	DescribeJobExecution(context.Context, *iot.DescribeJobExecutionInput, ...func(*iot.Options)) (*iot.DescribeJobExecutionOutput, error)
	DescribeJobTemplate(context.Context, *iot.DescribeJobTemplateInput, ...func(*iot.Options)) (*iot.DescribeJobTemplateOutput, error)
	DescribeManagedJobTemplate(context.Context, *iot.DescribeManagedJobTemplateInput, ...func(*iot.Options)) (*iot.DescribeManagedJobTemplateOutput, error)
	DescribeMitigationAction(context.Context, *iot.DescribeMitigationActionInput, ...func(*iot.Options)) (*iot.DescribeMitigationActionOutput, error)
	DescribeProvisioningTemplate(context.Context, *iot.DescribeProvisioningTemplateInput, ...func(*iot.Options)) (*iot.DescribeProvisioningTemplateOutput, error)
	DescribeProvisioningTemplateVersion(context.Context, *iot.DescribeProvisioningTemplateVersionInput, ...func(*iot.Options)) (*iot.DescribeProvisioningTemplateVersionOutput, error)
	DescribeRoleAlias(context.Context, *iot.DescribeRoleAliasInput, ...func(*iot.Options)) (*iot.DescribeRoleAliasOutput, error)
	DescribeScheduledAudit(context.Context, *iot.DescribeScheduledAuditInput, ...func(*iot.Options)) (*iot.DescribeScheduledAuditOutput, error)
	DescribeSecurityProfile(context.Context, *iot.DescribeSecurityProfileInput, ...func(*iot.Options)) (*iot.DescribeSecurityProfileOutput, error)
	DescribeStream(context.Context, *iot.DescribeStreamInput, ...func(*iot.Options)) (*iot.DescribeStreamOutput, error)
	DescribeThing(context.Context, *iot.DescribeThingInput, ...func(*iot.Options)) (*iot.DescribeThingOutput, error)
	DescribeThingGroup(context.Context, *iot.DescribeThingGroupInput, ...func(*iot.Options)) (*iot.DescribeThingGroupOutput, error)
	DescribeThingRegistrationTask(context.Context, *iot.DescribeThingRegistrationTaskInput, ...func(*iot.Options)) (*iot.DescribeThingRegistrationTaskOutput, error)
	DescribeThingType(context.Context, *iot.DescribeThingTypeInput, ...func(*iot.Options)) (*iot.DescribeThingTypeOutput, error)
	GetBehaviorModelTrainingSummaries(context.Context, *iot.GetBehaviorModelTrainingSummariesInput, ...func(*iot.Options)) (*iot.GetBehaviorModelTrainingSummariesOutput, error)
	GetBucketsAggregation(context.Context, *iot.GetBucketsAggregationInput, ...func(*iot.Options)) (*iot.GetBucketsAggregationOutput, error)
	GetCardinality(context.Context, *iot.GetCardinalityInput, ...func(*iot.Options)) (*iot.GetCardinalityOutput, error)
	GetEffectivePolicies(context.Context, *iot.GetEffectivePoliciesInput, ...func(*iot.Options)) (*iot.GetEffectivePoliciesOutput, error)
	GetIndexingConfiguration(context.Context, *iot.GetIndexingConfigurationInput, ...func(*iot.Options)) (*iot.GetIndexingConfigurationOutput, error)
	GetJobDocument(context.Context, *iot.GetJobDocumentInput, ...func(*iot.Options)) (*iot.GetJobDocumentOutput, error)
	GetLoggingOptions(context.Context, *iot.GetLoggingOptionsInput, ...func(*iot.Options)) (*iot.GetLoggingOptionsOutput, error)
	GetOTAUpdate(context.Context, *iot.GetOTAUpdateInput, ...func(*iot.Options)) (*iot.GetOTAUpdateOutput, error)
	GetPercentiles(context.Context, *iot.GetPercentilesInput, ...func(*iot.Options)) (*iot.GetPercentilesOutput, error)
	GetPolicy(context.Context, *iot.GetPolicyInput, ...func(*iot.Options)) (*iot.GetPolicyOutput, error)
	GetPolicyVersion(context.Context, *iot.GetPolicyVersionInput, ...func(*iot.Options)) (*iot.GetPolicyVersionOutput, error)
	GetRegistrationCode(context.Context, *iot.GetRegistrationCodeInput, ...func(*iot.Options)) (*iot.GetRegistrationCodeOutput, error)
	GetStatistics(context.Context, *iot.GetStatisticsInput, ...func(*iot.Options)) (*iot.GetStatisticsOutput, error)
	GetTopicRule(context.Context, *iot.GetTopicRuleInput, ...func(*iot.Options)) (*iot.GetTopicRuleOutput, error)
	GetTopicRuleDestination(context.Context, *iot.GetTopicRuleDestinationInput, ...func(*iot.Options)) (*iot.GetTopicRuleDestinationOutput, error)
	GetV2LoggingOptions(context.Context, *iot.GetV2LoggingOptionsInput, ...func(*iot.Options)) (*iot.GetV2LoggingOptionsOutput, error)
	ListActiveViolations(context.Context, *iot.ListActiveViolationsInput, ...func(*iot.Options)) (*iot.ListActiveViolationsOutput, error)
	ListAttachedPolicies(context.Context, *iot.ListAttachedPoliciesInput, ...func(*iot.Options)) (*iot.ListAttachedPoliciesOutput, error)
	ListAuditFindings(context.Context, *iot.ListAuditFindingsInput, ...func(*iot.Options)) (*iot.ListAuditFindingsOutput, error)
	ListAuditMitigationActionsExecutions(context.Context, *iot.ListAuditMitigationActionsExecutionsInput, ...func(*iot.Options)) (*iot.ListAuditMitigationActionsExecutionsOutput, error)
	ListAuditMitigationActionsTasks(context.Context, *iot.ListAuditMitigationActionsTasksInput, ...func(*iot.Options)) (*iot.ListAuditMitigationActionsTasksOutput, error)
	ListAuditSuppressions(context.Context, *iot.ListAuditSuppressionsInput, ...func(*iot.Options)) (*iot.ListAuditSuppressionsOutput, error)
	ListAuditTasks(context.Context, *iot.ListAuditTasksInput, ...func(*iot.Options)) (*iot.ListAuditTasksOutput, error)
	ListAuthorizers(context.Context, *iot.ListAuthorizersInput, ...func(*iot.Options)) (*iot.ListAuthorizersOutput, error)
	ListBillingGroups(context.Context, *iot.ListBillingGroupsInput, ...func(*iot.Options)) (*iot.ListBillingGroupsOutput, error)
	ListCACertificates(context.Context, *iot.ListCACertificatesInput, ...func(*iot.Options)) (*iot.ListCACertificatesOutput, error)
	ListCertificates(context.Context, *iot.ListCertificatesInput, ...func(*iot.Options)) (*iot.ListCertificatesOutput, error)
	ListCertificatesByCA(context.Context, *iot.ListCertificatesByCAInput, ...func(*iot.Options)) (*iot.ListCertificatesByCAOutput, error)
	ListCustomMetrics(context.Context, *iot.ListCustomMetricsInput, ...func(*iot.Options)) (*iot.ListCustomMetricsOutput, error)
	ListDetectMitigationActionsExecutions(context.Context, *iot.ListDetectMitigationActionsExecutionsInput, ...func(*iot.Options)) (*iot.ListDetectMitigationActionsExecutionsOutput, error)
	ListDetectMitigationActionsTasks(context.Context, *iot.ListDetectMitigationActionsTasksInput, ...func(*iot.Options)) (*iot.ListDetectMitigationActionsTasksOutput, error)
	ListDimensions(context.Context, *iot.ListDimensionsInput, ...func(*iot.Options)) (*iot.ListDimensionsOutput, error)
	ListDomainConfigurations(context.Context, *iot.ListDomainConfigurationsInput, ...func(*iot.Options)) (*iot.ListDomainConfigurationsOutput, error)
	ListFleetMetrics(context.Context, *iot.ListFleetMetricsInput, ...func(*iot.Options)) (*iot.ListFleetMetricsOutput, error)
	ListIndices(context.Context, *iot.ListIndicesInput, ...func(*iot.Options)) (*iot.ListIndicesOutput, error)
	ListJobExecutionsForJob(context.Context, *iot.ListJobExecutionsForJobInput, ...func(*iot.Options)) (*iot.ListJobExecutionsForJobOutput, error)
	ListJobExecutionsForThing(context.Context, *iot.ListJobExecutionsForThingInput, ...func(*iot.Options)) (*iot.ListJobExecutionsForThingOutput, error)
	ListJobTemplates(context.Context, *iot.ListJobTemplatesInput, ...func(*iot.Options)) (*iot.ListJobTemplatesOutput, error)
	ListJobs(context.Context, *iot.ListJobsInput, ...func(*iot.Options)) (*iot.ListJobsOutput, error)
	ListManagedJobTemplates(context.Context, *iot.ListManagedJobTemplatesInput, ...func(*iot.Options)) (*iot.ListManagedJobTemplatesOutput, error)
	ListMetricValues(context.Context, *iot.ListMetricValuesInput, ...func(*iot.Options)) (*iot.ListMetricValuesOutput, error)
	ListMitigationActions(context.Context, *iot.ListMitigationActionsInput, ...func(*iot.Options)) (*iot.ListMitigationActionsOutput, error)
	ListOTAUpdates(context.Context, *iot.ListOTAUpdatesInput, ...func(*iot.Options)) (*iot.ListOTAUpdatesOutput, error)
	ListOutgoingCertificates(context.Context, *iot.ListOutgoingCertificatesInput, ...func(*iot.Options)) (*iot.ListOutgoingCertificatesOutput, error)
	ListPolicies(context.Context, *iot.ListPoliciesInput, ...func(*iot.Options)) (*iot.ListPoliciesOutput, error)
	ListPolicyPrincipals(context.Context, *iot.ListPolicyPrincipalsInput, ...func(*iot.Options)) (*iot.ListPolicyPrincipalsOutput, error)
	ListPolicyVersions(context.Context, *iot.ListPolicyVersionsInput, ...func(*iot.Options)) (*iot.ListPolicyVersionsOutput, error)
	ListPrincipalPolicies(context.Context, *iot.ListPrincipalPoliciesInput, ...func(*iot.Options)) (*iot.ListPrincipalPoliciesOutput, error)
	ListPrincipalThings(context.Context, *iot.ListPrincipalThingsInput, ...func(*iot.Options)) (*iot.ListPrincipalThingsOutput, error)
	ListProvisioningTemplateVersions(context.Context, *iot.ListProvisioningTemplateVersionsInput, ...func(*iot.Options)) (*iot.ListProvisioningTemplateVersionsOutput, error)
	ListProvisioningTemplates(context.Context, *iot.ListProvisioningTemplatesInput, ...func(*iot.Options)) (*iot.ListProvisioningTemplatesOutput, error)
	ListRoleAliases(context.Context, *iot.ListRoleAliasesInput, ...func(*iot.Options)) (*iot.ListRoleAliasesOutput, error)
	ListScheduledAudits(context.Context, *iot.ListScheduledAuditsInput, ...func(*iot.Options)) (*iot.ListScheduledAuditsOutput, error)
	ListSecurityProfiles(context.Context, *iot.ListSecurityProfilesInput, ...func(*iot.Options)) (*iot.ListSecurityProfilesOutput, error)
	ListSecurityProfilesForTarget(context.Context, *iot.ListSecurityProfilesForTargetInput, ...func(*iot.Options)) (*iot.ListSecurityProfilesForTargetOutput, error)
	ListStreams(context.Context, *iot.ListStreamsInput, ...func(*iot.Options)) (*iot.ListStreamsOutput, error)
	ListTagsForResource(context.Context, *iot.ListTagsForResourceInput, ...func(*iot.Options)) (*iot.ListTagsForResourceOutput, error)
	ListTargetsForPolicy(context.Context, *iot.ListTargetsForPolicyInput, ...func(*iot.Options)) (*iot.ListTargetsForPolicyOutput, error)
	ListTargetsForSecurityProfile(context.Context, *iot.ListTargetsForSecurityProfileInput, ...func(*iot.Options)) (*iot.ListTargetsForSecurityProfileOutput, error)
	ListThingGroups(context.Context, *iot.ListThingGroupsInput, ...func(*iot.Options)) (*iot.ListThingGroupsOutput, error)
	ListThingGroupsForThing(context.Context, *iot.ListThingGroupsForThingInput, ...func(*iot.Options)) (*iot.ListThingGroupsForThingOutput, error)
	ListThingPrincipals(context.Context, *iot.ListThingPrincipalsInput, ...func(*iot.Options)) (*iot.ListThingPrincipalsOutput, error)
	ListThingRegistrationTaskReports(context.Context, *iot.ListThingRegistrationTaskReportsInput, ...func(*iot.Options)) (*iot.ListThingRegistrationTaskReportsOutput, error)
	ListThingRegistrationTasks(context.Context, *iot.ListThingRegistrationTasksInput, ...func(*iot.Options)) (*iot.ListThingRegistrationTasksOutput, error)
	ListThingTypes(context.Context, *iot.ListThingTypesInput, ...func(*iot.Options)) (*iot.ListThingTypesOutput, error)
	ListThings(context.Context, *iot.ListThingsInput, ...func(*iot.Options)) (*iot.ListThingsOutput, error)
	ListThingsInBillingGroup(context.Context, *iot.ListThingsInBillingGroupInput, ...func(*iot.Options)) (*iot.ListThingsInBillingGroupOutput, error)
	ListThingsInThingGroup(context.Context, *iot.ListThingsInThingGroupInput, ...func(*iot.Options)) (*iot.ListThingsInThingGroupOutput, error)
	ListTopicRuleDestinations(context.Context, *iot.ListTopicRuleDestinationsInput, ...func(*iot.Options)) (*iot.ListTopicRuleDestinationsOutput, error)
	ListTopicRules(context.Context, *iot.ListTopicRulesInput, ...func(*iot.Options)) (*iot.ListTopicRulesOutput, error)
	ListV2LoggingLevels(context.Context, *iot.ListV2LoggingLevelsInput, ...func(*iot.Options)) (*iot.ListV2LoggingLevelsOutput, error)
	ListViolationEvents(context.Context, *iot.ListViolationEventsInput, ...func(*iot.Options)) (*iot.ListViolationEventsOutput, error)
	SearchIndex(context.Context, *iot.SearchIndexInput, ...func(*iot.Options)) (*iot.SearchIndexOutput, error)
}
