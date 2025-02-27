// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hierarchical firewall policy rules let you create and enforce a consistent firewall policy across your organization. Rules can explicitly allow or deny connections or delegate evaluation to lower level policies.
//
// For more information see the [official documentation](https://cloud.google.com/vpc/docs/using-firewall-policies#create-rules)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultFirewallPolicy, err := compute.NewFirewallPolicy(ctx, "defaultFirewallPolicy", &compute.FirewallPolicyArgs{
// 			Parent:      pulumi.String("organizations/12345"),
// 			ShortName:   pulumi.String("my-policy"),
// 			Description: pulumi.String("Example Resource"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewFirewallPolicyRule(ctx, "defaultFirewallPolicyRule", &compute.FirewallPolicyRuleArgs{
// 			FirewallPolicy: defaultFirewallPolicy.ID(),
// 			Description:    pulumi.String("Example Resource"),
// 			Priority:       pulumi.Int(9000),
// 			EnableLogging:  pulumi.Bool(true),
// 			Action:         pulumi.String("allow"),
// 			Direction:      pulumi.String("EGRESS"),
// 			Disabled:       pulumi.Bool(false),
// 			Match: &compute.FirewallPolicyRuleMatchArgs{
// 				Layer4Configs: compute.FirewallPolicyRuleMatchLayer4ConfigArray{
// 					&compute.FirewallPolicyRuleMatchLayer4ConfigArgs{
// 						IpProtocol: pulumi.String("tcp"),
// 						Ports: pulumi.StringArray{
// 							pulumi.String("80"),
// 							pulumi.String("8080"),
// 						},
// 					},
// 				},
// 				DestIpRanges: pulumi.StringArray{
// 					pulumi.String("11.100.0.1/32"),
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
// FirewallPolicyRule can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/firewallPolicyRule:FirewallPolicyRule default locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/firewallPolicyRule:FirewallPolicyRule default {{firewall_policy}}/{{priority}}
// ```
type FirewallPolicyRule struct {
	pulumi.CustomResourceState

	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	Action pulumi.StringOutput `pulumi:"action"`
	// An optional description for this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction pulumi.StringOutput `pulumi:"direction"`
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// The firewall policy of the resource.
	FirewallPolicy pulumi.StringOutput `pulumi:"firewallPolicy"`
	// Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match FirewallPolicyRuleMatchOutput `pulumi:"match"`
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Calculation of the complexity of a single firewall policy rule.
	RuleTupleCount pulumi.IntOutput `pulumi:"ruleTupleCount"`
	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	TargetResources pulumi.StringArrayOutput `pulumi:"targetResources"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayOutput `pulumi:"targetServiceAccounts"`
}

// NewFirewallPolicyRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicyRule(ctx *pulumi.Context,
	name string, args *FirewallPolicyRuleArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.FirewallPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicy'")
	}
	if args.Match == nil {
		return nil, errors.New("invalid value for required argument 'Match'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	var resource FirewallPolicyRule
	err := ctx.RegisterResource("gcp:compute/firewallPolicyRule:FirewallPolicyRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicyRule gets an existing FirewallPolicyRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicyRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyRuleState, opts ...pulumi.ResourceOption) (*FirewallPolicyRule, error) {
	var resource FirewallPolicyRule
	err := ctx.ReadResource("gcp:compute/firewallPolicyRule:FirewallPolicyRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicyRule resources.
type firewallPolicyRuleState struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	Action *string `pulumi:"action"`
	// An optional description for this resource.
	Description *string `pulumi:"description"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction *string `pulumi:"direction"`
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled *bool `pulumi:"disabled"`
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging *bool `pulumi:"enableLogging"`
	// The firewall policy of the resource.
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
	Kind *string `pulumi:"kind"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match *FirewallPolicyRuleMatch `pulumi:"match"`
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority *int `pulumi:"priority"`
	// Calculation of the complexity of a single firewall policy rule.
	RuleTupleCount *int `pulumi:"ruleTupleCount"`
	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	TargetResources []string `pulumi:"targetResources"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts []string `pulumi:"targetServiceAccounts"`
}

