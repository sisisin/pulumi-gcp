// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// FeaturestoreEntitytype can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType default {{featurestore}}/entityTypes/{{name}}
// ```
type AiFeatureStoreEntityType struct {
	pulumi.CustomResourceState

	// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Used to perform consistent read-modify-write updates.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
	Featurestore pulumi.StringOutput `pulumi:"featurestore"`
	// A set of key/value label pairs to assign to this EntityType.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The default monitoring configuration for all Features under this EntityType.
	// If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
	// Structure is documented below.
	MonitoringConfig AiFeatureStoreEntityTypeMonitoringConfigPtrOutput `pulumi:"monitoringConfig"`
	// The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name pulumi.StringOutput `pulumi:"name"`
	// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAiFeatureStoreEntityType registers a new resource with the given unique name, arguments, and options.
func NewAiFeatureStoreEntityType(ctx *pulumi.Context,
	name string, args *AiFeatureStoreEntityTypeArgs, opts ...pulumi.ResourceOption) (*AiFeatureStoreEntityType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Featurestore == nil {
		return nil, errors.New("invalid value for required argument 'Featurestore'")
	}
	var resource AiFeatureStoreEntityType
	err := ctx.RegisterResource("gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiFeatureStoreEntityType gets an existing AiFeatureStoreEntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiFeatureStoreEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiFeatureStoreEntityTypeState, opts ...pulumi.ResourceOption) (*AiFeatureStoreEntityType, error) {
	var resource AiFeatureStoreEntityType
	err := ctx.ReadResource("gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiFeatureStoreEntityType resources.
type aiFeatureStoreEntityTypeState struct {
	// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime *string `pulumi:"createTime"`
	// Used to perform consistent read-modify-write updates.
	Etag *string `pulumi:"etag"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
	Featurestore *string `pulumi:"featurestore"`
	// A set of key/value label pairs to assign to this EntityType.
	Labels map[string]string `pulumi:"labels"`
	// The default monitoring configuration for all Features under this EntityType.
	// If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
	// Structure is documented below.
	MonitoringConfig *AiFeatureStoreEntityTypeMonitoringConfig `pulumi:"monitoringConfig"`
	// The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name *string `pulumi:"name"`
	// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime *string `pulumi:"updateTime"`
}

type AiFeatureStoreEntityTypeState struct {
	// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime pulumi.StringPtrInput
	// Used to perform consistent read-modify-write updates.
	Etag pulumi.StringPtrInput
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
	Featurestore pulumi.StringPtrInput
	// A set of key/value label pairs to assign to this EntityType.
	Labels pulumi.StringMapInput
	// The default monitoring configuration for all Features under this EntityType.
	// If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
	// Structure is documented below.
	MonitoringConfig AiFeatureStoreEntityTypeMonitoringConfigPtrInput
	// The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name pulumi.StringPtrInput
	// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime pulumi.StringPtrInput
}

func (AiFeatureStoreEntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreEntityTypeState)(nil)).Elem()
}

type aiFeatureStoreEntityTypeArgs struct {
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
	Featurestore string `pulumi:"featurestore"`
	// A set of key/value label pairs to assign to this EntityType.
	Labels map[string]string `pulumi:"labels"`
	// The default monitoring configuration for all Features under this EntityType.
	// If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
	// Structure is documented below.
	MonitoringConfig *AiFeatureStoreEntityTypeMonitoringConfig `pulumi:"monitoringConfig"`
	// The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AiFeatureStoreEntityType resource.
type AiFeatureStoreEntityTypeArgs struct {
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
	Featurestore pulumi.StringInput
	// A set of key/value label pairs to assign to this EntityType.
	Labels pulumi.StringMapInput
	// The default monitoring configuration for all Features under this EntityType.
	// If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
	// Structure is documented below.
	MonitoringConfig AiFeatureStoreEntityTypeMonitoringConfigPtrInput
	// The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name pulumi.StringPtrInput
}

func (AiFeatureStoreEntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreEntityTypeArgs)(nil)).Elem()
}

type AiFeatureStoreEntityTypeInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeOutput() AiFeatureStoreEntityTypeOutput
	ToAiFeatureStoreEntityTypeOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeOutput
}

func (*AiFeatureStoreEntityType) ElementType() reflect.Type {
	return reflect.TypeOf((*AiFeatureStoreEntityType)(nil))
}

func (i *AiFeatureStoreEntityType) ToAiFeatureStoreEntityTypeOutput() AiFeatureStoreEntityTypeOutput {
	return i.ToAiFeatureStoreEntityTypeOutputWithContext(context.Background())
}

func (i *AiFeatureStoreEntityType) ToAiFeatureStoreEntityTypeOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeOutput)
}

func (i *AiFeatureStoreEntityType) ToAiFeatureStoreEntityTypePtrOutput() AiFeatureStoreEntityTypePtrOutput {
	return i.ToAiFeatureStoreEntityTypePtrOutputWithContext(context.Background())
}

func (i *AiFeatureStoreEntityType) ToAiFeatureStoreEntityTypePtrOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypePtrOutput)
}

type AiFeatureStoreEntityTypePtrInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypePtrOutput() AiFeatureStoreEntityTypePtrOutput
	ToAiFeatureStoreEntityTypePtrOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypePtrOutput
}

type aiFeatureStoreEntityTypePtrType AiFeatureStoreEntityTypeArgs

func (*aiFeatureStoreEntityTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreEntityType)(nil))
}

func (i *aiFeatureStoreEntityTypePtrType) ToAiFeatureStoreEntityTypePtrOutput() AiFeatureStoreEntityTypePtrOutput {
	return i.ToAiFeatureStoreEntityTypePtrOutputWithContext(context.Background())
}

func (i *aiFeatureStoreEntityTypePtrType) ToAiFeatureStoreEntityTypePtrOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypePtrOutput)
}

// AiFeatureStoreEntityTypeArrayInput is an input type that accepts AiFeatureStoreEntityTypeArray and AiFeatureStoreEntityTypeArrayOutput values.
// You can construct a concrete instance of `AiFeatureStoreEntityTypeArrayInput` via:
//
//          AiFeatureStoreEntityTypeArray{ AiFeatureStoreEntityTypeArgs{...} }
type AiFeatureStoreEntityTypeArrayInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeArrayOutput() AiFeatureStoreEntityTypeArrayOutput
	ToAiFeatureStoreEntityTypeArrayOutputWithContext(context.Context) AiFeatureStoreEntityTypeArrayOutput
}

type AiFeatureStoreEntityTypeArray []AiFeatureStoreEntityTypeInput

func (AiFeatureStoreEntityTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreEntityType)(nil)).Elem()
}

func (i AiFeatureStoreEntityTypeArray) ToAiFeatureStoreEntityTypeArrayOutput() AiFeatureStoreEntityTypeArrayOutput {
	return i.ToAiFeatureStoreEntityTypeArrayOutputWithContext(context.Background())
}

func (i AiFeatureStoreEntityTypeArray) ToAiFeatureStoreEntityTypeArrayOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeArrayOutput)
}

// AiFeatureStoreEntityTypeMapInput is an input type that accepts AiFeatureStoreEntityTypeMap and AiFeatureStoreEntityTypeMapOutput values.
// You can construct a concrete instance of `AiFeatureStoreEntityTypeMapInput` via:
//
//          AiFeatureStoreEntityTypeMap{ "key": AiFeatureStoreEntityTypeArgs{...} }
type AiFeatureStoreEntityTypeMapInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeMapOutput() AiFeatureStoreEntityTypeMapOutput
	ToAiFeatureStoreEntityTypeMapOutputWithContext(context.Context) AiFeatureStoreEntityTypeMapOutput
}

type AiFeatureStoreEntityTypeMap map[string]AiFeatureStoreEntityTypeInput

func (AiFeatureStoreEntityTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreEntityType)(nil)).Elem()
}

func (i AiFeatureStoreEntityTypeMap) ToAiFeatureStoreEntityTypeMapOutput() AiFeatureStoreEntityTypeMapOutput {
	return i.ToAiFeatureStoreEntityTypeMapOutputWithContext(context.Background())
}

func (i AiFeatureStoreEntityTypeMap) ToAiFeatureStoreEntityTypeMapOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeMapOutput)
}

type AiFeatureStoreEntityTypeOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AiFeatureStoreEntityType)(nil))
}

func (o AiFeatureStoreEntityTypeOutput) ToAiFeatureStoreEntityTypeOutput() AiFeatureStoreEntityTypeOutput {
	return o
}

func (o AiFeatureStoreEntityTypeOutput) ToAiFeatureStoreEntityTypeOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeOutput {
	return o
}

func (o AiFeatureStoreEntityTypeOutput) ToAiFeatureStoreEntityTypePtrOutput() AiFeatureStoreEntityTypePtrOutput {
	return o.ToAiFeatureStoreEntityTypePtrOutputWithContext(context.Background())
}

func (o AiFeatureStoreEntityTypeOutput) ToAiFeatureStoreEntityTypePtrOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AiFeatureStoreEntityType) *AiFeatureStoreEntityType {
		return &v
	}).(AiFeatureStoreEntityTypePtrOutput)
}

type AiFeatureStoreEntityTypePtrOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreEntityType)(nil))
}

func (o AiFeatureStoreEntityTypePtrOutput) ToAiFeatureStoreEntityTypePtrOutput() AiFeatureStoreEntityTypePtrOutput {
	return o
}

func (o AiFeatureStoreEntityTypePtrOutput) ToAiFeatureStoreEntityTypePtrOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypePtrOutput {
	return o
}

func (o AiFeatureStoreEntityTypePtrOutput) Elem() AiFeatureStoreEntityTypeOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityType) AiFeatureStoreEntityType {
		if v != nil {
			return *v
		}
		var ret AiFeatureStoreEntityType
		return ret
	}).(AiFeatureStoreEntityTypeOutput)
}

type AiFeatureStoreEntityTypeArrayOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AiFeatureStoreEntityType)(nil))
}

func (o AiFeatureStoreEntityTypeArrayOutput) ToAiFeatureStoreEntityTypeArrayOutput() AiFeatureStoreEntityTypeArrayOutput {
	return o
}

func (o AiFeatureStoreEntityTypeArrayOutput) ToAiFeatureStoreEntityTypeArrayOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeArrayOutput {
	return o
}

func (o AiFeatureStoreEntityTypeArrayOutput) Index(i pulumi.IntInput) AiFeatureStoreEntityTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AiFeatureStoreEntityType {
		return vs[0].([]AiFeatureStoreEntityType)[vs[1].(int)]
	}).(AiFeatureStoreEntityTypeOutput)
}

type AiFeatureStoreEntityTypeMapOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AiFeatureStoreEntityType)(nil))
}

func (o AiFeatureStoreEntityTypeMapOutput) ToAiFeatureStoreEntityTypeMapOutput() AiFeatureStoreEntityTypeMapOutput {
	return o
}

func (o AiFeatureStoreEntityTypeMapOutput) ToAiFeatureStoreEntityTypeMapOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeMapOutput {
	return o
}

func (o AiFeatureStoreEntityTypeMapOutput) MapIndex(k pulumi.StringInput) AiFeatureStoreEntityTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AiFeatureStoreEntityType {
		return vs[0].(map[string]AiFeatureStoreEntityType)[vs[1].(string)]
	}).(AiFeatureStoreEntityTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypePtrOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeArrayOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeMapOutput{})
}