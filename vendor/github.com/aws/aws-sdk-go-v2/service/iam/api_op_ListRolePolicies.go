// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRolePoliciesInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The name of the role to list policies for.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListRolePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRolePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRolePoliciesInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListRolePolicies request.
type ListRolePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`

	// A list of policy names.
	//
	// PolicyNames is a required field
	PolicyNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s ListRolePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRolePolicies = "ListRolePolicies"

// ListRolePoliciesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the names of the inline policies that are embedded in the specified
// IAM role.
//
// An IAM role can also have managed policies attached to it. To list the managed
// policies that are attached to a role, use ListAttachedRolePolicies. For more
// information about policies, see Managed Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. If
// there are no inline policies embedded with the specified role, the operation
// returns an empty list.
//
//    // Example sending a request using ListRolePoliciesRequest.
//    req := client.ListRolePoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListRolePolicies
func (c *Client) ListRolePoliciesRequest(input *ListRolePoliciesInput) ListRolePoliciesRequest {
	op := &aws.Operation{
		Name:       opListRolePolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListRolePoliciesInput{}
	}

	req := c.newRequest(op, input, &ListRolePoliciesOutput{})
	return ListRolePoliciesRequest{Request: req, Input: input, Copy: c.ListRolePoliciesRequest}
}

// ListRolePoliciesRequest is the request type for the
// ListRolePolicies API operation.
type ListRolePoliciesRequest struct {
	*aws.Request
	Input *ListRolePoliciesInput
	Copy  func(*ListRolePoliciesInput) ListRolePoliciesRequest
}

// Send marshals and sends the ListRolePolicies API request.
func (r ListRolePoliciesRequest) Send(ctx context.Context) (*ListRolePoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRolePoliciesResponse{
		ListRolePoliciesOutput: r.Request.Data.(*ListRolePoliciesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRolePoliciesRequestPaginator returns a paginator for ListRolePolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRolePoliciesRequest(input)
//   p := iam.NewListRolePoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRolePoliciesPaginator(req ListRolePoliciesRequest) ListRolePoliciesPaginator {
	return ListRolePoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRolePoliciesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListRolePoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRolePoliciesPaginator struct {
	aws.Pager
}

func (p *ListRolePoliciesPaginator) CurrentPage() *ListRolePoliciesOutput {
	return p.Pager.CurrentPage().(*ListRolePoliciesOutput)
}

// ListRolePoliciesResponse is the response type for the
// ListRolePolicies API operation.
type ListRolePoliciesResponse struct {
	*ListRolePoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRolePolicies request.
func (r *ListRolePoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