type FirewallPolicyRuleState struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	Action pulumi.StringPtrInput
	// An optional description for this resource.
	Description pulumi.StringPtrInput
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction pulumi.StringPtrInput
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled pulumi.BoolPtrInput
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging pulumi.BoolPtrInput
	// The firewall policy of the resource.
	FirewallPolicy pulumi.StringPtrInput
	// Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
	Kind pulumi.StringPtrInput
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match FirewallPolicyRuleMatchPtrInput
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntPtrInput
	// Calculation of the complexity of a single firewall policy rule.
	RuleTupleCount pulumi.IntPtrInput
	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	TargetResources pulumi.StringArrayInput
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayInput
}

func (FirewallPolicyRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleState)(nil)).Elem()
}

type firewallPolicyRuleArgs struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	Action string `pulumi:"action"`
	// An optional description for this resource.
	Description *string `pulumi:"description"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction string `pulumi:"direction"`
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled *bool `pulumi:"disabled"`
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging *bool `pulumi:"enableLogging"`
	// The firewall policy of the resource.
	FirewallPolicy string `pulumi:"firewallPolicy"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match FirewallPolicyRuleMatch `pulumi:"match"`
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority int `pulumi:"priority"`
	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	TargetResources []string `pulumi:"targetResources"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts []string `pulumi:"targetServiceAccounts"`
}

// The set of arguments for constructing a FirewallPolicyRule resource.
type FirewallPolicyRuleArgs struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	Action pulumi.StringInput
	// An optional description for this resource.
	Description pulumi.StringPtrInput
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction pulumi.StringInput
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled pulumi.BoolPtrInput
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging pulumi.BoolPtrInput
	// The firewall policy of the resource.
	FirewallPolicy pulumi.StringInput
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match FirewallPolicyRuleMatchInput
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntInput
	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	TargetResources pulumi.StringArrayInput
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayInput
}

func (FirewallPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleArgs)(nil)).Elem()
}

type FirewallPolicyRuleInput interface {
	pulumi.Input

	ToFirewallPolicyRuleOutput() FirewallPolicyRuleOutput
	ToFirewallPolicyRuleOutputWithContext(ctx context.Context) FirewallPolicyRuleOutput
}

func (*FirewallPolicyRule) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRule)(nil))
}

func (i *FirewallPolicyRule) ToFirewallPolicyRuleOutput() FirewallPolicyRuleOutput {
	return i.ToFirewallPolicyRuleOutputWithContext(context.Background())
}

func (i *FirewallPolicyRule) ToFirewallPolicyRuleOutputWithContext(ctx context.Context) FirewallPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleOutput)
}

func (i *FirewallPolicyRule) ToFirewallPolicyRulePtrOutput() FirewallPolicyRulePtrOutput {
	return i.ToFirewallPolicyRulePtrOutputWithContext(context.Background())
}

func (i *FirewallPolicyRule) ToFirewallPolicyRulePtrOutputWithContext(ctx context.Context) FirewallPolicyRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRulePtrOutput)
}

type FirewallPolicyRulePtrInput interface {
	pulumi.Input

	ToFirewallPolicyRulePtrOutput() FirewallPolicyRulePtrOutput
	ToFirewallPolicyRulePtrOutputWithContext(ctx context.Context) FirewallPolicyRulePtrOutput
}

type firewallPolicyRulePtrType FirewallPolicyRuleArgs

func (*firewallPolicyRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRule)(nil))
}

func (i *firewallPolicyRulePtrType) ToFirewallPolicyRulePtrOutput() FirewallPolicyRulePtrOutput {
	return i.ToFirewallPolicyRulePtrOutputWithContext(context.Background())
}

func (i *firewallPolicyRulePtrType) ToFirewallPolicyRulePtrOutputWithContext(ctx context.Context) FirewallPolicyRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRulePtrOutput)
}

// FirewallPolicyRuleArrayInput is an input type that accepts FirewallPolicyRuleArray and FirewallPolicyRuleArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyRuleArrayInput` via:
//
//          FirewallPolicyRuleArray{ FirewallPolicyRuleArgs{...} }
type FirewallPolicyRuleArrayInput interface {
	pulumi.Input

	ToFirewallPolicyRuleArrayOutput() FirewallPolicyRuleArrayOutput
	ToFirewallPolicyRuleArrayOutputWithContext(context.Context) FirewallPolicyRuleArrayOutput
}

type FirewallPolicyRuleArray []FirewallPolicyRuleInput

func (FirewallPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicyRule)(nil)).Elem()
}

