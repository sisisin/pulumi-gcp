// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A policy is a collection of DNS rules applied to one or more Virtual
 * Private Cloud resources.
 *
 * To get more information about Policy, see:
 *
 * * [API documentation](https://cloud.google.com/dns/docs/reference/v1beta2/policies)
 * * How-to Guides
 *     * [Using DNS server policies](https://cloud.google.com/dns/zones/#using-dns-server-policies)
 *
 * ## Example Usage
 * ### Dns Policy Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const network_1 = new gcp.compute.Network("network-1", {autoCreateSubnetworks: false});
 * const network_2 = new gcp.compute.Network("network-2", {autoCreateSubnetworks: false});
 * const example_policy = new gcp.dns.Policy("example-policy", {
 *     enableInboundForwarding: true,
 *     enableLogging: true,
 *     alternativeNameServerConfig: {
 *         targetNameServers: [
 *             {
 *                 ipv4Address: "172.16.1.10",
 *                 forwardingPath: "private",
 *             },
 *             {
 *                 ipv4Address: "172.16.1.20",
 *             },
 *         ],
 *     },
 *     networks: [
 *         {
 *             networkUrl: network_1.id,
 *         },
 *         {
 *             networkUrl: network_2.id,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Policy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dns/policy:Policy default projects/{{project}}/policies/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dns/policy:Policy default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dns/policy:Policy default {{name}}
 * ```
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dns/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Sets an alternative name server for the associated networks.
     * When specified, all DNS queries are forwarded to a name server that you choose.
     * Names such as .internal are not available when an alternative name server is specified.
     * Structure is documented below.
     */
    public readonly alternativeNameServerConfig!: pulumi.Output<outputs.dns.PolicyAlternativeNameServerConfig | undefined>;
    /**
     * A textual description field. Defaults to 'Managed by Pulumi'.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Allows networks bound to this policy to receive DNS queries sent
     * by VMs or applications over VPN connections. When enabled, a
     * virtual IP address will be allocated from each of the sub-networks
     * that are bound to this policy.
     */
    public readonly enableInboundForwarding!: pulumi.Output<boolean | undefined>;
    /**
     * Controls whether logging is enabled for the networks bound to this policy.
     * Defaults to no logging if not set.
     */
    public readonly enableLogging!: pulumi.Output<boolean | undefined>;
    /**
     * User assigned name for this policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of network names specifying networks to which this policy is applied.
     * Structure is documented below.
     */
    public readonly networks!: pulumi.Output<outputs.dns.PolicyNetwork[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["alternativeNameServerConfig"] = state ? state.alternativeNameServerConfig : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableInboundForwarding"] = state ? state.enableInboundForwarding : undefined;
            inputs["enableLogging"] = state ? state.enableLogging : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networks"] = state ? state.networks : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            inputs["alternativeNameServerConfig"] = args ? args.alternativeNameServerConfig : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableInboundForwarding"] = args ? args.enableInboundForwarding : undefined;
            inputs["enableLogging"] = args ? args.enableLogging : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networks"] = args ? args.networks : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Sets an alternative name server for the associated networks.
     * When specified, all DNS queries are forwarded to a name server that you choose.
     * Names such as .internal are not available when an alternative name server is specified.
     * Structure is documented below.
     */
    alternativeNameServerConfig?: pulumi.Input<inputs.dns.PolicyAlternativeNameServerConfig>;
    /**
     * A textual description field. Defaults to 'Managed by Pulumi'.
     */
    description?: pulumi.Input<string>;
    /**
     * Allows networks bound to this policy to receive DNS queries sent
     * by VMs or applications over VPN connections. When enabled, a
     * virtual IP address will be allocated from each of the sub-networks
     * that are bound to this policy.
     */
    enableInboundForwarding?: pulumi.Input<boolean>;
    /**
     * Controls whether logging is enabled for the networks bound to this policy.
     * Defaults to no logging if not set.
     */
    enableLogging?: pulumi.Input<boolean>;
    /**
     * User assigned name for this policy.
     */
    name?: pulumi.Input<string>;
    /**
     * List of network names specifying networks to which this policy is applied.
     * Structure is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.dns.PolicyNetwork>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Sets an alternative name server for the associated networks.
     * When specified, all DNS queries are forwarded to a name server that you choose.
     * Names such as .internal are not available when an alternative name server is specified.
     * Structure is documented below.
     */
    alternativeNameServerConfig?: pulumi.Input<inputs.dns.PolicyAlternativeNameServerConfig>;
    /**
     * A textual description field. Defaults to 'Managed by Pulumi'.
     */
    description?: pulumi.Input<string>;
    /**
     * Allows networks bound to this policy to receive DNS queries sent
     * by VMs or applications over VPN connections. When enabled, a
     * virtual IP address will be allocated from each of the sub-networks
     * that are bound to this policy.
     */
    enableInboundForwarding?: pulumi.Input<boolean>;
    /**
     * Controls whether logging is enabled for the networks bound to this policy.
     * Defaults to no logging if not set.
     */
    enableLogging?: pulumi.Input<boolean>;
    /**
     * User assigned name for this policy.
     */
    name?: pulumi.Input<string>;
    /**
     * List of network names specifying networks to which this policy is applied.
     * Structure is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.dns.PolicyNetwork>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
