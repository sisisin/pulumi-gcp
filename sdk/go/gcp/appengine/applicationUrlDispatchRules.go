// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Rules to match an HTTP request and dispatch that request to a service.
//
// To get more information about ApplicationUrlDispatchRules, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)
//
// ## Example Usage
// ### App Engine Application Url Dispatch Rules Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/appengine"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := storage.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		object, err := storage.NewBucketObject(ctx, "object", &storage.BucketObjectArgs{
// 			Bucket: bucket.Name,
// 			Source: pulumi.NewFileAsset("./test-fixtures/appengine/hello-world.zip"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		adminV3, err := appengine.NewStandardAppVersion(ctx, "adminV3", &appengine.StandardAppVersionArgs{
// 			VersionId: pulumi.String("v3"),
// 			Service:   pulumi.String("admin"),
// 			Runtime:   pulumi.String("nodejs10"),
// 			Entrypoint: &appengine.StandardAppVersionEntrypointArgs{
// 				Shell: pulumi.String("node ./app.js"),
// 			},
// 			Deployment: &appengine.StandardAppVersionDeploymentArgs{
// 				Zip: &appengine.StandardAppVersionDeploymentZipArgs{
// 					SourceUrl: pulumi.All(bucket.Name, object.Name).ApplyT(func(_args []interface{}) (string, error) {
// 						bucketName := _args[0].(string)
// 						objectName := _args[1].(string)
// 						return fmt.Sprintf("%v%v%v%v", "https://storage.googleapis.com/", bucketName, "/", objectName), nil
// 					}).(pulumi.StringOutput),
// 				},
// 			},
// 			EnvVariables: pulumi.StringMap{
// 				"port": pulumi.String("8080"),
// 			},
// 			NoopOnDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appengine.NewApplicationUrlDispatchRules(ctx, "webService", &appengine.ApplicationUrlDispatchRulesArgs{
// 			DispatchRules: appengine.ApplicationUrlDispatchRulesDispatchRuleArray{
// 				&appengine.ApplicationUrlDispatchRulesDispatchRuleArgs{
// 					Domain:  pulumi.String("*"),
// 					Path:    pulumi.String("/*"),
// 					Service: pulumi.String("default"),
// 				},
// 				&appengine.ApplicationUrlDispatchRulesDispatchRuleArgs{
// 					Domain:  pulumi.String("*"),
// 					Path:    pulumi.String("/admin/*"),
// 					Service: adminV3.Service,
// 				},
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
// ApplicationUrlDispatchRules can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules default {{project}}
// ```
type ApplicationUrlDispatchRules struct {
	pulumi.CustomResourceState

	// Rules to match an HTTP request and dispatch that request to a service.
	// Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayOutput `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewApplicationUrlDispatchRules registers a new resource with the given unique name, arguments, and options.
func NewApplicationUrlDispatchRules(ctx *pulumi.Context,
	name string, args *ApplicationUrlDispatchRulesArgs, opts ...pulumi.ResourceOption) (*ApplicationUrlDispatchRules, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DispatchRules == nil {
		return nil, errors.New("invalid value for required argument 'DispatchRules'")
	}
	var resource ApplicationUrlDispatchRules
	err := ctx.RegisterResource("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationUrlDispatchRules gets an existing ApplicationUrlDispatchRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationUrlDispatchRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationUrlDispatchRulesState, opts ...pulumi.ResourceOption) (*ApplicationUrlDispatchRules, error) {
	var resource ApplicationUrlDispatchRules
	err := ctx.ReadResource("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationUrlDispatchRules resources.
type applicationUrlDispatchRulesState struct {
	// Rules to match an HTTP request and dispatch that request to a service.
	// Structure is documented below.
	DispatchRules []ApplicationUrlDispatchRulesDispatchRule `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ApplicationUrlDispatchRulesState struct {
	// Rules to match an HTTP request and dispatch that request to a service.
	// Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApplicationUrlDispatchRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationUrlDispatchRulesState)(nil)).Elem()
}

type applicationUrlDispatchRulesArgs struct {
	// Rules to match an HTTP request and dispatch that request to a service.
	// Structure is documented below.
	DispatchRules []ApplicationUrlDispatchRulesDispatchRule `pulumi:"dispatchRules"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ApplicationUrlDispatchRules resource.
type ApplicationUrlDispatchRulesArgs struct {
	// Rules to match an HTTP request and dispatch that request to a service.
	// Structure is documented below.
	DispatchRules ApplicationUrlDispatchRulesDispatchRuleArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApplicationUrlDispatchRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationUrlDispatchRulesArgs)(nil)).Elem()
}

type ApplicationUrlDispatchRulesInput interface {
	pulumi.Input

	ToApplicationUrlDispatchRulesOutput() ApplicationUrlDispatchRulesOutput
	ToApplicationUrlDispatchRulesOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesOutput
}

func (*ApplicationUrlDispatchRules) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUrlDispatchRules)(nil))
}

