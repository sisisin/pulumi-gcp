// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Machine Image resource. Machine images store all the configuration,
// metadata, permissions, and data from one or more disks required to create a
// Virtual machine (VM) instance.
//
// To get more information about MachineImage, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/machineImages)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/machine-images)
//
// ## Example Usage
// ### Machine Image Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		vm, err := compute.NewInstance(ctx, "vm", &compute.InstanceArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String("debian-cloud/debian-9"),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewMachineImage(ctx, "image", &compute.MachineImageArgs{
// 			SourceInstance: vm.SelfLink,
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Compute Machine Image Kms
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/kms"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		vm, err := compute.NewInstance(ctx, "vm", &compute.InstanceArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String("debian-cloud/debian-9"),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		keyRing, err := kms.NewKeyRing(ctx, "keyRing", &kms.KeyRingArgs{
// 			Location: pulumi.String("us"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		cryptoKey, err := kms.NewCryptoKey(ctx, "cryptoKey", &kms.CryptoKeyArgs{
// 			KeyRing: keyRing.ID(),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMMember(ctx, "kms_project_binding", &projects.IAMMemberArgs{
// 			Project: pulumi.String(project.ProjectId),
// 			Role:    pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
// 			Member:  pulumi.String(fmt.Sprintf("%v%v%v", "serviceAccount:service-", project.Number, "@compute-system.iam.gserviceaccount.com")),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewMachineImage(ctx, "image", &compute.MachineImageArgs{
// 			SourceInstance: vm.SelfLink,
// 			MachineImageEncryptionKey: &compute.MachineImageMachineImageEncryptionKeyArgs{
// 				KmsKeyName: cryptoKey.ID(),
// 			},
// 		}, pulumi.Provider(google_beta), pulumi.DependsOn([]pulumi.Resource{
// 			kms_project_binding,
// 		}))
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
// MachineImage can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/machineImage:MachineImage default projects/{{project}}/global/machineImages/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/machineImage:MachineImage default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/machineImage:MachineImage default {{name}}
// ```
type MachineImage struct {
	pulumi.CustomResourceState

	// A text description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrOutput `pulumi:"guestFlush"`
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey MachineImageMachineImageEncryptionKeyPtrOutput `pulumi:"machineImageEncryptionKey"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	SourceInstance pulumi.StringOutput `pulumi:"sourceInstance"`
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
}

// NewMachineImage registers a new resource with the given unique name, arguments, and options.
func NewMachineImage(ctx *pulumi.Context,
	name string, args *MachineImageArgs, opts ...pulumi.ResourceOption) (*MachineImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceInstance == nil {
		return nil, errors.New("invalid value for required argument 'SourceInstance'")
	}
	var resource MachineImage
	err := ctx.RegisterResource("gcp:compute/machineImage:MachineImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineImage gets an existing MachineImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineImageState, opts ...pulumi.ResourceOption) (*MachineImage, error) {
	var resource MachineImage
	err := ctx.ReadResource("gcp:compute/machineImage:MachineImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineImage resources.
type machineImageState struct {
	// A text description of the resource.
	Description *string `pulumi:"description"`
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush *bool `pulumi:"guestFlush"`
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey *MachineImageMachineImageEncryptionKey `pulumi:"machineImageEncryptionKey"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	SourceInstance *string `pulumi:"sourceInstance"`
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations []string `pulumi:"storageLocations"`
}

type MachineImageState struct {
	// A text description of the resource.
	Description pulumi.StringPtrInput
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrInput
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey MachineImageMachineImageEncryptionKeyPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	SourceInstance pulumi.StringPtrInput
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations pulumi.StringArrayInput
}

func (MachineImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageState)(nil)).Elem()
}

