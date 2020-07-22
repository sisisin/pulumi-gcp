// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentState, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:notebooks/environment:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * Use a container image to start the notebook instance.  Structure is documented below.
     */
    public readonly containerImage!: pulumi.Output<outputs.notebooks.EnvironmentContainerImage | undefined>;
    /**
     * Instance creation time
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A brief description of this environment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Display name of this environment for the UI.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * A reference to the zone where the machine resides.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name specified for the Environment instance.
     * Format: projects/{project_id}/locations/{location}/environments/{environmentId}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up.
     * The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
     */
    public readonly postStartupScript!: pulumi.Output<string | undefined>;
    /**
     * The name of the Google Cloud project that this VM image belongs to.
     * Format: projects/{project_id}
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Use a Compute Engine VM image to start the notebook instance.  Structure is documented below.
     */
    public readonly vmImage!: pulumi.Output<outputs.notebooks.EnvironmentVmImage | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentArgs | EnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EnvironmentState | undefined;
            inputs["containerImage"] = state ? state.containerImage : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["postStartupScript"] = state ? state.postStartupScript : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["vmImage"] = state ? state.vmImage : undefined;
        } else {
            const args = argsOrState as EnvironmentArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            inputs["containerImage"] = args ? args.containerImage : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["postStartupScript"] = args ? args.postStartupScript : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["vmImage"] = args ? args.vmImage : undefined;
            inputs["createTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Environment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Environment resources.
 */
export interface EnvironmentState {
    /**
     * Use a container image to start the notebook instance.  Structure is documented below.
     */
    readonly containerImage?: pulumi.Input<inputs.notebooks.EnvironmentContainerImage>;
    /**
     * Instance creation time
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * A brief description of this environment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Display name of this environment for the UI.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * A reference to the zone where the machine resides.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name specified for the Environment instance.
     * Format: projects/{project_id}/locations/{location}/environments/{environmentId}
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up.
     * The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
     */
    readonly postStartupScript?: pulumi.Input<string>;
    /**
     * The name of the Google Cloud project that this VM image belongs to.
     * Format: projects/{project_id}
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Use a Compute Engine VM image to start the notebook instance.  Structure is documented below.
     */
    readonly vmImage?: pulumi.Input<inputs.notebooks.EnvironmentVmImage>;
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * Use a container image to start the notebook instance.  Structure is documented below.
     */
    readonly containerImage?: pulumi.Input<inputs.notebooks.EnvironmentContainerImage>;
    /**
     * A brief description of this environment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Display name of this environment for the UI.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * A reference to the zone where the machine resides.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name specified for the Environment instance.
     * Format: projects/{project_id}/locations/{location}/environments/{environmentId}
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up.
     * The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
     */
    readonly postStartupScript?: pulumi.Input<string>;
    /**
     * The name of the Google Cloud project that this VM image belongs to.
     * Format: projects/{project_id}
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Use a Compute Engine VM image to start the notebook instance.  Structure is documented below.
     */
    readonly vmImage?: pulumi.Input<inputs.notebooks.EnvironmentVmImage>;
}