// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a TargetSslProxy resource, which is used by one or more
// global forwarding rule to route incoming SSL requests to a backend
// service.
//
// To get more information about TargetSslProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetSslProxies)
// * How-to Guides
//     * [Setting Up SSL proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/)
//
// ## Example Usage
// ### Target Ssl Proxy Basic
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultSSLCertificate, err := compute.NewSSLCertificate(ctx, "defaultSSLCertificate", &compute.SSLCertificateArgs{
// 			PrivateKey:  readFileOrPanic("path/to/private.key"),
// 			Certificate: readFileOrPanic("path/to/certificate.crt"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultHealthCheck, err := compute.NewHealthCheck(ctx, "defaultHealthCheck", &compute.HealthCheckArgs{
// 			CheckIntervalSec: pulumi.Int(1),
// 			TimeoutSec:       pulumi.Int(1),
// 			TcpHealthCheck: &compute.HealthCheckTcpHealthCheckArgs{
// 				Port: pulumi.Int(443),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultBackendService, err := compute.NewBackendService(ctx, "defaultBackendService", &compute.BackendServiceArgs{
// 			Protocol: pulumi.String("SSL"),
// 			HealthChecks: pulumi.String{
// 				defaultHealthCheck.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewTargetSSLProxy(ctx, "defaultTargetSSLProxy", &compute.TargetSSLProxyArgs{
// 			BackendService: defaultBackendService.ID(),
// 			SslCertificates: pulumi.StringArray{
// 				defaultSSLCertificate.ID(),
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
// TargetSslProxy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/targetSSLProxy:TargetSSLProxy default projects/{{project}}/global/targetSslProxies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetSSLProxy:TargetSSLProxy default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetSSLProxy:TargetSSLProxy default {{name}}
// ```
type TargetSSLProxy struct {
	pulumi.CustomResourceState

	// A reference to the BackendService resource.
	BackendService pulumi.StringOutput `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrOutput `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId pulumi.IntOutput `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one
	// SSL certificate must be specified.
	SslCertificates pulumi.StringArrayOutput `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetSslProxy resource. If not set, the TargetSslProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrOutput `pulumi:"sslPolicy"`
}

// NewTargetSSLProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetSSLProxy(ctx *pulumi.Context,
	name string, args *TargetSSLProxyArgs, opts ...pulumi.ResourceOption) (*TargetSSLProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendService == nil {
		return nil, errors.New("invalid value for required argument 'BackendService'")
	}
	if args.SslCertificates == nil {
		return nil, errors.New("invalid value for required argument 'SslCertificates'")
	}
	var resource TargetSSLProxy
	err := ctx.RegisterResource("gcp:compute/targetSSLProxy:TargetSSLProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetSSLProxy gets an existing TargetSSLProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetSSLProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetSSLProxyState, opts ...pulumi.ResourceOption) (*TargetSSLProxy, error) {
	var resource TargetSSLProxy
	err := ctx.ReadResource("gcp:compute/targetSSLProxy:TargetSSLProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetSSLProxy resources.
type targetSSLProxyState struct {
	// A reference to the BackendService resource.
	BackendService *string `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader *string `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId *int `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one
	// SSL certificate must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetSslProxy resource. If not set, the TargetSslProxy
	// resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
}

type TargetSSLProxyState struct {
	// A reference to the BackendService resource.
	BackendService pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrInput
	// The unique identifier for the resource.
	ProxyId pulumi.IntPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one
	// SSL certificate must be specified.
	SslCertificates pulumi.StringArrayInput
	// A reference to the SslPolicy resource that will be associated with
	// the TargetSslProxy resource. If not set, the TargetSslProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
}

func (TargetSSLProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetSSLProxyState)(nil)).Elem()
}

type targetSSLProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService string `pulumi:"backendService"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader *string `pulumi:"proxyHeader"`
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one
	// SSL certificate must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetSslProxy resource. If not set, the TargetSslProxy
	// resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
}

// The set of arguments for constructing a TargetSSLProxy resource.
type TargetSSLProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService pulumi.StringInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrInput
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one
	// SSL certificate must be specified.
	SslCertificates pulumi.StringArrayInput
	// A reference to the SslPolicy resource that will be associated with
	// the TargetSslProxy resource. If not set, the TargetSslProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
}

func (TargetSSLProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetSSLProxyArgs)(nil)).Elem()
}

type TargetSSLProxyInput interface {
	pulumi.Input

	ToTargetSSLProxyOutput() TargetSSLProxyOutput
	ToTargetSSLProxyOutputWithContext(ctx context.Context) TargetSSLProxyOutput
}

func (*TargetSSLProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetSSLProxy)(nil))
}

