// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A managed metastore service that serves metadata queries.
 *
 * ## Example Usage
 * ### Dataproc Metastore Service Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.dataproc.MetastoreService("default", {
 *     serviceId: "metastore-srv",
 *     location: "us-central1",
 *     port: 9080,
 *     tier: "DEVELOPER",
 *     maintenanceWindow: {
 *         hourOfDay: 2,
 *         dayOfWeek: "SUNDAY",
 *     },
 *     hiveMetastoreConfig: {
 *         version: "2.3.6",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * Service can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreService:MetastoreService default projects/{{project}}/locations/{{location}}/services/{{service_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreService:MetastoreService default {{project}}/{{location}}/{{service_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/metastoreService:MetastoreService default {{location}}/{{service_id}}
 * ```
 */
export class MetastoreService extends pulumi.CustomResource {
    /**
     * Get an existing MetastoreService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MetastoreServiceState, opts?: pulumi.CustomResourceOptions): MetastoreService {
        return new MetastoreService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/metastoreService:MetastoreService';

    /**
     * Returns true if the given object is an instance of MetastoreService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MetastoreService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MetastoreService.__pulumiType;
    }

    /**
     * A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
     */
    public /*out*/ readonly artifactGcsUri!: pulumi.Output<string>;
    /**
     * The URI of the endpoint used to access the metastore service.
     */
    public /*out*/ readonly endpointUri!: pulumi.Output<string>;
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     * Structure is documented below.
     */
    public readonly hiveMetastoreConfig!: pulumi.Output<outputs.dataproc.MetastoreServiceHiveMetastoreConfig | undefined>;
    /**
     * User-defined labels for the metastore service.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The  location where the autoscaling policy should reside.
     * The default value is `global`.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The one hour maintenance window of the metastore service.
     * This specifies when the service can be restarted for maintenance purposes in UTC time.
     * Structure is documented below.
     */
    public readonly maintenanceWindow!: pulumi.Output<outputs.dataproc.MetastoreServiceMaintenanceWindow | undefined>;
    /**
     * The relative resource name of the metastore service.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * "projects/{projectNumber}/global/networks/{network_id}".
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 63 characters.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * The current state of the metastore service.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current state of the metastore service, if available.
     */
    public /*out*/ readonly stateMessage!: pulumi.Output<string>;
    /**
     * The tier of the service.
     * Possible values are `DEVELOPER` and `ENTERPRISE`.
     */
    public readonly tier!: pulumi.Output<string>;

    /**
     * Create a MetastoreService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetastoreServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MetastoreServiceArgs | MetastoreServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MetastoreServiceState | undefined;
            inputs["artifactGcsUri"] = state ? state.artifactGcsUri : undefined;
            inputs["endpointUri"] = state ? state.endpointUri : undefined;
            inputs["hiveMetastoreConfig"] = state ? state.hiveMetastoreConfig : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maintenanceWindow"] = state ? state.maintenanceWindow : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["serviceId"] = state ? state.serviceId : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["stateMessage"] = state ? state.stateMessage : undefined;
            inputs["tier"] = state ? state.tier : undefined;
        } else {
            const args = argsOrState as MetastoreServiceArgs | undefined;
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            inputs["hiveMetastoreConfig"] = args ? args.hiveMetastoreConfig : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["serviceId"] = args ? args.serviceId : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["artifactGcsUri"] = undefined /*out*/;
            inputs["endpointUri"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["stateMessage"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MetastoreService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MetastoreService resources.
 */
export interface MetastoreServiceState {
    /**
     * A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
     */
    artifactGcsUri?: pulumi.Input<string>;
    /**
     * The URI of the endpoint used to access the metastore service.
     */
    endpointUri?: pulumi.Input<string>;
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     * Structure is documented below.
     */
    hiveMetastoreConfig?: pulumi.Input<inputs.dataproc.MetastoreServiceHiveMetastoreConfig>;
    /**
     * User-defined labels for the metastore service.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The  location where the autoscaling policy should reside.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * The one hour maintenance window of the metastore service.
     * This specifies when the service can be restarted for maintenance purposes in UTC time.
     * Structure is documented below.
     */
    maintenanceWindow?: pulumi.Input<inputs.dataproc.MetastoreServiceMaintenanceWindow>;
    /**
     * The relative resource name of the metastore service.
     */
    name?: pulumi.Input<string>;
    /**
     * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * "projects/{projectNumber}/global/networks/{network_id}".
     */
    network?: pulumi.Input<string>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 63 characters.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The current state of the metastore service.
     */
    state?: pulumi.Input<string>;
    /**
     * Additional information about the current state of the metastore service, if available.
     */
    stateMessage?: pulumi.Input<string>;
    /**
     * The tier of the service.
     * Possible values are `DEVELOPER` and `ENTERPRISE`.
     */
    tier?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MetastoreService resource.
 */
export interface MetastoreServiceArgs {
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     * Structure is documented below.
     */
    hiveMetastoreConfig?: pulumi.Input<inputs.dataproc.MetastoreServiceHiveMetastoreConfig>;
    /**
     * User-defined labels for the metastore service.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The  location where the autoscaling policy should reside.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * The one hour maintenance window of the metastore service.
     * This specifies when the service can be restarted for maintenance purposes in UTC time.
     * Structure is documented below.
     */
    maintenanceWindow?: pulumi.Input<inputs.dataproc.MetastoreServiceMaintenanceWindow>;
    /**
     * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * "projects/{projectNumber}/global/networks/{network_id}".
     */
    network?: pulumi.Input<string>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 63 characters.
     */
    serviceId: pulumi.Input<string>;
    /**
     * The tier of the service.
     * Possible values are `DEVELOPER` and `ENTERPRISE`.
     */
    tier?: pulumi.Input<string>;
}