type machineImageArgs struct {
	// A text description of the resource.
	Description *string `pulumi:"description"`
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush *bool `pulumi:"guestFlush"`
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey *MachineImageMachineImageEncryptionKey `pulumi:"machineImageEncryptionKey"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	SourceInstance string `pulumi:"sourceInstance"`
}

// The set of arguments for constructing a MachineImage resource.
type MachineImageArgs struct {
	// A text description of the resource.
	Description pulumi.StringPtrInput
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush pulumi.BoolPtrInput
	// Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image)
	// Structure is documented below.
	MachineImageEncryptionKey MachineImageMachineImageEncryptionKeyPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
	SourceInstance pulumi.StringInput
}

func (MachineImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageArgs)(nil)).Elem()
}

type MachineImageInput interface {
	pulumi.Input

	ToMachineImageOutput() MachineImageOutput
	ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput
}

func (*MachineImage) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineImage)(nil))
}

func (i *MachineImage) ToMachineImageOutput() MachineImageOutput {
	return i.ToMachineImageOutputWithContext(context.Background())
}

func (i *MachineImage) ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageOutput)
}

func (i *MachineImage) ToMachineImagePtrOutput() MachineImagePtrOutput {
	return i.ToMachineImagePtrOutputWithContext(context.Background())
}

func (i *MachineImage) ToMachineImagePtrOutputWithContext(ctx context.Context) MachineImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImagePtrOutput)
}

type MachineImagePtrInput interface {
	pulumi.Input

	ToMachineImagePtrOutput() MachineImagePtrOutput
	ToMachineImagePtrOutputWithContext(ctx context.Context) MachineImagePtrOutput
}

type machineImagePtrType MachineImageArgs

func (*machineImagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImage)(nil))
}

func (i *machineImagePtrType) ToMachineImagePtrOutput() MachineImagePtrOutput {
	return i.ToMachineImagePtrOutputWithContext(context.Background())
}

func (i *machineImagePtrType) ToMachineImagePtrOutputWithContext(ctx context.Context) MachineImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImagePtrOutput)
}

// MachineImageArrayInput is an input type that accepts MachineImageArray and MachineImageArrayOutput values.
// You can construct a concrete instance of `MachineImageArrayInput` via:
//
//          MachineImageArray{ MachineImageArgs{...} }
type MachineImageArrayInput interface {
	pulumi.Input

	ToMachineImageArrayOutput() MachineImageArrayOutput
	ToMachineImageArrayOutputWithContext(context.Context) MachineImageArrayOutput
}

type MachineImageArray []MachineImageInput

func (MachineImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineImage)(nil)).Elem()
}

func (i MachineImageArray) ToMachineImageArrayOutput() MachineImageArrayOutput {
	return i.ToMachineImageArrayOutputWithContext(context.Background())
}

func (i MachineImageArray) ToMachineImageArrayOutputWithContext(ctx context.Context) MachineImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageArrayOutput)
}

// MachineImageMapInput is an input type that accepts MachineImageMap and MachineImageMapOutput values.
// You can construct a concrete instance of `MachineImageMapInput` via:
//
//          MachineImageMap{ "key": MachineImageArgs{...} }
type MachineImageMapInput interface {
	pulumi.Input

	ToMachineImageMapOutput() MachineImageMapOutput
	ToMachineImageMapOutputWithContext(context.Context) MachineImageMapOutput
}

type MachineImageMap map[string]MachineImageInput

func (MachineImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineImage)(nil)).Elem()
}

func (i MachineImageMap) ToMachineImageMapOutput() MachineImageMapOutput {
	return i.ToMachineImageMapOutputWithContext(context.Background())
}

func (i MachineImageMap) ToMachineImageMapOutputWithContext(ctx context.Context) MachineImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageMapOutput)
}

type MachineImageOutput struct{ *pulumi.OutputState }

func (MachineImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineImage)(nil))
}

func (o MachineImageOutput) ToMachineImageOutput() MachineImageOutput {
	return o
}

func (o MachineImageOutput) ToMachineImageOutputWithContext(ctx context.Context) MachineImageOutput {
	return o
}

func (o MachineImageOutput) ToMachineImagePtrOutput() MachineImagePtrOutput {
	return o.ToMachineImagePtrOutputWithContext(context.Background())
}

func (o MachineImageOutput) ToMachineImagePtrOutputWithContext(ctx context.Context) MachineImagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineImage) *MachineImage {
		return &v
	}).(MachineImagePtrOutput)
}

type MachineImagePtrOutput struct{ *pulumi.OutputState }

func (MachineImagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineImage)(nil))
}

func (o MachineImagePtrOutput) ToMachineImagePtrOutput() MachineImagePtrOutput {
	return o
}

func (o MachineImagePtrOutput) ToMachineImagePtrOutputWithContext(ctx context.Context) MachineImagePtrOutput {
	return o
}

func (o MachineImagePtrOutput) Elem() MachineImageOutput {
	return o.ApplyT(func(v *MachineImage) MachineImage {
		if v != nil {
			return *v
		}
		var ret MachineImage
		return ret
	}).(MachineImageOutput)
}

type MachineImageArrayOutput struct{ *pulumi.OutputState }

func (MachineImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineImage)(nil))
}

func (o MachineImageArrayOutput) ToMachineImageArrayOutput() MachineImageArrayOutput {
	return o
}

func (o MachineImageArrayOutput) ToMachineImageArrayOutputWithContext(ctx context.Context) MachineImageArrayOutput {
	return o
}

func (o MachineImageArrayOutput) Index(i pulumi.IntInput) MachineImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineImage {
		return vs[0].([]MachineImage)[vs[1].(int)]
	}).(MachineImageOutput)
}

type MachineImageMapOutput struct{ *pulumi.OutputState }

func (MachineImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MachineImage)(nil))
}

func (o MachineImageMapOutput) ToMachineImageMapOutput() MachineImageMapOutput {
	return o
}

func (o MachineImageMapOutput) ToMachineImageMapOutputWithContext(ctx context.Context) MachineImageMapOutput {
	return o
}

func (o MachineImageMapOutput) MapIndex(k pulumi.StringInput) MachineImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MachineImage {
		return vs[0].(map[string]MachineImage)[vs[1].(string)]
	}).(MachineImageOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineImageOutput{})
	pulumi.RegisterOutputType(MachineImagePtrOutput{})
	pulumi.RegisterOutputType(MachineImageArrayOutput{})
	pulumi.RegisterOutputType(MachineImageMapOutput{})
}
