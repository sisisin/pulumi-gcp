// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Certificate Authority Service CaPool. Each of these resources serves a different use case:
//
// * `certificateauthority.CaPoolIamPolicy`: Authoritative. Sets the IAM policy for the capool and replaces any existing policy already attached.
// * `certificateauthority.CaPoolIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the capool are preserved.
// * `certificateauthority.CaPoolIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the capool are preserved.
//
// > **Note:** `certificateauthority.CaPoolIamPolicy` **cannot** be used in conjunction with `certificateauthority.CaPoolIamBinding` and `certificateauthority.CaPoolIamMember` or they will fight over what your policy should be.
//
// > **Note:** `certificateauthority.CaPoolIamBinding` resources **can be** used in conjunction with `certificateauthority.CaPoolIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_privateca\_ca\_pool\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/privateca.certificateManager",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = certificateauthority.NewCaPoolIamPolicy(ctx, "policy", &certificateauthority.CaPoolIamPolicyArgs{
// 			CaPool:     pulumi.Any(google_privateca_ca_pool.Default.Id),
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
// ## google\_privateca\_ca\_pool\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := certificateauthority.NewCaPoolIamBinding(ctx, "binding", &certificateauthority.CaPoolIamBindingArgs{
// 			CaPool: pulumi.Any(google_privateca_ca_pool.Default.Id),
// 			Role:   pulumi.String("roles/privateca.certificateManager"),
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
// ## google\_privateca\_ca\_pool\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := certificateauthority.NewCaPoolIamMember(ctx, "member", &certificateauthority.CaPoolIamMemberArgs{
// 			CaPool: pulumi.Any(google_privateca_ca_pool.Default.Id),
// 			Role:   pulumi.String("roles/privateca.certificateManager"),
// 			Member: pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/caPools/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} Any variables not passed in the import command will be taken from the provider configuration. Certificate Authority Service capool IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding editor projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type CaPoolIamBinding struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CaPool    pulumi.StringOutput                `pulumi:"caPool"`
	Condition CaPoolIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewCaPoolIamBinding registers a new resource with the given unique name, arguments, and options.
