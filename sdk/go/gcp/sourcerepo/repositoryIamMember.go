// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sourcerepo

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Pub/Sub Topic. Each of these resources serves a different use case:
//
// * `pubsub.TopicIAMPolicy`: Authoritative. Sets the IAM policy for the topic and replaces any existing policy already attached.
// * `pubsub.TopicIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the topic are preserved.
// * `pubsub.TopicIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the topic are preserved.
//
// > **Note:** `pubsub.TopicIAMPolicy` **cannot** be used in conjunction with `pubsub.TopicIAMBinding` and `pubsub.TopicIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `pubsub.TopicIAMBinding` resources **can be** used in conjunction with `pubsub.TopicIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_pubsub\_topic\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewTopicIAMPolicy(ctx, "policy", &pubsub.TopicIAMPolicyArgs{
// 			Project:    pulumi.Any(google_pubsub_topic.Example.Project),
// 			Topic:      pulumi.Any(google_pubsub_topic.Example.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_topic\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewTopicIAMBinding(ctx, "binding", &pubsub.TopicIAMBindingArgs{
// 			Project: pulumi.Any(google_pubsub_topic.Example.Project),
// 			Topic:   pulumi.Any(google_pubsub_topic.Example.Name),
// 			Role:    pulumi.String("roles/viewer"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_topic\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewTopicIAMMember(ctx, "member", &pubsub.TopicIAMMemberArgs{
// 			Project: pulumi.Any(google_pubsub_topic.Example.Project),
// 			Topic:   pulumi.Any(google_pubsub_topic.Example.Name),
// 			Role:    pulumi.String("roles/viewer"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/topics/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Pub/Sub topic IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:sourcerepo/repositoryIamMember:RepositoryIamMember editor "projects/{{project}}/topics/{{topic}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:sourcerepo/repositoryIamMember:RepositoryIamMember editor "projects/{{project}}/topics/{{topic}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:sourcerepo/repositoryIamMember:RepositoryIamMember editor projects/{{project}}/topics/{{topic}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RepositoryIamMember struct {
	pulumi.CustomResourceState

	Condition RepositoryIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    pulumi.StringOutput `pulumi:"project"`
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRepositoryIamMember registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamMember(ctx *pulumi.Context,
	name string, args *RepositoryIamMemberArgs, opts ...pulumi.ResourceOption) (*RepositoryIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource RepositoryIamMember
	err := ctx.RegisterResource("gcp:sourcerepo/repositoryIamMember:RepositoryIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryIamMember gets an existing RepositoryIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryIamMemberState, opts ...pulumi.ResourceOption) (*RepositoryIamMember, error) {
	var resource RepositoryIamMember
	err := ctx.ReadResource("gcp:sourcerepo/repositoryIamMember:RepositoryIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryIamMember resources.
type repositoryIamMemberState struct {
	Condition *RepositoryIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    *string `pulumi:"project"`
	Repository *string `pulumi:"repository"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type RepositoryIamMemberState struct {
	Condition RepositoryIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    pulumi.StringPtrInput
	Repository pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (RepositoryIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamMemberState)(nil)).Elem()
}

type repositoryIamMemberArgs struct {
	Condition *RepositoryIamMemberCondition `pulumi:"condition"`
	Member    string                        `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    *string `pulumi:"project"`
	Repository string  `pulumi:"repository"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a RepositoryIamMember resource.
type RepositoryIamMemberArgs struct {
	Condition RepositoryIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    pulumi.StringPtrInput
	Repository pulumi.StringInput
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (RepositoryIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamMemberArgs)(nil)).Elem()
}

type RepositoryIamMemberInput interface {
	pulumi.Input

	ToRepositoryIamMemberOutput() RepositoryIamMemberOutput
	ToRepositoryIamMemberOutputWithContext(ctx context.Context) RepositoryIamMemberOutput
}

func (*RepositoryIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamMember)(nil))
}

func (i *RepositoryIamMember) ToRepositoryIamMemberOutput() RepositoryIamMemberOutput {
	return i.ToRepositoryIamMemberOutputWithContext(context.Background())
}

func (i *RepositoryIamMember) ToRepositoryIamMemberOutputWithContext(ctx context.Context) RepositoryIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberOutput)
}

func (i *RepositoryIamMember) ToRepositoryIamMemberPtrOutput() RepositoryIamMemberPtrOutput {
	return i.ToRepositoryIamMemberPtrOutputWithContext(context.Background())
}

func (i *RepositoryIamMember) ToRepositoryIamMemberPtrOutputWithContext(ctx context.Context) RepositoryIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberPtrOutput)
}

type RepositoryIamMemberPtrInput interface {
	pulumi.Input

	ToRepositoryIamMemberPtrOutput() RepositoryIamMemberPtrOutput
	ToRepositoryIamMemberPtrOutputWithContext(ctx context.Context) RepositoryIamMemberPtrOutput
}

type repositoryIamMemberPtrType RepositoryIamMemberArgs

func (*repositoryIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamMember)(nil))
}

func (i *repositoryIamMemberPtrType) ToRepositoryIamMemberPtrOutput() RepositoryIamMemberPtrOutput {
	return i.ToRepositoryIamMemberPtrOutputWithContext(context.Background())
}

func (i *repositoryIamMemberPtrType) ToRepositoryIamMemberPtrOutputWithContext(ctx context.Context) RepositoryIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberPtrOutput)
}

// RepositoryIamMemberArrayInput is an input type that accepts RepositoryIamMemberArray and RepositoryIamMemberArrayOutput values.
// You can construct a concrete instance of `RepositoryIamMemberArrayInput` via:
//
//          RepositoryIamMemberArray{ RepositoryIamMemberArgs{...} }
type RepositoryIamMemberArrayInput interface {
	pulumi.Input

	ToRepositoryIamMemberArrayOutput() RepositoryIamMemberArrayOutput
	ToRepositoryIamMemberArrayOutputWithContext(context.Context) RepositoryIamMemberArrayOutput
}

type RepositoryIamMemberArray []RepositoryIamMemberInput

func (RepositoryIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryIamMember)(nil)).Elem()
}

func (i RepositoryIamMemberArray) ToRepositoryIamMemberArrayOutput() RepositoryIamMemberArrayOutput {
	return i.ToRepositoryIamMemberArrayOutputWithContext(context.Background())
}

func (i RepositoryIamMemberArray) ToRepositoryIamMemberArrayOutputWithContext(ctx context.Context) RepositoryIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberArrayOutput)
}

// RepositoryIamMemberMapInput is an input type that accepts RepositoryIamMemberMap and RepositoryIamMemberMapOutput values.
// You can construct a concrete instance of `RepositoryIamMemberMapInput` via:
//
//          RepositoryIamMemberMap{ "key": RepositoryIamMemberArgs{...} }
type RepositoryIamMemberMapInput interface {
	pulumi.Input

	ToRepositoryIamMemberMapOutput() RepositoryIamMemberMapOutput
	ToRepositoryIamMemberMapOutputWithContext(context.Context) RepositoryIamMemberMapOutput
}

type RepositoryIamMemberMap map[string]RepositoryIamMemberInput

func (RepositoryIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryIamMember)(nil)).Elem()
}

func (i RepositoryIamMemberMap) ToRepositoryIamMemberMapOutput() RepositoryIamMemberMapOutput {
	return i.ToRepositoryIamMemberMapOutputWithContext(context.Background())
}

func (i RepositoryIamMemberMap) ToRepositoryIamMemberMapOutputWithContext(ctx context.Context) RepositoryIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberMapOutput)
}

type RepositoryIamMemberOutput struct{ *pulumi.OutputState }

func (RepositoryIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamMember)(nil))
}

func (o RepositoryIamMemberOutput) ToRepositoryIamMemberOutput() RepositoryIamMemberOutput {
	return o
}

func (o RepositoryIamMemberOutput) ToRepositoryIamMemberOutputWithContext(ctx context.Context) RepositoryIamMemberOutput {
	return o
}

func (o RepositoryIamMemberOutput) ToRepositoryIamMemberPtrOutput() RepositoryIamMemberPtrOutput {
	return o.ToRepositoryIamMemberPtrOutputWithContext(context.Background())
}

func (o RepositoryIamMemberOutput) ToRepositoryIamMemberPtrOutputWithContext(ctx context.Context) RepositoryIamMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RepositoryIamMember) *RepositoryIamMember {
		return &v
	}).(RepositoryIamMemberPtrOutput)
}

type RepositoryIamMemberPtrOutput struct{ *pulumi.OutputState }

func (RepositoryIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamMember)(nil))
}

func (o RepositoryIamMemberPtrOutput) ToRepositoryIamMemberPtrOutput() RepositoryIamMemberPtrOutput {
	return o
}

func (o RepositoryIamMemberPtrOutput) ToRepositoryIamMemberPtrOutputWithContext(ctx context.Context) RepositoryIamMemberPtrOutput {
	return o
}

func (o RepositoryIamMemberPtrOutput) Elem() RepositoryIamMemberOutput {
	return o.ApplyT(func(v *RepositoryIamMember) RepositoryIamMember {
		if v != nil {
			return *v
		}
		var ret RepositoryIamMember
		return ret
	}).(RepositoryIamMemberOutput)
}

type RepositoryIamMemberArrayOutput struct{ *pulumi.OutputState }

func (RepositoryIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryIamMember)(nil))
}

func (o RepositoryIamMemberArrayOutput) ToRepositoryIamMemberArrayOutput() RepositoryIamMemberArrayOutput {
	return o
}

func (o RepositoryIamMemberArrayOutput) ToRepositoryIamMemberArrayOutputWithContext(ctx context.Context) RepositoryIamMemberArrayOutput {
	return o
}

func (o RepositoryIamMemberArrayOutput) Index(i pulumi.IntInput) RepositoryIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepositoryIamMember {
		return vs[0].([]RepositoryIamMember)[vs[1].(int)]
	}).(RepositoryIamMemberOutput)
}

type RepositoryIamMemberMapOutput struct{ *pulumi.OutputState }

func (RepositoryIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RepositoryIamMember)(nil))
}

func (o RepositoryIamMemberMapOutput) ToRepositoryIamMemberMapOutput() RepositoryIamMemberMapOutput {
	return o
}

func (o RepositoryIamMemberMapOutput) ToRepositoryIamMemberMapOutputWithContext(ctx context.Context) RepositoryIamMemberMapOutput {
	return o
}

func (o RepositoryIamMemberMapOutput) MapIndex(k pulumi.StringInput) RepositoryIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RepositoryIamMember {
		return vs[0].(map[string]RepositoryIamMember)[vs[1].(string)]
	}).(RepositoryIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryIamMemberOutput{})
	pulumi.RegisterOutputType(RepositoryIamMemberPtrOutput{})
	pulumi.RegisterOutputType(RepositoryIamMemberArrayOutput{})
	pulumi.RegisterOutputType(RepositoryIamMemberMapOutput{})
}
