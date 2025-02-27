// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy TunnelInstance. Each of these resources serves a different use case:
//
// * `iap.TunnelInstanceIAMPolicy`: Authoritative. Sets the IAM policy for the tunnelinstance and replaces any existing policy already attached.
// * `iap.TunnelInstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunnelinstance are preserved.
// * `iap.TunnelInstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunnelinstance are preserved.
//
// > **Note:** `iap.TunnelInstanceIAMPolicy` **cannot** be used in conjunction with `iap.TunnelInstanceIAMBinding` and `iap.TunnelInstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.TunnelInstanceIAMBinding` resources **can be** used in conjunction with `iap.TunnelInstanceIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_tunnel\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.tunnelResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewTunnelInstanceIAMPolicy(ctx, "policy", &iap.TunnelInstanceIAMPolicyArgs{
// 			Project:    pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:       pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance:   pulumi.Any(google_compute_instance.Tunnelvm.Name),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.tunnelResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewTunnelInstanceIAMPolicy(ctx, "policy", &iap.TunnelInstanceIAMPolicyArgs{
// 			Project:    pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:       pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance:   pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_tunnel\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewTunnelInstanceIAMBinding(ctx, "binding", &iap.TunnelInstanceIAMBindingArgs{
// 			Project:  pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:     pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance: pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewTunnelInstanceIAMBinding(ctx, "binding", &iap.TunnelInstanceIAMBindingArgs{
// 			Project:  pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:     pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance: pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &iap.TunnelInstanceIAMBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_tunnel\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewTunnelInstanceIAMMember(ctx, "member", &iap.TunnelInstanceIAMMemberArgs{
// 			Project:  pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:     pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance: pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewTunnelInstanceIAMMember(ctx, "member", &iap.TunnelInstanceIAMMemberArgs{
// 			Project:  pulumi.Any(google_compute_instance.Tunnelvm.Project),
// 			Zone:     pulumi.Any(google_compute_instance.Tunnelvm.Zone),
// 			Instance: pulumi.Any(google_compute_instance.Tunnelvm.Name),
// 			Role:     pulumi.String("roles/iap.tunnelResourceAccessor"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 			Condition: &iap.TunnelInstanceIAMMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{name}} * projects/{{project}}/zones/{{zone}}/instances/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy tunnelinstance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor "projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor "projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TunnelInstanceIAMMember struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringOutput `pulumi:"instance"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTunnelInstanceIAMMember registers a new resource with the given unique name, arguments, and options.
func NewTunnelInstanceIAMMember(ctx *pulumi.Context,
	name string, args *TunnelInstanceIAMMemberArgs, opts ...pulumi.ResourceOption) (*TunnelInstanceIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource TunnelInstanceIAMMember
	err := ctx.RegisterResource("gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTunnelInstanceIAMMember gets an existing TunnelInstanceIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTunnelInstanceIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TunnelInstanceIAMMemberState, opts ...pulumi.ResourceOption) (*TunnelInstanceIAMMember, error) {
	var resource TunnelInstanceIAMMember
	err := ctx.ReadResource("gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TunnelInstanceIAMMember resources.
type tunnelInstanceIAMMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *TunnelInstanceIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Instance *string `pulumi:"instance"`
	Member   *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	Zone *string `pulumi:"zone"`
}

type TunnelInstanceIAMMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	Zone pulumi.StringPtrInput
}

func (TunnelInstanceIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelInstanceIAMMemberState)(nil)).Elem()
}

type tunnelInstanceIAMMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *TunnelInstanceIAMMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Instance string `pulumi:"instance"`
	Member   string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string  `pulumi:"role"`
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a TunnelInstanceIAMMember resource.
type TunnelInstanceIAMMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition TunnelInstanceIAMMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Instance pulumi.StringInput
	Member   pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	Zone pulumi.StringPtrInput
}

func (TunnelInstanceIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelInstanceIAMMemberArgs)(nil)).Elem()
}

