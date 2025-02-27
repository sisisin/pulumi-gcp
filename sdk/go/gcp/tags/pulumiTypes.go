// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tags

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagKeyIamBindingCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// TagKeyIamBindingConditionInput is an input type that accepts TagKeyIamBindingConditionArgs and TagKeyIamBindingConditionOutput values.
// You can construct a concrete instance of `TagKeyIamBindingConditionInput` via:
//
//          TagKeyIamBindingConditionArgs{...}
type TagKeyIamBindingConditionInput interface {
	pulumi.Input

	ToTagKeyIamBindingConditionOutput() TagKeyIamBindingConditionOutput
	ToTagKeyIamBindingConditionOutputWithContext(context.Context) TagKeyIamBindingConditionOutput
}

type TagKeyIamBindingConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (TagKeyIamBindingConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagKeyIamBindingCondition)(nil)).Elem()
}

func (i TagKeyIamBindingConditionArgs) ToTagKeyIamBindingConditionOutput() TagKeyIamBindingConditionOutput {
	return i.ToTagKeyIamBindingConditionOutputWithContext(context.Background())
}

func (i TagKeyIamBindingConditionArgs) ToTagKeyIamBindingConditionOutputWithContext(ctx context.Context) TagKeyIamBindingConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamBindingConditionOutput)
}

func (i TagKeyIamBindingConditionArgs) ToTagKeyIamBindingConditionPtrOutput() TagKeyIamBindingConditionPtrOutput {
	return i.ToTagKeyIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i TagKeyIamBindingConditionArgs) ToTagKeyIamBindingConditionPtrOutputWithContext(ctx context.Context) TagKeyIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamBindingConditionOutput).ToTagKeyIamBindingConditionPtrOutputWithContext(ctx)
}

// TagKeyIamBindingConditionPtrInput is an input type that accepts TagKeyIamBindingConditionArgs, TagKeyIamBindingConditionPtr and TagKeyIamBindingConditionPtrOutput values.
// You can construct a concrete instance of `TagKeyIamBindingConditionPtrInput` via:
//
//          TagKeyIamBindingConditionArgs{...}
//
//  or:
//
//          nil
type TagKeyIamBindingConditionPtrInput interface {
	pulumi.Input

	ToTagKeyIamBindingConditionPtrOutput() TagKeyIamBindingConditionPtrOutput
	ToTagKeyIamBindingConditionPtrOutputWithContext(context.Context) TagKeyIamBindingConditionPtrOutput
}

type tagKeyIamBindingConditionPtrType TagKeyIamBindingConditionArgs

func TagKeyIamBindingConditionPtr(v *TagKeyIamBindingConditionArgs) TagKeyIamBindingConditionPtrInput {
	return (*tagKeyIamBindingConditionPtrType)(v)
}

func (*tagKeyIamBindingConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamBindingCondition)(nil)).Elem()
}

func (i *tagKeyIamBindingConditionPtrType) ToTagKeyIamBindingConditionPtrOutput() TagKeyIamBindingConditionPtrOutput {
	return i.ToTagKeyIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i *tagKeyIamBindingConditionPtrType) ToTagKeyIamBindingConditionPtrOutputWithContext(ctx context.Context) TagKeyIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamBindingConditionPtrOutput)
}

type TagKeyIamBindingConditionOutput struct{ *pulumi.OutputState }

func (TagKeyIamBindingConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagKeyIamBindingCondition)(nil)).Elem()
}

func (o TagKeyIamBindingConditionOutput) ToTagKeyIamBindingConditionOutput() TagKeyIamBindingConditionOutput {
	return o
}

func (o TagKeyIamBindingConditionOutput) ToTagKeyIamBindingConditionOutputWithContext(ctx context.Context) TagKeyIamBindingConditionOutput {
	return o
}

func (o TagKeyIamBindingConditionOutput) ToTagKeyIamBindingConditionPtrOutput() TagKeyIamBindingConditionPtrOutput {
	return o.ToTagKeyIamBindingConditionPtrOutputWithContext(context.Background())
}

func (o TagKeyIamBindingConditionOutput) ToTagKeyIamBindingConditionPtrOutputWithContext(ctx context.Context) TagKeyIamBindingConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagKeyIamBindingCondition) *TagKeyIamBindingCondition {
		return &v
	}).(TagKeyIamBindingConditionPtrOutput)
}

