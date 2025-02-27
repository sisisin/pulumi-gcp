// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A repository for storing artifacts
 *
 * To get more information about Repository, see:
 *
 * * [API documentation](https://cloud.google.com/artifact-registry/docs/reference/rest/v1beta2/projects.locations.repositories)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/artifact-registry/docs/overview)
 *
 * ## Example Usage
 * ### Artifact Registry Repository Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my_repo = new gcp.artifactregistry.Repository("my-repo", {
 *     location: "us-central1",
 *     repositoryId: "my-repository",
 *     description: "example docker repository",
 *     format: "DOCKER",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Artifact Registry Repository Cmek
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my_repo = new gcp.artifactregistry.Repository("my-repo", {
 *     location: "us-central1",
 *     repositoryId: "my-repository",
 *     description: "example docker repository with cmek",
 *     format: "DOCKER",
 *     kmsKeyName: "kms-key",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Artifact Registry Repository Iam
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my_repo = new gcp.artifactregistry.Repository("my-repo", {
 *     location: "us-central1",
 *     repositoryId: "my-repository",
 *     description: "example docker repository with iam",
 *     format: "DOCKER",
 * }, {
 *     provider: google_beta,
 * });
 * const test_account = new gcp.serviceaccount.Account("test-account", {
 *     accountId: "my-account",
 *     displayName: "Test Service Account",
 * }, {
 *     provider: google_beta,
 * });
 * const test_iam = new gcp.artifactregistry.RepositoryIamMember("test-iam", {
 *     location: my_repo.location,
 *     repository: my_repo.name,
 *     role: "roles/artifactregistry.reader",
 *     member: pulumi.interpolate`serviceAccount:${test_account.email}`,
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * Repository can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repository:Repository default projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repository:Repository default {{project}}/{{location}}/{{repository_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repository:Repository default {{location}}/{{repository_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repository:Repository default {{repository_id}}
 * ```
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryState, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:artifactregistry/repository:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * The time when the repository was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The user-provided description of the repository.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The format of packages that are stored in the repository. You can only create
     * alpha formats if you are a member of the [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
     * - DOCKER
     * - MAVEN ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - NPM ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - PYTHON ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - APT ([alpha](https://cloud.google.com/products#product-launch-stages))
     * - YUM ([alpha](https://cloud.google.com/products#product-launch-stages))
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * The Cloud KMS resource name of the customer managed encryption key that’s
     * used to encrypt the contents of the Repository. Has the form:
     * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
     * This value may not be changed after the Repository has been created.
     */
    public readonly kmsKeyName!: pulumi.Output<string | undefined>;
    /**
     * Labels with user-defined metadata.
     * This field may contain up to 64 entries. Label keys and values may be no
     * longer than 63 characters. Label keys must begin with a lowercase letter
     * and may only contain lowercase letters, numeric characters, underscores,
     * and dashes.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the location this repository is located in.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1"
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The last part of the repository name, for example:
     * "repo1"
     */
    public readonly repositoryId!: pulumi.Output<string>;
    /**
     * The time when the repository was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryArgs | RepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["kmsKeyName"] = state ? state.kmsKeyName : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["repositoryId"] = state ? state.repositoryId : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            if ((!args || args.format === undefined) && !opts.urn) {
                throw new Error("Missing required property 'format'");
            }
            if ((!args || args.repositoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["repositoryId"] = args ? args.repositoryId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Repository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * The time when the repository was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The user-provided description of the repository.
     */
    description?: pulumi.Input<string>;
    /**
     * The format of packages that are stored in the repository. You can only create
     * alpha formats if you are a member of the [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
     * - DOCKER
     * - MAVEN ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - NPM ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - PYTHON ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - APT ([alpha](https://cloud.google.com/products#product-launch-stages))
     * - YUM ([alpha](https://cloud.google.com/products#product-launch-stages))
     */
    format?: pulumi.Input<string>;
    /**
     * The Cloud KMS resource name of the customer managed encryption key that’s
     * used to encrypt the contents of the Repository. Has the form:
     * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
     * This value may not be changed after the Repository has been created.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Labels with user-defined metadata.
     * This field may contain up to 64 entries. Label keys and values may be no
     * longer than 63 characters. Label keys must begin with a lowercase letter
     * and may only contain lowercase letters, numeric characters, underscores,
     * and dashes.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location this repository is located in.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1"
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The last part of the repository name, for example:
     * "repo1"
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The time when the repository was last updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * The user-provided description of the repository.
     */
    description?: pulumi.Input<string>;
    /**
     * The format of packages that are stored in the repository. You can only create
     * alpha formats if you are a member of the [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
     * - DOCKER
     * - MAVEN ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - NPM ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - PYTHON ([Preview](https://cloud.google.com/products#product-launch-stages))
     * - APT ([alpha](https://cloud.google.com/products#product-launch-stages))
     * - YUM ([alpha](https://cloud.google.com/products#product-launch-stages))
     */
    format: pulumi.Input<string>;
    /**
     * The Cloud KMS resource name of the customer managed encryption key that’s
     * used to encrypt the contents of the Repository. Has the form:
     * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
     * This value may not be changed after the Repository has been created.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Labels with user-defined metadata.
     * This field may contain up to 64 entries. Label keys and values may be no
     * longer than 63 characters. Label keys must begin with a lowercase letter
     * and may only contain lowercase letters, numeric characters, underscores,
     * and dashes.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location this repository is located in.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The last part of the repository name, for example:
     * "repo1"
     */
    repositoryId: pulumi.Input<string>;
}