func (i *TargetSSLProxy) ToTargetSSLProxyOutput() TargetSSLProxyOutput {
	return i.ToTargetSSLProxyOutputWithContext(context.Background())
}

func (i *TargetSSLProxy) ToTargetSSLProxyOutputWithContext(ctx context.Context) TargetSSLProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetSSLProxyOutput)
}

func (i *TargetSSLProxy) ToTargetSSLProxyPtrOutput() TargetSSLProxyPtrOutput {
	return i.ToTargetSSLProxyPtrOutputWithContext(context.Background())
}

func (i *TargetSSLProxy) ToTargetSSLProxyPtrOutputWithContext(ctx context.Context) TargetSSLProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetSSLProxyPtrOutput)
}

type TargetSSLProxyPtrInput interface {
	pulumi.Input

	ToTargetSSLProxyPtrOutput() TargetSSLProxyPtrOutput
	ToTargetSSLProxyPtrOutputWithContext(ctx context.Context) TargetSSLProxyPtrOutput
}

type targetSSLProxyPtrType TargetSSLProxyArgs

func (*targetSSLProxyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetSSLProxy)(nil))
}

func (i *targetSSLProxyPtrType) ToTargetSSLProxyPtrOutput() TargetSSLProxyPtrOutput {
	return i.ToTargetSSLProxyPtrOutputWithContext(context.Background())
}

func (i *targetSSLProxyPtrType) ToTargetSSLProxyPtrOutputWithContext(ctx context.Context) TargetSSLProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetSSLProxyPtrOutput)
}

// TargetSSLProxyArrayInput is an input type that accepts TargetSSLProxyArray and TargetSSLProxyArrayOutput values.
// You can construct a concrete instance of `TargetSSLProxyArrayInput` via:
//
//          TargetSSLProxyArray{ TargetSSLProxyArgs{...} }
type TargetSSLProxyArrayInput interface {
	pulumi.Input

	ToTargetSSLProxyArrayOutput() TargetSSLProxyArrayOutput
	ToTargetSSLProxyArrayOutputWithContext(context.Context) TargetSSLProxyArrayOutput
}

type TargetSSLProxyArray []TargetSSLProxyInput

func (TargetSSLProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetSSLProxy)(nil)).Elem()
}

func (i TargetSSLProxyArray) ToTargetSSLProxyArrayOutput() TargetSSLProxyArrayOutput {
	return i.ToTargetSSLProxyArrayOutputWithContext(context.Background())
}

func (i TargetSSLProxyArray) ToTargetSSLProxyArrayOutputWithContext(ctx context.Context) TargetSSLProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetSSLProxyArrayOutput)
}

// TargetSSLProxyMapInput is an input type that accepts TargetSSLProxyMap and TargetSSLProxyMapOutput values.
// You can construct a concrete instance of `TargetSSLProxyMapInput` via:
//
//          TargetSSLProxyMap{ "key": TargetSSLProxyArgs{...} }
type TargetSSLProxyMapInput interface {
	pulumi.Input

	ToTargetSSLProxyMapOutput() TargetSSLProxyMapOutput
	ToTargetSSLProxyMapOutputWithContext(context.Context) TargetSSLProxyMapOutput
}

type TargetSSLProxyMap map[string]TargetSSLProxyInput

