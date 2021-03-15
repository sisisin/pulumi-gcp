// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AuthorityIamPolicy struct {
	pulumi.CustomResourceState

	CertificateAuthority pulumi.StringOutput `pulumi:"certificateAuthority"`
	Etag                 pulumi.StringOutput `pulumi:"etag"`
	PolicyData           pulumi.StringOutput `pulumi:"policyData"`
}

// NewAuthorityIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAuthorityIamPolicy(ctx *pulumi.Context,
	name string, args *AuthorityIamPolicyArgs, opts ...pulumi.ResourceOption) (*AuthorityIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateAuthority == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthority'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource AuthorityIamPolicy
	err := ctx.RegisterResource("gcp:certificateauthority/authorityIamPolicy:AuthorityIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorityIamPolicy gets an existing AuthorityIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorityIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorityIamPolicyState, opts ...pulumi.ResourceOption) (*AuthorityIamPolicy, error) {
	var resource AuthorityIamPolicy
	err := ctx.ReadResource("gcp:certificateauthority/authorityIamPolicy:AuthorityIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorityIamPolicy resources.
type authorityIamPolicyState struct {
	CertificateAuthority *string `pulumi:"certificateAuthority"`
	Etag                 *string `pulumi:"etag"`
	PolicyData           *string `pulumi:"policyData"`
}

type AuthorityIamPolicyState struct {
	CertificateAuthority pulumi.StringPtrInput
	Etag                 pulumi.StringPtrInput
	PolicyData           pulumi.StringPtrInput
}

func (AuthorityIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorityIamPolicyState)(nil)).Elem()
}

type authorityIamPolicyArgs struct {
	CertificateAuthority string `pulumi:"certificateAuthority"`
	PolicyData           string `pulumi:"policyData"`
}

// The set of arguments for constructing a AuthorityIamPolicy resource.
type AuthorityIamPolicyArgs struct {
	CertificateAuthority pulumi.StringInput
	PolicyData           pulumi.StringInput
}

func (AuthorityIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorityIamPolicyArgs)(nil)).Elem()
}

type AuthorityIamPolicyInput interface {
	pulumi.Input

	ToAuthorityIamPolicyOutput() AuthorityIamPolicyOutput
	ToAuthorityIamPolicyOutputWithContext(ctx context.Context) AuthorityIamPolicyOutput
}

func (*AuthorityIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorityIamPolicy)(nil))
}

func (i *AuthorityIamPolicy) ToAuthorityIamPolicyOutput() AuthorityIamPolicyOutput {
	return i.ToAuthorityIamPolicyOutputWithContext(context.Background())
}

func (i *AuthorityIamPolicy) ToAuthorityIamPolicyOutputWithContext(ctx context.Context) AuthorityIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityIamPolicyOutput)
}

func (i *AuthorityIamPolicy) ToAuthorityIamPolicyPtrOutput() AuthorityIamPolicyPtrOutput {
	return i.ToAuthorityIamPolicyPtrOutputWithContext(context.Background())
}

func (i *AuthorityIamPolicy) ToAuthorityIamPolicyPtrOutputWithContext(ctx context.Context) AuthorityIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityIamPolicyPtrOutput)
}

type AuthorityIamPolicyPtrInput interface {
	pulumi.Input

	ToAuthorityIamPolicyPtrOutput() AuthorityIamPolicyPtrOutput
	ToAuthorityIamPolicyPtrOutputWithContext(ctx context.Context) AuthorityIamPolicyPtrOutput
}

type authorityIamPolicyPtrType AuthorityIamPolicyArgs

func (*authorityIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorityIamPolicy)(nil))
}

func (i *authorityIamPolicyPtrType) ToAuthorityIamPolicyPtrOutput() AuthorityIamPolicyPtrOutput {
	return i.ToAuthorityIamPolicyPtrOutputWithContext(context.Background())
}

func (i *authorityIamPolicyPtrType) ToAuthorityIamPolicyPtrOutputWithContext(ctx context.Context) AuthorityIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityIamPolicyPtrOutput)
}

// AuthorityIamPolicyArrayInput is an input type that accepts AuthorityIamPolicyArray and AuthorityIamPolicyArrayOutput values.
// You can construct a concrete instance of `AuthorityIamPolicyArrayInput` via:
//
//          AuthorityIamPolicyArray{ AuthorityIamPolicyArgs{...} }
type AuthorityIamPolicyArrayInput interface {
	pulumi.Input

	ToAuthorityIamPolicyArrayOutput() AuthorityIamPolicyArrayOutput
	ToAuthorityIamPolicyArrayOutputWithContext(context.Context) AuthorityIamPolicyArrayOutput
}

type AuthorityIamPolicyArray []AuthorityIamPolicyInput

