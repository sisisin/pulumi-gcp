// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Service is a discrete, autonomous, and network-accessible unit,
// designed to solve an individual concern (Wikipedia). In Cloud Monitoring,
// a Service acts as the root resource under which operational aspects of
// the service are accessible
//
// To get more information about Service, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
// * How-to Guides
//     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
//     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
//
// ## Example Usage
// ### Monitoring Service Custom
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/monitoring"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := monitoring.NewCustomService(ctx, "custom", &monitoring.CustomServiceArgs{
// 			DisplayName: pulumi.String("My Custom Service custom-srv"),
// 			ServiceId:   pulumi.String("custom-srv"),
// 			Telemetry: &monitoring.CustomServiceTelemetryArgs{
// 				ResourceName: pulumi.String("//product.googleapis.com/foo/foo/services/test"),
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
// Service can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:monitoring/customService:CustomService default {{name}}
// ```
type CustomService struct {
	pulumi.CustomResourceState

	// Name used for UI elements listing this Service.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry CustomServiceTelemetryPtrOutput `pulumi:"telemetry"`
}

// NewCustomService registers a new resource with the given unique name, arguments, and options.
func NewCustomService(ctx *pulumi.Context,
	name string, args *CustomServiceArgs, opts ...pulumi.ResourceOption) (*CustomService, error) {
	if args == nil {
		args = &CustomServiceArgs{}
	}

	var resource CustomService
	err := ctx.RegisterResource("gcp:monitoring/customService:CustomService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomService gets an existing CustomService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomServiceState, opts ...pulumi.ResourceOption) (*CustomService, error) {
	var resource CustomService
	err := ctx.ReadResource("gcp:monitoring/customService:CustomService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomService resources.
type customServiceState struct {
	// Name used for UI elements listing this Service.
	DisplayName *string `pulumi:"displayName"`
	// The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId *string `pulumi:"serviceId"`
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry *CustomServiceTelemetry `pulumi:"telemetry"`
}

type CustomServiceState struct {
	// Name used for UI elements listing this Service.
	DisplayName pulumi.StringPtrInput
	// The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId pulumi.StringPtrInput
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry CustomServiceTelemetryPtrInput
}

func (CustomServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*customServiceState)(nil)).Elem()
}

type customServiceArgs struct {
	// Name used for UI elements listing this Service.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId *string `pulumi:"serviceId"`
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry *CustomServiceTelemetry `pulumi:"telemetry"`
}

// The set of arguments for constructing a CustomService resource.
type CustomServiceArgs struct {
	// Name used for UI elements listing this Service.
	DisplayName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId pulumi.StringPtrInput
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry CustomServiceTelemetryPtrInput
}

func (CustomServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customServiceArgs)(nil)).Elem()
}

type CustomServiceInput interface {
	pulumi.Input

	ToCustomServiceOutput() CustomServiceOutput
	ToCustomServiceOutputWithContext(ctx context.Context) CustomServiceOutput
}

func (*CustomService) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomService)(nil))
}

func (i *CustomService) ToCustomServiceOutput() CustomServiceOutput {
	return i.ToCustomServiceOutputWithContext(context.Background())
}

func (i *CustomService) ToCustomServiceOutputWithContext(ctx context.Context) CustomServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomServiceOutput)
}

func (i *CustomService) ToCustomServicePtrOutput() CustomServicePtrOutput {
	return i.ToCustomServicePtrOutputWithContext(context.Background())
}

func (i *CustomService) ToCustomServicePtrOutputWithContext(ctx context.Context) CustomServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomServicePtrOutput)
}

type CustomServicePtrInput interface {
	pulumi.Input

	ToCustomServicePtrOutput() CustomServicePtrOutput
	ToCustomServicePtrOutputWithContext(ctx context.Context) CustomServicePtrOutput
}

type customServicePtrType CustomServiceArgs

func (*customServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomService)(nil))
}

func (i *customServicePtrType) ToCustomServicePtrOutput() CustomServicePtrOutput {
	return i.ToCustomServicePtrOutputWithContext(context.Background())
}

func (i *customServicePtrType) ToCustomServicePtrOutputWithContext(ctx context.Context) CustomServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomServicePtrOutput)
}

