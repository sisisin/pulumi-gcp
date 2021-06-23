// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The autoscaling policy used by the cluster. Only resource names including projectid and location (region) are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/` Note that the policy must be in the same project and Dataproc region.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigGetArgs()
        {
        }
    }
}