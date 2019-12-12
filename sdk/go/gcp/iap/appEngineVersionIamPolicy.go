// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_app_engine_version_iam_policy.html.markdown.
type AppEngineVersionIamPolicy struct {
	s *pulumi.ResourceState
}

// NewAppEngineVersionIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAppEngineVersionIamPolicy(ctx *pulumi.Context,
	name string, args *AppEngineVersionIamPolicyArgs, opts ...pulumi.ResourceOpt) (*AppEngineVersionIamPolicy, error) {
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil || args.VersionId == nil {
		return nil, errors.New("missing required argument 'VersionId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appId"] = nil
		inputs["policyData"] = nil
		inputs["project"] = nil
		inputs["service"] = nil
		inputs["versionId"] = nil
	} else {
		inputs["appId"] = args.AppId
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
		inputs["service"] = args.Service
		inputs["versionId"] = args.VersionId
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AppEngineVersionIamPolicy{s: s}, nil
}

// GetAppEngineVersionIamPolicy gets an existing AppEngineVersionIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineVersionIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AppEngineVersionIamPolicyState, opts ...pulumi.ResourceOpt) (*AppEngineVersionIamPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appId"] = state.AppId
		inputs["etag"] = state.Etag
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
		inputs["service"] = state.Service
		inputs["versionId"] = state.VersionId
	}
	s, err := ctx.ReadResource("gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AppEngineVersionIamPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AppEngineVersionIamPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AppEngineVersionIamPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
func (r *AppEngineVersionIamPolicy) AppId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["appId"])
}

// (Computed) The etag of the IAM policy.
func (r *AppEngineVersionIamPolicy) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (r *AppEngineVersionIamPolicy) PolicyData() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policyData"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *AppEngineVersionIamPolicy) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
func (r *AppEngineVersionIamPolicy) Service() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["service"])
}

// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
func (r *AppEngineVersionIamPolicy) VersionId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["versionId"])
}

// Input properties used for looking up and filtering AppEngineVersionIamPolicy resources.
type AppEngineVersionIamPolicyState struct {
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
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service interface{}
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId interface{}
}

// The set of arguments for constructing a AppEngineVersionIamPolicy resource.
type AppEngineVersionIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service interface{}
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId interface{}
}