func (o TagKeyIamBindingConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TagKeyIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TagKeyIamBindingConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v TagKeyIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o TagKeyIamBindingConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v TagKeyIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type TagKeyIamBindingConditionPtrOutput struct{ *pulumi.OutputState }

func (TagKeyIamBindingConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamBindingCondition)(nil)).Elem()
}

func (o TagKeyIamBindingConditionPtrOutput) ToTagKeyIamBindingConditionPtrOutput() TagKeyIamBindingConditionPtrOutput {
	return o
}

func (o TagKeyIamBindingConditionPtrOutput) ToTagKeyIamBindingConditionPtrOutputWithContext(ctx context.Context) TagKeyIamBindingConditionPtrOutput {
	return o
}

func (o TagKeyIamBindingConditionPtrOutput) Elem() TagKeyIamBindingConditionOutput {
	return o.ApplyT(func(v *TagKeyIamBindingCondition) TagKeyIamBindingCondition {
		if v != nil {
			return *v
		}
		var ret TagKeyIamBindingCondition
		return ret
	}).(TagKeyIamBindingConditionOutput)
}

func (o TagKeyIamBindingConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagKeyIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o TagKeyIamBindingConditionPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagKeyIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Expression
	}).(pulumi.StringPtrOutput)
}

func (o TagKeyIamBindingConditionPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagKeyIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

type TagKeyIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// TagKeyIamMemberConditionInput is an input type that accepts TagKeyIamMemberConditionArgs and TagKeyIamMemberConditionOutput values.
// You can construct a concrete instance of `TagKeyIamMemberConditionInput` via:
//
//          TagKeyIamMemberConditionArgs{...}
type TagKeyIamMemberConditionInput interface {
	pulumi.Input

	ToTagKeyIamMemberConditionOutput() TagKeyIamMemberConditionOutput
	ToTagKeyIamMemberConditionOutputWithContext(context.Context) TagKeyIamMemberConditionOutput
}

type TagKeyIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (TagKeyIamMemberConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagKeyIamMemberCondition)(nil)).Elem()
}

func (i TagKeyIamMemberConditionArgs) ToTagKeyIamMemberConditionOutput() TagKeyIamMemberConditionOutput {
	return i.ToTagKeyIamMemberConditionOutputWithContext(context.Background())
}

func (i TagKeyIamMemberConditionArgs) ToTagKeyIamMemberConditionOutputWithContext(ctx context.Context) TagKeyIamMemberConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamMemberConditionOutput)
}

func (i TagKeyIamMemberConditionArgs) ToTagKeyIamMemberConditionPtrOutput() TagKeyIamMemberConditionPtrOutput {
	return i.ToTagKeyIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i TagKeyIamMemberConditionArgs) ToTagKeyIamMemberConditionPtrOutputWithContext(ctx context.Context) TagKeyIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamMemberConditionOutput).ToTagKeyIamMemberConditionPtrOutputWithContext(ctx)
}

// TagKeyIamMemberConditionPtrInput is an input type that accepts TagKeyIamMemberConditionArgs, TagKeyIamMemberConditionPtr and TagKeyIamMemberConditionPtrOutput values.
// You can construct a concrete instance of `TagKeyIamMemberConditionPtrInput` via:
//
//          TagKeyIamMemberConditionArgs{...}
//
//  or:
//
//          nil
type TagKeyIamMemberConditionPtrInput interface {
	pulumi.Input

	ToTagKeyIamMemberConditionPtrOutput() TagKeyIamMemberConditionPtrOutput
	ToTagKeyIamMemberConditionPtrOutputWithContext(context.Context) TagKeyIamMemberConditionPtrOutput
}

type tagKeyIamMemberConditionPtrType TagKeyIamMemberConditionArgs

func TagKeyIamMemberConditionPtr(v *TagKeyIamMemberConditionArgs) TagKeyIamMemberConditionPtrInput {
	return (*tagKeyIamMemberConditionPtrType)(v)
}

func (*tagKeyIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamMemberCondition)(nil)).Elem()
}

func (i *tagKeyIamMemberConditionPtrType) ToTagKeyIamMemberConditionPtrOutput() TagKeyIamMemberConditionPtrOutput {
	return i.ToTagKeyIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *tagKeyIamMemberConditionPtrType) ToTagKeyIamMemberConditionPtrOutputWithContext(ctx context.Context) TagKeyIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamMemberConditionPtrOutput)
}

type TagKeyIamMemberConditionOutput struct{ *pulumi.OutputState }

func (TagKeyIamMemberConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagKeyIamMemberCondition)(nil)).Elem()
}

