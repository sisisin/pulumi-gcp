// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class ClusterNodePoolNetworkConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ) Whether to create a new range for pod IPs in this node pool. Defaults are provided for `pod_range` and `pod_ipv4_cidr_block` if they are not specified.
        /// </summary>
        [Input("createPodRange")]
        public Input<bool>? CreatePodRange { get; set; }

        /// <summary>
        /// ) The IP address range for pod IPs in this node pool. Only applicable if createPodRange is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.
        /// </summary>
        [Input("podIpv4CidrBlock")]
        public Input<string>? PodIpv4CidrBlock { get; set; }

        /// <summary>
        /// ) The ID of the secondary range for pod IPs. If `create_pod_range` is true, this ID is used for the new range. If `create_pod_range` is false, uses an existing secondary range with this ID.
        /// </summary>
        [Input("podRange", required: true)]
        public Input<string> PodRange { get; set; } = null!;

        public ClusterNodePoolNetworkConfigGetArgs()
        {
        }
    }
}
