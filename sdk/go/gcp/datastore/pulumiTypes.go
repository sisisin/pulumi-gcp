// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datastore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataStoreIndexProperty struct {
	// The direction the index should optimize for sorting.
	// Possible values are `ASCENDING` and `DESCENDING`.
	Direction string `pulumi:"direction"`
	// The property name to index.
	Name string `pulumi:"name"`
}

// DataStoreIndexPropertyInput is an input type that accepts DataStoreIndexPropertyArgs and DataStoreIndexPropertyOutput values.
// You can construct a concrete instance of `DataStoreIndexPropertyInput` via:
//
//          DataStoreIndexPropertyArgs{...}
type DataStoreIndexPropertyInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput
	ToDataStoreIndexPropertyOutputWithContext(context.Context) DataStoreIndexPropertyOutput
}

type DataStoreIndexPropertyArgs struct {
	// The direction the index should optimize for sorting.
	// Possible values are `ASCENDING` and `DESCENDING`.
	Direction pulumi.StringInput `pulumi:"direction"`
	// The property name to index.
	Name pulumi.StringInput `pulumi:"name"`
}

func (DataStoreIndexPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexProperty)(nil)).Elem()
}

func (i DataStoreIndexPropertyArgs) ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput {
	return i.ToDataStoreIndexPropertyOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyArgs) ToDataStoreIndexPropertyOutputWithContext(ctx context.Context) DataStoreIndexPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyOutput)
}

// DataStoreIndexPropertyArrayInput is an input type that accepts DataStoreIndexPropertyArray and DataStoreIndexPropertyArrayOutput values.
// You can construct a concrete instance of `DataStoreIndexPropertyArrayInput` via:
//
//          DataStoreIndexPropertyArray{ DataStoreIndexPropertyArgs{...} }
type DataStoreIndexPropertyArrayInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput
	ToDataStoreIndexPropertyArrayOutputWithContext(context.Context) DataStoreIndexPropertyArrayOutput
}

type DataStoreIndexPropertyArray []DataStoreIndexPropertyInput

func (DataStoreIndexPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexProperty)(nil)).Elem()
}

func (i DataStoreIndexPropertyArray) ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput {
	return i.ToDataStoreIndexPropertyArrayOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyArray) ToDataStoreIndexPropertyArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyArrayOutput)
}

type DataStoreIndexPropertyOutput struct{ *pulumi.OutputState }

func (DataStoreIndexPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexProperty)(nil)).Elem()
}

func (o DataStoreIndexPropertyOutput) ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput {
	return o
}

func (o DataStoreIndexPropertyOutput) ToDataStoreIndexPropertyOutputWithContext(ctx context.Context) DataStoreIndexPropertyOutput {
	return o
}

// The direction the index should optimize for sorting.
// Possible values are `ASCENDING` and `DESCENDING`.
func (o DataStoreIndexPropertyOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v DataStoreIndexProperty) string { return v.Direction }).(pulumi.StringOutput)
}

// The property name to index.
func (o DataStoreIndexPropertyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataStoreIndexProperty) string { return v.Name }).(pulumi.StringOutput)
}

type DataStoreIndexPropertyArrayOutput struct{ *pulumi.OutputState }

func (DataStoreIndexPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexProperty)(nil)).Elem()
}

func (o DataStoreIndexPropertyArrayOutput) ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput {
	return o
}

func (o DataStoreIndexPropertyArrayOutput) ToDataStoreIndexPropertyArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyArrayOutput {
	return o
}

func (o DataStoreIndexPropertyArrayOutput) Index(i pulumi.IntInput) DataStoreIndexPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataStoreIndexProperty {
		return vs[0].([]DataStoreIndexProperty)[vs[1].(int)]
	}).(DataStoreIndexPropertyOutput)
}

func init() {
	pulumi.RegisterOutputType(DataStoreIndexPropertyOutput{})
	pulumi.RegisterOutputType(DataStoreIndexPropertyArrayOutput{})
}