type TunnelInstanceIAMMemberInput interface {
	pulumi.Input

	ToTunnelInstanceIAMMemberOutput() TunnelInstanceIAMMemberOutput
	ToTunnelInstanceIAMMemberOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberOutput
}

func (*TunnelInstanceIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((*TunnelInstanceIAMMember)(nil))
}

func (i *TunnelInstanceIAMMember) ToTunnelInstanceIAMMemberOutput() TunnelInstanceIAMMemberOutput {
	return i.ToTunnelInstanceIAMMemberOutputWithContext(context.Background())
}

func (i *TunnelInstanceIAMMember) ToTunnelInstanceIAMMemberOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMMemberOutput)
}

func (i *TunnelInstanceIAMMember) ToTunnelInstanceIAMMemberPtrOutput() TunnelInstanceIAMMemberPtrOutput {
	return i.ToTunnelInstanceIAMMemberPtrOutputWithContext(context.Background())
}

func (i *TunnelInstanceIAMMember) ToTunnelInstanceIAMMemberPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMMemberPtrOutput)
}

type TunnelInstanceIAMMemberPtrInput interface {
	pulumi.Input

	ToTunnelInstanceIAMMemberPtrOutput() TunnelInstanceIAMMemberPtrOutput
	ToTunnelInstanceIAMMemberPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberPtrOutput
}

type tunnelInstanceIAMMemberPtrType TunnelInstanceIAMMemberArgs

func (*tunnelInstanceIAMMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelInstanceIAMMember)(nil))
}

func (i *tunnelInstanceIAMMemberPtrType) ToTunnelInstanceIAMMemberPtrOutput() TunnelInstanceIAMMemberPtrOutput {
	return i.ToTunnelInstanceIAMMemberPtrOutputWithContext(context.Background())
}

func (i *tunnelInstanceIAMMemberPtrType) ToTunnelInstanceIAMMemberPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMMemberPtrOutput)
}

// TunnelInstanceIAMMemberArrayInput is an input type that accepts TunnelInstanceIAMMemberArray and TunnelInstanceIAMMemberArrayOutput values.
// You can construct a concrete instance of `TunnelInstanceIAMMemberArrayInput` via:
//
//          TunnelInstanceIAMMemberArray{ TunnelInstanceIAMMemberArgs{...} }
type TunnelInstanceIAMMemberArrayInput interface {
	pulumi.Input

	ToTunnelInstanceIAMMemberArrayOutput() TunnelInstanceIAMMemberArrayOutput
	ToTunnelInstanceIAMMemberArrayOutputWithContext(context.Context) TunnelInstanceIAMMemberArrayOutput
}

type TunnelInstanceIAMMemberArray []TunnelInstanceIAMMemberInput

func (TunnelInstanceIAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TunnelInstanceIAMMember)(nil)).Elem()
}

func (i TunnelInstanceIAMMemberArray) ToTunnelInstanceIAMMemberArrayOutput() TunnelInstanceIAMMemberArrayOutput {
	return i.ToTunnelInstanceIAMMemberArrayOutputWithContext(context.Background())
}

func (i TunnelInstanceIAMMemberArray) ToTunnelInstanceIAMMemberArrayOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMMemberArrayOutput)
}

// TunnelInstanceIAMMemberMapInput is an input type that accepts TunnelInstanceIAMMemberMap and TunnelInstanceIAMMemberMapOutput values.
// You can construct a concrete instance of `TunnelInstanceIAMMemberMapInput` via:
//
//          TunnelInstanceIAMMemberMap{ "key": TunnelInstanceIAMMemberArgs{...} }
type TunnelInstanceIAMMemberMapInput interface {
	pulumi.Input

	ToTunnelInstanceIAMMemberMapOutput() TunnelInstanceIAMMemberMapOutput
	ToTunnelInstanceIAMMemberMapOutputWithContext(context.Context) TunnelInstanceIAMMemberMapOutput
}

type TunnelInstanceIAMMemberMap map[string]TunnelInstanceIAMMemberInput

func (TunnelInstanceIAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TunnelInstanceIAMMember)(nil)).Elem()
}

