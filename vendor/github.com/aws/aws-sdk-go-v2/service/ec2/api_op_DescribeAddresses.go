// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAddressesInput struct {
	_ struct{} `type:"structure"`

	// [EC2-VPC] Information about the allocation IDs.
	AllocationIds []string `locationName:"AllocationId" locationNameList:"AllocationId" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters. Filter names and values are case-sensitive.
	//
	//    * allocation-id - [EC2-VPC] The allocation ID for the address.
	//
	//    * association-id - [EC2-VPC] The association ID for the address.
	//
	//    * domain - Indicates whether the address is for use in EC2-Classic (standard)
	//    or in a VPC (vpc).
	//
	//    * instance-id - The ID of the instance the address is associated with,
	//    if any.
	//
	//    * network-border-group - The location from where the IP address is advertised.
	//
	//    * network-interface-id - [EC2-VPC] The ID of the network interface that
	//    the address is associated with, if any.
	//
	//    * network-interface-owner-id - The AWS account ID of the owner.
	//
	//    * private-ip-address - [EC2-VPC] The private IP address associated with
	//    the Elastic IP address.
	//
	//    * public-ip - The Elastic IP address.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more Elastic IP addresses.
	//
	// Default: Describes all your Elastic IP addresses.
	PublicIps []string `locationName:"PublicIp" locationNameList:"PublicIp" type:"list"`
}

// String returns the string representation
func (s DescribeAddressesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeAddressesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Elastic IP addresses.
	Addresses []Address `locationName:"addressesSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeAddressesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAddresses = "DescribeAddresses"

// DescribeAddressesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified Elastic IP addresses or all of your Elastic IP addresses.
//
// An Elastic IP address is for use in either the EC2-Classic platform or in
// a VPC. For more information, see Elastic IP Addresses (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeAddressesRequest.
//    req := client.DescribeAddressesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAddresses
func (c *Client) DescribeAddressesRequest(input *DescribeAddressesInput) DescribeAddressesRequest {
	op := &aws.Operation{
		Name:       opDescribeAddresses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAddressesInput{}
	}

	req := c.newRequest(op, input, &DescribeAddressesOutput{})
	return DescribeAddressesRequest{Request: req, Input: input, Copy: c.DescribeAddressesRequest}
}

// DescribeAddressesRequest is the request type for the
// DescribeAddresses API operation.
type DescribeAddressesRequest struct {
	*aws.Request
	Input *DescribeAddressesInput
	Copy  func(*DescribeAddressesInput) DescribeAddressesRequest
}

// Send marshals and sends the DescribeAddresses API request.
func (r DescribeAddressesRequest) Send(ctx context.Context) (*DescribeAddressesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAddressesResponse{
		DescribeAddressesOutput: r.Request.Data.(*DescribeAddressesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAddressesResponse is the response type for the
// DescribeAddresses API operation.
type DescribeAddressesResponse struct {
	*DescribeAddressesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAddresses request.
func (r *DescribeAddressesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
