// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

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
//  $ pulumi import gcp:pubsub/topicIAMMember:TopicIAMMember editor "projects/{{project}}/topics/{{topic}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:pubsub/topicIAMMember:TopicIAMMember editor "projects/{{project}}/topics/{{topic}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:pubsub/topicIAMMember:TopicIAMMember editor projects/{{project}}/topics/{{topic}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TopicIAMMember struct {
	pulumi.CustomResourceState

	Condition TopicIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Topic pulumi.StringOutput `pulumi:"topic"`
}

// NewTopicIAMMember registers a new resource with the given unique name, arguments, and options.
func NewTopicIAMMember(ctx *pulumi.Context,
	name string, args *TopicIAMMemberArgs, opts ...pulumi.ResourceOption) (*TopicIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Topic == nil {
		return nil, errors.New("invalid value for required argument 'Topic'")
	}
	var resource TopicIAMMember
	err := ctx.RegisterResource("gcp:pubsub/topicIAMMember:TopicIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicIAMMember gets an existing TopicIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicIAMMemberState, opts ...pulumi.ResourceOption) (*TopicIAMMember, error) {
	var resource TopicIAMMember
	err := ctx.ReadResource("gcp:pubsub/topicIAMMember:TopicIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicIAMMember resources.
type topicIAMMemberState struct {
	Condition *TopicIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Topic *string `pulumi:"topic"`
}

type TopicIAMMemberState struct {
	Condition TopicIAMMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Topic pulumi.StringPtrInput
}

func (TopicIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicIAMMemberState)(nil)).Elem()
}

type topicIAMMemberArgs struct {
	Condition *TopicIAMMemberCondition `pulumi:"condition"`
	Member    string                   `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Topic string `pulumi:"topic"`
}

// The set of arguments for constructing a TopicIAMMember resource.
type TopicIAMMemberArgs struct {
	Condition TopicIAMMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Topic pulumi.StringInput
}

func (TopicIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicIAMMemberArgs)(nil)).Elem()
}

type TopicIAMMemberInput interface {
	pulumi.Input

	ToTopicIAMMemberOutput() TopicIAMMemberOutput
	ToTopicIAMMemberOutputWithContext(ctx context.Context) TopicIAMMemberOutput
}

func (*TopicIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicIAMMember)(nil))
}

func (i *TopicIAMMember) ToTopicIAMMemberOutput() TopicIAMMemberOutput {
	return i.ToTopicIAMMemberOutputWithContext(context.Background())
}

func (i *TopicIAMMember) ToTopicIAMMemberOutputWithContext(ctx context.Context) TopicIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicIAMMemberOutput)
}

func (i *TopicIAMMember) ToTopicIAMMemberPtrOutput() TopicIAMMemberPtrOutput {
	return i.ToTopicIAMMemberPtrOutputWithContext(context.Background())
}

func (i *TopicIAMMember) ToTopicIAMMemberPtrOutputWithContext(ctx context.Context) TopicIAMMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicIAMMemberPtrOutput)
}

type TopicIAMMemberPtrInput interface {
	pulumi.Input

	ToTopicIAMMemberPtrOutput() TopicIAMMemberPtrOutput
	ToTopicIAMMemberPtrOutputWithContext(ctx context.Context) TopicIAMMemberPtrOutput
}

type topicIAMMemberPtrType TopicIAMMemberArgs

func (*topicIAMMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicIAMMember)(nil))
}

func (i *topicIAMMemberPtrType) ToTopicIAMMemberPtrOutput() TopicIAMMemberPtrOutput {
	return i.ToTopicIAMMemberPtrOutputWithContext(context.Background())
}

func (i *topicIAMMemberPtrType) ToTopicIAMMemberPtrOutputWithContext(ctx context.Context) TopicIAMMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicIAMMemberPtrOutput)
}

// TopicIAMMemberArrayInput is an input type that accepts TopicIAMMemberArray and TopicIAMMemberArrayOutput values.
// You can construct a concrete instance of `TopicIAMMemberArrayInput` via:
//
//          TopicIAMMemberArray{ TopicIAMMemberArgs{...} }
type TopicIAMMemberArrayInput interface {
	pulumi.Input

	ToTopicIAMMemberArrayOutput() TopicIAMMemberArrayOutput
	ToTopicIAMMemberArrayOutputWithContext(context.Context) TopicIAMMemberArrayOutput
}

type TopicIAMMemberArray []TopicIAMMemberInput

func (TopicIAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TopicIAMMember)(nil)).Elem()
}

func (i TopicIAMMemberArray) ToTopicIAMMemberArrayOutput() TopicIAMMemberArrayOutput {
	return i.ToTopicIAMMemberArrayOutputWithContext(context.Background())
}

func (i TopicIAMMemberArray) ToTopicIAMMemberArrayOutputWithContext(ctx context.Context) TopicIAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicIAMMemberArrayOutput)
}

// TopicIAMMemberMapInput is an input type that accepts TopicIAMMemberMap and TopicIAMMemberMapOutput values.
// You can construct a concrete instance of `TopicIAMMemberMapInput` via:
//
//          TopicIAMMemberMap{ "key": TopicIAMMemberArgs{...} }
type TopicIAMMemberMapInput interface {
	pulumi.Input

	ToTopicIAMMemberMapOutput() TopicIAMMemberMapOutput
	ToTopicIAMMemberMapOutputWithContext(context.Context) TopicIAMMemberMapOutput
}

type TopicIAMMemberMap map[string]TopicIAMMemberInput

func (TopicIAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TopicIAMMember)(nil)).Elem()
}

func (i TopicIAMMemberMap) ToTopicIAMMemberMapOutput() TopicIAMMemberMapOutput {
	return i.ToTopicIAMMemberMapOutputWithContext(context.Background())
}

func (i TopicIAMMemberMap) ToTopicIAMMemberMapOutputWithContext(ctx context.Context) TopicIAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicIAMMemberMapOutput)
}

type TopicIAMMemberOutput struct{ *pulumi.OutputState }

func (TopicIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicIAMMember)(nil))
}

func (o TopicIAMMemberOutput) ToTopicIAMMemberOutput() TopicIAMMemberOutput {
	return o
}

func (o TopicIAMMemberOutput) ToTopicIAMMemberOutputWithContext(ctx context.Context) TopicIAMMemberOutput {
	return o
}

func (o TopicIAMMemberOutput) ToTopicIAMMemberPtrOutput() TopicIAMMemberPtrOutput {
	return o.ToTopicIAMMemberPtrOutputWithContext(context.Background())
}

func (o TopicIAMMemberOutput) ToTopicIAMMemberPtrOutputWithContext(ctx context.Context) TopicIAMMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TopicIAMMember) *TopicIAMMember {
		return &v
	}).(TopicIAMMemberPtrOutput)
}

type TopicIAMMemberPtrOutput struct{ *pulumi.OutputState }

func (TopicIAMMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicIAMMember)(nil))
}

func (o TopicIAMMemberPtrOutput) ToTopicIAMMemberPtrOutput() TopicIAMMemberPtrOutput {
	return o
}

func (o TopicIAMMemberPtrOutput) ToTopicIAMMemberPtrOutputWithContext(ctx context.Context) TopicIAMMemberPtrOutput {
	return o
}

func (o TopicIAMMemberPtrOutput) Elem() TopicIAMMemberOutput {
	return o.ApplyT(func(v *TopicIAMMember) TopicIAMMember {
		if v != nil {
			return *v
		}
		var ret TopicIAMMember
		return ret
	}).(TopicIAMMemberOutput)
}

type TopicIAMMemberArrayOutput struct{ *pulumi.OutputState }

func (TopicIAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicIAMMember)(nil))
}

func (o TopicIAMMemberArrayOutput) ToTopicIAMMemberArrayOutput() TopicIAMMemberArrayOutput {
	return o
}

func (o TopicIAMMemberArrayOutput) ToTopicIAMMemberArrayOutputWithContext(ctx context.Context) TopicIAMMemberArrayOutput {
	return o
}

func (o TopicIAMMemberArrayOutput) Index(i pulumi.IntInput) TopicIAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicIAMMember {
		return vs[0].([]TopicIAMMember)[vs[1].(int)]
	}).(TopicIAMMemberOutput)
}

type TopicIAMMemberMapOutput struct{ *pulumi.OutputState }

func (TopicIAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TopicIAMMember)(nil))
}

func (o TopicIAMMemberMapOutput) ToTopicIAMMemberMapOutput() TopicIAMMemberMapOutput {
	return o
}

func (o TopicIAMMemberMapOutput) ToTopicIAMMemberMapOutputWithContext(ctx context.Context) TopicIAMMemberMapOutput {
	return o
}

func (o TopicIAMMemberMapOutput) MapIndex(k pulumi.StringInput) TopicIAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TopicIAMMember {
		return vs[0].(map[string]TopicIAMMember)[vs[1].(string)]
	}).(TopicIAMMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicIAMMemberOutput{})
	pulumi.RegisterOutputType(TopicIAMMemberPtrOutput{})
	pulumi.RegisterOutputType(TopicIAMMemberArrayOutput{})
	pulumi.RegisterOutputType(TopicIAMMemberMapOutput{})
}
