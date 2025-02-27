// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Compute Engine MachineImage. Each of these resources serves a different use case:
 *
 * * `gcp.compute.MachineImageIamPolicy`: Authoritative. Sets the IAM policy for the machineimage and replaces any existing policy already attached.
 * * `gcp.compute.MachineImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the machineimage are preserved.
 * * `gcp.compute.MachineImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the machineimage are preserved.
 *
 * > **Note:** `gcp.compute.MachineImageIamPolicy` **cannot** be used in conjunction with `gcp.compute.MachineImageIamBinding` and `gcp.compute.MachineImageIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.compute.MachineImageIamBinding` resources **can be** used in conjunction with `gcp.compute.MachineImageIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_compute\_machine\_image\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/compute.admin",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.compute.MachineImageIamPolicy("policy", {
 *     project: google_compute_machine_image.image.project,
 *     machineImage: google_compute_machine_image.image.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/compute.admin",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expires_after_2019_12_31",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const policy = new gcp.compute.MachineImageIamPolicy("policy", {
 *     project: google_compute_machine_image.image.project,
 *     machineImage: google_compute_machine_image.image.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * ## google\_compute\_machine\_image\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.compute.MachineImageIamBinding("binding", {
 *     project: google_compute_machine_image.image.project,
 *     machineImage: google_compute_machine_image.image.name,
 *     role: "roles/compute.admin",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.compute.MachineImageIamBinding("binding", {
 *     project: google_compute_machine_image.image.project,
 *     machineImage: google_compute_machine_image.image.name,
 *     role: "roles/compute.admin",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * ## google\_compute\_machine\_image\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.compute.MachineImageIamMember("member", {
 *     project: google_compute_machine_image.image.project,
 *     machineImage: google_compute_machine_image.image.name,
 *     role: "roles/compute.admin",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.compute.MachineImageIamMember("member", {
 *     project: google_compute_machine_image.image.project,
 *     machineImage: google_compute_machine_image.image.name,
 *     role: "roles/compute.admin",
 *     member: "user:jane@example.com",
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/global/machineImages/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine machineimage IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/machineImageIamBinding:MachineImageIamBinding editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/machineImageIamBinding:MachineImageIamBinding editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/machineImageIamBinding:MachineImageIamBinding editor projects/{{project}}/global/machineImages/{{machine_image}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class MachineImageIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing MachineImageIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MachineImageIamBindingState, opts?: pulumi.CustomResourceOptions): MachineImageIamBinding {
        return new MachineImageIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/machineImageIamBinding:MachineImageIamBinding';

    /**
     * Returns true if the given object is an instance of MachineImageIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachineImageIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachineImageIamBinding.__pulumiType;
    }

    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.compute.MachineImageIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly machineImage!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a MachineImageIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MachineImageIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MachineImageIamBindingArgs | MachineImageIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MachineImageIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["machineImage"] = state ? state.machineImage : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as MachineImageIamBindingArgs | undefined;
            if ((!args || args.machineImage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineImage'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["machineImage"] = args ? args.machineImage : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MachineImageIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MachineImageIamBinding resources.
 */
export interface MachineImageIamBindingState {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.compute.MachineImageIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    machineImage?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MachineImageIamBinding resource.
 */
export interface MachineImageIamBindingArgs {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.compute.MachineImageIamBindingCondition>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    machineImage: pulumi.Input<string>;
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