// CustomServiceArrayInput is an input type that accepts CustomServiceArray and CustomServiceArrayOutput values.
// You can construct a concrete instance of `CustomServiceArrayInput` via:
//
//          CustomServiceArray{ CustomServiceArgs{...} }
type CustomServiceArrayInput interface {
	pulumi.Input

	ToCustomServiceArrayOutput() CustomServiceArrayOutput
	ToCustomServiceArrayOutputWithContext(context.Context) CustomServiceArrayOutput
}

type CustomServiceArray []CustomServiceInput

func (CustomServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomService)(nil)).Elem()
}

func (i CustomServiceArray) ToCustomServiceArrayOutput() CustomServiceArrayOutput {
	return i.ToCustomServiceArrayOutputWithContext(context.Background())
}

func (i CustomServiceArray) ToCustomServiceArrayOutputWithContext(ctx context.Context) CustomServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomServiceArrayOutput)
}

// CustomServiceMapInput is an input type that accepts CustomServiceMap and CustomServiceMapOutput values.
// You can construct a concrete instance of `CustomServiceMapInput` via:
//
//          CustomServiceMap{ "key": CustomServiceArgs{...} }
type CustomServiceMapInput interface {
	pulumi.Input

	ToCustomServiceMapOutput() CustomServiceMapOutput
	ToCustomServiceMapOutputWithContext(context.Context) CustomServiceMapOutput
}

type CustomServiceMap map[string]CustomServiceInput

func (CustomServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomService)(nil)).Elem()
}

func (i CustomServiceMap) ToCustomServiceMapOutput() CustomServiceMapOutput {
	return i.ToCustomServiceMapOutputWithContext(context.Background())
}

func (i CustomServiceMap) ToCustomServiceMapOutputWithContext(ctx context.Context) CustomServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomServiceMapOutput)
}

type CustomServiceOutput struct{ *pulumi.OutputState }

func (CustomServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomService)(nil))
}

func (o CustomServiceOutput) ToCustomServiceOutput() CustomServiceOutput {
	return o
}

func (o CustomServiceOutput) ToCustomServiceOutputWithContext(ctx context.Context) CustomServiceOutput {
	return o
}

func (o CustomServiceOutput) ToCustomServicePtrOutput() CustomServicePtrOutput {
	return o.ToCustomServicePtrOutputWithContext(context.Background())
}

func (o CustomServiceOutput) ToCustomServicePtrOutputWithContext(ctx context.Context) CustomServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomService) *CustomService {
		return &v
	}).(CustomServicePtrOutput)
}

type CustomServicePtrOutput struct{ *pulumi.OutputState }

func (CustomServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomService)(nil))
}

func (o CustomServicePtrOutput) ToCustomServicePtrOutput() CustomServicePtrOutput {
	return o
}

func (o CustomServicePtrOutput) ToCustomServicePtrOutputWithContext(ctx context.Context) CustomServicePtrOutput {
	return o
}

func (o CustomServicePtrOutput) Elem() CustomServiceOutput {
	return o.ApplyT(func(v *CustomService) CustomService {
		if v != nil {
			return *v
		}
		var ret CustomService
		return ret
	}).(CustomServiceOutput)
}

type CustomServiceArrayOutput struct{ *pulumi.OutputState }

func (CustomServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomService)(nil))
}

func (o CustomServiceArrayOutput) ToCustomServiceArrayOutput() CustomServiceArrayOutput {
	return o
}

func (o CustomServiceArrayOutput) ToCustomServiceArrayOutputWithContext(ctx context.Context) CustomServiceArrayOutput {
	return o
}

func (o CustomServiceArrayOutput) Index(i pulumi.IntInput) CustomServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomService {
		return vs[0].([]CustomService)[vs[1].(int)]
	}).(CustomServiceOutput)
}

type CustomServiceMapOutput struct{ *pulumi.OutputState }

func (CustomServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomService)(nil))
}

func (o CustomServiceMapOutput) ToCustomServiceMapOutput() CustomServiceMapOutput {
	return o
}

func (o CustomServiceMapOutput) ToCustomServiceMapOutputWithContext(ctx context.Context) CustomServiceMapOutput {
	return o
}

func (o CustomServiceMapOutput) MapIndex(k pulumi.StringInput) CustomServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomService {
		return vs[0].(map[string]CustomService)[vs[1].(string)]
	}).(CustomServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomServiceOutput{})
	pulumi.RegisterOutputType(CustomServicePtrOutput{})
	pulumi.RegisterOutputType(CustomServiceArrayOutput{})
	pulumi.RegisterOutputType(CustomServiceMapOutput{})
}
