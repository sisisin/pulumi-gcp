// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identityplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Inbound SAML configuration for a Identity Toolkit tenant.
//
// You must enable the
// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
// the marketplace prior to using this resource.
//
// ## Example Usage
// ### Identity Platform Tenant Inbound Saml Config Basic
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/identityplatform"
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
// 		tenant, err := identityplatform.NewTenant(ctx, "tenant", &identityplatform.TenantArgs{
// 			DisplayName: pulumi.String("tenant"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = identityplatform.NewTenantInboundSamlConfig(ctx, "tenantSamlConfig", &identityplatform.TenantInboundSamlConfigArgs{
// 			DisplayName: pulumi.String("Display Name"),
// 			Tenant:      tenant.Name,
// 			IdpConfig: &identityplatform.TenantInboundSamlConfigIdpConfigArgs{
// 				IdpEntityId: pulumi.String("tf-idp"),
// 				SignRequest: pulumi.Bool(true),
// 				SsoUrl:      pulumi.String("https://example.com"),
// 				IdpCertificates: identityplatform.TenantInboundSamlConfigIdpConfigIdpCertificateArray{
// 					&identityplatform.TenantInboundSamlConfigIdpConfigIdpCertificateArgs{
// 						X509Certificate: readFileOrPanic("test-fixtures/rsa_cert.pem"),
// 					},
// 				},
// 			},
// 			SpConfig: &identityplatform.TenantInboundSamlConfigSpConfigArgs{
// 				SpEntityId:  pulumi.String("tf-sp"),
// 				CallbackUri: pulumi.String("https://example.com"),
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
// TenantInboundSamlConfig can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default {{project}}/{{tenant}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default {{tenant}}/{{name}}
// ```
type TenantInboundSamlConfig struct {
	pulumi.CustomResourceState

	// Human friendly display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig TenantInboundSamlConfigIdpConfigOutput `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig TenantInboundSamlConfigSpConfigOutput `pulumi:"spConfig"`
	// The name of the tenant where this inbound SAML config resource exists
	Tenant pulumi.StringOutput `pulumi:"tenant"`
}

// NewTenantInboundSamlConfig registers a new resource with the given unique name, arguments, and options.
func NewTenantInboundSamlConfig(ctx *pulumi.Context,
	name string, args *TenantInboundSamlConfigArgs, opts ...pulumi.ResourceOption) (*TenantInboundSamlConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IdpConfig == nil {
		return nil, errors.New("invalid value for required argument 'IdpConfig'")
	}
	if args.SpConfig == nil {
		return nil, errors.New("invalid value for required argument 'SpConfig'")
	}
	if args.Tenant == nil {
		return nil, errors.New("invalid value for required argument 'Tenant'")
	}
	var resource TenantInboundSamlConfig
	err := ctx.RegisterResource("gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenantInboundSamlConfig gets an existing TenantInboundSamlConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenantInboundSamlConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantInboundSamlConfigState, opts ...pulumi.ResourceOption) (*TenantInboundSamlConfig, error) {
	var resource TenantInboundSamlConfig
	err := ctx.ReadResource("gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TenantInboundSamlConfig resources.
type tenantInboundSamlConfigState struct {
	// Human friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig *TenantInboundSamlConfigIdpConfig `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig *TenantInboundSamlConfigSpConfig `pulumi:"spConfig"`
	// The name of the tenant where this inbound SAML config resource exists
	Tenant *string `pulumi:"tenant"`
}

type TenantInboundSamlConfigState struct {
	// Human friendly display name.
	DisplayName pulumi.StringPtrInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig TenantInboundSamlConfigIdpConfigPtrInput
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig TenantInboundSamlConfigSpConfigPtrInput
	// The name of the tenant where this inbound SAML config resource exists
	Tenant pulumi.StringPtrInput
}

func (TenantInboundSamlConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantInboundSamlConfigState)(nil)).Elem()
}

type tenantInboundSamlConfigArgs struct {
	// Human friendly display name.
	DisplayName string `pulumi:"displayName"`
	// If this config allows users to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig TenantInboundSamlConfigIdpConfig `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig TenantInboundSamlConfigSpConfig `pulumi:"spConfig"`
	// The name of the tenant where this inbound SAML config resource exists
	Tenant string `pulumi:"tenant"`
}

// The set of arguments for constructing a TenantInboundSamlConfig resource.
type TenantInboundSamlConfigArgs struct {
	// Human friendly display name.
	DisplayName pulumi.StringInput
	// If this config allows users to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdpConfig TenantInboundSamlConfigIdpConfigInput
	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig TenantInboundSamlConfigSpConfigInput
	// The name of the tenant where this inbound SAML config resource exists
	Tenant pulumi.StringInput
}

func (TenantInboundSamlConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantInboundSamlConfigArgs)(nil)).Elem()
}

