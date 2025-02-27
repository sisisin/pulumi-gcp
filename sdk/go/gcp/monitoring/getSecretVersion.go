// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: gcp.monitoring.getSecretVersion has been deprecated in favor of gcp.secretmanager.getSecretVersion
func GetSecretVersion(ctx *pulumi.Context, args *GetSecretVersionArgs, opts ...pulumi.InvokeOption) (*GetSecretVersionResult, error) {
	var rv GetSecretVersionResult
	err := ctx.Invoke("gcp:monitoring/getSecretVersion:getSecretVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretVersion.
type GetSecretVersionArgs struct {
	Project *string `pulumi:"project"`
	Secret  string  `pulumi:"secret"`
	Version *string `pulumi:"version"`
}

// A collection of values returned by getSecretVersion.
type GetSecretVersionResult struct {
	CreateTime  string `pulumi:"createTime"`
	DestroyTime string `pulumi:"destroyTime"`
	Enabled     bool   `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Name       string `pulumi:"name"`
	Project    string `pulumi:"project"`
	Secret     string `pulumi:"secret"`
	SecretData string `pulumi:"secretData"`
	Version    string `pulumi:"version"`
}

func GetSecretVersionOutput(ctx *pulumi.Context, args GetSecretVersionOutputArgs, opts ...pulumi.InvokeOption) GetSecretVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretVersionResult, error) {
			args := v.(GetSecretVersionArgs)
			r, err := GetSecretVersion(ctx, &args, opts...)
			return *r, err
		}).(GetSecretVersionResultOutput)
}

// A collection of arguments for invoking getSecretVersion.
type GetSecretVersionOutputArgs struct {
	Project pulumi.StringPtrInput `pulumi:"project"`
	Secret  pulumi.StringInput    `pulumi:"secret"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (GetSecretVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretVersionArgs)(nil)).Elem()
}

// A collection of values returned by getSecretVersion.
type GetSecretVersionResultOutput struct{ *pulumi.OutputState }

func (GetSecretVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretVersionResult)(nil)).Elem()
}

func (o GetSecretVersionResultOutput) ToGetSecretVersionResultOutput() GetSecretVersionResultOutput {
	return o
}

func (o GetSecretVersionResultOutput) ToGetSecretVersionResultOutputWithContext(ctx context.Context) GetSecretVersionResultOutput {
	return o
}

func (o GetSecretVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) DestroyTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.DestroyTime }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSecretVersionResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecretVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Secret }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) SecretData() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.SecretData }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretVersionResultOutput{})
}
