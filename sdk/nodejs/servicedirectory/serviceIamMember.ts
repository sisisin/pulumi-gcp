// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Service Directory Service. Each of these resources serves a different use case:
 *
 * * `gcp.servicedirectory.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
 * * `gcp.servicedirectory.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
 * * `gcp.servicedirectory.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
 *
 * > **Note:** `gcp.servicedirectory.ServiceIamPolicy` **cannot** be used in conjunction with `gcp.servicedirectory.ServiceIamBinding` and `gcp.servicedirectory.ServiceIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.servicedirectory.ServiceIamBinding` resources **can be** used in conjunction with `gcp.servicedirectory.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_service\_directory\_service\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.servicedirectory.ServiceIamPolicy("policy", {policyData: admin.then(admin => admin.policyData)});
 * ```
 *
 * ## google\_service\_directory\_service\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.servicedirectory.ServiceIamBinding("binding", {
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_service\_directory\_service\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.servicedirectory.ServiceIamMember("member", {
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} * {{project}}/{{location}}/{{namespace_id}}/{{service_id}} * {{location}}/{{namespace_id}}/{{service_id}} Any variables not passed in the import command will be taken from the provider configuration. Service Directory service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class ServiceIamMember extends pulumi.CustomResource {
    /**
     * Get an existing ServiceIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceIamMemberState, opts?: pulumi.CustomResourceOptions): ServiceIamMember {
        return new ServiceIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:servicedirectory/serviceIamMember:ServiceIamMember';

    /**
     * Returns true if the given object is an instance of ServiceIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceIamMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.servicedirectory.ServiceIamMemberCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a ServiceIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceIamMemberArgs | ServiceIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceIamMemberState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ServiceIamMemberArgs | undefined;
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceIamMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceIamMember resources.
 */
export interface ServiceIamMemberState {
    condition?: pulumi.Input<inputs.servicedirectory.ServiceIamMemberCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    member?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceIamMember resource.
 */
export interface ServiceIamMemberArgs {
    condition?: pulumi.Input<inputs.servicedirectory.ServiceIamMemberCondition>;
    member: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
