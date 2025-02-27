// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
 *
 * * `gcp.dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
 * * `gcp.dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
 * * `gcp.dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
 *
 * > **Note:** `gcp.dataproc.JobIAMPolicy` **cannot** be used in conjunction with `gcp.dataproc.JobIAMBinding` and `gcp.dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `gcp.dataproc.JobIAMPolicy` replaces the entire policy.
 *
 * > **Note:** `gcp.dataproc.JobIAMBinding` resources **can be** used in conjunction with `gcp.dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_dataproc\_job\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const editor = new gcp.dataproc.JobIAMPolicy("editor", {
 *     project: "your-project",
 *     region: "your-region",
 *     jobId: "your-dataproc-job",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_dataproc\_job\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.dataproc.JobIAMBinding("editor", {
 *     jobId: "your-dataproc-job",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_dataproc\_job\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.dataproc.JobIAMMember("editor", {
 *     jobId: "your-dataproc-job",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## Import
 *
 * Job IAM resources can be imported using the project, region, job id, role and/or member.
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/jobIAMBinding:JobIAMBinding editor "projects/{project}/regions/{region}/jobs/{job_id}"
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/jobIAMBinding:JobIAMBinding editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor"
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/jobIAMBinding:JobIAMBinding editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com"
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class JobIAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing JobIAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobIAMBindingState, opts?: pulumi.CustomResourceOptions): JobIAMBinding {
        return new JobIAMBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/jobIAMBinding:JobIAMBinding';

    /**
     * Returns true if the given object is an instance of JobIAMBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobIAMBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobIAMBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.dataproc.JobIAMBindingCondition | undefined>;
    /**
     * (Computed) The etag of the jobs's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly jobId!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The project in which the job belongs. If it
     * is not provided, the provider will use a default.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region in which the job belongs. If it
     * is not provided, the provider will use a default.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a JobIAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobIAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobIAMBindingArgs | JobIAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobIAMBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["jobId"] = state ? state.jobId : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as JobIAMBindingArgs | undefined;
            if ((!args || args.jobId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobId'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["jobId"] = args ? args.jobId : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(JobIAMBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering JobIAMBinding resources.
 */
export interface JobIAMBindingState {
    condition?: pulumi.Input<inputs.dataproc.JobIAMBindingCondition>;
    /**
     * (Computed) The etag of the jobs's IAM policy.
     */
    etag?: pulumi.Input<string>;
    jobId?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project in which the job belongs. If it
     * is not provided, the provider will use a default.
     */
    project?: pulumi.Input<string>;
    /**
     * The region in which the job belongs. If it
     * is not provided, the provider will use a default.
     */
    region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a JobIAMBinding resource.
 */
export interface JobIAMBindingArgs {
    condition?: pulumi.Input<inputs.dataproc.JobIAMBindingCondition>;
    jobId: pulumi.Input<string>;
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project in which the job belongs. If it
     * is not provided, the provider will use a default.
     */
    project?: pulumi.Input<string>;
    /**
     * The region in which the job belongs. If it
     * is not provided, the provider will use a default.
     */
    region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
