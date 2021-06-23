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
    public sealed class GetResourcePolicyInstanceSchedulePolicyResult
    {
        public readonly string ExpirationTime;
        public readonly string StartTime;
        public readonly string TimeZone;
        public readonly ImmutableArray<Outputs.GetResourcePolicyInstanceSchedulePolicyVmStartScheduleResult> VmStartSchedules;
        public readonly ImmutableArray<Outputs.GetResourcePolicyInstanceSchedulePolicyVmStopScheduleResult> VmStopSchedules;

        [OutputConstructor]
        private GetResourcePolicyInstanceSchedulePolicyResult(
            string expirationTime,

            string startTime,

            string timeZone,

            ImmutableArray<Outputs.GetResourcePolicyInstanceSchedulePolicyVmStartScheduleResult> vmStartSchedules,

            ImmutableArray<Outputs.GetResourcePolicyInstanceSchedulePolicyVmStopScheduleResult> vmStopSchedules)
        {
            ExpirationTime = expirationTime;
            StartTime = startTime;
            TimeZone = timeZone;
            VmStartSchedules = vmStartSchedules;
            VmStopSchedules = vmStopSchedules;
        }
    }
}