func (o TagKeyIamMemberConditionOutput) ToTagKeyIamMemberConditionOutput() TagKeyIamMemberConditionOutput {
	return o
}

func (o TagKeyIamMemberConditionOutput) ToTagKeyIamMemberConditionOutputWithContext(ctx context.Context) TagKeyIamMemberConditionOutput {
	return o
}

func (o TagKeyIamMemberConditionOutput) ToTagKeyIamMemberConditionPtrOutput() TagKeyIamMemberConditionPtrOutput {
	return o.ToTagKeyIamMemberConditionPtrOutputWithContext(context.Background())
}

func (o TagKeyIamMemberConditionOutput) ToTagKeyIamMemberConditionPtrOutputWithContext(ctx context.Context) TagKeyIamMemberConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagKeyIamMemberCondition) *TagKeyIamMemberCondition {
		return &v
	}).(TagKeyIamMemberConditionPtrOutput)
}

func (o TagKeyIamMemberConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TagKeyIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TagKeyIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v TagKeyIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o TagKeyIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v TagKeyIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type TagKeyIamMemberConditionPtrOutput struct{ *pulumi.OutputState }

func (TagKeyIamMemberConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamMemberCondition)(nil)).Elem()
}

func (o TagKeyIamMemberConditionPtrOutput) ToTagKeyIamMemberConditionPtrOutput() TagKeyIamMemberConditionPtrOutput {
	return o
}

func (o TagKeyIamMemberConditionPtrOutput) ToTagKeyIamMemberConditionPtrOutputWithContext(ctx context.Context) TagKeyIamMemberConditionPtrOutput {
	return o
}

func (o TagKeyIamMemberConditionPtrOutput) Elem() TagKeyIamMemberConditionOutput {
	return o.ApplyT(func(v *TagKeyIamMemberCondition) TagKeyIamMemberCondition {
		if v != nil {
			return *v
		}
		var ret TagKeyIamMemberCondition
		return ret
	}).(TagKeyIamMemberConditionOutput)
}

func (o TagKeyIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagKeyIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o TagKeyIamMemberConditionPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagKeyIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Expression
	}).(pulumi.StringPtrOutput)
}

func (o TagKeyIamMemberConditionPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagKeyIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

type TagValueIamBindingCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// TagValueIamBindingConditionInput is an input type that accepts TagValueIamBindingConditionArgs and TagValueIamBindingConditionOutput values.
// You can construct a concrete instance of `TagValueIamBindingConditionInput` via:
//
//          TagValueIamBindingConditionArgs{...}
type TagValueIamBindingConditionInput interface {
	pulumi.Input

	ToTagValueIamBindingConditionOutput() TagValueIamBindingConditionOutput
	ToTagValueIamBindingConditionOutputWithContext(context.Context) TagValueIamBindingConditionOutput
}

type TagValueIamBindingConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (TagValueIamBindingConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagValueIamBindingCondition)(nil)).Elem()
}

func (i TagValueIamBindingConditionArgs) ToTagValueIamBindingConditionOutput() TagValueIamBindingConditionOutput {
	return i.ToTagValueIamBindingConditionOutputWithContext(context.Background())
}

func (i TagValueIamBindingConditionArgs) ToTagValueIamBindingConditionOutputWithContext(ctx context.Context) TagValueIamBindingConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamBindingConditionOutput)
}

func (i TagValueIamBindingConditionArgs) ToTagValueIamBindingConditionPtrOutput() TagValueIamBindingConditionPtrOutput {
	return i.ToTagValueIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i TagValueIamBindingConditionArgs) ToTagValueIamBindingConditionPtrOutputWithContext(ctx context.Context) TagValueIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamBindingConditionOutput).ToTagValueIamBindingConditionPtrOutputWithContext(ctx)
}

// TagValueIamBindingConditionPtrInput is an input type that accepts TagValueIamBindingConditionArgs, TagValueIamBindingConditionPtr and TagValueIamBindingConditionPtrOutput values.
// You can construct a concrete instance of `TagValueIamBindingConditionPtrInput` via:
//
//          TagValueIamBindingConditionArgs{...}
//
//  or:
//
//          nil
type TagValueIamBindingConditionPtrInput interface {
	pulumi.Input

	ToTagValueIamBindingConditionPtrOutput() TagValueIamBindingConditionPtrOutput
	ToTagValueIamBindingConditionPtrOutputWithContext(context.Context) TagValueIamBindingConditionPtrOutput
}

type tagValueIamBindingConditionPtrType TagValueIamBindingConditionArgs

