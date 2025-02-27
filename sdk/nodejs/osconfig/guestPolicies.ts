// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * An OS Config resource representing a guest configuration policy. These policies represent
 * the desired state for VM instance guest environments including packages to install or remove,
 * package repository configurations, and software to install.
 *
 * To get more information about GuestPolicies, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/osconfig/rest)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/os-config-management)
 *
 * ## Example Usage
 * ### Os Config Guest Policies Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const foobar = new gcp.compute.Instance("foobar", {
 *     machineType: "e2-medium",
 *     zone: "us-central1-a",
 *     canIpForward: false,
 *     tags: [
 *         "foo",
 *         "bar",
 *     ],
 *     bootDisk: {
 *         initializeParams: {
 *             image: myImage.then(myImage => myImage.selfLink),
 *         },
 *     },
 *     networkInterfaces: [{
 *         network: "default",
 *     }],
 *     metadata: {
 *         foo: "bar",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const guestPolicies = new gcp.osconfig.GuestPolicies("guestPolicies", {
 *     guestPolicyId: "guest-policy",
 *     assignment: {
 *         instances: [foobar.id],
 *     },
 *     packages: [{
 *         name: "my-package",
 *         desiredState: "UPDATED",
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Os Config Guest Policies Packages
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const guestPolicies = new gcp.osconfig.GuestPolicies("guestPolicies", {
 *     guestPolicyId: "guest-policy",
 *     assignment: {
 *         groupLabels: [
 *             {
 *                 labels: {
 *                     color: "red",
 *                     env: "test",
 *                 },
 *             },
 *             {
 *                 labels: {
 *                     color: "blue",
 *                     env: "test",
 *                 },
 *             },
 *         ],
 *     },
 *     packages: [
 *         {
 *             name: "my-package",
 *             desiredState: "INSTALLED",
 *         },
 *         {
 *             name: "bad-package-1",
 *             desiredState: "REMOVED",
 *         },
 *         {
 *             name: "bad-package-2",
 *             desiredState: "REMOVED",
 *             manager: "APT",
 *         },
 *     ],
 *     packageRepositories: [
 *         {
 *             apt: {
 *                 uri: "https://packages.cloud.google.com/apt",
 *                 archiveType: "DEB",
 *                 distribution: "cloud-sdk-stretch",
 *                 components: ["main"],
 *             },
 *         },
 *         {
 *             yum: {
 *                 id: "google-cloud-sdk",
 *                 displayName: "Google Cloud SDK",
 *                 baseUrl: "https://packages.cloud.google.com/yum/repos/cloud-sdk-el7-x86_64",
 *                 gpgKeys: [
 *                     "https://packages.cloud.google.com/yum/doc/yum-key.gpg",
 *                     "https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg",
 *                 ],
 *             },
 *         },
 *     ],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Os Config Guest Policies Recipes
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const guestPolicies = new gcp.osconfig.GuestPolicies("guestPolicies", {
 *     guestPolicyId: "guest-policy",
 *     assignment: {
 *         zones: [
 *             "us-east1-b",
 *             "us-east1-d",
 *         ],
 *     },
 *     recipes: [{
 *         name: "guest-policy-recipe",
 *         desiredState: "INSTALLED",
 *         artifacts: [{
 *             id: "guest-policy-artifact-id",
 *             gcs: {
 *                 bucket: "my-bucket",
 *                 object: "executable.msi",
 *                 generation: 1546030865175603,
 *             },
 *         }],
 *         installSteps: [{
 *             msiInstallation: {
 *                 artifactId: "guest-policy-artifact-id",
 *             },
 *         }],
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * GuestPolicies can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default projects/{{project}}/guestPolicies/{{guest_policy_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default {{project}}/{{guest_policy_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default {{guest_policy_id}}
 * ```
 */
export class GuestPolicies extends pulumi.CustomResource {
    /**
     * Get an existing GuestPolicies resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GuestPoliciesState, opts?: pulumi.CustomResourceOptions): GuestPolicies {
        return new GuestPolicies(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:osconfig/guestPolicies:GuestPolicies';

    /**
     * Returns true if the given object is an instance of GuestPolicies.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GuestPolicies {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GuestPolicies.__pulumiType;
    }

    /**
     * Specifies the VM instances that are assigned to this policy. This allows you to target sets
     * or groups of VM instances by different parameters such as labels, names, OS, or zones.
     * If left empty, all VM instances underneath this policy are targeted.
     * At the same level in the resource hierarchy (that is within a project), the service prevents
     * the creation of multiple policies that conflict with each other.
     * For more information, see how the service
     * [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
     * Structure is documented below.
     */
    public readonly assignment!: pulumi.Output<outputs.osconfig.GuestPoliciesAssignment>;
    /**
     * Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
     * "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the guest policy. Length of the description is limited to 1024 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The etag for this guest policy. If this is provided on update, it must match the server's etag.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The logical name of the guest policy in the project with the following restrictions:
     * * Must contain only lowercase letters, numbers, and hyphens.
     * * Must start with a letter.
     * * Must be between 1-63 characters.
     * * Must end with a number or a letter.
     * * Must be unique within the project.
     */
    public readonly guestPolicyId!: pulumi.Output<string>;
    /**
     * Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
     * Names are also used to identify resources which helps to determine whether guest policies have conflicts.
     * This means that requests to create multiple recipes with the same name and version are rejected since they
     * could potentially have conflicting assignments.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A list of package repositories to configure on the VM instance.
     * This is done before any other configs are applied so they can use these repos.
     * Package repositories are only configured if the corresponding package manager(s) are available.
     * Structure is documented below.
     */
    public readonly packageRepositories!: pulumi.Output<outputs.osconfig.GuestPoliciesPackageRepository[] | undefined>;
    /**
     * The software packages to be managed by this policy.
     * Structure is documented below.
     */
    public readonly packages!: pulumi.Output<outputs.osconfig.GuestPoliciesPackage[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A list of Recipes to install on the VM instance.
     * Structure is documented below.
     */
    public readonly recipes!: pulumi.Output<outputs.osconfig.GuestPoliciesRecipe[] | undefined>;
    /**
     * Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
     * "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a GuestPolicies resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GuestPoliciesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GuestPoliciesArgs | GuestPoliciesState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GuestPoliciesState | undefined;
            inputs["assignment"] = state ? state.assignment : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["guestPolicyId"] = state ? state.guestPolicyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["packageRepositories"] = state ? state.packageRepositories : undefined;
            inputs["packages"] = state ? state.packages : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["recipes"] = state ? state.recipes : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as GuestPoliciesArgs | undefined;
            if ((!args || args.assignment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assignment'");
            }
            if ((!args || args.guestPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'guestPolicyId'");
            }
            inputs["assignment"] = args ? args.assignment : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["guestPolicyId"] = args ? args.guestPolicyId : undefined;
            inputs["packageRepositories"] = args ? args.packageRepositories : undefined;
            inputs["packages"] = args ? args.packages : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["recipes"] = args ? args.recipes : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GuestPolicies.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GuestPolicies resources.
 */
export interface GuestPoliciesState {
    /**
     * Specifies the VM instances that are assigned to this policy. This allows you to target sets
     * or groups of VM instances by different parameters such as labels, names, OS, or zones.
     * If left empty, all VM instances underneath this policy are targeted.
     * At the same level in the resource hierarchy (that is within a project), the service prevents
     * the creation of multiple policies that conflict with each other.
     * For more information, see how the service
     * [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
     * Structure is documented below.
     */
    assignment?: pulumi.Input<inputs.osconfig.GuestPoliciesAssignment>;
    /**
     * Time this guest policy was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
     * "2014-10-02T15:01:23.045123456Z".
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of the guest policy. Length of the description is limited to 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The etag for this guest policy. If this is provided on update, it must match the server's etag.
     */
    etag?: pulumi.Input<string>;
    /**
     * The logical name of the guest policy in the project with the following restrictions:
     * * Must contain only lowercase letters, numbers, and hyphens.
     * * Must start with a letter.
     * * Must be between 1-63 characters.
     * * Must end with a number or a letter.
     * * Must be unique within the project.
     */
    guestPolicyId?: pulumi.Input<string>;
    /**
     * Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
     * Names are also used to identify resources which helps to determine whether guest policies have conflicts.
     * This means that requests to create multiple recipes with the same name and version are rejected since they
     * could potentially have conflicting assignments.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of package repositories to configure on the VM instance.
     * This is done before any other configs are applied so they can use these repos.
     * Package repositories are only configured if the corresponding package manager(s) are available.
     * Structure is documented below.
     */
    packageRepositories?: pulumi.Input<pulumi.Input<inputs.osconfig.GuestPoliciesPackageRepository>[]>;
    /**
     * The software packages to be managed by this policy.
     * Structure is documented below.
     */
    packages?: pulumi.Input<pulumi.Input<inputs.osconfig.GuestPoliciesPackage>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A list of Recipes to install on the VM instance.
     * Structure is documented below.
     */
    recipes?: pulumi.Input<pulumi.Input<inputs.osconfig.GuestPoliciesRecipe>[]>;
    /**
     * Last time this guest policy was updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example:
     * "2014-10-02T15:01:23.045123456Z".
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GuestPolicies resource.
 */
export interface GuestPoliciesArgs {
    /**
     * Specifies the VM instances that are assigned to this policy. This allows you to target sets
     * or groups of VM instances by different parameters such as labels, names, OS, or zones.
     * If left empty, all VM instances underneath this policy are targeted.
     * At the same level in the resource hierarchy (that is within a project), the service prevents
     * the creation of multiple policies that conflict with each other.
     * For more information, see how the service
     * [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
     * Structure is documented below.
     */
    assignment: pulumi.Input<inputs.osconfig.GuestPoliciesAssignment>;
    /**
     * Description of the guest policy. Length of the description is limited to 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The etag for this guest policy. If this is provided on update, it must match the server's etag.
     */
    etag?: pulumi.Input<string>;
    /**
     * The logical name of the guest policy in the project with the following restrictions:
     * * Must contain only lowercase letters, numbers, and hyphens.
     * * Must start with a letter.
     * * Must be between 1-63 characters.
     * * Must end with a number or a letter.
     * * Must be unique within the project.
     */
    guestPolicyId: pulumi.Input<string>;
    /**
     * A list of package repositories to configure on the VM instance.
     * This is done before any other configs are applied so they can use these repos.
     * Package repositories are only configured if the corresponding package manager(s) are available.
     * Structure is documented below.
     */
    packageRepositories?: pulumi.Input<pulumi.Input<inputs.osconfig.GuestPoliciesPackageRepository>[]>;
    /**
     * The software packages to be managed by this policy.
     * Structure is documented below.
     */
    packages?: pulumi.Input<pulumi.Input<inputs.osconfig.GuestPoliciesPackage>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A list of Recipes to install on the VM instance.
     * Structure is documented below.
     */
    recipes?: pulumi.Input<pulumi.Input<inputs.osconfig.GuestPoliciesRecipe>[]>;
}
