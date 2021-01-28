/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package authorizer

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.
// TODO(muvaf): We can generate one-time boilerplate for these hooks but currently
// ACK doesn't support not generating if file exists.

// GenerateGetAuthorizerInput returns input for read
// operation.
func GenerateGetAuthorizerInput(cr *svcapitypes.Authorizer) *svcsdk.GetAuthorizerInput {
	res := preGenerateGetAuthorizerInput(cr, &svcsdk.GetAuthorizerInput{})

	if cr.Status.AtProvider.AuthorizerID != nil {
		res.SetAuthorizerId(*cr.Status.AtProvider.AuthorizerID)
	}

	return postGenerateGetAuthorizerInput(cr, res)
}

// GenerateAuthorizer returns the current state in the form of *svcapitypes.Authorizer.
func GenerateAuthorizer(resp *svcsdk.GetAuthorizerOutput) *svcapitypes.Authorizer {
	cr := &svcapitypes.Authorizer{}

	if resp.AuthorizerId != nil {
		cr.Status.AtProvider.AuthorizerID = resp.AuthorizerId
	}

	return cr
}

// GenerateCreateAuthorizerInput returns a create input.
func GenerateCreateAuthorizerInput(cr *svcapitypes.Authorizer) *svcsdk.CreateAuthorizerInput {
	res := preGenerateCreateAuthorizerInput(cr, &svcsdk.CreateAuthorizerInput{})

	if cr.Spec.ForProvider.AuthorizerCredentialsARN != nil {
		res.SetAuthorizerCredentialsArn(*cr.Spec.ForProvider.AuthorizerCredentialsARN)
	}
	if cr.Spec.ForProvider.AuthorizerPayloadFormatVersion != nil {
		res.SetAuthorizerPayloadFormatVersion(*cr.Spec.ForProvider.AuthorizerPayloadFormatVersion)
	}
	if cr.Spec.ForProvider.AuthorizerResultTtlInSeconds != nil {
		res.SetAuthorizerResultTtlInSeconds(*cr.Spec.ForProvider.AuthorizerResultTtlInSeconds)
	}
	if cr.Spec.ForProvider.AuthorizerType != nil {
		res.SetAuthorizerType(*cr.Spec.ForProvider.AuthorizerType)
	}
	if cr.Spec.ForProvider.AuthorizerURI != nil {
		res.SetAuthorizerUri(*cr.Spec.ForProvider.AuthorizerURI)
	}
	if cr.Spec.ForProvider.EnableSimpleResponses != nil {
		res.SetEnableSimpleResponses(*cr.Spec.ForProvider.EnableSimpleResponses)
	}
	if cr.Spec.ForProvider.IDentitySource != nil {
		f6 := []*string{}
		for _, f6iter := range cr.Spec.ForProvider.IDentitySource {
			var f6elem string
			f6elem = *f6iter
			f6 = append(f6, &f6elem)
		}
		res.SetIdentitySource(f6)
	}
	if cr.Spec.ForProvider.IDentityValidationExpression != nil {
		res.SetIdentityValidationExpression(*cr.Spec.ForProvider.IDentityValidationExpression)
	}
	if cr.Spec.ForProvider.JWTConfiguration != nil {
		f8 := &svcsdk.JWTConfiguration{}
		if cr.Spec.ForProvider.JWTConfiguration.Audience != nil {
			f8f0 := []*string{}
			for _, f8f0iter := range cr.Spec.ForProvider.JWTConfiguration.Audience {
				var f8f0elem string
				f8f0elem = *f8f0iter
				f8f0 = append(f8f0, &f8f0elem)
			}
			f8.SetAudience(f8f0)
		}
		if cr.Spec.ForProvider.JWTConfiguration.Issuer != nil {
			f8.SetIssuer(*cr.Spec.ForProvider.JWTConfiguration.Issuer)
		}
		res.SetJwtConfiguration(f8)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}

	return postGenerateCreateAuthorizerInput(cr, res)
}

// GenerateDeleteAuthorizerInput returns a deletion input.
func GenerateDeleteAuthorizerInput(cr *svcapitypes.Authorizer) *svcsdk.DeleteAuthorizerInput {
	res := preGenerateDeleteAuthorizerInput(cr, &svcsdk.DeleteAuthorizerInput{})

	if cr.Status.AtProvider.AuthorizerID != nil {
		res.SetAuthorizerId(*cr.Status.AtProvider.AuthorizerID)
	}

	return postGenerateDeleteAuthorizerInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}