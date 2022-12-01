// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

//go:generate mockgen -package=mocks -destination=../mocks/scheduler.go -source=scheduler.go SchedulerClient
type SchedulerClient interface {
	GetSchedule(context.Context, *scheduler.GetScheduleInput, ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error)
	GetScheduleGroup(context.Context, *scheduler.GetScheduleGroupInput, ...func(*scheduler.Options)) (*scheduler.GetScheduleGroupOutput, error)
	ListScheduleGroups(context.Context, *scheduler.ListScheduleGroupsInput, ...func(*scheduler.Options)) (*scheduler.ListScheduleGroupsOutput, error)
	ListSchedules(context.Context, *scheduler.ListSchedulesInput, ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error)
	ListTagsForResource(context.Context, *scheduler.ListTagsForResourceInput, ...func(*scheduler.Options)) (*scheduler.ListTagsForResourceOutput, error)
}
