// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

//go:generate mockgen -package=mocks -destination=../mocks/cognitoidentityprovider.go -source=cognitoidentityprovider.go CognitoidentityproviderClient
type CognitoidentityproviderClient interface {
	DescribeIdentityProvider(context.Context, *cognitoidentityprovider.DescribeIdentityProviderInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
	DescribeResourceServer(context.Context, *cognitoidentityprovider.DescribeResourceServerInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeResourceServerOutput, error)
	DescribeRiskConfiguration(context.Context, *cognitoidentityprovider.DescribeRiskConfigurationInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error)
	DescribeUserImportJob(context.Context, *cognitoidentityprovider.DescribeUserImportJobInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	DescribeUserPool(context.Context, *cognitoidentityprovider.DescribeUserPoolInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	DescribeUserPoolClient(context.Context, *cognitoidentityprovider.DescribeUserPoolClientInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	DescribeUserPoolDomain(context.Context, *cognitoidentityprovider.DescribeUserPoolDomainInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error)
	GetCSVHeader(context.Context, *cognitoidentityprovider.GetCSVHeaderInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	GetDevice(context.Context, *cognitoidentityprovider.GetDeviceInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetDeviceOutput, error)
	GetGroup(context.Context, *cognitoidentityprovider.GetGroupInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error)
	GetIdentityProviderByIdentifier(context.Context, *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error)
	GetSigningCertificate(context.Context, *cognitoidentityprovider.GetSigningCertificateInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetSigningCertificateOutput, error)
	GetUICustomization(context.Context, *cognitoidentityprovider.GetUICustomizationInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUICustomizationOutput, error)
	GetUser(context.Context, *cognitoidentityprovider.GetUserInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserOutput, error)
	GetUserAttributeVerificationCode(context.Context, *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	GetUserPoolMfaConfig(context.Context, *cognitoidentityprovider.GetUserPoolMfaConfigInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error)
	ListDevices(context.Context, *cognitoidentityprovider.ListDevicesInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListDevicesOutput, error)
	ListGroups(context.Context, *cognitoidentityprovider.ListGroupsInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListGroupsOutput, error)
	ListIdentityProviders(context.Context, *cognitoidentityprovider.ListIdentityProvidersInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
	ListResourceServers(context.Context, *cognitoidentityprovider.ListResourceServersInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListResourceServersOutput, error)
	ListTagsForResource(context.Context, *cognitoidentityprovider.ListTagsForResourceInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListTagsForResourceOutput, error)
	ListUserImportJobs(context.Context, *cognitoidentityprovider.ListUserImportJobsInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	ListUserPoolClients(context.Context, *cognitoidentityprovider.ListUserPoolClientsInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	ListUserPools(context.Context, *cognitoidentityprovider.ListUserPoolsInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	ListUsers(context.Context, *cognitoidentityprovider.ListUsersInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersInGroup(context.Context, *cognitoidentityprovider.ListUsersInGroupInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
}
