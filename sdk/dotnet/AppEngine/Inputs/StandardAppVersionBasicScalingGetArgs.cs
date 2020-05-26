// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class StandardAppVersionBasicScalingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Duration of time after the last request that an instance must wait before the instance is shut down.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
        /// </summary>
        [Input("idleTimeout")]
        public Input<string>? IdleTimeout { get; set; }

        /// <summary>
        /// Maximum number of instances to create for this version. Must be in the range [1.0, 200.0].
        /// </summary>
        [Input("maxInstances", required: true)]
        public Input<int> MaxInstances { get; set; } = null!;

        public StandardAppVersionBasicScalingGetArgs()
        {
        }
    }
}