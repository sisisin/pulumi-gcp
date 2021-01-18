// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub.Inputs
{

    public sealed class LiteTopicPartitionConfigCapacityGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Subscribe throughput capacity per partition in MiB/s. Must be &gt;= 4 and &lt;= 16.
        /// </summary>
        [Input("publishMibPerSec", required: true)]
        public Input<int> PublishMibPerSec { get; set; } = null!;

        /// <summary>
        /// Publish throughput capacity per partition in MiB/s. Must be &gt;= 4 and &lt;= 16.
        /// </summary>
        [Input("subscribeMibPerSec", required: true)]
        public Input<int> SubscribeMibPerSec { get; set; } = null!;

        public LiteTopicPartitionConfigCapacityGetArgs()
        {
        }
    }
}