func (i TunnelInstanceIAMMemberMap) ToTunnelInstanceIAMMemberMapOutput() TunnelInstanceIAMMemberMapOutput {
	return i.ToTunnelInstanceIAMMemberMapOutputWithContext(context.Background())
}

func (i TunnelInstanceIAMMemberMap) ToTunnelInstanceIAMMemberMapOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelInstanceIAMMemberMapOutput)
}

type TunnelInstanceIAMMemberOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TunnelInstanceIAMMember)(nil))
}

func (o TunnelInstanceIAMMemberOutput) ToTunnelInstanceIAMMemberOutput() TunnelInstanceIAMMemberOutput {
	return o
}

func (o TunnelInstanceIAMMemberOutput) ToTunnelInstanceIAMMemberOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberOutput {
	return o
}

func (o TunnelInstanceIAMMemberOutput) ToTunnelInstanceIAMMemberPtrOutput() TunnelInstanceIAMMemberPtrOutput {
	return o.ToTunnelInstanceIAMMemberPtrOutputWithContext(context.Background())
}

func (o TunnelInstanceIAMMemberOutput) ToTunnelInstanceIAMMemberPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TunnelInstanceIAMMember) *TunnelInstanceIAMMember {
		return &v
	}).(TunnelInstanceIAMMemberPtrOutput)
}

type TunnelInstanceIAMMemberPtrOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelInstanceIAMMember)(nil))
}

func (o TunnelInstanceIAMMemberPtrOutput) ToTunnelInstanceIAMMemberPtrOutput() TunnelInstanceIAMMemberPtrOutput {
	return o
}

func (o TunnelInstanceIAMMemberPtrOutput) ToTunnelInstanceIAMMemberPtrOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberPtrOutput {
	return o
}

func (o TunnelInstanceIAMMemberPtrOutput) Elem() TunnelInstanceIAMMemberOutput {
	return o.ApplyT(func(v *TunnelInstanceIAMMember) TunnelInstanceIAMMember {
		if v != nil {
			return *v
		}
		var ret TunnelInstanceIAMMember
		return ret
	}).(TunnelInstanceIAMMemberOutput)
}

type TunnelInstanceIAMMemberArrayOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TunnelInstanceIAMMember)(nil))
}

func (o TunnelInstanceIAMMemberArrayOutput) ToTunnelInstanceIAMMemberArrayOutput() TunnelInstanceIAMMemberArrayOutput {
	return o
}

func (o TunnelInstanceIAMMemberArrayOutput) ToTunnelInstanceIAMMemberArrayOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberArrayOutput {
	return o
}

func (o TunnelInstanceIAMMemberArrayOutput) Index(i pulumi.IntInput) TunnelInstanceIAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TunnelInstanceIAMMember {
		return vs[0].([]TunnelInstanceIAMMember)[vs[1].(int)]
	}).(TunnelInstanceIAMMemberOutput)
}

type TunnelInstanceIAMMemberMapOutput struct{ *pulumi.OutputState }

func (TunnelInstanceIAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TunnelInstanceIAMMember)(nil))
}

func (o TunnelInstanceIAMMemberMapOutput) ToTunnelInstanceIAMMemberMapOutput() TunnelInstanceIAMMemberMapOutput {
	return o
}

func (o TunnelInstanceIAMMemberMapOutput) ToTunnelInstanceIAMMemberMapOutputWithContext(ctx context.Context) TunnelInstanceIAMMemberMapOutput {
	return o
}

func (o TunnelInstanceIAMMemberMapOutput) MapIndex(k pulumi.StringInput) TunnelInstanceIAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TunnelInstanceIAMMember {
		return vs[0].(map[string]TunnelInstanceIAMMember)[vs[1].(string)]
	}).(TunnelInstanceIAMMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(TunnelInstanceIAMMemberOutput{})
	pulumi.RegisterOutputType(TunnelInstanceIAMMemberPtrOutput{})
	pulumi.RegisterOutputType(TunnelInstanceIAMMemberArrayOutput{})
	pulumi.RegisterOutputType(TunnelInstanceIAMMemberMapOutput{})
}