func TagValueIamBindingConditionPtr(v *TagValueIamBindingConditionArgs) TagValueIamBindingConditionPtrInput {
	return (*tagValueIamBindingConditionPtrType)(v)
}

func (*tagValueIamBindingConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValueIamBindingCondition)(nil)).Elem()
}

func (i *tagValueIamBindingConditionPtrType) ToTagValueIamBindingConditionPtrOutput() TagValueIamBindingConditionPtrOutput {
	return i.ToTagValueIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i *tagValueIamBindingConditionPtrType) ToTagValueIamBindingConditionPtrOutputWithContext(ctx context.Context) TagValueIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamBindingConditionPtrOutput)
}

type TagValueIamBindingConditionOutput struct{ *pulumi.OutputState }

func (TagValueIamBindingConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagValueIamBindingCondition)(nil)).Elem()
}

func (o TagValueIamBindingConditionOutput) ToTagValueIamBindingConditionOutput() TagValueIamBindingConditionOutput {
	return o
}

func (o TagValueIamBindingConditionOutput) ToTagValueIamBindingConditionOutputWithContext(ctx context.Context) TagValueIamBindingConditionOutput {
	return o
}

func (o TagValueIamBindingConditionOutput) ToTagValueIamBindingConditionPtrOutput() TagValueIamBindingConditionPtrOutput {
	return o.ToTagValueIamBindingConditionPtrOutputWithContext(context.Background())
}

func (o TagValueIamBindingConditionOutput) ToTagValueIamBindingConditionPtrOutputWithContext(ctx context.Context) TagValueIamBindingConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagValueIamBindingCondition) *TagValueIamBindingCondition {
		return &v
	}).(TagValueIamBindingConditionPtrOutput)
}

func (o TagValueIamBindingConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TagValueIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TagValueIamBindingConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v TagValueIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o TagValueIamBindingConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v TagValueIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type TagValueIamBindingConditionPtrOutput struct{ *pulumi.OutputState }

func (TagValueIamBindingConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValueIamBindingCondition)(nil)).Elem()
}

func (o TagValueIamBindingConditionPtrOutput) ToTagValueIamBindingConditionPtrOutput() TagValueIamBindingConditionPtrOutput {
	return o
}

func (o TagValueIamBindingConditionPtrOutput) ToTagValueIamBindingConditionPtrOutputWithContext(ctx context.Context) TagValueIamBindingConditionPtrOutput {
	return o
}

func (o TagValueIamBindingConditionPtrOutput) Elem() TagValueIamBindingConditionOutput {
	return o.ApplyT(func(v *TagValueIamBindingCondition) TagValueIamBindingCondition {
		if v != nil {
			return *v
		}
		var ret TagValueIamBindingCondition
		return ret
	}).(TagValueIamBindingConditionOutput)
}

func (o TagValueIamBindingConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagValueIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o TagValueIamBindingConditionPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagValueIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Expression
	}).(pulumi.StringPtrOutput)
}

func (o TagValueIamBindingConditionPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagValueIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

type TagValueIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// TagValueIamMemberConditionInput is an input type that accepts TagValueIamMemberConditionArgs and TagValueIamMemberConditionOutput values.
// You can construct a concrete instance of `TagValueIamMemberConditionInput` via:
//
//          TagValueIamMemberConditionArgs{...}
type TagValueIamMemberConditionInput interface {
	pulumi.Input

	ToTagValueIamMemberConditionOutput() TagValueIamMemberConditionOutput
	ToTagValueIamMemberConditionOutputWithContext(context.Context) TagValueIamMemberConditionOutput
}

type TagValueIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (TagValueIamMemberConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagValueIamMemberCondition)(nil)).Elem()
}

func (i TagValueIamMemberConditionArgs) ToTagValueIamMemberConditionOutput() TagValueIamMemberConditionOutput {
	return i.ToTagValueIamMemberConditionOutputWithContext(context.Background())
}

func (i TagValueIamMemberConditionArgs) ToTagValueIamMemberConditionOutputWithContext(ctx context.Context) TagValueIamMemberConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamMemberConditionOutput)
}

func (i TagValueIamMemberConditionArgs) ToTagValueIamMemberConditionPtrOutput() TagValueIamMemberConditionPtrOutput {
	return i.ToTagValueIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i TagValueIamMemberConditionArgs) ToTagValueIamMemberConditionPtrOutputWithContext(ctx context.Context) TagValueIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamMemberConditionOutput).ToTagValueIamMemberConditionPtrOutputWithContext(ctx)
}

