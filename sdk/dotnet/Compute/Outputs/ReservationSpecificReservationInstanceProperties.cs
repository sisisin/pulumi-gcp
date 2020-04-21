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
    public sealed class ReservationSpecificReservationInstanceProperties
    {
        public readonly ImmutableArray<Outputs.ReservationSpecificReservationInstancePropertiesGuestAccelerator> GuestAccelerators;
        public readonly ImmutableArray<Outputs.ReservationSpecificReservationInstancePropertiesLocalSsd> LocalSsds;
        public readonly string MachineType;
        public readonly string? MinCpuPlatform;

        [OutputConstructor]
        private ReservationSpecificReservationInstanceProperties(
            ImmutableArray<Outputs.ReservationSpecificReservationInstancePropertiesGuestAccelerator> guestAccelerators,

            ImmutableArray<Outputs.ReservationSpecificReservationInstancePropertiesLocalSsd> localSsds,

            string machineType,

            string? minCpuPlatform)
        {
            GuestAccelerators = guestAccelerators;
            LocalSsds = localSsds;
            MachineType = machineType;
            MinCpuPlatform = minCpuPlatform;
        }
    }
}