func (AuthorityIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AuthorityIamPolicy)(nil))
}

func (i AuthorityIamPolicyArray) ToAuthorityIamPolicyArrayOutput() AuthorityIamPolicyArrayOutput {
	return i.ToAuthorityIamPolicyArrayOutputWithContext(context.Background())
}

func (i AuthorityIamPolicyArray) ToAuthorityIamPolicyArrayOutputWithContext(ctx context.Context) AuthorityIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityIamPolicyArrayOutput)
}

// AuthorityIamPolicyMapInput is an input type that accepts AuthorityIamPolicyMap and AuthorityIamPolicyMapOutput values.
// You can construct a concrete instance of `AuthorityIamPolicyMapInput` via:
//
//          AuthorityIamPolicyMap{ "key": AuthorityIamPolicyArgs{...} }
type AuthorityIamPolicyMapInput interface {
	pulumi.Input

	ToAuthorityIamPolicyMapOutput() AuthorityIamPolicyMapOutput
	ToAuthorityIamPolicyMapOutputWithContext(context.Context) AuthorityIamPolicyMapOutput
}

type AuthorityIamPolicyMap map[string]AuthorityIamPolicyInput

func (AuthorityIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AuthorityIamPolicy)(nil))
}

func (i AuthorityIamPolicyMap) ToAuthorityIamPolicyMapOutput() AuthorityIamPolicyMapOutput {
	return i.ToAuthorityIamPolicyMapOutputWithContext(context.Background())
}

func (i AuthorityIamPolicyMap) ToAuthorityIamPolicyMapOutputWithContext(ctx context.Context) AuthorityIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityIamPolicyMapOutput)
}

type AuthorityIamPolicyOutput struct {
	*pulumi.OutputState
}

func (AuthorityIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorityIamPolicy)(nil))
}

func (o AuthorityIamPolicyOutput) ToAuthorityIamPolicyOutput() AuthorityIamPolicyOutput {
	return o
}

func (o AuthorityIamPolicyOutput) ToAuthorityIamPolicyOutputWithContext(ctx context.Context) AuthorityIamPolicyOutput {
	return o
}

func (o AuthorityIamPolicyOutput) ToAuthorityIamPolicyPtrOutput() AuthorityIamPolicyPtrOutput {
	return o.ToAuthorityIamPolicyPtrOutputWithContext(context.Background())
}

func (o AuthorityIamPolicyOutput) ToAuthorityIamPolicyPtrOutputWithContext(ctx context.Context) AuthorityIamPolicyPtrOutput {
	return o.ApplyT(func(v AuthorityIamPolicy) *AuthorityIamPolicy {
		return &v
	}).(AuthorityIamPolicyPtrOutput)
}

type AuthorityIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (AuthorityIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorityIamPolicy)(nil))
}

func (o AuthorityIamPolicyPtrOutput) ToAuthorityIamPolicyPtrOutput() AuthorityIamPolicyPtrOutput {
	return o
}

func (o AuthorityIamPolicyPtrOutput) ToAuthorityIamPolicyPtrOutputWithContext(ctx context.Context) AuthorityIamPolicyPtrOutput {
	return o
}

type AuthorityIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (AuthorityIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorityIamPolicy)(nil))
}

func (o AuthorityIamPolicyArrayOutput) ToAuthorityIamPolicyArrayOutput() AuthorityIamPolicyArrayOutput {
	return o
}

func (o AuthorityIamPolicyArrayOutput) ToAuthorityIamPolicyArrayOutputWithContext(ctx context.Context) AuthorityIamPolicyArrayOutput {
	return o
}

func (o AuthorityIamPolicyArrayOutput) Index(i pulumi.IntInput) AuthorityIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthorityIamPolicy {
		return vs[0].([]AuthorityIamPolicy)[vs[1].(int)]
	}).(AuthorityIamPolicyOutput)
}

type AuthorityIamPolicyMapOutput struct{ *pulumi.OutputState }

func (AuthorityIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthorityIamPolicy)(nil))
}

func (o AuthorityIamPolicyMapOutput) ToAuthorityIamPolicyMapOutput() AuthorityIamPolicyMapOutput {
	return o
}

func (o AuthorityIamPolicyMapOutput) ToAuthorityIamPolicyMapOutputWithContext(ctx context.Context) AuthorityIamPolicyMapOutput {
	return o
}

func (o AuthorityIamPolicyMapOutput) MapIndex(k pulumi.StringInput) AuthorityIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthorityIamPolicy {
		return vs[0].(map[string]AuthorityIamPolicy)[vs[1].(string)]
	}).(AuthorityIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorityIamPolicyOutput{})
	pulumi.RegisterOutputType(AuthorityIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(AuthorityIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(AuthorityIamPolicyMapOutput{})
}