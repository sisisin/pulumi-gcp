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
    public sealed class RegionInstanceGroupManagerUpdatePolicy
    {
        /// <summary>
        /// - The instance redistribution policy for regional managed instance groups. Valid values are: `"PROACTIVE"`, `"NONE"`. If `PROACTIVE` (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If `NONE`, proactive redistribution is disabled.
        /// </summary>
        public readonly string? InstanceRedistributionType;
        /// <summary>
        /// , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. It has to be either 0 or at least equal to the number of zones.  If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
        /// </summary>
        public readonly int? MaxSurgeFixed;
        /// <summary>
        /// , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
        /// </summary>
        public readonly int? MaxSurgePercent;
        /// <summary>
        /// , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
        /// </summary>
        public readonly int? MaxUnavailableFixed;
        /// <summary>
        /// , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
        /// </summary>
        public readonly int? MaxUnavailablePercent;
        /// <summary>
        /// , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
        /// - - -
        /// </summary>
        public readonly int? MinReadySec;
        /// <summary>
        /// - Minimal action to be taken on an instance. You can specify either `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `RESTART`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
        /// </summary>
        public readonly string MinimalAction;
        /// <summary>
        /// - The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private RegionInstanceGroupManagerUpdatePolicy(
            string? instanceRedistributionType,

            int? maxSurgeFixed,

            int? maxSurgePercent,

            int? maxUnavailableFixed,

            int? maxUnavailablePercent,

            int? minReadySec,

            string minimalAction,

            string type)
        {
            InstanceRedistributionType = instanceRedistributionType;
            MaxSurgeFixed = maxSurgeFixed;
            MaxSurgePercent = maxSurgePercent;
            MaxUnavailableFixed = maxUnavailableFixed;
            MaxUnavailablePercent = maxUnavailablePercent;
            MinReadySec = minReadySec;
            MinimalAction = minimalAction;
            Type = type;
        }
    }
}