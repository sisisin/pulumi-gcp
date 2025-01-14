// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * To get more information about RuntimeConfigs, see:
 *
 * * [API documentation](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/v1beta1/projects.configs)
 * * How-to Guides
 *     * [Runtime Configurator Fundamentals](https://cloud.google.com/deployment-manager/runtime-configurator/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const run_service = pulumi.output(gcp.runtimeconfig.getConfig({
 *     name: "my-service",
 * }));
 * ```
 */
export function getConfig(args: GetConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:runtimeconfig/getConfig:getConfig", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getConfig.
 */
export interface GetConfigArgs {
    /**
     * The name of the Runtime Configurator configuration.
     */
    name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: string;
}

/**
 * A collection of values returned by getConfig.
 */
export interface GetConfigResult {
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly project?: string;
}
