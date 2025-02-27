// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a organization-level logging sink. For more information see:
// * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/organizations.sinks)
// * How-to Guides
//     * [Exporting Logs](https://cloud.google.com/logging/docs/export)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/logging"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.NewBucket(ctx, "log_bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logging.NewOrganizationSink(ctx, "my_sink", &logging.OrganizationSinkArgs{
// 			Description: pulumi.String("some explanation on what this is"),
// 			OrgId:       pulumi.String("123456789"),
// 			Destination: log_bucket.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "storage.googleapis.com/", name), nil
// 			}).(pulumi.StringOutput),
// 			Filter: pulumi.String("resource.type = gce_instance AND severity >= WARNING"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMMember(ctx, "log_writer", &projects.IAMMemberArgs{
// 			Role:   pulumi.String("roles/storage.objectCreator"),
// 			Member: my_sink.WriterIdentity,
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
// Organization-level logging sinks can be imported using this format
//
// ```sh
//  $ pulumi import gcp:logging/organizationSink:OrganizationSink my_sink organizations/{{organization_id}}/sinks/{{sink_id}}
// ```
type OrganizationSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions OrganizationSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions OrganizationSinkExclusionArrayOutput `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren pulumi.BoolPtrOutput `pulumi:"includeChildren"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringOutput `pulumi:"name"`
	// The numeric ID of the organization to be exported to the sink.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewOrganizationSink registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSink(ctx *pulumi.Context,
	name string, args *OrganizationSinkArgs, opts ...pulumi.ResourceOption) (*OrganizationSink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	var resource OrganizationSink
	err := ctx.RegisterResource("gcp:logging/organizationSink:OrganizationSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSink gets an existing OrganizationSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSinkState, opts ...pulumi.ResourceOption) (*OrganizationSink, error) {
	var resource OrganizationSink
	err := ctx.ReadResource("gcp:logging/organizationSink:OrganizationSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSink resources.
type organizationSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *OrganizationSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination *string `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []OrganizationSinkExclusion `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren *bool `pulumi:"includeChildren"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
	// The numeric ID of the organization to be exported to the sink.
	OrgId *string `pulumi:"orgId"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity *string `pulumi:"writerIdentity"`
}

type OrganizationSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions OrganizationSinkBigqueryOptionsPtrInput
	// A description of this exclusion.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringPtrInput
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions OrganizationSinkExclusionArrayInput
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren pulumi.BoolPtrInput
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
	// The numeric ID of the organization to be exported to the sink.
	OrgId pulumi.StringPtrInput
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringPtrInput
}

func (OrganizationSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSinkState)(nil)).Elem()
}

type organizationSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *OrganizationSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination string `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []OrganizationSinkExclusion `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren *bool `pulumi:"includeChildren"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
	// The numeric ID of the organization to be exported to the sink.
	OrgId string `pulumi:"orgId"`
}

// The set of arguments for constructing a OrganizationSink resource.
type OrganizationSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions OrganizationSinkBigqueryOptionsPtrInput
	// A description of this exclusion.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringInput
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions OrganizationSinkExclusionArrayInput
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// Whether or not to include children organizations in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren pulumi.BoolPtrInput
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
	// The numeric ID of the organization to be exported to the sink.
	OrgId pulumi.StringInput
}

func (OrganizationSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSinkArgs)(nil)).Elem()
}

type OrganizationSinkInput interface {
	pulumi.Input

	ToOrganizationSinkOutput() OrganizationSinkOutput
	ToOrganizationSinkOutputWithContext(ctx context.Context) OrganizationSinkOutput
}

func (*OrganizationSink) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSink)(nil))
}

func (i *OrganizationSink) ToOrganizationSinkOutput() OrganizationSinkOutput {
	return i.ToOrganizationSinkOutputWithContext(context.Background())
}

func (i *OrganizationSink) ToOrganizationSinkOutputWithContext(ctx context.Context) OrganizationSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkOutput)
}

func (i *OrganizationSink) ToOrganizationSinkPtrOutput() OrganizationSinkPtrOutput {
	return i.ToOrganizationSinkPtrOutputWithContext(context.Background())
}

func (i *OrganizationSink) ToOrganizationSinkPtrOutputWithContext(ctx context.Context) OrganizationSinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkPtrOutput)
}

type OrganizationSinkPtrInput interface {
	pulumi.Input

	ToOrganizationSinkPtrOutput() OrganizationSinkPtrOutput
	ToOrganizationSinkPtrOutputWithContext(ctx context.Context) OrganizationSinkPtrOutput
}

type organizationSinkPtrType OrganizationSinkArgs

func (*organizationSinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationSink)(nil))
}

func (i *organizationSinkPtrType) ToOrganizationSinkPtrOutput() OrganizationSinkPtrOutput {
	return i.ToOrganizationSinkPtrOutputWithContext(context.Background())
}

func (i *organizationSinkPtrType) ToOrganizationSinkPtrOutputWithContext(ctx context.Context) OrganizationSinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkPtrOutput)
}

// OrganizationSinkArrayInput is an input type that accepts OrganizationSinkArray and OrganizationSinkArrayOutput values.
// You can construct a concrete instance of `OrganizationSinkArrayInput` via:
//
//          OrganizationSinkArray{ OrganizationSinkArgs{...} }
type OrganizationSinkArrayInput interface {
	pulumi.Input

	ToOrganizationSinkArrayOutput() OrganizationSinkArrayOutput
	ToOrganizationSinkArrayOutputWithContext(context.Context) OrganizationSinkArrayOutput
}

type OrganizationSinkArray []OrganizationSinkInput

func (OrganizationSinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationSink)(nil)).Elem()
}

func (i OrganizationSinkArray) ToOrganizationSinkArrayOutput() OrganizationSinkArrayOutput {
	return i.ToOrganizationSinkArrayOutputWithContext(context.Background())
}

func (i OrganizationSinkArray) ToOrganizationSinkArrayOutputWithContext(ctx context.Context) OrganizationSinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkArrayOutput)
}

// OrganizationSinkMapInput is an input type that accepts OrganizationSinkMap and OrganizationSinkMapOutput values.
// You can construct a concrete instance of `OrganizationSinkMapInput` via:
//
//          OrganizationSinkMap{ "key": OrganizationSinkArgs{...} }
type OrganizationSinkMapInput interface {
	pulumi.Input

	ToOrganizationSinkMapOutput() OrganizationSinkMapOutput
	ToOrganizationSinkMapOutputWithContext(context.Context) OrganizationSinkMapOutput
}

type OrganizationSinkMap map[string]OrganizationSinkInput

func (OrganizationSinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationSink)(nil)).Elem()
}

func (i OrganizationSinkMap) ToOrganizationSinkMapOutput() OrganizationSinkMapOutput {
	return i.ToOrganizationSinkMapOutputWithContext(context.Background())
}

func (i OrganizationSinkMap) ToOrganizationSinkMapOutputWithContext(ctx context.Context) OrganizationSinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSinkMapOutput)
}

type OrganizationSinkOutput struct{ *pulumi.OutputState }

func (OrganizationSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSink)(nil))
}

func (o OrganizationSinkOutput) ToOrganizationSinkOutput() OrganizationSinkOutput {
	return o
}

func (o OrganizationSinkOutput) ToOrganizationSinkOutputWithContext(ctx context.Context) OrganizationSinkOutput {
	return o
}

func (o OrganizationSinkOutput) ToOrganizationSinkPtrOutput() OrganizationSinkPtrOutput {
	return o.ToOrganizationSinkPtrOutputWithContext(context.Background())
}

func (o OrganizationSinkOutput) ToOrganizationSinkPtrOutputWithContext(ctx context.Context) OrganizationSinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationSink) *OrganizationSink {
		return &v
	}).(OrganizationSinkPtrOutput)
}

type OrganizationSinkPtrOutput struct{ *pulumi.OutputState }

func (OrganizationSinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationSink)(nil))
}

func (o OrganizationSinkPtrOutput) ToOrganizationSinkPtrOutput() OrganizationSinkPtrOutput {
	return o
}

func (o OrganizationSinkPtrOutput) ToOrganizationSinkPtrOutputWithContext(ctx context.Context) OrganizationSinkPtrOutput {
	return o
}

func (o OrganizationSinkPtrOutput) Elem() OrganizationSinkOutput {
	return o.ApplyT(func(v *OrganizationSink) OrganizationSink {
		if v != nil {
			return *v
		}
		var ret OrganizationSink
		return ret
	}).(OrganizationSinkOutput)
}

type OrganizationSinkArrayOutput struct{ *pulumi.OutputState }

func (OrganizationSinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrganizationSink)(nil))
}

func (o OrganizationSinkArrayOutput) ToOrganizationSinkArrayOutput() OrganizationSinkArrayOutput {
	return o
}

func (o OrganizationSinkArrayOutput) ToOrganizationSinkArrayOutputWithContext(ctx context.Context) OrganizationSinkArrayOutput {
	return o
}

func (o OrganizationSinkArrayOutput) Index(i pulumi.IntInput) OrganizationSinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OrganizationSink {
		return vs[0].([]OrganizationSink)[vs[1].(int)]
	}).(OrganizationSinkOutput)
}

type OrganizationSinkMapOutput struct{ *pulumi.OutputState }

func (OrganizationSinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OrganizationSink)(nil))
}

func (o OrganizationSinkMapOutput) ToOrganizationSinkMapOutput() OrganizationSinkMapOutput {
	return o
}

func (o OrganizationSinkMapOutput) ToOrganizationSinkMapOutputWithContext(ctx context.Context) OrganizationSinkMapOutput {
	return o
}

func (o OrganizationSinkMapOutput) MapIndex(k pulumi.StringInput) OrganizationSinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OrganizationSink {
		return vs[0].(map[string]OrganizationSink)[vs[1].(string)]
	}).(OrganizationSinkOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationSinkOutput{})
	pulumi.RegisterOutputType(OrganizationSinkPtrOutput{})
	pulumi.RegisterOutputType(OrganizationSinkArrayOutput{})
	pulumi.RegisterOutputType(OrganizationSinkMapOutput{})
}