func NewCaPoolIamBinding(ctx *pulumi.Context,
	name string, args *CaPoolIamBindingArgs, opts ...pulumi.ResourceOption) (*CaPoolIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPool == nil {
		return nil, errors.New("invalid value for required argument 'CaPool'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource CaPoolIamBinding
	err := ctx.RegisterResource("gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaPoolIamBinding gets an existing CaPoolIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaPoolIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaPoolIamBindingState, opts ...pulumi.ResourceOption) (*CaPoolIamBinding, error) {
	var resource CaPoolIamBinding
	err := ctx.ReadResource("gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaPoolIamBinding resources.
type caPoolIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool    *string                    `pulumi:"caPool"`
	Condition *CaPoolIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type CaPoolIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool    pulumi.StringPtrInput
	Condition CaPoolIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (CaPoolIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamBindingState)(nil)).Elem()
}

type caPoolIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool    string                     `pulumi:"caPool"`
	Condition *CaPoolIamBindingCondition `pulumi:"condition"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a CaPoolIamBinding resource.
type CaPoolIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool    pulumi.StringInput
	Condition CaPoolIamBindingConditionPtrInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (CaPoolIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamBindingArgs)(nil)).Elem()
}

type CaPoolIamBindingInput interface {
	pulumi.Input

	ToCaPoolIamBindingOutput() CaPoolIamBindingOutput
	ToCaPoolIamBindingOutputWithContext(ctx context.Context) CaPoolIamBindingOutput
}

func (*CaPoolIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*CaPoolIamBinding)(nil))
}

func (i *CaPoolIamBinding) ToCaPoolIamBindingOutput() CaPoolIamBindingOutput {
	return i.ToCaPoolIamBindingOutputWithContext(context.Background())
}

func (i *CaPoolIamBinding) ToCaPoolIamBindingOutputWithContext(ctx context.Context) CaPoolIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamBindingOutput)
}

func (i *CaPoolIamBinding) ToCaPoolIamBindingPtrOutput() CaPoolIamBindingPtrOutput {
	return i.ToCaPoolIamBindingPtrOutputWithContext(context.Background())
}

func (i *CaPoolIamBinding) ToCaPoolIamBindingPtrOutputWithContext(ctx context.Context) CaPoolIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamBindingPtrOutput)
}

type CaPoolIamBindingPtrInput interface {
	pulumi.Input

	ToCaPoolIamBindingPtrOutput() CaPoolIamBindingPtrOutput
	ToCaPoolIamBindingPtrOutputWithContext(ctx context.Context) CaPoolIamBindingPtrOutput
}

type caPoolIamBindingPtrType CaPoolIamBindingArgs

func (*caPoolIamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPoolIamBinding)(nil))
}

func (i *caPoolIamBindingPtrType) ToCaPoolIamBindingPtrOutput() CaPoolIamBindingPtrOutput {
	return i.ToCaPoolIamBindingPtrOutputWithContext(context.Background())
}

func (i *caPoolIamBindingPtrType) ToCaPoolIamBindingPtrOutputWithContext(ctx context.Context) CaPoolIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamBindingPtrOutput)
}

// CaPoolIamBindingArrayInput is an input type that accepts CaPoolIamBindingArray and CaPoolIamBindingArrayOutput values.
// You can construct a concrete instance of `CaPoolIamBindingArrayInput` via:
//
//          CaPoolIamBindingArray{ CaPoolIamBindingArgs{...} }
type CaPoolIamBindingArrayInput interface {
	pulumi.Input

	ToCaPoolIamBindingArrayOutput() CaPoolIamBindingArrayOutput
	ToCaPoolIamBindingArrayOutputWithContext(context.Context) CaPoolIamBindingArrayOutput
}

type CaPoolIamBindingArray []CaPoolIamBindingInput

func (CaPoolIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPoolIamBinding)(nil)).Elem()
}

func (i CaPoolIamBindingArray) ToCaPoolIamBindingArrayOutput() CaPoolIamBindingArrayOutput {
	return i.ToCaPoolIamBindingArrayOutputWithContext(context.Background())
}

func (i CaPoolIamBindingArray) ToCaPoolIamBindingArrayOutputWithContext(ctx context.Context) CaPoolIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamBindingArrayOutput)
}

// CaPoolIamBindingMapInput is an input type that accepts CaPoolIamBindingMap and CaPoolIamBindingMapOutput values.
// You can construct a concrete instance of `CaPoolIamBindingMapInput` via:
//
//          CaPoolIamBindingMap{ "key": CaPoolIamBindingArgs{...} }
type CaPoolIamBindingMapInput interface {
	pulumi.Input

	ToCaPoolIamBindingMapOutput() CaPoolIamBindingMapOutput
	ToCaPoolIamBindingMapOutputWithContext(context.Context) CaPoolIamBindingMapOutput
}

type CaPoolIamBindingMap map[string]CaPoolIamBindingInput

func (CaPoolIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPoolIamBinding)(nil)).Elem()
}

func (i CaPoolIamBindingMap) ToCaPoolIamBindingMapOutput() CaPoolIamBindingMapOutput {
	return i.ToCaPoolIamBindingMapOutputWithContext(context.Background())
}

func (i CaPoolIamBindingMap) ToCaPoolIamBindingMapOutputWithContext(ctx context.Context) CaPoolIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamBindingMapOutput)
}

type CaPoolIamBindingOutput struct{ *pulumi.OutputState }

func (CaPoolIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaPoolIamBinding)(nil))
}

func (o CaPoolIamBindingOutput) ToCaPoolIamBindingOutput() CaPoolIamBindingOutput {
	return o
}

func (o CaPoolIamBindingOutput) ToCaPoolIamBindingOutputWithContext(ctx context.Context) CaPoolIamBindingOutput {
	return o
}

func (o CaPoolIamBindingOutput) ToCaPoolIamBindingPtrOutput() CaPoolIamBindingPtrOutput {
	return o.ToCaPoolIamBindingPtrOutputWithContext(context.Background())
}

func (o CaPoolIamBindingOutput) ToCaPoolIamBindingPtrOutputWithContext(ctx context.Context) CaPoolIamBindingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CaPoolIamBinding) *CaPoolIamBinding {
		return &v
	}).(CaPoolIamBindingPtrOutput)
}

type CaPoolIamBindingPtrOutput struct{ *pulumi.OutputState }

func (CaPoolIamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPoolIamBinding)(nil))
}

func (o CaPoolIamBindingPtrOutput) ToCaPoolIamBindingPtrOutput() CaPoolIamBindingPtrOutput {
	return o
}

func (o CaPoolIamBindingPtrOutput) ToCaPoolIamBindingPtrOutputWithContext(ctx context.Context) CaPoolIamBindingPtrOutput {
	return o
}

func (o CaPoolIamBindingPtrOutput) Elem() CaPoolIamBindingOutput {
	return o.ApplyT(func(v *CaPoolIamBinding) CaPoolIamBinding {
		if v != nil {
			return *v
		}
		var ret CaPoolIamBinding
		return ret
	}).(CaPoolIamBindingOutput)
}

type CaPoolIamBindingArrayOutput struct{ *pulumi.OutputState }

func (CaPoolIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CaPoolIamBinding)(nil))
}

func (o CaPoolIamBindingArrayOutput) ToCaPoolIamBindingArrayOutput() CaPoolIamBindingArrayOutput {
	return o
}

func (o CaPoolIamBindingArrayOutput) ToCaPoolIamBindingArrayOutputWithContext(ctx context.Context) CaPoolIamBindingArrayOutput {
	return o
}

func (o CaPoolIamBindingArrayOutput) Index(i pulumi.IntInput) CaPoolIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CaPoolIamBinding {
		return vs[0].([]CaPoolIamBinding)[vs[1].(int)]
	}).(CaPoolIamBindingOutput)
}

type CaPoolIamBindingMapOutput struct{ *pulumi.OutputState }

func (CaPoolIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CaPoolIamBinding)(nil))
}

func (o CaPoolIamBindingMapOutput) ToCaPoolIamBindingMapOutput() CaPoolIamBindingMapOutput {
	return o
}

func (o CaPoolIamBindingMapOutput) ToCaPoolIamBindingMapOutputWithContext(ctx context.Context) CaPoolIamBindingMapOutput {
	return o
}

func (o CaPoolIamBindingMapOutput) MapIndex(k pulumi.StringInput) CaPoolIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CaPoolIamBinding {
		return vs[0].(map[string]CaPoolIamBinding)[vs[1].(string)]
	}).(CaPoolIamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(CaPoolIamBindingOutput{})
	pulumi.RegisterOutputType(CaPoolIamBindingPtrOutput{})
	pulumi.RegisterOutputType(CaPoolIamBindingArrayOutput{})
	pulumi.RegisterOutputType(CaPoolIamBindingMapOutput{})
}