// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class ImageRawDiskArgs : Pulumi.ResourceArgs
    {
        [Input("containerType")]
        public Input<string>? ContainerType { get; set; }

        [Input("sha1")]
        public Input<string>? Sha1 { get; set; }

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public ImageRawDiskArgs()
        {
        }
    }
}