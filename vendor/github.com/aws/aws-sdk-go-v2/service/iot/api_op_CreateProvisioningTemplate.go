// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateProvisioningTemplateInput struct {
	_ struct{} `type:"structure"`

	// The description of the fleet provisioning template.
	Description *string `locationName:"description" type:"string"`

	// True to enable the fleet provisioning template, otherwise false.
	Enabled *bool `locationName:"enabled" type:"boolean"`

	// The role ARN for the role associated with the fleet provisioning template.
	// This IoT role grants permission to provision a device.
	//
	// ProvisioningRoleArn is a required field
	ProvisioningRoleArn *string `locationName:"provisioningRoleArn" min:"20" type:"string" required:"true"`

	// Metadata which can be used to manage the fleet provisioning template.
	//
	// For URI Request parameters use format: ...key1=value1&key2=value2...
	//
	// For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	//
	// For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags []Tag `locationName:"tags" type:"list"`

	// The JSON formatted contents of the fleet provisioning template.
	//
	// TemplateBody is a required field
	TemplateBody *string `locationName:"templateBody" type:"string" required:"true"`

	// The name of the fleet provisioning template.
	//
	// TemplateName is a required field
	TemplateName *string `locationName:"templateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProvisioningTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProvisioningTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProvisioningTemplateInput"}

	if s.ProvisioningRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisioningRoleArn"))
	}
	if s.ProvisioningRoleArn != nil && len(*s.ProvisioningRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisioningRoleArn", 20))
	}

	if s.TemplateBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateBody"))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProvisioningTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enabled", protocol.BoolValue(v), metadata)
	}
	if s.ProvisioningRoleArn != nil {
		v := *s.ProvisioningRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "provisioningRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TemplateBody != nil {
		v := *s.TemplateBody

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateBody", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateProvisioningTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The default version of the fleet provisioning template.
	DefaultVersionId *int64 `locationName:"defaultVersionId" type:"integer"`

	// The ARN that identifies the provisioning template.
	TemplateArn *string `locationName:"templateArn" type:"string"`

	// The name of the fleet provisioning template.
	TemplateName *string `locationName:"templateName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateProvisioningTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProvisioningTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DefaultVersionId != nil {
		v := *s.DefaultVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultVersionId", protocol.Int64Value(v), metadata)
	}
	if s.TemplateArn != nil {
		v := *s.TemplateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateProvisioningTemplate = "CreateProvisioningTemplate"

// CreateProvisioningTemplateRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a fleet provisioning template.
//
//    // Example sending a request using CreateProvisioningTemplateRequest.
//    req := client.CreateProvisioningTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateProvisioningTemplateRequest(input *CreateProvisioningTemplateInput) CreateProvisioningTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateProvisioningTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/provisioning-templates",
	}

	if input == nil {
		input = &CreateProvisioningTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateProvisioningTemplateOutput{})
	return CreateProvisioningTemplateRequest{Request: req, Input: input, Copy: c.CreateProvisioningTemplateRequest}
}

// CreateProvisioningTemplateRequest is the request type for the
// CreateProvisioningTemplate API operation.
type CreateProvisioningTemplateRequest struct {
	*aws.Request
	Input *CreateProvisioningTemplateInput
	Copy  func(*CreateProvisioningTemplateInput) CreateProvisioningTemplateRequest
}

// Send marshals and sends the CreateProvisioningTemplate API request.
func (r CreateProvisioningTemplateRequest) Send(ctx context.Context) (*CreateProvisioningTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProvisioningTemplateResponse{
		CreateProvisioningTemplateOutput: r.Request.Data.(*CreateProvisioningTemplateOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProvisioningTemplateResponse is the response type for the
// CreateProvisioningTemplate API operation.
type CreateProvisioningTemplateResponse struct {
	*CreateProvisioningTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProvisioningTemplate request.
func (r *CreateProvisioningTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
