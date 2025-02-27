// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package firebase

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Google Cloud Firebase web application instance
//
// To get more information about WebApp, see:
//
// * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps)
// * How-to Guides
//     * [Official Documentation](https://firebase.google.com/)
//
// ## Example Usage
//
// ## Import
//
// WebApp can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:firebase/webApp:WebApp default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:firebase/webApp:WebApp default {{project}} {{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:firebase/webApp:WebApp default {{name}}
// ```
type WebApp struct {
	pulumi.CustomResourceState

	// Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The user-assigned display name of the App.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewWebApp registers a new resource with the given unique name, arguments, and options.
func NewWebApp(ctx *pulumi.Context,
	name string, args *WebAppArgs, opts ...pulumi.ResourceOption) (*WebApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource WebApp
	err := ctx.RegisterResource("gcp:firebase/webApp:WebApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebApp gets an existing WebApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppState, opts ...pulumi.ResourceOption) (*WebApp, error) {
	var resource WebApp
	err := ctx.ReadResource("gcp:firebase/webApp:WebApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebApp resources.
type webAppState struct {
	// Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId *string `pulumi:"appId"`
	// The user-assigned display name of the App.
	DisplayName *string `pulumi:"displayName"`
	// The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type WebAppState struct {
	// Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId pulumi.StringPtrInput
	// The user-assigned display name of the App.
	DisplayName pulumi.StringPtrInput
	// The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppState)(nil)).Elem()
}

type webAppArgs struct {
	// The user-assigned display name of the App.
	DisplayName string `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a WebApp resource.
type WebAppArgs struct {
	// The user-assigned display name of the App.
	DisplayName pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppArgs)(nil)).Elem()
}

type WebAppInput interface {
	pulumi.Input

	ToWebAppOutput() WebAppOutput
	ToWebAppOutputWithContext(ctx context.Context) WebAppOutput
}

func (*WebApp) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApp)(nil))
}

func (i *WebApp) ToWebAppOutput() WebAppOutput {
	return i.ToWebAppOutputWithContext(context.Background())
}

func (i *WebApp) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppOutput)
}

func (i *WebApp) ToWebAppPtrOutput() WebAppPtrOutput {
	return i.ToWebAppPtrOutputWithContext(context.Background())
}

func (i *WebApp) ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPtrOutput)
}

type WebAppPtrInput interface {
	pulumi.Input

	ToWebAppPtrOutput() WebAppPtrOutput
	ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput
}

type webAppPtrType WebAppArgs

func (*webAppPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApp)(nil))
}

func (i *webAppPtrType) ToWebAppPtrOutput() WebAppPtrOutput {
	return i.ToWebAppPtrOutputWithContext(context.Background())
}

func (i *webAppPtrType) ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPtrOutput)
}

// WebAppArrayInput is an input type that accepts WebAppArray and WebAppArrayOutput values.
// You can construct a concrete instance of `WebAppArrayInput` via:
//
//          WebAppArray{ WebAppArgs{...} }
type WebAppArrayInput interface {
	pulumi.Input

	ToWebAppArrayOutput() WebAppArrayOutput
	ToWebAppArrayOutputWithContext(context.Context) WebAppArrayOutput
}

type WebAppArray []WebAppInput

func (WebAppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebApp)(nil)).Elem()
}

func (i WebAppArray) ToWebAppArrayOutput() WebAppArrayOutput {
	return i.ToWebAppArrayOutputWithContext(context.Background())
}

func (i WebAppArray) ToWebAppArrayOutputWithContext(ctx context.Context) WebAppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppArrayOutput)
}

// WebAppMapInput is an input type that accepts WebAppMap and WebAppMapOutput values.
// You can construct a concrete instance of `WebAppMapInput` via:
//
//          WebAppMap{ "key": WebAppArgs{...} }
type WebAppMapInput interface {
	pulumi.Input

	ToWebAppMapOutput() WebAppMapOutput
	ToWebAppMapOutputWithContext(context.Context) WebAppMapOutput
}

type WebAppMap map[string]WebAppInput

func (WebAppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebApp)(nil)).Elem()
}

func (i WebAppMap) ToWebAppMapOutput() WebAppMapOutput {
	return i.ToWebAppMapOutputWithContext(context.Background())
}

func (i WebAppMap) ToWebAppMapOutputWithContext(ctx context.Context) WebAppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppMapOutput)
}

type WebAppOutput struct{ *pulumi.OutputState }

func (WebAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApp)(nil))
}

func (o WebAppOutput) ToWebAppOutput() WebAppOutput {
	return o
}

func (o WebAppOutput) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return o
}

func (o WebAppOutput) ToWebAppPtrOutput() WebAppPtrOutput {
	return o.ToWebAppPtrOutputWithContext(context.Background())
}

func (o WebAppOutput) ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebApp) *WebApp {
		return &v
	}).(WebAppPtrOutput)
}

type WebAppPtrOutput struct{ *pulumi.OutputState }

func (WebAppPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApp)(nil))
}

func (o WebAppPtrOutput) ToWebAppPtrOutput() WebAppPtrOutput {
	return o
}

func (o WebAppPtrOutput) ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput {
	return o
}

func (o WebAppPtrOutput) Elem() WebAppOutput {
	return o.ApplyT(func(v *WebApp) WebApp {
		if v != nil {
			return *v
		}
		var ret WebApp
		return ret
	}).(WebAppOutput)
}

type WebAppArrayOutput struct{ *pulumi.OutputState }

func (WebAppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApp)(nil))
}

func (o WebAppArrayOutput) ToWebAppArrayOutput() WebAppArrayOutput {
	return o
}

func (o WebAppArrayOutput) ToWebAppArrayOutputWithContext(ctx context.Context) WebAppArrayOutput {
	return o
}

func (o WebAppArrayOutput) Index(i pulumi.IntInput) WebAppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebApp {
		return vs[0].([]WebApp)[vs[1].(int)]
	}).(WebAppOutput)
}

type WebAppMapOutput struct{ *pulumi.OutputState }

func (WebAppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebApp)(nil))
}

func (o WebAppMapOutput) ToWebAppMapOutput() WebAppMapOutput {
	return o
}

func (o WebAppMapOutput) ToWebAppMapOutputWithContext(ctx context.Context) WebAppMapOutput {
	return o
}

func (o WebAppMapOutput) MapIndex(k pulumi.StringInput) WebAppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebApp {
		return vs[0].(map[string]WebApp)[vs[1].(string)]
	}).(WebAppOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppOutput{})
	pulumi.RegisterOutputType(WebAppPtrOutput{})
	pulumi.RegisterOutputType(WebAppArrayOutput{})
	pulumi.RegisterOutputType(WebAppMapOutput{})
}
