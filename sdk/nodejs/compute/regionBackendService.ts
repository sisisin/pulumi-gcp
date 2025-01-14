// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A Region Backend Service defines a regionally-scoped group of virtual
 * machines that will serve traffic for load balancing.
 *
 * To get more information about RegionBackendService, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/latest/regionBackendServices)
 * * How-to Guides
 *     * [Internal TCP/UDP Load Balancing](https://cloud.google.com/compute/docs/load-balancing/internal/)
 *
 * > **Warning:** All arguments including `iap.oauth2_client_secret` and `iap.oauth2_client_secret_sha256` will be stored in the raw
 * state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * ## Example Usage
 * ### Region Backend Service Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultHealthCheck = new gcp.compute.HealthCheck("defaultHealthCheck", {
 *     checkIntervalSec: 1,
 *     timeoutSec: 1,
 *     tcpHealthCheck: {
 *         port: "80",
 *     },
 * });
 * const defaultRegionBackendService = new gcp.compute.RegionBackendService("defaultRegionBackendService", {
 *     region: "us-central1",
 *     healthChecks: [defaultHealthCheck.id],
 *     connectionDrainingTimeoutSec: 10,
 *     sessionAffinity: "CLIENT_IP",
 * });
 * ```
 * ### Region Backend Service Cache
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultRegionHealthCheck = new gcp.compute.RegionHealthCheck("defaultRegionHealthCheck", {
 *     region: "us-central1",
 *     httpHealthCheck: {
 *         port: 80,
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const defaultRegionBackendService = new gcp.compute.RegionBackendService("defaultRegionBackendService", {
 *     region: "us-central1",
 *     healthChecks: [defaultRegionHealthCheck.id],
 *     enableCdn: true,
 *     cdnPolicy: {
 *         cacheMode: "CACHE_ALL_STATIC",
 *         defaultTtl: 3600,
 *         clientTtl: 7200,
 *         maxTtl: 10800,
 *         negativeCaching: true,
 *         signedUrlCacheMaxAgeSec: 7200,
 *     },
 *     loadBalancingScheme: "EXTERNAL",
 *     protocol: "HTTP",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Region Backend Service Ilb Round Robin
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const healthCheck = new gcp.compute.HealthCheck("healthCheck", {httpHealthCheck: {
 *     port: 80,
 * }});
 * const _default = new gcp.compute.RegionBackendService("default", {
 *     region: "us-central1",
 *     healthChecks: [healthCheck.id],
 *     protocol: "HTTP",
 *     loadBalancingScheme: "INTERNAL_MANAGED",
 *     localityLbPolicy: "ROUND_ROBIN",
 * });
 * ```
 * ### Region Backend Service External
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const healthCheck = new gcp.compute.RegionHealthCheck("healthCheck", {
 *     region: "us-central1",
 *     tcpHealthCheck: {
 *         port: 80,
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const _default = new gcp.compute.RegionBackendService("default", {
 *     region: "us-central1",
 *     healthChecks: [healthCheck.id],
 *     protocol: "TCP",
 *     loadBalancingScheme: "EXTERNAL",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Region Backend Service Ilb Ring Hash
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const healthCheck = new gcp.compute.HealthCheck("healthCheck", {httpHealthCheck: {
 *     port: 80,
 * }});
 * const _default = new gcp.compute.RegionBackendService("default", {
 *     region: "us-central1",
 *     healthChecks: [healthCheck.id],
 *     loadBalancingScheme: "INTERNAL_MANAGED",
 *     localityLbPolicy: "RING_HASH",
 *     sessionAffinity: "HTTP_COOKIE",
 *     protocol: "HTTP",
 *     circuitBreakers: {
 *         maxConnections: 10,
 *     },
 *     consistentHash: {
 *         httpCookie: {
 *             ttl: {
 *                 seconds: 11,
 *                 nanos: 1111,
 *             },
 *             name: "mycookie",
 *         },
 *     },
 *     outlierDetection: {
 *         consecutiveErrors: 2,
 *     },
 * });
 * ```
 * ### Region Backend Service Balancing Mode
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const debianImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {
 *     autoCreateSubnetworks: false,
 *     routingMode: "REGIONAL",
 * });
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.1.2.0/24",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * const instanceTemplate = new gcp.compute.InstanceTemplate("instanceTemplate", {
 *     machineType: "e2-medium",
 *     networkInterfaces: [{
 *         network: defaultNetwork.id,
 *         subnetwork: defaultSubnetwork.id,
 *     }],
 *     disks: [{
 *         sourceImage: debianImage.then(debianImage => debianImage.selfLink),
 *         autoDelete: true,
 *         boot: true,
 *     }],
 *     tags: [
 *         "allow-ssh",
 *         "load-balanced-backend",
 *     ],
 * });
 * const rigm = new gcp.compute.RegionInstanceGroupManager("rigm", {
 *     region: "us-central1",
 *     versions: [{
 *         instanceTemplate: instanceTemplate.id,
 *         name: "primary",
 *     }],
 *     baseInstanceName: "internal-glb",
 *     targetSize: 1,
 * });
 * const defaultRegionHealthCheck = new gcp.compute.RegionHealthCheck("defaultRegionHealthCheck", {
 *     region: "us-central1",
 *     httpHealthCheck: {
 *         portSpecification: "USE_SERVING_PORT",
 *     },
 * });
 * const defaultRegionBackendService = new gcp.compute.RegionBackendService("defaultRegionBackendService", {
 *     loadBalancingScheme: "INTERNAL_MANAGED",
 *     backends: [{
 *         group: rigm.instanceGroup,
 *         balancingMode: "UTILIZATION",
 *         capacityScaler: 1,
 *     }],
 *     region: "us-central1",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 *     healthChecks: [defaultRegionHealthCheck.id],
 * });
 * ```
 *
 * ## Import
 *
 * RegionBackendService can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionBackendService:RegionBackendService default projects/{{project}}/regions/{{region}}/backendServices/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionBackendService:RegionBackendService default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionBackendService:RegionBackendService default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionBackendService:RegionBackendService default {{name}}
 * ```
 */
export class RegionBackendService extends pulumi.CustomResource {
    /**
     * Get an existing RegionBackendService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionBackendServiceState, opts?: pulumi.CustomResourceOptions): RegionBackendService {
        return new RegionBackendService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionBackendService:RegionBackendService';

    /**
     * Returns true if the given object is an instance of RegionBackendService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionBackendService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionBackendService.__pulumiType;
    }

    /**
     * Lifetime of cookies in seconds if sessionAffinity is
     * GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
     * only until the end of the browser session (or equivalent). The
     * maximum allowed value for TTL is one day.
     * When the load balancing scheme is INTERNAL, this field is not used.
     */
    public readonly affinityCookieTtlSec!: pulumi.Output<number | undefined>;
    /**
     * The set of backends that serve this RegionBackendService.
     * Structure is documented below.
     */
    public readonly backends!: pulumi.Output<outputs.compute.RegionBackendServiceBackend[] | undefined>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    public readonly cdnPolicy!: pulumi.Output<outputs.compute.RegionBackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
     * and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    public readonly circuitBreakers!: pulumi.Output<outputs.compute.RegionBackendServiceCircuitBreakers | undefined>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    public readonly connectionDrainingTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing.
     * This field only applies when all of the following are true -
     */
    public readonly consistentHash!: pulumi.Output<outputs.compute.RegionBackendServiceConsistentHash | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     * Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If true, enable Cloud CDN for this RegionBackendService.
     */
    public readonly enableCdn!: pulumi.Output<boolean | undefined>;
    /**
     * Policy for failovers.
     * Structure is documented below.
     */
    public readonly failoverPolicy!: pulumi.Output<outputs.compute.RegionBackendServiceFailoverPolicy | undefined>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The set of URLs to HealthCheck resources for health checking
     * this RegionBackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     */
    public readonly healthChecks!: pulumi.Output<string | undefined>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    public readonly iap!: pulumi.Output<outputs.compute.RegionBackendServiceIap | undefined>;
    /**
     * Indicates what kind of load balancing this regional backend service
     * will be used for. A backend service created for one type of load
     * balancing cannot be used with the other(s).
     * Default value is `INTERNAL`.
     * Possible values are `EXTERNAL`, `INTERNAL`, and `INTERNAL_MANAGED`.
     */
    public readonly loadBalancingScheme!: pulumi.Output<string | undefined>;
    /**
     * The load balancing algorithm used within the scope of the locality.
     * The possible values are -
     * * ROUND_ROBIN - This is a simple policy in which each healthy backend
     * is selected in round robin order.
     * * LEAST_REQUEST - An O(1) algorithm which selects two random healthy
     * hosts and picks the host which has fewer active requests.
     * * RING_HASH - The ring/modulo hash load balancer implements consistent
     * hashing to backends. The algorithm has the property that the
     * addition/removal of a host from a set of N hosts only affects
     * 1/N of the requests.
     * * RANDOM - The load balancer selects a random healthy host.
     * * ORIGINAL_DESTINATION - Backend host is selected based on the client
     * connection metadata, i.e., connections are opened
     * to the same address as the destination address of
     * the incoming connection before the connection
     * was redirected to the load balancer.
     * * MAGLEV - used as a drop in replacement for the ring hash load balancer.
     * Maglev is not as stable as ring hash but has faster table lookup
     * build times and host selection times. For more information about
     * Maglev, refer to https://ai.google/research/pubs/pub44824
     * This field is applicable only when the `loadBalancingScheme` is set to
     * INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Possible values are `ROUND_ROBIN`, `LEAST_REQUEST`, `RING_HASH`, `RANDOM`, `ORIGINAL_DESTINATION`, and `MAGLEV`.
     */
    public readonly localityLbPolicy!: pulumi.Output<string | undefined>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service.
     * If logging is enabled, logs will be exported to Stackdriver.
     * Structure is documented below.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.RegionBackendServiceLogConfig>;
    /**
     * Name of the cookie.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of the network to which this backend service belongs.
     * This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    public readonly network!: pulumi.Output<string | undefined>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * This field is applicable only when the `loadBalancingScheme` is set
     * to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    public readonly outlierDetection!: pulumi.Output<outputs.compute.RegionBackendServiceOutlierDetection | undefined>;
    /**
     * A named port on a backend instance group representing the port for
     * communication to the backend VMs in that group. Required when the
     * loadBalancingScheme is EXTERNAL, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED
     * and the backends are instance groups. The named port must be defined on each
     * backend instance group. This parameter has no meaning if the backends are NEGs. API sets a
     * default of "http" if not given.
     * Must be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).
     */
    public readonly portName!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The protocol this RegionBackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API.
     * Possible values are `HTTP`, `HTTPS`, `HTTP2`, `SSL`, `TCP`, `UDP`, `GRPC`, and `UNSPECIFIED`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The Region in which the created backend service should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, and `HTTP_COOKIE`.
     */
    public readonly sessionAffinity!: pulumi.Output<string>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    public readonly timeoutSec!: pulumi.Output<number>;

    /**
     * Create a RegionBackendService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegionBackendServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionBackendServiceArgs | RegionBackendServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegionBackendServiceState | undefined;
            inputs["affinityCookieTtlSec"] = state ? state.affinityCookieTtlSec : undefined;
            inputs["backends"] = state ? state.backends : undefined;
            inputs["cdnPolicy"] = state ? state.cdnPolicy : undefined;
            inputs["circuitBreakers"] = state ? state.circuitBreakers : undefined;
            inputs["connectionDrainingTimeoutSec"] = state ? state.connectionDrainingTimeoutSec : undefined;
            inputs["consistentHash"] = state ? state.consistentHash : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableCdn"] = state ? state.enableCdn : undefined;
            inputs["failoverPolicy"] = state ? state.failoverPolicy : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["healthChecks"] = state ? state.healthChecks : undefined;
            inputs["iap"] = state ? state.iap : undefined;
            inputs["loadBalancingScheme"] = state ? state.loadBalancingScheme : undefined;
            inputs["localityLbPolicy"] = state ? state.localityLbPolicy : undefined;
            inputs["logConfig"] = state ? state.logConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["outlierDetection"] = state ? state.outlierDetection : undefined;
            inputs["portName"] = state ? state.portName : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sessionAffinity"] = state ? state.sessionAffinity : undefined;
            inputs["timeoutSec"] = state ? state.timeoutSec : undefined;
        } else {
            const args = argsOrState as RegionBackendServiceArgs | undefined;
            inputs["affinityCookieTtlSec"] = args ? args.affinityCookieTtlSec : undefined;
            inputs["backends"] = args ? args.backends : undefined;
            inputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            inputs["circuitBreakers"] = args ? args.circuitBreakers : undefined;
            inputs["connectionDrainingTimeoutSec"] = args ? args.connectionDrainingTimeoutSec : undefined;
            inputs["consistentHash"] = args ? args.consistentHash : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableCdn"] = args ? args.enableCdn : undefined;
            inputs["failoverPolicy"] = args ? args.failoverPolicy : undefined;
            inputs["healthChecks"] = args ? args.healthChecks : undefined;
            inputs["iap"] = args ? args.iap : undefined;
            inputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            inputs["localityLbPolicy"] = args ? args.localityLbPolicy : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["outlierDetection"] = args ? args.outlierDetection : undefined;
            inputs["portName"] = args ? args.portName : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegionBackendService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionBackendService resources.
 */
export interface RegionBackendServiceState {
    /**
     * Lifetime of cookies in seconds if sessionAffinity is
     * GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
     * only until the end of the browser session (or equivalent). The
     * maximum allowed value for TTL is one day.
     * When the load balancing scheme is INTERNAL, this field is not used.
     */
    affinityCookieTtlSec?: pulumi.Input<number>;
    /**
     * The set of backends that serve this RegionBackendService.
     * Structure is documented below.
     */
    backends?: pulumi.Input<pulumi.Input<inputs.compute.RegionBackendServiceBackend>[]>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    cdnPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
     * and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    circuitBreakers?: pulumi.Input<inputs.compute.RegionBackendServiceCircuitBreakers>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    connectionDrainingTimeoutSec?: pulumi.Input<number>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing.
     * This field only applies when all of the following are true -
     */
    consistentHash?: pulumi.Input<inputs.compute.RegionBackendServiceConsistentHash>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     * Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this RegionBackendService.
     */
    enableCdn?: pulumi.Input<boolean>;
    /**
     * Policy for failovers.
     * Structure is documented below.
     */
    failoverPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceFailoverPolicy>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The set of URLs to HealthCheck resources for health checking
     * this RegionBackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     */
    healthChecks?: pulumi.Input<string>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    iap?: pulumi.Input<inputs.compute.RegionBackendServiceIap>;
    /**
     * Indicates what kind of load balancing this regional backend service
     * will be used for. A backend service created for one type of load
     * balancing cannot be used with the other(s).
     * Default value is `INTERNAL`.
     * Possible values are `EXTERNAL`, `INTERNAL`, and `INTERNAL_MANAGED`.
     */
    loadBalancingScheme?: pulumi.Input<string>;
    /**
     * The load balancing algorithm used within the scope of the locality.
     * The possible values are -
     * * ROUND_ROBIN - This is a simple policy in which each healthy backend
     * is selected in round robin order.
     * * LEAST_REQUEST - An O(1) algorithm which selects two random healthy
     * hosts and picks the host which has fewer active requests.
     * * RING_HASH - The ring/modulo hash load balancer implements consistent
     * hashing to backends. The algorithm has the property that the
     * addition/removal of a host from a set of N hosts only affects
     * 1/N of the requests.
     * * RANDOM - The load balancer selects a random healthy host.
     * * ORIGINAL_DESTINATION - Backend host is selected based on the client
     * connection metadata, i.e., connections are opened
     * to the same address as the destination address of
     * the incoming connection before the connection
     * was redirected to the load balancer.
     * * MAGLEV - used as a drop in replacement for the ring hash load balancer.
     * Maglev is not as stable as ring hash but has faster table lookup
     * build times and host selection times. For more information about
     * Maglev, refer to https://ai.google/research/pubs/pub44824
     * This field is applicable only when the `loadBalancingScheme` is set to
     * INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Possible values are `ROUND_ROBIN`, `LEAST_REQUEST`, `RING_HASH`, `RANDOM`, `ORIGINAL_DESTINATION`, and `MAGLEV`.
     */
    localityLbPolicy?: pulumi.Input<string>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service.
     * If logging is enabled, logs will be exported to Stackdriver.
     * Structure is documented below.
     */
    logConfig?: pulumi.Input<inputs.compute.RegionBackendServiceLogConfig>;
    /**
     * Name of the cookie.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of the network to which this backend service belongs.
     * This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    network?: pulumi.Input<string>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * This field is applicable only when the `loadBalancingScheme` is set
     * to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    outlierDetection?: pulumi.Input<inputs.compute.RegionBackendServiceOutlierDetection>;
    /**
     * A named port on a backend instance group representing the port for
     * communication to the backend VMs in that group. Required when the
     * loadBalancingScheme is EXTERNAL, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED
     * and the backends are instance groups. The named port must be defined on each
     * backend instance group. This parameter has no meaning if the backends are NEGs. API sets a
     * default of "http" if not given.
     * Must be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).
     */
    portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The protocol this RegionBackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API.
     * Possible values are `HTTP`, `HTTPS`, `HTTP2`, `SSL`, `TCP`, `UDP`, `GRPC`, and `UNSPECIFIED`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The Region in which the created backend service should reside.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, and `HTTP_COOKIE`.
     */
    sessionAffinity?: pulumi.Input<string>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    timeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RegionBackendService resource.
 */
export interface RegionBackendServiceArgs {
    /**
     * Lifetime of cookies in seconds if sessionAffinity is
     * GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
     * only until the end of the browser session (or equivalent). The
     * maximum allowed value for TTL is one day.
     * When the load balancing scheme is INTERNAL, this field is not used.
     */
    affinityCookieTtlSec?: pulumi.Input<number>;
    /**
     * The set of backends that serve this RegionBackendService.
     * Structure is documented below.
     */
    backends?: pulumi.Input<pulumi.Input<inputs.compute.RegionBackendServiceBackend>[]>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    cdnPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
     * and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    circuitBreakers?: pulumi.Input<inputs.compute.RegionBackendServiceCircuitBreakers>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    connectionDrainingTimeoutSec?: pulumi.Input<number>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing.
     * This field only applies when all of the following are true -
     */
    consistentHash?: pulumi.Input<inputs.compute.RegionBackendServiceConsistentHash>;
    /**
     * An optional description of this resource.
     * Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this RegionBackendService.
     */
    enableCdn?: pulumi.Input<boolean>;
    /**
     * Policy for failovers.
     * Structure is documented below.
     */
    failoverPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceFailoverPolicy>;
    /**
     * The set of URLs to HealthCheck resources for health checking
     * this RegionBackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     */
    healthChecks?: pulumi.Input<string>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    iap?: pulumi.Input<inputs.compute.RegionBackendServiceIap>;
    /**
     * Indicates what kind of load balancing this regional backend service
     * will be used for. A backend service created for one type of load
     * balancing cannot be used with the other(s).
     * Default value is `INTERNAL`.
     * Possible values are `EXTERNAL`, `INTERNAL`, and `INTERNAL_MANAGED`.
     */
    loadBalancingScheme?: pulumi.Input<string>;
    /**
     * The load balancing algorithm used within the scope of the locality.
     * The possible values are -
     * * ROUND_ROBIN - This is a simple policy in which each healthy backend
     * is selected in round robin order.
     * * LEAST_REQUEST - An O(1) algorithm which selects two random healthy
     * hosts and picks the host which has fewer active requests.
     * * RING_HASH - The ring/modulo hash load balancer implements consistent
     * hashing to backends. The algorithm has the property that the
     * addition/removal of a host from a set of N hosts only affects
     * 1/N of the requests.
     * * RANDOM - The load balancer selects a random healthy host.
     * * ORIGINAL_DESTINATION - Backend host is selected based on the client
     * connection metadata, i.e., connections are opened
     * to the same address as the destination address of
     * the incoming connection before the connection
     * was redirected to the load balancer.
     * * MAGLEV - used as a drop in replacement for the ring hash load balancer.
     * Maglev is not as stable as ring hash but has faster table lookup
     * build times and host selection times. For more information about
     * Maglev, refer to https://ai.google/research/pubs/pub44824
     * This field is applicable only when the `loadBalancingScheme` is set to
     * INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Possible values are `ROUND_ROBIN`, `LEAST_REQUEST`, `RING_HASH`, `RANDOM`, `ORIGINAL_DESTINATION`, and `MAGLEV`.
     */
    localityLbPolicy?: pulumi.Input<string>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service.
     * If logging is enabled, logs will be exported to Stackdriver.
     * Structure is documented below.
     */
    logConfig?: pulumi.Input<inputs.compute.RegionBackendServiceLogConfig>;
    /**
     * Name of the cookie.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of the network to which this backend service belongs.
     * This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    network?: pulumi.Input<string>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * This field is applicable only when the `loadBalancingScheme` is set
     * to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    outlierDetection?: pulumi.Input<inputs.compute.RegionBackendServiceOutlierDetection>;
    /**
     * A named port on a backend instance group representing the port for
     * communication to the backend VMs in that group. Required when the
     * loadBalancingScheme is EXTERNAL, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED
     * and the backends are instance groups. The named port must be defined on each
     * backend instance group. This parameter has no meaning if the backends are NEGs. API sets a
     * default of "http" if not given.
     * Must be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).
     */
    portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The protocol this RegionBackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API.
     * Possible values are `HTTP`, `HTTPS`, `HTTP2`, `SSL`, `TCP`, `UDP`, `GRPC`, and `UNSPECIFIED`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The Region in which the created backend service should reside.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, and `HTTP_COOKIE`.
     */
    sessionAffinity?: pulumi.Input<string>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    timeoutSec?: pulumi.Input<number>;
}
