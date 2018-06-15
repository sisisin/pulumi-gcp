// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a Forwarding Rule within GCE. This binds an ip and port range to a target pool. For more
 * information see [the official
 * documentation](https://cloud.google.com/compute/docs/load-balancing/network/forwarding-rules) and
 * [API](https://cloud.google.com/compute/docs/reference/latest/forwardingRules).
 */
export class ForwardingRule extends pulumi.CustomResource {
    /**
     * Get an existing ForwardingRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ForwardingRuleState): ForwardingRule {
        return new ForwardingRule(name, <any>state, { id });
    }

    /**
     * BackendService resource to receive the
     * matched traffic. Only used for internal load balancing.
     */
    public readonly backendService: pulumi.Output<string | undefined>;
    /**
     * Textual description field.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The static IP. (if not set, an ephemeral IP is
     * used).
     */
    public readonly ipAddress: pulumi.Output<string>;
    /**
     * The IP protocol to route, one of "TCP" "UDP" "AH"
     * "ESP" or "SCTP" for external load balancing, "TCP" or "UDP" for internal
     * (default "TCP").
     */
    public readonly ipProtocol: pulumi.Output<string>;
    /**
     * Type of load balancing to use. Can be
     * set to "INTERNAL" or "EXTERNAL" (default "EXTERNAL").
     */
    public readonly loadBalancingScheme: pulumi.Output<string | undefined>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Network name or self_link that the load balanced IP
     * should belong to. Only used for internal load balancing. If it is not
     * provided, the default network is used.
     */
    public readonly network: pulumi.Output<string>;
    /**
     * A range e.g. "1024-2048" or a single port "1024"
     * (defaults to all ports!). Only used for external load balancing.
     * Some types of forwarding targets have constraints on the acceptable ports:
     * * Target HTTP proxy: 80, 8080
     * * Target HTTPS proxy: 443
     * * Target TCP proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target SSL proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target VPN gateway: 500, 4500
     */
    public readonly portRange: pulumi.Output<string | undefined>;
    /**
     * A list of ports (maximum of 5) to use for internal load
     * balancing. Packets addressed to these ports will be forwarded to the backends
     * configured with this forwarding rule. Required for internal load balancing.
     */
    public readonly ports: pulumi.Output<string[] | undefined>;
    /**
     * The ID of project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The Region in which the created address should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    /**
     * Subnetwork that the load balanced IP should belong
     * to. Only used for internal load balancing. Must be specified if the network
     * is in custom subnet mode.
     */
    public readonly subnetwork: pulumi.Output<string>;
    /**
     * URL of target pool. Required for external load
     * balancing.
     */
    public readonly target: pulumi.Output<string | undefined>;

    /**
     * Create a ForwardingRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ForwardingRuleArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: ForwardingRuleArgs | ForwardingRuleState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ForwardingRuleState = argsOrState as ForwardingRuleState | undefined;
            inputs["backendService"] = state ? state.backendService : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            inputs["loadBalancingScheme"] = state ? state.loadBalancingScheme : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["portRange"] = state ? state.portRange : undefined;
            inputs["ports"] = state ? state.ports : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as ForwardingRuleArgs | undefined;
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            inputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["portRange"] = args ? args.portRange : undefined;
            inputs["ports"] = args ? args.ports : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/forwardingRule:ForwardingRule", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ForwardingRule resources.
 */
export interface ForwardingRuleState {
    /**
     * BackendService resource to receive the
     * matched traffic. Only used for internal load balancing.
     */
    readonly backendService?: pulumi.Input<string>;
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The static IP. (if not set, an ephemeral IP is
     * used).
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The IP protocol to route, one of "TCP" "UDP" "AH"
     * "ESP" or "SCTP" for external load balancing, "TCP" or "UDP" for internal
     * (default "TCP").
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * Type of load balancing to use. Can be
     * set to "INTERNAL" or "EXTERNAL" (default "EXTERNAL").
     */
    readonly loadBalancingScheme?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Network name or self_link that the load balanced IP
     * should belong to. Only used for internal load balancing. If it is not
     * provided, the default network is used.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * A range e.g. "1024-2048" or a single port "1024"
     * (defaults to all ports!). Only used for external load balancing.
     * Some types of forwarding targets have constraints on the acceptable ports:
     * * Target HTTP proxy: 80, 8080
     * * Target HTTPS proxy: 443
     * * Target TCP proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target SSL proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target VPN gateway: 500, 4500
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * A list of ports (maximum of 5) to use for internal load
     * balancing. Packets addressed to these ports will be forwarded to the backends
     * configured with this forwarding rule. Required for internal load balancing.
     */
    readonly ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created address should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Subnetwork that the load balanced IP should belong
     * to. Only used for internal load balancing. Must be specified if the network
     * is in custom subnet mode.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * URL of target pool. Required for external load
     * balancing.
     */
    readonly target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ForwardingRule resource.
 */
export interface ForwardingRuleArgs {
    /**
     * BackendService resource to receive the
     * matched traffic. Only used for internal load balancing.
     */
    readonly backendService?: pulumi.Input<string>;
    /**
     * Textual description field.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The static IP. (if not set, an ephemeral IP is
     * used).
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The IP protocol to route, one of "TCP" "UDP" "AH"
     * "ESP" or "SCTP" for external load balancing, "TCP" or "UDP" for internal
     * (default "TCP").
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * Type of load balancing to use. Can be
     * set to "INTERNAL" or "EXTERNAL" (default "EXTERNAL").
     */
    readonly loadBalancingScheme?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE. Changing
     * this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Network name or self_link that the load balanced IP
     * should belong to. Only used for internal load balancing. If it is not
     * provided, the default network is used.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * A range e.g. "1024-2048" or a single port "1024"
     * (defaults to all ports!). Only used for external load balancing.
     * Some types of forwarding targets have constraints on the acceptable ports:
     * * Target HTTP proxy: 80, 8080
     * * Target HTTPS proxy: 443
     * * Target TCP proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target SSL proxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222
     * * Target VPN gateway: 500, 4500
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * A list of ports (maximum of 5) to use for internal load
     * balancing. Packets addressed to these ports will be forwarded to the backends
     * configured with this forwarding rule. Required for internal load balancing.
     */
    readonly ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created address should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Subnetwork that the load balanced IP should belong
     * to. Only used for internal load balancing. Must be specified if the network
     * is in custom subnet mode.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * URL of target pool. Required for external load
     * balancing.
     */
    readonly target?: pulumi.Input<string>;
}