func (i *ApplicationUrlDispatchRules) ToApplicationUrlDispatchRulesOutput() ApplicationUrlDispatchRulesOutput {
	return i.ToApplicationUrlDispatchRulesOutputWithContext(context.Background())
}

func (i *ApplicationUrlDispatchRules) ToApplicationUrlDispatchRulesOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUrlDispatchRulesOutput)
}

func (i *ApplicationUrlDispatchRules) ToApplicationUrlDispatchRulesPtrOutput() ApplicationUrlDispatchRulesPtrOutput {
	return i.ToApplicationUrlDispatchRulesPtrOutputWithContext(context.Background())
}

func (i *ApplicationUrlDispatchRules) ToApplicationUrlDispatchRulesPtrOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUrlDispatchRulesPtrOutput)
}

type ApplicationUrlDispatchRulesPtrInput interface {
	pulumi.Input

	ToApplicationUrlDispatchRulesPtrOutput() ApplicationUrlDispatchRulesPtrOutput
	ToApplicationUrlDispatchRulesPtrOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesPtrOutput
}

type applicationUrlDispatchRulesPtrType ApplicationUrlDispatchRulesArgs

func (*applicationUrlDispatchRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUrlDispatchRules)(nil))
}

func (i *applicationUrlDispatchRulesPtrType) ToApplicationUrlDispatchRulesPtrOutput() ApplicationUrlDispatchRulesPtrOutput {
	return i.ToApplicationUrlDispatchRulesPtrOutputWithContext(context.Background())
}

func (i *applicationUrlDispatchRulesPtrType) ToApplicationUrlDispatchRulesPtrOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUrlDispatchRulesPtrOutput)
}

// ApplicationUrlDispatchRulesArrayInput is an input type that accepts ApplicationUrlDispatchRulesArray and ApplicationUrlDispatchRulesArrayOutput values.
// You can construct a concrete instance of `ApplicationUrlDispatchRulesArrayInput` via:
//
//          ApplicationUrlDispatchRulesArray{ ApplicationUrlDispatchRulesArgs{...} }
type ApplicationUrlDispatchRulesArrayInput interface {
	pulumi.Input

	ToApplicationUrlDispatchRulesArrayOutput() ApplicationUrlDispatchRulesArrayOutput
	ToApplicationUrlDispatchRulesArrayOutputWithContext(context.Context) ApplicationUrlDispatchRulesArrayOutput
}

type ApplicationUrlDispatchRulesArray []ApplicationUrlDispatchRulesInput

func (ApplicationUrlDispatchRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationUrlDispatchRules)(nil)).Elem()
}

func (i ApplicationUrlDispatchRulesArray) ToApplicationUrlDispatchRulesArrayOutput() ApplicationUrlDispatchRulesArrayOutput {
	return i.ToApplicationUrlDispatchRulesArrayOutputWithContext(context.Background())
}

func (i ApplicationUrlDispatchRulesArray) ToApplicationUrlDispatchRulesArrayOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUrlDispatchRulesArrayOutput)
}

// ApplicationUrlDispatchRulesMapInput is an input type that accepts ApplicationUrlDispatchRulesMap and ApplicationUrlDispatchRulesMapOutput values.
// You can construct a concrete instance of `ApplicationUrlDispatchRulesMapInput` via:
//
//          ApplicationUrlDispatchRulesMap{ "key": ApplicationUrlDispatchRulesArgs{...} }
type ApplicationUrlDispatchRulesMapInput interface {
	pulumi.Input

	ToApplicationUrlDispatchRulesMapOutput() ApplicationUrlDispatchRulesMapOutput
	ToApplicationUrlDispatchRulesMapOutputWithContext(context.Context) ApplicationUrlDispatchRulesMapOutput
}

type ApplicationUrlDispatchRulesMap map[string]ApplicationUrlDispatchRulesInput

func (ApplicationUrlDispatchRulesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationUrlDispatchRules)(nil)).Elem()
}

