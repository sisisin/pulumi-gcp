// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class ReservationSpecificReservationInstancePropertiesLocalSsd
    {
        public readonly int DiskSizeGb;
        public readonly string? Interface;

        [OutputConstructor]
        private ReservationSpecificReservationInstancePropertiesLocalSsd(
            int diskSizeGb,

            string? @interface)
        {
            DiskSizeGb = diskSizeGb;
            Interface = @interface;
        }
    }
}