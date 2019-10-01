// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_web_type_app_engine_iam_policy.html.markdown.
type WebTypeAppEngingIamPolicy struct {
	s *pulumi.ResourceState
}

// NewWebTypeAppEngingIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebTypeAppEngingIamPolicy(ctx *pulumi.Context,
	name string, args *WebTypeAppEngingIamPolicyArgs, opts ...pulumi.ResourceOpt) (*WebTypeAppEngingIamPolicy, error) {
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appId"] = nil
		inputs["policyData"] = nil
		inputs["project"] = nil
	} else {
		inputs["appId"] = args.AppId
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &WebTypeAppEngingIamPolicy{s: s}, nil
}

// GetWebTypeAppEngingIamPolicy gets an existing WebTypeAppEngingIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeAppEngingIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *WebTypeAppEngingIamPolicyState, opts ...pulumi.ResourceOpt) (*WebTypeAppEngingIamPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appId"] = state.AppId
		inputs["etag"] = state.Etag
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
	}
	s, err := ctx.ReadResource("gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &WebTypeAppEngingIamPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *WebTypeAppEngingIamPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *WebTypeAppEngingIamPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
func (r *WebTypeAppEngingIamPolicy) AppId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["appId"])
}

// (Computed) The etag of the IAM policy.
func (r *WebTypeAppEngingIamPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (r *WebTypeAppEngingIamPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *WebTypeAppEngingIamPolicy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Input properties used for looking up and filtering WebTypeAppEngingIamPolicy resources.
type WebTypeAppEngingIamPolicyState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
}

// The set of arguments for constructing a WebTypeAppEngingIamPolicy resource.
type WebTypeAppEngingIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
}