// TagValueIamMemberConditionPtrInput is an input type that accepts TagValueIamMemberConditionArgs, TagValueIamMemberConditionPtr and TagValueIamMemberConditionPtrOutput values.
// You can construct a concrete instance of `TagValueIamMemberConditionPtrInput` via:
//
//          TagValueIamMemberConditionArgs{...}
//
//  or:
//
//          nil
type TagValueIamMemberConditionPtrInput interface {
	pulumi.Input

	ToTagValueIamMemberConditionPtrOutput() TagValueIamMemberConditionPtrOutput
	ToTagValueIamMemberConditionPtrOutputWithContext(context.Context) TagValueIamMemberConditionPtrOutput
}

type tagValueIamMemberConditionPtrType TagValueIamMemberConditionArgs

func TagValueIamMemberConditionPtr(v *TagValueIamMemberConditionArgs) TagValueIamMemberConditionPtrInput {
	return (*tagValueIamMemberConditionPtrType)(v)
}

func (*tagValueIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValueIamMemberCondition)(nil)).Elem()
}

func (i *tagValueIamMemberConditionPtrType) ToTagValueIamMemberConditionPtrOutput() TagValueIamMemberConditionPtrOutput {
	return i.ToTagValueIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *tagValueIamMemberConditionPtrType) ToTagValueIamMemberConditionPtrOutputWithContext(ctx context.Context) TagValueIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamMemberConditionPtrOutput)
}

type TagValueIamMemberConditionOutput struct{ *pulumi.OutputState }

func (TagValueIamMemberConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagValueIamMemberCondition)(nil)).Elem()
}

func (o TagValueIamMemberConditionOutput) ToTagValueIamMemberConditionOutput() TagValueIamMemberConditionOutput {
	return o
}

func (o TagValueIamMemberConditionOutput) ToTagValueIamMemberConditionOutputWithContext(ctx context.Context) TagValueIamMemberConditionOutput {
	return o
}

func (o TagValueIamMemberConditionOutput) ToTagValueIamMemberConditionPtrOutput() TagValueIamMemberConditionPtrOutput {
	return o.ToTagValueIamMemberConditionPtrOutputWithContext(context.Background())
}

func (o TagValueIamMemberConditionOutput) ToTagValueIamMemberConditionPtrOutputWithContext(ctx context.Context) TagValueIamMemberConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagValueIamMemberCondition) *TagValueIamMemberCondition {
		return &v
	}).(TagValueIamMemberConditionPtrOutput)
}

func (o TagValueIamMemberConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TagValueIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TagValueIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v TagValueIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o TagValueIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v TagValueIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type TagValueIamMemberConditionPtrOutput struct{ *pulumi.OutputState }

func (TagValueIamMemberConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValueIamMemberCondition)(nil)).Elem()
}

func (o TagValueIamMemberConditionPtrOutput) ToTagValueIamMemberConditionPtrOutput() TagValueIamMemberConditionPtrOutput {
	return o
}

func (o TagValueIamMemberConditionPtrOutput) ToTagValueIamMemberConditionPtrOutputWithContext(ctx context.Context) TagValueIamMemberConditionPtrOutput {
	return o
}

func (o TagValueIamMemberConditionPtrOutput) Elem() TagValueIamMemberConditionOutput {
	return o.ApplyT(func(v *TagValueIamMemberCondition) TagValueIamMemberCondition {
		if v != nil {
			return *v
		}
		var ret TagValueIamMemberCondition
		return ret
	}).(TagValueIamMemberConditionOutput)
}

func (o TagValueIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagValueIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o TagValueIamMemberConditionPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagValueIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Expression
	}).(pulumi.StringPtrOutput)
}

func (o TagValueIamMemberConditionPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagValueIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TagKeyIamBindingConditionOutput{})
	pulumi.RegisterOutputType(TagKeyIamBindingConditionPtrOutput{})
	pulumi.RegisterOutputType(TagKeyIamMemberConditionOutput{})
	pulumi.RegisterOutputType(TagKeyIamMemberConditionPtrOutput{})
	pulumi.RegisterOutputType(TagValueIamBindingConditionOutput{})
	pulumi.RegisterOutputType(TagValueIamBindingConditionPtrOutput{})
	pulumi.RegisterOutputType(TagValueIamMemberConditionOutput{})
	pulumi.RegisterOutputType(TagValueIamMemberConditionPtrOutput{})
}