type TenantInboundSamlConfigInput interface {
	pulumi.Input

	ToTenantInboundSamlConfigOutput() TenantInboundSamlConfigOutput
	ToTenantInboundSamlConfigOutputWithContext(ctx context.Context) TenantInboundSamlConfigOutput
}

func (*TenantInboundSamlConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*TenantInboundSamlConfig)(nil))
}

func (i *TenantInboundSamlConfig) ToTenantInboundSamlConfigOutput() TenantInboundSamlConfigOutput {
	return i.ToTenantInboundSamlConfigOutputWithContext(context.Background())
}

func (i *TenantInboundSamlConfig) ToTenantInboundSamlConfigOutputWithContext(ctx context.Context) TenantInboundSamlConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantInboundSamlConfigOutput)
}

func (i *TenantInboundSamlConfig) ToTenantInboundSamlConfigPtrOutput() TenantInboundSamlConfigPtrOutput {
	return i.ToTenantInboundSamlConfigPtrOutputWithContext(context.Background())
}

func (i *TenantInboundSamlConfig) ToTenantInboundSamlConfigPtrOutputWithContext(ctx context.Context) TenantInboundSamlConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantInboundSamlConfigPtrOutput)
}

type TenantInboundSamlConfigPtrInput interface {
	pulumi.Input

	ToTenantInboundSamlConfigPtrOutput() TenantInboundSamlConfigPtrOutput
	ToTenantInboundSamlConfigPtrOutputWithContext(ctx context.Context) TenantInboundSamlConfigPtrOutput
}

type tenantInboundSamlConfigPtrType TenantInboundSamlConfigArgs

func (*tenantInboundSamlConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantInboundSamlConfig)(nil))
}

func (i *tenantInboundSamlConfigPtrType) ToTenantInboundSamlConfigPtrOutput() TenantInboundSamlConfigPtrOutput {
	return i.ToTenantInboundSamlConfigPtrOutputWithContext(context.Background())
}

func (i *tenantInboundSamlConfigPtrType) ToTenantInboundSamlConfigPtrOutputWithContext(ctx context.Context) TenantInboundSamlConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantInboundSamlConfigPtrOutput)
}

// TenantInboundSamlConfigArrayInput is an input type that accepts TenantInboundSamlConfigArray and TenantInboundSamlConfigArrayOutput values.
// You can construct a concrete instance of `TenantInboundSamlConfigArrayInput` via:
//
//          TenantInboundSamlConfigArray{ TenantInboundSamlConfigArgs{...} }
type TenantInboundSamlConfigArrayInput interface {
	pulumi.Input

	ToTenantInboundSamlConfigArrayOutput() TenantInboundSamlConfigArrayOutput
	ToTenantInboundSamlConfigArrayOutputWithContext(context.Context) TenantInboundSamlConfigArrayOutput
}

type TenantInboundSamlConfigArray []TenantInboundSamlConfigInput

func (TenantInboundSamlConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantInboundSamlConfig)(nil)).Elem()
}

func (i TenantInboundSamlConfigArray) ToTenantInboundSamlConfigArrayOutput() TenantInboundSamlConfigArrayOutput {
	return i.ToTenantInboundSamlConfigArrayOutputWithContext(context.Background())
}

func (i TenantInboundSamlConfigArray) ToTenantInboundSamlConfigArrayOutputWithContext(ctx context.Context) TenantInboundSamlConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantInboundSamlConfigArrayOutput)
}

// TenantInboundSamlConfigMapInput is an input type that accepts TenantInboundSamlConfigMap and TenantInboundSamlConfigMapOutput values.
// You can construct a concrete instance of `TenantInboundSamlConfigMapInput` via:
//
//          TenantInboundSamlConfigMap{ "key": TenantInboundSamlConfigArgs{...} }
type TenantInboundSamlConfigMapInput interface {
	pulumi.Input

	ToTenantInboundSamlConfigMapOutput() TenantInboundSamlConfigMapOutput
	ToTenantInboundSamlConfigMapOutputWithContext(context.Context) TenantInboundSamlConfigMapOutput
}

type TenantInboundSamlConfigMap map[string]TenantInboundSamlConfigInput

