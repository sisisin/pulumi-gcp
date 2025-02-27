// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtasks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A named resource to which messages are sent by publishers.
//
// > **Warning:** This resource requires an App Engine application to be created on the
// project you're provisioning it on. If you haven't already enabled it, you
// can create a `appengine.Application` resource to do so. This
// resource's location will be the same as the App Engine location specified.
//
// ## Example Usage
// ### Queue Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/cloudtasks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudtasks.NewQueue(ctx, "_default", &cloudtasks.QueueArgs{
// 			Location: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloud Tasks Queue Advanced
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/cloudtasks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudtasks.NewQueue(ctx, "advancedConfiguration", &cloudtasks.QueueArgs{
// 			AppEngineRoutingOverride: &cloudtasks.QueueAppEngineRoutingOverrideArgs{
// 				Instance: pulumi.String("test"),
// 				Service:  pulumi.String("worker"),
// 				Version:  pulumi.String("1.0"),
// 			},
// 			Location: pulumi.String("us-central1"),
// 			RateLimits: &cloudtasks.QueueRateLimitsArgs{
// 				MaxConcurrentDispatches: pulumi.Int(3),
// 				MaxDispatchesPerSecond:  pulumi.Float64(2),
// 			},
// 			RetryConfig: &cloudtasks.QueueRetryConfigArgs{
// 				MaxAttempts:      pulumi.Int(5),
// 				MaxBackoff:       pulumi.String("3s"),
// 				MaxDoublings:     pulumi.Int(1),
// 				MaxRetryDuration: pulumi.String("4s"),
// 				MinBackoff:       pulumi.String("2s"),
// 			},
// 			StackdriverLoggingConfig: &cloudtasks.QueueStackdriverLoggingConfigArgs{
// 				SamplingRatio: pulumi.Float64(0.9),
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
// Queue can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudtasks/queue:Queue default projects/{{project}}/locations/{{location}}/queues/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudtasks/queue:Queue default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudtasks/queue:Queue default {{location}}/{{name}}
// ```
type Queue struct {
	pulumi.CustomResourceState

	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride QueueAppEngineRoutingOverridePtrOutput `pulumi:"appEngineRoutingOverride"`
	// The location of the queue
	Location pulumi.StringOutput `pulumi:"location"`
	// The queue name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits QueueRateLimitsOutput `pulumi:"rateLimits"`
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig QueueRetryConfigOutput `pulumi:"retryConfig"`
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig QueueStackdriverLoggingConfigPtrOutput `pulumi:"stackdriverLoggingConfig"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource Queue
	err := ctx.RegisterResource("gcp:cloudtasks/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("gcp:cloudtasks/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride *QueueAppEngineRoutingOverride `pulumi:"appEngineRoutingOverride"`
	// The location of the queue
	Location *string `pulumi:"location"`
	// The queue name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits *QueueRateLimits `pulumi:"rateLimits"`
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig *QueueRetryConfig `pulumi:"retryConfig"`
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig *QueueStackdriverLoggingConfig `pulumi:"stackdriverLoggingConfig"`
}

type QueueState struct {
	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride QueueAppEngineRoutingOverridePtrInput
	// The location of the queue
	Location pulumi.StringPtrInput
	// The queue name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits QueueRateLimitsPtrInput
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig QueueRetryConfigPtrInput
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig QueueStackdriverLoggingConfigPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride *QueueAppEngineRoutingOverride `pulumi:"appEngineRoutingOverride"`
	// The location of the queue
	Location string `pulumi:"location"`
	// The queue name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits *QueueRateLimits `pulumi:"rateLimits"`
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig *QueueRetryConfig `pulumi:"retryConfig"`
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig *QueueStackdriverLoggingConfig `pulumi:"stackdriverLoggingConfig"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// Overrides for task-level appEngineRouting. These settings apply only
	// to App Engine tasks in this queue
	// Structure is documented below.
	AppEngineRoutingOverride QueueAppEngineRoutingOverridePtrInput
	// The location of the queue
	Location pulumi.StringInput
	// The queue name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Rate limits for task dispatches.
	// The queue's actual dispatch rate is the result of:
	// * Number of tasks in the queue
	// * User-specified throttling: rateLimits, retryConfig, and the queue's state.
	// * System throttling due to 429 (Too Many Requests) or 503 (Service
	//   Unavailable) responses from the worker, high error rates, or to
	//   smooth sudden large traffic spikes.
	//   Structure is documented below.
	RateLimits QueueRateLimitsPtrInput
	// Settings that determine the retry behavior.
	// Structure is documented below.
	RetryConfig QueueRetryConfigPtrInput
	// Configuration options for writing logs to Stackdriver Logging.
	// Structure is documented below.
	StackdriverLoggingConfig QueueStackdriverLoggingConfigPtrInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil))
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

func (i *Queue) ToQueuePtrOutput() QueuePtrOutput {
	return i.ToQueuePtrOutputWithContext(context.Background())
}

func (i *Queue) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePtrOutput)
}

