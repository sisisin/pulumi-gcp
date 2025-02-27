// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A VPC network is a virtual version of the traditional physical networks
 * that exist within and between physical data centers. A VPC network
 * provides connectivity for your Compute Engine virtual machine (VM)
 * instances, Container Engine containers, App Engine Flex services, and
 * other network-related resources.
 *
 * Each GCP project contains one or more VPC networks. Each VPC network is a
 * global entity spanning all GCP regions. This global VPC network allows VM
 * instances and other resources to communicate with each other via internal,
 * private IP addresses.
 *
 * Each VPC network is subdivided into subnets, and each subnet is contained
 * within a single region. You can have more than one subnet in a region for
 * a given VPC network. Each subnet has a contiguous private RFC1918 IP
 * space. You create instances, containers, and the like in these subnets.
 * When you create an instance, you must create it in a subnet, and the
 * instance draws its internal IP address from that subnet.
 *
 * Virtual machine (VM) instances in a VPC network can communicate with
 * instances in all other subnets of the same VPC network, regardless of
 * region, using their RFC1918 private IP addresses. You can isolate portions
 * of the network, even entire subnets, using firewall rules.
 *
 * To get more information about Subnetwork, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/subnetworks)
 * * How-to Guides
 *     * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
 *     * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)
 *
 * ## Example Usage
 * ### Subnetwork Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const custom_test = new gcp.compute.Network("custom-test", {autoCreateSubnetworks: false});
 * const network_with_private_secondary_ip_ranges = new gcp.compute.Subnetwork("network-with-private-secondary-ip-ranges", {
 *     ipCidrRange: "10.2.0.0/16",
 *     region: "us-central1",
 *     network: custom_test.id,
 *     secondaryIpRanges: [{
 *         rangeName: "tf-test-secondary-range-update1",
 *         ipCidrRange: "192.168.10.0/24",
 *     }],
 * });
 * ```
 * ### Subnetwork Logging Config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const custom_test = new gcp.compute.Network("custom-test", {autoCreateSubnetworks: false});
 * const subnet_with_logging = new gcp.compute.Subnetwork("subnet-with-logging", {
 *     ipCidrRange: "10.2.0.0/16",
 *     region: "us-central1",
 *     network: custom_test.id,
 *     logConfig: {
 *         aggregationInterval: "INTERVAL_10_MIN",
 *         flowSampling: 0.5,
 *         metadata: "INCLUDE_ALL_METADATA",
 *     },
 * });
 * ```
 * ### Subnetwork Internal L7lb
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const custom_test = new gcp.compute.Network("custom-test", {autoCreateSubnetworks: false}, {
 *     provider: google_beta,
 * });
 * const network_for_l7lb = new gcp.compute.Subnetwork("network-for-l7lb", {
 *     ipCidrRange: "10.0.0.0/22",
 *     region: "us-central1",
 *     purpose: "INTERNAL_HTTPS_LOAD_BALANCER",
 *     role: "ACTIVE",
 *     network: custom_test.id,
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * Subnetwork can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/subnetwork:Subnetwork default projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/subnetwork:Subnetwork default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/subnetwork:Subnetwork default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/subnetwork:Subnetwork default {{name}}
 * ```
 */
export class Subnetwork extends pulumi.CustomResource {
    /**
     * Get an existing Subnetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetworkState, opts?: pulumi.CustomResourceOptions): Subnetwork {
        return new Subnetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/subnetwork:Subnetwork';

    /**
     * Returns true if the given object is an instance of Subnetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnetwork.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource. This field can be set only at resource
     * creation time.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Fingerprint of this resource. This field is used internally during updates of this resource.
     *
     * @deprecated This field is not useful for users, and has been removed as an output.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The gateway address for default routes to reach destination addresses outside this subnetwork.
     */
    public /*out*/ readonly gatewayAddress!: pulumi.Output<string>;
    /**
     * The range of IP addresses belonging to this subnetwork secondary
     * range. Provide this property when you create the subnetwork.
     * Ranges must be unique and non-overlapping with all primary and
     * secondary IP ranges within a network. Only IPv4 is supported.
     */
    public readonly ipCidrRange!: pulumi.Output<string>;
    /**
     * Denotes the logging options for the subnetwork flow logs. If logging is enabled
     * logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
     * subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
     * Structure is documented below.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.SubnetworkLogConfig | undefined>;
    /**
     * The name of the resource, provided by the client when initially
     * creating the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?` which
     * means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network this subnet belongs to.
     * Only networks that are in the distributed mode can have subnetworks.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * When enabled, VMs in this subnetwork without external IP addresses can
     * access Google APIs and services by using Private Google Access.
     */
    public readonly privateIpGoogleAccess!: pulumi.Output<boolean | undefined>;
    /**
     * The private IPv6 google access type for the VMs in this subnet.
     */
    public readonly privateIpv6GoogleAccess!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The purpose of the resource. This field can be either PRIVATE
     * or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
     * INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
     * reserved for Internal HTTP(S) Load Balancing. If unspecified, the
     * purpose defaults to PRIVATE.
     * If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set `role`.
     */
    public readonly purpose!: pulumi.Output<string>;
    /**
     * The GCP region for this subnetwork.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The role of subnetwork. Currently, this field is only used when
     * purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
     * or BACKUP. An ACTIVE subnetwork is one that is currently being used
     * for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
     * is ready to be promoted to ACTIVE or is currently draining.
     * Possible values are `ACTIVE` and `BACKUP`.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * An array of configurations for secondary IP ranges for VM instances
     * contained in this subnetwork. The primary IP of such VM must belong
     * to the primary ipCidrRange of the subnetwork. The alias IPs may belong
     * to either primary or secondary ranges.
     * Structure is documented below.
     */
    public readonly secondaryIpRanges!: pulumi.Output<outputs.compute.SubnetworkSecondaryIpRange[]>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a Subnetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetworkArgs | SubnetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubnetworkState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["gatewayAddress"] = state ? state.gatewayAddress : undefined;
            inputs["ipCidrRange"] = state ? state.ipCidrRange : undefined;
            inputs["logConfig"] = state ? state.logConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["privateIpGoogleAccess"] = state ? state.privateIpGoogleAccess : undefined;
            inputs["privateIpv6GoogleAccess"] = state ? state.privateIpv6GoogleAccess : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["purpose"] = state ? state.purpose : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["secondaryIpRanges"] = state ? state.secondaryIpRanges : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as SubnetworkArgs | undefined;
            if ((!args || args.ipCidrRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipCidrRange'");
            }
            if ((!args || args.network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'network'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ipCidrRange"] = args ? args.ipCidrRange : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["privateIpGoogleAccess"] = args ? args.privateIpGoogleAccess : undefined;
            inputs["privateIpv6GoogleAccess"] = args ? args.privateIpv6GoogleAccess : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["purpose"] = args ? args.purpose : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["secondaryIpRanges"] = args ? args.secondaryIpRanges : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["gatewayAddress"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Subnetwork.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subnetwork resources.
 */
export interface SubnetworkState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource. This field can be set only at resource
     * creation time.
     */
    description?: pulumi.Input<string>;
    /**
     * Fingerprint of this resource. This field is used internally during updates of this resource.
     *
     * @deprecated This field is not useful for users, and has been removed as an output.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The gateway address for default routes to reach destination addresses outside this subnetwork.
     */
    gatewayAddress?: pulumi.Input<string>;
    /**
     * The range of IP addresses belonging to this subnetwork secondary
     * range. Provide this property when you create the subnetwork.
     * Ranges must be unique and non-overlapping with all primary and
     * secondary IP ranges within a network. Only IPv4 is supported.
     */
    ipCidrRange?: pulumi.Input<string>;
    /**
     * Denotes the logging options for the subnetwork flow logs. If logging is enabled
     * logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
     * subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
     * Structure is documented below.
     */
    logConfig?: pulumi.Input<inputs.compute.SubnetworkLogConfig>;
    /**
     * The name of the resource, provided by the client when initially
     * creating the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?` which
     * means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The network this subnet belongs to.
     * Only networks that are in the distributed mode can have subnetworks.
     */
    network?: pulumi.Input<string>;
    /**
     * When enabled, VMs in this subnetwork without external IP addresses can
     * access Google APIs and services by using Private Google Access.
     */
    privateIpGoogleAccess?: pulumi.Input<boolean>;
    /**
     * The private IPv6 google access type for the VMs in this subnet.
     */
    privateIpv6GoogleAccess?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The purpose of the resource. This field can be either PRIVATE
     * or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
     * INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
     * reserved for Internal HTTP(S) Load Balancing. If unspecified, the
     * purpose defaults to PRIVATE.
     * If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set `role`.
     */
    purpose?: pulumi.Input<string>;
    /**
     * The GCP region for this subnetwork.
     */
    region?: pulumi.Input<string>;
    /**
     * The role of subnetwork. Currently, this field is only used when
     * purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
     * or BACKUP. An ACTIVE subnetwork is one that is currently being used
     * for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
     * is ready to be promoted to ACTIVE or is currently draining.
     * Possible values are `ACTIVE` and `BACKUP`.
     */
    role?: pulumi.Input<string>;
    /**
     * An array of configurations for secondary IP ranges for VM instances
     * contained in this subnetwork. The primary IP of such VM must belong
     * to the primary ipCidrRange of the subnetwork. The alias IPs may belong
     * to either primary or secondary ranges.
     * Structure is documented below.
     */
    secondaryIpRanges?: pulumi.Input<pulumi.Input<inputs.compute.SubnetworkSecondaryIpRange>[]>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subnetwork resource.
 */
export interface SubnetworkArgs {
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource. This field can be set only at resource
     * creation time.
     */
    description?: pulumi.Input<string>;
    /**
     * The range of IP addresses belonging to this subnetwork secondary
     * range. Provide this property when you create the subnetwork.
     * Ranges must be unique and non-overlapping with all primary and
     * secondary IP ranges within a network. Only IPv4 is supported.
     */
    ipCidrRange: pulumi.Input<string>;
    /**
     * Denotes the logging options for the subnetwork flow logs. If logging is enabled
     * logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
     * subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
     * Structure is documented below.
     */
    logConfig?: pulumi.Input<inputs.compute.SubnetworkLogConfig>;
    /**
     * The name of the resource, provided by the client when initially
     * creating the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?` which
     * means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The network this subnet belongs to.
     * Only networks that are in the distributed mode can have subnetworks.
     */
    network: pulumi.Input<string>;
    /**
     * When enabled, VMs in this subnetwork without external IP addresses can
     * access Google APIs and services by using Private Google Access.
     */
    privateIpGoogleAccess?: pulumi.Input<boolean>;
    /**
     * The private IPv6 google access type for the VMs in this subnet.
     */
    privateIpv6GoogleAccess?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The purpose of the resource. This field can be either PRIVATE
     * or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
     * INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
     * reserved for Internal HTTP(S) Load Balancing. If unspecified, the
     * purpose defaults to PRIVATE.
     * If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set `role`.
     */
    purpose?: pulumi.Input<string>;
    /**
     * The GCP region for this subnetwork.
     */
    region?: pulumi.Input<string>;
    /**
     * The role of subnetwork. Currently, this field is only used when
     * purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
     * or BACKUP. An ACTIVE subnetwork is one that is currently being used
     * for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
     * is ready to be promoted to ACTIVE or is currently draining.
     * Possible values are `ACTIVE` and `BACKUP`.
     */
    role?: pulumi.Input<string>;
    /**
     * An array of configurations for secondary IP ranges for VM instances
     * contained in this subnetwork. The primary IP of such VM must belong
     * to the primary ipCidrRange of the subnetwork. The alias IPs may belong
     * to either primary or secondary ranges.
     * Structure is documented below.
     */
    secondaryIpRanges?: pulumi.Input<pulumi.Input<inputs.compute.SubnetworkSecondaryIpRange>[]>;
}
