// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeKeyPairsRequest
type DescribeKeyPairsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * fingerprint - The fingerprint of the key pair.
	//
	//    * key-name - The name of the key pair.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The key pair names.
	//
	// Default: Describes all your key pairs.
	KeyNames []string `locationName:"KeyName" locationNameList:"KeyName" type:"list"`
}

// String returns the string representation
func (s DescribeKeyPairsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeKeyPairsResult
type DescribeKeyPairsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the key pairs.
	KeyPairs []KeyPairInfo `locationName:"keySet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeKeyPairsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeKeyPairs = "DescribeKeyPairs"

// DescribeKeyPairsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified key pairs or all of your key pairs.
//
// For more information about key pairs, see Key Pairs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeKeyPairsRequest.
//    req := client.DescribeKeyPairsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeKeyPairs
func (c *Client) DescribeKeyPairsRequest(input *DescribeKeyPairsInput) DescribeKeyPairsRequest {
	op := &aws.Operation{
		Name:       opDescribeKeyPairs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeKeyPairsInput{}
	}

	req := c.newRequest(op, input, &DescribeKeyPairsOutput{})
	return DescribeKeyPairsRequest{Request: req, Input: input, Copy: c.DescribeKeyPairsRequest}
}

// DescribeKeyPairsRequest is the request type for the
// DescribeKeyPairs API operation.
type DescribeKeyPairsRequest struct {
	*aws.Request
	Input *DescribeKeyPairsInput
	Copy  func(*DescribeKeyPairsInput) DescribeKeyPairsRequest
}

// Send marshals and sends the DescribeKeyPairs API request.
func (r DescribeKeyPairsRequest) Send(ctx context.Context) (*DescribeKeyPairsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeKeyPairsResponse{
		DescribeKeyPairsOutput: r.Request.Data.(*DescribeKeyPairsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeKeyPairsResponse is the response type for the
// DescribeKeyPairs API operation.
type DescribeKeyPairsResponse struct {
	*DescribeKeyPairsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeKeyPairs request.
func (r *DescribeKeyPairsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