func (i FirewallPolicyRuleArray) ToFirewallPolicyRuleArrayOutput() FirewallPolicyRuleArrayOutput {
	return i.ToFirewallPolicyRuleArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyRuleArray) ToFirewallPolicyRuleArrayOutputWithContext(ctx context.Context) FirewallPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleArrayOutput)
}

// FirewallPolicyRuleMapInput is an input type that accepts FirewallPolicyRuleMap and FirewallPolicyRuleMapOutput values.
// You can construct a concrete instance of `FirewallPolicyRuleMapInput` via:
//
//          FirewallPolicyRuleMap{ "key": FirewallPolicyRuleArgs{...} }
type FirewallPolicyRuleMapInput interface {
	pulumi.Input

	ToFirewallPolicyRuleMapOutput() FirewallPolicyRuleMapOutput
	ToFirewallPolicyRuleMapOutputWithContext(context.Context) FirewallPolicyRuleMapOutput
}

type FirewallPolicyRuleMap map[string]FirewallPolicyRuleInput

func (FirewallPolicyRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicyRule)(nil)).Elem()
}

func (i FirewallPolicyRuleMap) ToFirewallPolicyRuleMapOutput() FirewallPolicyRuleMapOutput {
	return i.ToFirewallPolicyRuleMapOutputWithContext(context.Background())
}

func (i FirewallPolicyRuleMap) ToFirewallPolicyRuleMapOutputWithContext(ctx context.Context) FirewallPolicyRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleMapOutput)
}

type FirewallPolicyRuleOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRule)(nil))
}

func (o FirewallPolicyRuleOutput) ToFirewallPolicyRuleOutput() FirewallPolicyRuleOutput {
	return o
}

func (o FirewallPolicyRuleOutput) ToFirewallPolicyRuleOutputWithContext(ctx context.Context) FirewallPolicyRuleOutput {
	return o
}

func (o FirewallPolicyRuleOutput) ToFirewallPolicyRulePtrOutput() FirewallPolicyRulePtrOutput {
	return o.ToFirewallPolicyRulePtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleOutput) ToFirewallPolicyRulePtrOutputWithContext(ctx context.Context) FirewallPolicyRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyRule) *FirewallPolicyRule {
		return &v
	}).(FirewallPolicyRulePtrOutput)
}

type FirewallPolicyRulePtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRule)(nil))
}

func (o FirewallPolicyRulePtrOutput) ToFirewallPolicyRulePtrOutput() FirewallPolicyRulePtrOutput {
	return o
}

func (o FirewallPolicyRulePtrOutput) ToFirewallPolicyRulePtrOutputWithContext(ctx context.Context) FirewallPolicyRulePtrOutput {
	return o
}

func (o FirewallPolicyRulePtrOutput) Elem() FirewallPolicyRuleOutput {
	return o.ApplyT(func(v *FirewallPolicyRule) FirewallPolicyRule {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyRule
		return ret
	}).(FirewallPolicyRuleOutput)
}

type FirewallPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallPolicyRule)(nil))
}

func (o FirewallPolicyRuleArrayOutput) ToFirewallPolicyRuleArrayOutput() FirewallPolicyRuleArrayOutput {
	return o
}

func (o FirewallPolicyRuleArrayOutput) ToFirewallPolicyRuleArrayOutputWithContext(ctx context.Context) FirewallPolicyRuleArrayOutput {
	return o
}

func (o FirewallPolicyRuleArrayOutput) Index(i pulumi.IntInput) FirewallPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallPolicyRule {
		return vs[0].([]FirewallPolicyRule)[vs[1].(int)]
	}).(FirewallPolicyRuleOutput)
}

type FirewallPolicyRuleMapOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallPolicyRule)(nil))
}

func (o FirewallPolicyRuleMapOutput) ToFirewallPolicyRuleMapOutput() FirewallPolicyRuleMapOutput {
	return o
}

func (o FirewallPolicyRuleMapOutput) ToFirewallPolicyRuleMapOutputWithContext(ctx context.Context) FirewallPolicyRuleMapOutput {
	return o
}

func (o FirewallPolicyRuleMapOutput) MapIndex(k pulumi.StringInput) FirewallPolicyRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallPolicyRule {
		return vs[0].(map[string]FirewallPolicyRule)[vs[1].(string)]
	}).(FirewallPolicyRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyRuleOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRulePtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleMapOutput{})
}