func (i ApplicationUrlDispatchRulesMap) ToApplicationUrlDispatchRulesMapOutput() ApplicationUrlDispatchRulesMapOutput {
	return i.ToApplicationUrlDispatchRulesMapOutputWithContext(context.Background())
}

func (i ApplicationUrlDispatchRulesMap) ToApplicationUrlDispatchRulesMapOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUrlDispatchRulesMapOutput)
}

type ApplicationUrlDispatchRulesOutput struct{ *pulumi.OutputState }

func (ApplicationUrlDispatchRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUrlDispatchRules)(nil))
}

func (o ApplicationUrlDispatchRulesOutput) ToApplicationUrlDispatchRulesOutput() ApplicationUrlDispatchRulesOutput {
	return o
}

func (o ApplicationUrlDispatchRulesOutput) ToApplicationUrlDispatchRulesOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesOutput {
	return o
}

func (o ApplicationUrlDispatchRulesOutput) ToApplicationUrlDispatchRulesPtrOutput() ApplicationUrlDispatchRulesPtrOutput {
	return o.ToApplicationUrlDispatchRulesPtrOutputWithContext(context.Background())
}

func (o ApplicationUrlDispatchRulesOutput) ToApplicationUrlDispatchRulesPtrOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationUrlDispatchRules) *ApplicationUrlDispatchRules {
		return &v
	}).(ApplicationUrlDispatchRulesPtrOutput)
}

type ApplicationUrlDispatchRulesPtrOutput struct{ *pulumi.OutputState }

func (ApplicationUrlDispatchRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUrlDispatchRules)(nil))
}

func (o ApplicationUrlDispatchRulesPtrOutput) ToApplicationUrlDispatchRulesPtrOutput() ApplicationUrlDispatchRulesPtrOutput {
	return o
}

func (o ApplicationUrlDispatchRulesPtrOutput) ToApplicationUrlDispatchRulesPtrOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesPtrOutput {
	return o
}

func (o ApplicationUrlDispatchRulesPtrOutput) Elem() ApplicationUrlDispatchRulesOutput {
	return o.ApplyT(func(v *ApplicationUrlDispatchRules) ApplicationUrlDispatchRules {
		if v != nil {
			return *v
		}
		var ret ApplicationUrlDispatchRules
		return ret
	}).(ApplicationUrlDispatchRulesOutput)
}

type ApplicationUrlDispatchRulesArrayOutput struct{ *pulumi.OutputState }

func (ApplicationUrlDispatchRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUrlDispatchRules)(nil))
}

func (o ApplicationUrlDispatchRulesArrayOutput) ToApplicationUrlDispatchRulesArrayOutput() ApplicationUrlDispatchRulesArrayOutput {
	return o
}

func (o ApplicationUrlDispatchRulesArrayOutput) ToApplicationUrlDispatchRulesArrayOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesArrayOutput {
	return o
}

func (o ApplicationUrlDispatchRulesArrayOutput) Index(i pulumi.IntInput) ApplicationUrlDispatchRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationUrlDispatchRules {
		return vs[0].([]ApplicationUrlDispatchRules)[vs[1].(int)]
	}).(ApplicationUrlDispatchRulesOutput)
}

type ApplicationUrlDispatchRulesMapOutput struct{ *pulumi.OutputState }

func (ApplicationUrlDispatchRulesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationUrlDispatchRules)(nil))
}

func (o ApplicationUrlDispatchRulesMapOutput) ToApplicationUrlDispatchRulesMapOutput() ApplicationUrlDispatchRulesMapOutput {
	return o
}

func (o ApplicationUrlDispatchRulesMapOutput) ToApplicationUrlDispatchRulesMapOutputWithContext(ctx context.Context) ApplicationUrlDispatchRulesMapOutput {
	return o
}

func (o ApplicationUrlDispatchRulesMapOutput) MapIndex(k pulumi.StringInput) ApplicationUrlDispatchRulesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationUrlDispatchRules {
		return vs[0].(map[string]ApplicationUrlDispatchRules)[vs[1].(string)]
	}).(ApplicationUrlDispatchRulesOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationUrlDispatchRulesOutput{})
	pulumi.RegisterOutputType(ApplicationUrlDispatchRulesPtrOutput{})
	pulumi.RegisterOutputType(ApplicationUrlDispatchRulesArrayOutput{})
	pulumi.RegisterOutputType(ApplicationUrlDispatchRulesMapOutput{})
}
