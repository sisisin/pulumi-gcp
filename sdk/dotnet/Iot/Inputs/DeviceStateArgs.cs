// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iot.Inputs
{

    public sealed class DeviceStateArgs : Pulumi.ResourceArgs
    {
        [Input("binaryData")]
        public Input<string>? BinaryData { get; set; }

        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public DeviceStateArgs()
        {
        }
    }
}
