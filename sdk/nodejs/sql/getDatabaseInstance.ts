// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a Cloud SQL instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const qa = pulumi.output(gcp.sql.getDatabaseInstance({
 *     name: "test-sql-instance",
 * }));
 * ```
 */
export function getDatabaseInstance(args: GetDatabaseInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseInstanceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:sql/getDatabaseInstance:getDatabaseInstance", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseInstance.
 */
export interface GetDatabaseInstanceArgs {
    /**
     * The name of the instance.
     */
    name: string;
    /**
     * The ID of the project in which the resource belongs.
     */
    project?: string;
}

/**
 * A collection of values returned by getDatabaseInstance.
 */
export interface GetDatabaseInstanceResult {
    readonly clones: outputs.sql.GetDatabaseInstanceClone[];
    readonly connectionName: string;
    readonly databaseVersion: string;
    readonly deletionProtection: boolean;
    readonly encryptionKeyName: string;
    readonly firstIpAddress: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddresses: outputs.sql.GetDatabaseInstanceIpAddress[];
    readonly masterInstanceName: string;
    readonly name: string;
    readonly privateIpAddress: string;
    readonly project?: string;
    readonly publicIpAddress: string;
    readonly region: string;
    readonly replicaConfigurations: outputs.sql.GetDatabaseInstanceReplicaConfiguration[];
    readonly restoreBackupContexts: outputs.sql.GetDatabaseInstanceRestoreBackupContext[];
    readonly rootPassword: string;
    readonly selfLink: string;
    readonly serverCaCerts: outputs.sql.GetDatabaseInstanceServerCaCert[];
    readonly serviceAccountEmailAddress: string;
    readonly settings: outputs.sql.GetDatabaseInstanceSetting[];
}