func (TenantInboundSamlConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantInboundSamlConfig)(nil)).Elem()
}

func (i TenantInboundSamlConfigMap) ToTenantInboundSamlConfigMapOutput() TenantInboundSamlConfigMapOutput {
	return i.ToTenantInboundSamlConfigMapOutputWithContext(context.Background())
}

func (i TenantInboundSamlConfigMap) ToTenantInboundSamlConfigMapOutputWithContext(ctx context.Context) TenantInboundSamlConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantInboundSamlConfigMapOutput)
}

type TenantInboundSamlConfigOutput struct{ *pulumi.OutputState }

func (TenantInboundSamlConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TenantInboundSamlConfig)(nil))
}

func (o TenantInboundSamlConfigOutput) ToTenantInboundSamlConfigOutput() TenantInboundSamlConfigOutput {
	return o
}

func (o TenantInboundSamlConfigOutput) ToTenantInboundSamlConfigOutputWithContext(ctx context.Context) TenantInboundSamlConfigOutput {
	return o
}

func (o TenantInboundSamlConfigOutput) ToTenantInboundSamlConfigPtrOutput() TenantInboundSamlConfigPtrOutput {
	return o.ToTenantInboundSamlConfigPtrOutputWithContext(context.Background())
}

func (o TenantInboundSamlConfigOutput) ToTenantInboundSamlConfigPtrOutputWithContext(ctx context.Context) TenantInboundSamlConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TenantInboundSamlConfig) *TenantInboundSamlConfig {
		return &v
	}).(TenantInboundSamlConfigPtrOutput)
}

type TenantInboundSamlConfigPtrOutput struct{ *pulumi.OutputState }

func (TenantInboundSamlConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantInboundSamlConfig)(nil))
}

func (o TenantInboundSamlConfigPtrOutput) ToTenantInboundSamlConfigPtrOutput() TenantInboundSamlConfigPtrOutput {
	return o
}

func (o TenantInboundSamlConfigPtrOutput) ToTenantInboundSamlConfigPtrOutputWithContext(ctx context.Context) TenantInboundSamlConfigPtrOutput {
	return o
}

func (o TenantInboundSamlConfigPtrOutput) Elem() TenantInboundSamlConfigOutput {
	return o.ApplyT(func(v *TenantInboundSamlConfig) TenantInboundSamlConfig {
		if v != nil {
			return *v
		}
		var ret TenantInboundSamlConfig
		return ret
	}).(TenantInboundSamlConfigOutput)
}

type TenantInboundSamlConfigArrayOutput struct{ *pulumi.OutputState }

func (TenantInboundSamlConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TenantInboundSamlConfig)(nil))
}

func (o TenantInboundSamlConfigArrayOutput) ToTenantInboundSamlConfigArrayOutput() TenantInboundSamlConfigArrayOutput {
	return o
}

func (o TenantInboundSamlConfigArrayOutput) ToTenantInboundSamlConfigArrayOutputWithContext(ctx context.Context) TenantInboundSamlConfigArrayOutput {
	return o
}

func (o TenantInboundSamlConfigArrayOutput) Index(i pulumi.IntInput) TenantInboundSamlConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TenantInboundSamlConfig {
		return vs[0].([]TenantInboundSamlConfig)[vs[1].(int)]
	}).(TenantInboundSamlConfigOutput)
}

type TenantInboundSamlConfigMapOutput struct{ *pulumi.OutputState }

func (TenantInboundSamlConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TenantInboundSamlConfig)(nil))
}

func (o TenantInboundSamlConfigMapOutput) ToTenantInboundSamlConfigMapOutput() TenantInboundSamlConfigMapOutput {
	return o
}

func (o TenantInboundSamlConfigMapOutput) ToTenantInboundSamlConfigMapOutputWithContext(ctx context.Context) TenantInboundSamlConfigMapOutput {
	return o
}

func (o TenantInboundSamlConfigMapOutput) MapIndex(k pulumi.StringInput) TenantInboundSamlConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TenantInboundSamlConfig {
		return vs[0].(map[string]TenantInboundSamlConfig)[vs[1].(string)]
	}).(TenantInboundSamlConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(TenantInboundSamlConfigOutput{})
	pulumi.RegisterOutputType(TenantInboundSamlConfigPtrOutput{})
	pulumi.RegisterOutputType(TenantInboundSamlConfigArrayOutput{})
	pulumi.RegisterOutputType(TenantInboundSamlConfigMapOutput{})
}
