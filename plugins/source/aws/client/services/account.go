// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
)

//go:generate mockgen -package=mocks -destination=../mocks/account.go -source=account.go AccountClient
type AccountClient interface {
	GetAlternateContact(context.Context, *account.GetAlternateContactInput, ...func(*account.Options)) (*account.GetAlternateContactOutput, error)
	GetContactInformation(context.Context, *account.GetContactInformationInput, ...func(*account.Options)) (*account.GetContactInformationOutput, error)
}
