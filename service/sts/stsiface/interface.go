package stsiface

import (
	"github.com/awslabs/aws-sdk-go/service/sts"
)

type STSAPI interface {
	AssumeRole(*sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error)

	AssumeRoleWithSAML(*sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error)

	AssumeRoleWithWebIdentity(*sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error)

	DecodeAuthorizationMessage(*sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error)

	GetFederationToken(*sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error)

	GetSessionToken(*sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error)
}