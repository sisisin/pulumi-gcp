// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a job resource within a Dataproc cluster within GCE. For more information see
 * [the official dataproc documentation](https://cloud.google.com/dataproc/).
 *
 * !> **Note:** This resource does not support 'update' and changing any attributes will cause the resource to be recreated.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const mycluster = new gcp.dataproc.Cluster("mycluster", {region: "us-central1"});
 * // Submit an example spark job to a dataproc cluster
 * const spark = new gcp.dataproc.Job("spark", {
 *     region: mycluster.region,
 *     forceDelete: true,
 *     placement: {
 *         clusterName: mycluster.name,
 *     },
 *     sparkConfig: {
 *         mainClass: "org.apache.spark.examples.SparkPi",
 *         jarFileUris: ["file:///usr/lib/spark/examples/jars/spark-examples.jar"],
 *         args: ["1000"],
 *         properties: {
 *             "spark.logConf": "true",
 *         },
 *         loggingConfig: {
 *             driverLogLevels: {
 *                 root: "INFO",
 *             },
 *         },
 *     },
 * });
 * // Submit an example pyspark job to a dataproc cluster
 * const pyspark = new gcp.dataproc.Job("pyspark", {
 *     region: mycluster.region,
 *     forceDelete: true,
 *     placement: {
 *         clusterName: mycluster.name,
 *     },
 *     pysparkConfig: {
 *         mainPythonFileUri: "gs://dataproc-examples-2f10d78d114f6aaec76462e3c310f31f/src/pyspark/hello-world/hello-world.py",
 *         properties: {
 *             "spark.logConf": "true",
 *         },
 *     },
 * });
 * export const sparkStatus = spark.statuses.apply(statuses => statuses[0].state);
 * export const pysparkStatus = pyspark.statuses.apply(statuses => statuses[0].state);
 * ```
 *
 * ## Import
 *
 * This resource does not support import.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
     */
    public /*out*/ readonly driverControlsFilesUri!: pulumi.Output<string>;
    /**
     * A URI pointing to the location of the stdout of the job's driver program.
     */
    public /*out*/ readonly driverOutputResourceUri!: pulumi.Output<string>;
    /**
     * By default, you can only delete inactive jobs within
     * Dataproc. Setting this to true, and calling destroy, will ensure that the
     * job is first cancelled before issuing the delete.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * The config of Hadoop job
     */
    public readonly hadoopConfig!: pulumi.Output<outputs.dataproc.JobHadoopConfig | undefined>;
    /**
     * The config of hive job
     */
    public readonly hiveConfig!: pulumi.Output<outputs.dataproc.JobHiveConfig | undefined>;
    /**
     * The list of labels (key/value pairs) to add to the job.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The config of pag job.
     */
    public readonly pigConfig!: pulumi.Output<outputs.dataproc.JobPigConfig | undefined>;
    /**
     * The config of job placement.
     */
    public readonly placement!: pulumi.Output<outputs.dataproc.JobPlacement>;
    /**
     * The project in which the `cluster` can be found and jobs
     * subsequently run against. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The config of pySpark job.
     */
    public readonly pysparkConfig!: pulumi.Output<outputs.dataproc.JobPysparkConfig | undefined>;
    /**
     * The reference of the job
     */
    public readonly reference!: pulumi.Output<outputs.dataproc.JobReference>;
    /**
     * The Cloud Dataproc region. This essentially determines which clusters are available
     * for this job to be submitted to. If not specified, defaults to `global`.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * Optional. Job scheduling configuration.
     */
    public readonly scheduling!: pulumi.Output<outputs.dataproc.JobScheduling | undefined>;
    /**
     * The config of the Spark job.
     */
    public readonly sparkConfig!: pulumi.Output<outputs.dataproc.JobSparkConfig | undefined>;
    /**
     * The config of SparkSql job
     */
    public readonly sparksqlConfig!: pulumi.Output<outputs.dataproc.JobSparksqlConfig | undefined>;
    /**
     * The status of the job.
     */
    public /*out*/ readonly statuses!: pulumi.Output<outputs.dataproc.JobStatus[]>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobState | undefined;
            inputs["driverControlsFilesUri"] = state ? state.driverControlsFilesUri : undefined;
            inputs["driverOutputResourceUri"] = state ? state.driverOutputResourceUri : undefined;
            inputs["forceDelete"] = state ? state.forceDelete : undefined;
            inputs["hadoopConfig"] = state ? state.hadoopConfig : undefined;
            inputs["hiveConfig"] = state ? state.hiveConfig : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["pigConfig"] = state ? state.pigConfig : undefined;
            inputs["placement"] = state ? state.placement : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pysparkConfig"] = state ? state.pysparkConfig : undefined;
            inputs["reference"] = state ? state.reference : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["scheduling"] = state ? state.scheduling : undefined;
            inputs["sparkConfig"] = state ? state.sparkConfig : undefined;
            inputs["sparksqlConfig"] = state ? state.sparksqlConfig : undefined;
            inputs["statuses"] = state ? state.statuses : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if ((!args || args.placement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'placement'");
            }
            inputs["forceDelete"] = args ? args.forceDelete : undefined;
            inputs["hadoopConfig"] = args ? args.hadoopConfig : undefined;
            inputs["hiveConfig"] = args ? args.hiveConfig : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["pigConfig"] = args ? args.pigConfig : undefined;
            inputs["placement"] = args ? args.placement : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pysparkConfig"] = args ? args.pysparkConfig : undefined;
            inputs["reference"] = args ? args.reference : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["scheduling"] = args ? args.scheduling : undefined;
            inputs["sparkConfig"] = args ? args.sparkConfig : undefined;
            inputs["sparksqlConfig"] = args ? args.sparksqlConfig : undefined;
            inputs["driverControlsFilesUri"] = undefined /*out*/;
            inputs["driverOutputResourceUri"] = undefined /*out*/;
            inputs["statuses"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
     */
    driverControlsFilesUri?: pulumi.Input<string>;
    /**
     * A URI pointing to the location of the stdout of the job's driver program.
     */
    driverOutputResourceUri?: pulumi.Input<string>;
    /**
     * By default, you can only delete inactive jobs within
     * Dataproc. Setting this to true, and calling destroy, will ensure that the
     * job is first cancelled before issuing the delete.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The config of Hadoop job
     */
    hadoopConfig?: pulumi.Input<inputs.dataproc.JobHadoopConfig>;
    /**
     * The config of hive job
     */
    hiveConfig?: pulumi.Input<inputs.dataproc.JobHiveConfig>;
    /**
     * The list of labels (key/value pairs) to add to the job.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The config of pag job.
     */
    pigConfig?: pulumi.Input<inputs.dataproc.JobPigConfig>;
    /**
     * The config of job placement.
     */
    placement?: pulumi.Input<inputs.dataproc.JobPlacement>;
    /**
     * The project in which the `cluster` can be found and jobs
     * subsequently run against. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The config of pySpark job.
     */
    pysparkConfig?: pulumi.Input<inputs.dataproc.JobPysparkConfig>;
    /**
     * The reference of the job
     */
    reference?: pulumi.Input<inputs.dataproc.JobReference>;
    /**
     * The Cloud Dataproc region. This essentially determines which clusters are available
     * for this job to be submitted to. If not specified, defaults to `global`.
     */
    region?: pulumi.Input<string>;
    /**
     * Optional. Job scheduling configuration.
     */
    scheduling?: pulumi.Input<inputs.dataproc.JobScheduling>;
    /**
     * The config of the Spark job.
     */
    sparkConfig?: pulumi.Input<inputs.dataproc.JobSparkConfig>;
    /**
     * The config of SparkSql job
     */
    sparksqlConfig?: pulumi.Input<inputs.dataproc.JobSparksqlConfig>;
    /**
     * The status of the job.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.dataproc.JobStatus>[]>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * By default, you can only delete inactive jobs within
     * Dataproc. Setting this to true, and calling destroy, will ensure that the
     * job is first cancelled before issuing the delete.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The config of Hadoop job
     */
    hadoopConfig?: pulumi.Input<inputs.dataproc.JobHadoopConfig>;
    /**
     * The config of hive job
     */
    hiveConfig?: pulumi.Input<inputs.dataproc.JobHiveConfig>;
    /**
     * The list of labels (key/value pairs) to add to the job.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The config of pag job.
     */
    pigConfig?: pulumi.Input<inputs.dataproc.JobPigConfig>;
    /**
     * The config of job placement.
     */
    placement: pulumi.Input<inputs.dataproc.JobPlacement>;
    /**
     * The project in which the `cluster` can be found and jobs
     * subsequently run against. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The config of pySpark job.
     */
    pysparkConfig?: pulumi.Input<inputs.dataproc.JobPysparkConfig>;
    /**
     * The reference of the job
     */
    reference?: pulumi.Input<inputs.dataproc.JobReference>;
    /**
     * The Cloud Dataproc region. This essentially determines which clusters are available
     * for this job to be submitted to. If not specified, defaults to `global`.
     */
    region?: pulumi.Input<string>;
    /**
     * Optional. Job scheduling configuration.
     */
    scheduling?: pulumi.Input<inputs.dataproc.JobScheduling>;
    /**
     * The config of the Spark job.
     */
    sparkConfig?: pulumi.Input<inputs.dataproc.JobSparkConfig>;
    /**
     * The config of SparkSql job
     */
    sparksqlConfig?: pulumi.Input<inputs.dataproc.JobSparksqlConfig>;
}
