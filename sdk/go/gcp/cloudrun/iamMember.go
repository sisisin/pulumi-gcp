// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_run_service_iam_member.html.markdown.
type IamMember struct {
	s *pulumi.ResourceState
}

// NewIamMember registers a new resource with the given unique name, arguments, and options.
func NewIamMember(ctx *pulumi.Context,
	name string, args *IamMemberArgs, opts ...pulumi.ResourceOpt) (*IamMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["location"] = nil
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
		inputs["service"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["location"] = args.Location
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["role"] = args.Role
		inputs["service"] = args.Service
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:cloudrun/iamMember:IamMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IamMember{s: s}, nil
}

// GetIamMember gets an existing IamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IamMemberState, opts ...pulumi.ResourceOpt) (*IamMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["location"] = state.Location
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["role"] = state.Role
		inputs["service"] = state.Service
	}
	s, err := ctx.ReadResource("gcp:cloudrun/iamMember:IamMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IamMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IamMember) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IamMember) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *IamMember) Condition() pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the IAM policy.
func (r *IamMember) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
func (r *IamMember) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

func (r *IamMember) Member() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["member"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *IamMember) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *IamMember) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// Used to find the parent resource to bind the IAM policy to
func (r *IamMember) Service() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["service"])
}

// Input properties used for looking up and filtering IamMember resources.
type IamMemberState struct {
	Condition interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location interface{}
	Member interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// Used to find the parent resource to bind the IAM policy to
	Service interface{}
}

// The set of arguments for constructing a IamMember resource.
type IamMemberArgs struct {
	Condition interface{}
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location interface{}
	Member interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// Used to find the parent resource to bind the IAM policy to
	Service interface{}
}