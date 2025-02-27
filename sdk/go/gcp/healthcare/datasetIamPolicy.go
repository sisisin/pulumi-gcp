// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare dataset. Each of these resources serves a different use case:
//
// * `healthcare.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
// * `healthcare.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
// * `healthcare.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
//
// > **Note:** `healthcare.DatasetIamPolicy` **cannot** be used in conjunction with `healthcare.DatasetIamBinding` and `healthcare.DatasetIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.DatasetIamBinding` resources **can be** used in conjunction with `healthcare.DatasetIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_healthcare\_dataset\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewDatasetIamPolicy(ctx, "dataset", &healthcare.DatasetIamPolicyArgs{
// 			DatasetId:  pulumi.String("your-dataset-id"),
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
// ## google\_healthcare\_dataset\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewDatasetIamBinding(ctx, "dataset", &healthcare.DatasetIamBindingArgs{
// 			DatasetId: pulumi.String("your-dataset-id"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_healthcare\_dataset\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewDatasetIamMember(ctx, "dataset", &healthcare.DatasetIamMemberArgs{
// 			DatasetId: pulumi.String("your-dataset-id"),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Role:      pulumi.String("roles/editor"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `dataset_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/datasetIamPolicy:DatasetIamPolicy dataset_iam "your-project-id/location-name/dataset-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `dataset_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/datasetIamPolicy:DatasetIamPolicy dataset_iam "your-project-id/location-name/dataset-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `dataset_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/datasetIamPolicy:DatasetIamPolicy dataset_iam your-project-id/location-name/dataset-name
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type DatasetIamPolicy struct {
	pulumi.CustomResourceState

	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewDatasetIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamPolicy(ctx *pulumi.Context,
	name string, args *DatasetIamPolicyArgs, opts ...pulumi.ResourceOption) (*DatasetIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource DatasetIamPolicy
	err := ctx.RegisterResource("gcp:healthcare/datasetIamPolicy:DatasetIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetIamPolicy gets an existing DatasetIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetIamPolicyState, opts ...pulumi.ResourceOption) (*DatasetIamPolicy, error) {
	var resource DatasetIamPolicy
	err := ctx.ReadResource("gcp:healthcare/datasetIamPolicy:DatasetIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamPolicy resources.
type datasetIamPolicyState struct {
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId *string `pulumi:"datasetId"`
	// (Computed) The etag of the dataset's IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type DatasetIamPolicyState struct {
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the dataset's IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (DatasetIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamPolicyState)(nil)).Elem()
}

type datasetIamPolicyArgs struct {
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId string `pulumi:"datasetId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a DatasetIamPolicy resource.
type DatasetIamPolicyArgs struct {
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (DatasetIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamPolicyArgs)(nil)).Elem()
}

type DatasetIamPolicyInput interface {
	pulumi.Input

	ToDatasetIamPolicyOutput() DatasetIamPolicyOutput
	ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput
}

func (*DatasetIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamPolicy)(nil))
}

func (i *DatasetIamPolicy) ToDatasetIamPolicyOutput() DatasetIamPolicyOutput {
	return i.ToDatasetIamPolicyOutputWithContext(context.Background())
}

func (i *DatasetIamPolicy) ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyOutput)
}

func (i *DatasetIamPolicy) ToDatasetIamPolicyPtrOutput() DatasetIamPolicyPtrOutput {
	return i.ToDatasetIamPolicyPtrOutputWithContext(context.Background())
}

func (i *DatasetIamPolicy) ToDatasetIamPolicyPtrOutputWithContext(ctx context.Context) DatasetIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyPtrOutput)
}

type DatasetIamPolicyPtrInput interface {
	pulumi.Input

	ToDatasetIamPolicyPtrOutput() DatasetIamPolicyPtrOutput
	ToDatasetIamPolicyPtrOutputWithContext(ctx context.Context) DatasetIamPolicyPtrOutput
}

type datasetIamPolicyPtrType DatasetIamPolicyArgs

func (*datasetIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamPolicy)(nil))
}

func (i *datasetIamPolicyPtrType) ToDatasetIamPolicyPtrOutput() DatasetIamPolicyPtrOutput {
	return i.ToDatasetIamPolicyPtrOutputWithContext(context.Background())
}

func (i *datasetIamPolicyPtrType) ToDatasetIamPolicyPtrOutputWithContext(ctx context.Context) DatasetIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyPtrOutput)
}

// DatasetIamPolicyArrayInput is an input type that accepts DatasetIamPolicyArray and DatasetIamPolicyArrayOutput values.
// You can construct a concrete instance of `DatasetIamPolicyArrayInput` via:
//
//          DatasetIamPolicyArray{ DatasetIamPolicyArgs{...} }
type DatasetIamPolicyArrayInput interface {
	pulumi.Input

	ToDatasetIamPolicyArrayOutput() DatasetIamPolicyArrayOutput
	ToDatasetIamPolicyArrayOutputWithContext(context.Context) DatasetIamPolicyArrayOutput
}

type DatasetIamPolicyArray []DatasetIamPolicyInput

func (DatasetIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetIamPolicy)(nil)).Elem()
}

func (i DatasetIamPolicyArray) ToDatasetIamPolicyArrayOutput() DatasetIamPolicyArrayOutput {
	return i.ToDatasetIamPolicyArrayOutputWithContext(context.Background())
}

func (i DatasetIamPolicyArray) ToDatasetIamPolicyArrayOutputWithContext(ctx context.Context) DatasetIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyArrayOutput)
}

// DatasetIamPolicyMapInput is an input type that accepts DatasetIamPolicyMap and DatasetIamPolicyMapOutput values.
// You can construct a concrete instance of `DatasetIamPolicyMapInput` via:
//
//          DatasetIamPolicyMap{ "key": DatasetIamPolicyArgs{...} }
type DatasetIamPolicyMapInput interface {
	pulumi.Input

	ToDatasetIamPolicyMapOutput() DatasetIamPolicyMapOutput
	ToDatasetIamPolicyMapOutputWithContext(context.Context) DatasetIamPolicyMapOutput
}

type DatasetIamPolicyMap map[string]DatasetIamPolicyInput

func (DatasetIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetIamPolicy)(nil)).Elem()
}

func (i DatasetIamPolicyMap) ToDatasetIamPolicyMapOutput() DatasetIamPolicyMapOutput {
	return i.ToDatasetIamPolicyMapOutputWithContext(context.Background())
}

func (i DatasetIamPolicyMap) ToDatasetIamPolicyMapOutputWithContext(ctx context.Context) DatasetIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyMapOutput)
}

type DatasetIamPolicyOutput struct{ *pulumi.OutputState }

func (DatasetIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamPolicy)(nil))
}

func (o DatasetIamPolicyOutput) ToDatasetIamPolicyOutput() DatasetIamPolicyOutput {
	return o
}

func (o DatasetIamPolicyOutput) ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput {
	return o
}

func (o DatasetIamPolicyOutput) ToDatasetIamPolicyPtrOutput() DatasetIamPolicyPtrOutput {
	return o.ToDatasetIamPolicyPtrOutputWithContext(context.Background())
}

func (o DatasetIamPolicyOutput) ToDatasetIamPolicyPtrOutputWithContext(ctx context.Context) DatasetIamPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetIamPolicy) *DatasetIamPolicy {
		return &v
	}).(DatasetIamPolicyPtrOutput)
}

type DatasetIamPolicyPtrOutput struct{ *pulumi.OutputState }

func (DatasetIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetIamPolicy)(nil))
}

func (o DatasetIamPolicyPtrOutput) ToDatasetIamPolicyPtrOutput() DatasetIamPolicyPtrOutput {
	return o
}

func (o DatasetIamPolicyPtrOutput) ToDatasetIamPolicyPtrOutputWithContext(ctx context.Context) DatasetIamPolicyPtrOutput {
	return o
}

func (o DatasetIamPolicyPtrOutput) Elem() DatasetIamPolicyOutput {
	return o.ApplyT(func(v *DatasetIamPolicy) DatasetIamPolicy {
		if v != nil {
			return *v
		}
		var ret DatasetIamPolicy
		return ret
	}).(DatasetIamPolicyOutput)
}

type DatasetIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (DatasetIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetIamPolicy)(nil))
}

func (o DatasetIamPolicyArrayOutput) ToDatasetIamPolicyArrayOutput() DatasetIamPolicyArrayOutput {
	return o
}

func (o DatasetIamPolicyArrayOutput) ToDatasetIamPolicyArrayOutputWithContext(ctx context.Context) DatasetIamPolicyArrayOutput {
	return o
}

func (o DatasetIamPolicyArrayOutput) Index(i pulumi.IntInput) DatasetIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatasetIamPolicy {
		return vs[0].([]DatasetIamPolicy)[vs[1].(int)]
	}).(DatasetIamPolicyOutput)
}

type DatasetIamPolicyMapOutput struct{ *pulumi.OutputState }

func (DatasetIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatasetIamPolicy)(nil))
}

func (o DatasetIamPolicyMapOutput) ToDatasetIamPolicyMapOutput() DatasetIamPolicyMapOutput {
	return o
}

func (o DatasetIamPolicyMapOutput) ToDatasetIamPolicyMapOutputWithContext(ctx context.Context) DatasetIamPolicyMapOutput {
	return o
}

func (o DatasetIamPolicyMapOutput) MapIndex(k pulumi.StringInput) DatasetIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatasetIamPolicy {
		return vs[0].(map[string]DatasetIamPolicy)[vs[1].(string)]
	}).(DatasetIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasetIamPolicyOutput{})
	pulumi.RegisterOutputType(DatasetIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(DatasetIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(DatasetIamPolicyMapOutput{})
}