type QueuePtrInput interface {
	pulumi.Input

	ToQueuePtrOutput() QueuePtrOutput
	ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput
}

type queuePtrType QueueArgs

func (*queuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil))
}

func (i *queuePtrType) ToQueuePtrOutput() QueuePtrOutput {
	return i.ToQueuePtrOutputWithContext(context.Background())
}

func (i *queuePtrType) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePtrOutput)
}

// QueueArrayInput is an input type that accepts QueueArray and QueueArrayOutput values.
// You can construct a concrete instance of `QueueArrayInput` via:
//
//          QueueArray{ QueueArgs{...} }
type QueueArrayInput interface {
	pulumi.Input

	ToQueueArrayOutput() QueueArrayOutput
	ToQueueArrayOutputWithContext(context.Context) QueueArrayOutput
}

type QueueArray []QueueInput

func (QueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (i QueueArray) ToQueueArrayOutput() QueueArrayOutput {
	return i.ToQueueArrayOutputWithContext(context.Background())
}

func (i QueueArray) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueArrayOutput)
}

// QueueMapInput is an input type that accepts QueueMap and QueueMapOutput values.
// You can construct a concrete instance of `QueueMapInput` via:
//
//          QueueMap{ "key": QueueArgs{...} }
type QueueMapInput interface {
	pulumi.Input

	ToQueueMapOutput() QueueMapOutput
	ToQueueMapOutputWithContext(context.Context) QueueMapOutput
}

type QueueMap map[string]QueueInput

func (QueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (i QueueMap) ToQueueMapOutput() QueueMapOutput {
	return i.ToQueueMapOutputWithContext(context.Background())
}

func (i QueueMap) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueMapOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil))
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func (o QueueOutput) ToQueuePtrOutput() QueuePtrOutput {
	return o.ToQueuePtrOutputWithContext(context.Background())
}

func (o QueueOutput) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Queue) *Queue {
		return &v
	}).(QueuePtrOutput)
}

type QueuePtrOutput struct{ *pulumi.OutputState }

func (QueuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil))
}

func (o QueuePtrOutput) ToQueuePtrOutput() QueuePtrOutput {
	return o
}

func (o QueuePtrOutput) ToQueuePtrOutputWithContext(ctx context.Context) QueuePtrOutput {
	return o
}

func (o QueuePtrOutput) Elem() QueueOutput {
	return o.ApplyT(func(v *Queue) Queue {
		if v != nil {
			return *v
		}
		var ret Queue
		return ret
	}).(QueueOutput)
}

type QueueArrayOutput struct{ *pulumi.OutputState }

func (QueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Queue)(nil))
}

func (o QueueArrayOutput) ToQueueArrayOutput() QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) Index(i pulumi.IntInput) QueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Queue {
		return vs[0].([]Queue)[vs[1].(int)]
	}).(QueueOutput)
}

type QueueMapOutput struct{ *pulumi.OutputState }

func (QueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Queue)(nil))
}

func (o QueueMapOutput) ToQueueMapOutput() QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return o
}

func (o QueueMapOutput) MapIndex(k pulumi.StringInput) QueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Queue {
		return vs[0].(map[string]Queue)[vs[1].(string)]
	}).(QueueOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueuePtrOutput{})
	pulumi.RegisterOutputType(QueueArrayOutput{})
	pulumi.RegisterOutputType(QueueMapOutput{})
}
