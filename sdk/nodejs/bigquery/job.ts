// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Jobs are actions that BigQuery runs on your behalf to load data, export data, query data, or copy data.
 * Once a BigQuery job is created, it cannot be changed or deleted.
 *
 * To get more information about Job, see:
 *
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/jobs)
 * * How-to Guides
 *     * [BigQuery Jobs Intro](https://cloud.google.com/bigquery/docs/jobs-overview)
 *
 * ## Example Usage
 * ### Bigquery Job Query
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bar = new gcp.bigquery.Dataset("bar", {
 *     datasetId: "job_query_dataset",
 *     friendlyName: "test",
 *     description: "This is a test description",
 *     location: "US",
 * });
 * const foo = new gcp.bigquery.Table("foo", {
 *     deletionProtection: false,
 *     datasetId: bar.datasetId,
 *     tableId: "job_query_table",
 * });
 * const job = new gcp.bigquery.Job("job", {
 *     jobId: "job_query",
 *     labels: {
 *         "example-label": "example-value",
 *     },
 *     query: {
 *         query: "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
 *         destinationTable: {
 *             projectId: foo.project,
 *             datasetId: foo.datasetId,
 *             tableId: foo.tableId,
 *         },
 *         allowLargeResults: true,
 *         flattenResults: true,
 *         scriptOptions: {
 *             keyResultStatement: "LAST",
 *         },
 *     },
 * });
 * ```
 * ### Bigquery Job Query Table Reference
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bar = new gcp.bigquery.Dataset("bar", {
 *     datasetId: "job_query_dataset",
 *     friendlyName: "test",
 *     description: "This is a test description",
 *     location: "US",
 * });
 * const foo = new gcp.bigquery.Table("foo", {
 *     deletionProtection: false,
 *     datasetId: bar.datasetId,
 *     tableId: "job_query_table",
 * });
 * const job = new gcp.bigquery.Job("job", {
 *     jobId: "job_query",
 *     labels: {
 *         "example-label": "example-value",
 *     },
 *     query: {
 *         query: "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
 *         destinationTable: {
 *             tableId: foo.id,
 *         },
 *         defaultDataset: {
 *             datasetId: bar.id,
 *         },
 *         allowLargeResults: true,
 *         flattenResults: true,
 *         scriptOptions: {
 *             keyResultStatement: "LAST",
 *         },
 *     },
 * });
 * ```
 * ### Bigquery Job Load
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bar = new gcp.bigquery.Dataset("bar", {
 *     datasetId: "job_load_dataset",
 *     friendlyName: "test",
 *     description: "This is a test description",
 *     location: "US",
 * });
 * const foo = new gcp.bigquery.Table("foo", {
 *     deletionProtection: false,
 *     datasetId: bar.datasetId,
 *     tableId: "job_load_table",
 * });
 * const job = new gcp.bigquery.Job("job", {
 *     jobId: "job_load",
 *     labels: {
 *         my_job: "load",
 *     },
 *     load: {
 *         sourceUris: ["gs://cloud-samples-data/bigquery/us-states/us-states-by-date.csv"],
 *         destinationTable: {
 *             projectId: foo.project,
 *             datasetId: foo.datasetId,
 *             tableId: foo.tableId,
 *         },
 *         skipLeadingRows: 1,
 *         schemaUpdateOptions: [
 *             "ALLOW_FIELD_RELAXATION",
 *             "ALLOW_FIELD_ADDITION",
 *         ],
 *         writeDisposition: "WRITE_APPEND",
 *         autodetect: true,
 *     },
 * });
 * ```
 * ### Bigquery Job Extract
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const source_oneDataset = new gcp.bigquery.Dataset("source-oneDataset", {
 *     datasetId: "job_extract_dataset",
 *     friendlyName: "test",
 *     description: "This is a test description",
 *     location: "US",
 * });
 * const source_oneTable = new gcp.bigquery.Table("source-oneTable", {
 *     deletionProtection: false,
 *     datasetId: source_oneDataset.datasetId,
 *     tableId: "job_extract_table",
 *     schema: `[
 *   {
 *     "name": "name",
 *     "type": "STRING",
 *     "mode": "NULLABLE"
 *   },
 *   {
 *     "name": "post_abbr",
 *     "type": "STRING",
 *     "mode": "NULLABLE"
 *   },
 *   {
 *     "name": "date",
 *     "type": "DATE",
 *     "mode": "NULLABLE"
 *   }
 * ]
 * `,
 * });
 * const dest = new gcp.storage.Bucket("dest", {forceDestroy: true});
 * const job = new gcp.bigquery.Job("job", {
 *     jobId: "job_extract",
 *     extract: {
 *         destinationUris: [pulumi.interpolate`${dest.url}/extract`],
 *         sourceTable: {
 *             projectId: source_oneTable.project,
 *             datasetId: source_oneTable.datasetId,
 *             tableId: source_oneTable.tableId,
 *         },
 *         destinationFormat: "NEWLINE_DELIMITED_JSON",
 *         compression: "GZIP",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Job can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}/location/{{location}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}/{{location}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/job:Job default {{job_id}}/{{location}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/job:Job default {{job_id}}
 * ```
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
    public static readonly __pulumiType = 'gcp:bigquery/job:Job';

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
     * Copies a table.
     * Structure is documented below.
     */
    public readonly copy!: pulumi.Output<outputs.bigquery.JobCopy | undefined>;
    /**
     * Configures an extract job.
     * Structure is documented below.
     */
    public readonly extract!: pulumi.Output<outputs.bigquery.JobExtract | undefined>;
    /**
     * The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
     */
    public readonly jobId!: pulumi.Output<string>;
    /**
     * Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
     */
    public readonly jobTimeoutMs!: pulumi.Output<string | undefined>;
    /**
     * The type of the job.
     */
    public /*out*/ readonly jobType!: pulumi.Output<string>;
    /**
     * The labels associated with this job. You can use these to organize and group your jobs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Configures a load job.
     * Structure is documented below.
     */
    public readonly load!: pulumi.Output<outputs.bigquery.JobLoad | undefined>;
    /**
     * The geographic location of the job. The default value is US.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Configures a query job.
     * Structure is documented below.
     */
    public readonly query!: pulumi.Output<outputs.bigquery.JobQuery | undefined>;
    /**
     * The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
     */
    public /*out*/ readonly statuses!: pulumi.Output<outputs.bigquery.JobStatus[]>;
    /**
     * Email address of the user who ran the job.
     */
    public /*out*/ readonly userEmail!: pulumi.Output<string>;

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
            inputs["copy"] = state ? state.copy : undefined;
            inputs["extract"] = state ? state.extract : undefined;
            inputs["jobId"] = state ? state.jobId : undefined;
            inputs["jobTimeoutMs"] = state ? state.jobTimeoutMs : undefined;
            inputs["jobType"] = state ? state.jobType : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["load"] = state ? state.load : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["query"] = state ? state.query : undefined;
            inputs["statuses"] = state ? state.statuses : undefined;
            inputs["userEmail"] = state ? state.userEmail : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if ((!args || args.jobId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobId'");
            }
            inputs["copy"] = args ? args.copy : undefined;
            inputs["extract"] = args ? args.extract : undefined;
            inputs["jobId"] = args ? args.jobId : undefined;
            inputs["jobTimeoutMs"] = args ? args.jobTimeoutMs : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["load"] = args ? args.load : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["query"] = args ? args.query : undefined;
            inputs["jobType"] = undefined /*out*/;
            inputs["statuses"] = undefined /*out*/;
            inputs["userEmail"] = undefined /*out*/;
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
     * Copies a table.
     * Structure is documented below.
     */
    copy?: pulumi.Input<inputs.bigquery.JobCopy>;
    /**
     * Configures an extract job.
     * Structure is documented below.
     */
    extract?: pulumi.Input<inputs.bigquery.JobExtract>;
    /**
     * The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
     */
    jobId?: pulumi.Input<string>;
    /**
     * Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
     */
    jobTimeoutMs?: pulumi.Input<string>;
    /**
     * The type of the job.
     */
    jobType?: pulumi.Input<string>;
    /**
     * The labels associated with this job. You can use these to organize and group your jobs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configures a load job.
     * Structure is documented below.
     */
    load?: pulumi.Input<inputs.bigquery.JobLoad>;
    /**
     * The geographic location of the job. The default value is US.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Configures a query job.
     * Structure is documented below.
     */
    query?: pulumi.Input<inputs.bigquery.JobQuery>;
    /**
     * The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.bigquery.JobStatus>[]>;
    /**
     * Email address of the user who ran the job.
     */
    userEmail?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Copies a table.
     * Structure is documented below.
     */
    copy?: pulumi.Input<inputs.bigquery.JobCopy>;
    /**
     * Configures an extract job.
     * Structure is documented below.
     */
    extract?: pulumi.Input<inputs.bigquery.JobExtract>;
    /**
     * The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
     */
    jobId: pulumi.Input<string>;
    /**
     * Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
     */
    jobTimeoutMs?: pulumi.Input<string>;
    /**
     * The labels associated with this job. You can use these to organize and group your jobs.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configures a load job.
     * Structure is documented below.
     */
    load?: pulumi.Input<inputs.bigquery.JobLoad>;
    /**
     * The geographic location of the job. The default value is US.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Configures a query job.
     * Structure is documented below.
     */
    query?: pulumi.Input<inputs.bigquery.JobQuery>;
}
