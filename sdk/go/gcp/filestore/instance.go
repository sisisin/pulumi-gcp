// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package filestore

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Google Cloud Filestore instance.
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
//     * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
//     * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)
//
// ## Example Usage
// ### Filestore Instance Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/filestore"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := filestore.NewInstance(ctx, "instance", &filestore.InstanceArgs{
// 			FileShares: &filestore.InstanceFileSharesArgs{
// 				CapacityGb: pulumi.Int(2660),
// 				Name:       pulumi.String("share1"),
// 			},
// 			Networks: filestore.InstanceNetworkArray{
// 				&filestore.InstanceNetworkArgs{
// 					Modes: pulumi.StringArray{
// 						pulumi.String("MODE_IPV4"),
// 					},
// 					Network: pulumi.String("default"),
// 				},
// 			},
// 			Tier: pulumi.String("PREMIUM"),
// 			Zone: pulumi.String("us-central1-b"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filestore Instance Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/filestore"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := filestore.NewInstance(ctx, "instance", &filestore.InstanceArgs{
// 			Zone: pulumi.String("us-central1-b"),
// 			Tier: pulumi.String("BASIC_SSD"),
// 			FileShares: &filestore.InstanceFileSharesArgs{
// 				CapacityGb: pulumi.Int(2660),
// 				Name:       pulumi.String("share1"),
// 				NfsExportOptions: filestore.InstanceFileSharesNfsExportOptionArray{
// 					&filestore.InstanceFileSharesNfsExportOptionArgs{
// 						IpRanges: pulumi.StringArray{
// 							pulumi.String("10.0.0.0/24"),
// 						},
// 						AccessMode: pulumi.String("READ_WRITE"),
// 						SquashMode: pulumi.String("NO_ROOT_SQUASH"),
// 					},
// 					&filestore.InstanceFileSharesNfsExportOptionArgs{
// 						IpRanges: pulumi.StringArray{
// 							pulumi.String("10.10.0.0/24"),
// 						},
// 						AccessMode: pulumi.String("READ_ONLY"),
// 						SquashMode: pulumi.String("ROOT_SQUASH"),
// 						AnonUid:    pulumi.Int(123),
// 						AnonGid:    pulumi.Int(456),
// 					},
// 				},
// 			},
// 			Networks: filestore.InstanceNetworkArray{
// 				&filestore.InstanceNetworkArgs{
// 					Network: pulumi.String("default"),
// 					Modes: pulumi.StringArray{
// 						pulumi.String("MODE_IPV4"),
// 					},
// 					ConnectMode: pulumi.String("DIRECT_PEERING"),
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
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
// Instance can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:filestore/instance:Instance default projects/{{project}}/locations/{{zone}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:filestore/instance:Instance default {{project}}/{{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:filestore/instance:Instance default {{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:filestore/instance:Instance default {{name}}
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A description of the instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares InstanceFileSharesOutput `pulumi:"fileShares"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the fileshare (16 characters or less)
	Name pulumi.StringOutput `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks InstanceNetworkArrayOutput `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier pulumi.StringOutput `pulumi:"tier"`
	// The name of the Filestore zone of the instance.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileShares == nil {
		return nil, errors.New("invalid value for required argument 'FileShares'")
	}
	if args.Networks == nil {
		return nil, errors.New("invalid value for required argument 'Networks'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:filestore/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:filestore/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// A description of the instance.
	Description *string `pulumi:"description"`
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag *string `pulumi:"etag"`
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares *InstanceFileShares `pulumi:"fileShares"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The name of the fileshare (16 characters or less)
	Name *string `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks []InstanceNetwork `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier *string `pulumi:"tier"`
	// The name of the Filestore zone of the instance.
	Zone *string `pulumi:"zone"`
}

type InstanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// A description of the instance.
	Description pulumi.StringPtrInput
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag pulumi.StringPtrInput
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares InstanceFileSharesPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// The name of the fileshare (16 characters or less)
	Name pulumi.StringPtrInput
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks InstanceNetworkArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier pulumi.StringPtrInput
	// The name of the Filestore zone of the instance.
	Zone pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// A description of the instance.
	Description *string `pulumi:"description"`
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares InstanceFileShares `pulumi:"fileShares"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The name of the fileshare (16 characters or less)
	Name *string `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks []InstanceNetwork `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier string `pulumi:"tier"`
	// The name of the Filestore zone of the instance.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// A description of the instance.
	Description pulumi.StringPtrInput
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares InstanceFileSharesInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// The name of the fileshare (16 characters or less)
	Name pulumi.StringPtrInput
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks InstanceNetworkArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service tier of the instance.
	// Possible values are `TIER_UNSPECIFIED`, `STANDARD`, `PREMIUM`, `BASIC_HDD`, `BASIC_SSD`, and `HIGH_SCALE_SSD`.
	Tier pulumi.StringInput
	// The name of the Filestore zone of the instance.
	Zone pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *Instance) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

type InstancePtrInput interface {
	pulumi.Input

	ToInstancePtrOutput() InstancePtrOutput
	ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput
}

type instancePtrType InstanceArgs

func (*instancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (i *instancePtrType) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *instancePtrType) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o.ToInstancePtrOutputWithContext(context.Background())
}

func (o InstanceOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Instance) *Instance {
		return &v
	}).(InstancePtrOutput)
}

type InstancePtrOutput struct{ *pulumi.OutputState }

func (InstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (o InstancePtrOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) Elem() InstanceOutput {
	return o.ApplyT(func(v *Instance) Instance {
		if v != nil {
			return *v
		}
		var ret Instance
		return ret
	}).(InstanceOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Instance)(nil))
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Instance {
		return vs[0].([]Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Instance)(nil))
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Instance {
		return vs[0].(map[string]Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstancePtrOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
