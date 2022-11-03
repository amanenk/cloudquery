// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

//go:generate mockgen -package=mocks -destination=../mocks/codepipeline.go -source=codepipeline.go CodepipelineClient
type CodepipelineClient interface {
	GetActionType(context.Context, *codepipeline.GetActionTypeInput, ...func(*codepipeline.Options)) (*codepipeline.GetActionTypeOutput, error)
	GetJobDetails(context.Context, *codepipeline.GetJobDetailsInput, ...func(*codepipeline.Options)) (*codepipeline.GetJobDetailsOutput, error)
	GetPipeline(context.Context, *codepipeline.GetPipelineInput, ...func(*codepipeline.Options)) (*codepipeline.GetPipelineOutput, error)
	GetPipelineExecution(context.Context, *codepipeline.GetPipelineExecutionInput, ...func(*codepipeline.Options)) (*codepipeline.GetPipelineExecutionOutput, error)
	GetPipelineState(context.Context, *codepipeline.GetPipelineStateInput, ...func(*codepipeline.Options)) (*codepipeline.GetPipelineStateOutput, error)
	GetThirdPartyJobDetails(context.Context, *codepipeline.GetThirdPartyJobDetailsInput, ...func(*codepipeline.Options)) (*codepipeline.GetThirdPartyJobDetailsOutput, error)
	ListActionExecutions(context.Context, *codepipeline.ListActionExecutionsInput, ...func(*codepipeline.Options)) (*codepipeline.ListActionExecutionsOutput, error)
	ListActionTypes(context.Context, *codepipeline.ListActionTypesInput, ...func(*codepipeline.Options)) (*codepipeline.ListActionTypesOutput, error)
	ListPipelineExecutions(context.Context, *codepipeline.ListPipelineExecutionsInput, ...func(*codepipeline.Options)) (*codepipeline.ListPipelineExecutionsOutput, error)
	ListPipelines(context.Context, *codepipeline.ListPipelinesInput, ...func(*codepipeline.Options)) (*codepipeline.ListPipelinesOutput, error)
	ListTagsForResource(context.Context, *codepipeline.ListTagsForResourceInput, ...func(*codepipeline.Options)) (*codepipeline.ListTagsForResourceOutput, error)
	ListWebhooks(context.Context, *codepipeline.ListWebhooksInput, ...func(*codepipeline.Options)) (*codepipeline.ListWebhooksOutput, error)
}