func (TargetSSLProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetSSLProxy)(nil)).Elem()
}

func (i TargetSSLProxyMap) ToTargetSSLProxyMapOutput() TargetSSLProxyMapOutput {
	return i.ToTargetSSLProxyMapOutputWithContext(context.Background())
}

func (i TargetSSLProxyMap) ToTargetSSLProxyMapOutputWithContext(ctx context.Context) TargetSSLProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetSSLProxyMapOutput)
}

type TargetSSLProxyOutput struct{ *pulumi.OutputState }

func (TargetSSLProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetSSLProxy)(nil))
}

func (o TargetSSLProxyOutput) ToTargetSSLProxyOutput() TargetSSLProxyOutput {
	return o
}

func (o TargetSSLProxyOutput) ToTargetSSLProxyOutputWithContext(ctx context.Context) TargetSSLProxyOutput {
	return o
}

func (o TargetSSLProxyOutput) ToTargetSSLProxyPtrOutput() TargetSSLProxyPtrOutput {
	return o.ToTargetSSLProxyPtrOutputWithContext(context.Background())
}

func (o TargetSSLProxyOutput) ToTargetSSLProxyPtrOutputWithContext(ctx context.Context) TargetSSLProxyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TargetSSLProxy) *TargetSSLProxy {
		return &v
	}).(TargetSSLProxyPtrOutput)
}

type TargetSSLProxyPtrOutput struct{ *pulumi.OutputState }

func (TargetSSLProxyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetSSLProxy)(nil))
}

func (o TargetSSLProxyPtrOutput) ToTargetSSLProxyPtrOutput() TargetSSLProxyPtrOutput {
	return o
}

func (o TargetSSLProxyPtrOutput) ToTargetSSLProxyPtrOutputWithContext(ctx context.Context) TargetSSLProxyPtrOutput {
	return o
}

func (o TargetSSLProxyPtrOutput) Elem() TargetSSLProxyOutput {
	return o.ApplyT(func(v *TargetSSLProxy) TargetSSLProxy {
		if v != nil {
			return *v
		}
		var ret TargetSSLProxy
		return ret
	}).(TargetSSLProxyOutput)
}

type TargetSSLProxyArrayOutput struct{ *pulumi.OutputState }

func (TargetSSLProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetSSLProxy)(nil))
}

func (o TargetSSLProxyArrayOutput) ToTargetSSLProxyArrayOutput() TargetSSLProxyArrayOutput {
	return o
}

func (o TargetSSLProxyArrayOutput) ToTargetSSLProxyArrayOutputWithContext(ctx context.Context) TargetSSLProxyArrayOutput {
	return o
}

func (o TargetSSLProxyArrayOutput) Index(i pulumi.IntInput) TargetSSLProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetSSLProxy {
		return vs[0].([]TargetSSLProxy)[vs[1].(int)]
	}).(TargetSSLProxyOutput)
}

type TargetSSLProxyMapOutput struct{ *pulumi.OutputState }

func (TargetSSLProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TargetSSLProxy)(nil))
}

func (o TargetSSLProxyMapOutput) ToTargetSSLProxyMapOutput() TargetSSLProxyMapOutput {
	return o
}

func (o TargetSSLProxyMapOutput) ToTargetSSLProxyMapOutputWithContext(ctx context.Context) TargetSSLProxyMapOutput {
	return o
}

func (o TargetSSLProxyMapOutput) MapIndex(k pulumi.StringInput) TargetSSLProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TargetSSLProxy {
		return vs[0].(map[string]TargetSSLProxy)[vs[1].(string)]
	}).(TargetSSLProxyOutput)
}

func init() {
	pulumi.RegisterOutputType(TargetSSLProxyOutput{})
	pulumi.RegisterOutputType(TargetSSLProxyPtrOutput{})
	pulumi.RegisterOutputType(TargetSSLProxyArrayOutput{})
	pulumi.RegisterOutputType(TargetSSLProxyMapOutput{})
}
