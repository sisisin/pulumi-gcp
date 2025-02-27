// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class WorkerPoolWorkerConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the disk attached to the worker, in GB. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If `0` is specified, Cloud Build will use a standard disk size.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        /// <summary>
        /// Machine type of a worker, such as `n1-standard-1`. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use `n1-standard-1`.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// If true, workers are created without any public address, which prevents network egress to public IPs.
        /// </summary>
        [Input("noExternalIp")]
        public Input<bool>? NoExternalIp { get; set; }

        public WorkerPoolWorkerConfigGetArgs()
        {
        }
    }
}
