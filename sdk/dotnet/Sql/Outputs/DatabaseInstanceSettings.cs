// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql.Outputs
{

    [OutputType]
    public sealed class DatabaseInstanceSettings
    {
        /// <summary>
        /// This specifies when the instance should be
        /// active. Can be either `ALWAYS`, `NEVER` or `ON_DEMAND`.
        /// </summary>
        public readonly string? ActivationPolicy;
        /// <summary>
        /// This property is only applicable to First Generation instances.
        /// First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen)
        /// for information on how to upgrade to Second Generation instances.
        /// A list of Google App Engine (GAE) project names that are allowed to access this instance.
        /// </summary>
        public readonly ImmutableArray<string> AuthorizedGaeApplications;
        /// <summary>
        /// The availability type of the Cloud SQL
        /// instance, high availability (`REGIONAL`) or single zone (`ZONAL`).' For MySQL
        /// instances, ensure that `settings.backup_configuration.enabled` and
        /// `settings.backup_configuration.binary_log_enabled` are both set to `true`.
        /// </summary>
        public readonly string? AvailabilityType;
        public readonly Outputs.DatabaseInstanceSettingsBackupConfiguration? BackupConfiguration;
        /// <summary>
        /// The name of server instance collation.
        /// </summary>
        public readonly string? Collation;
        /// <summary>
        /// This property is only applicable to First Generation instances.
        /// First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen)
        /// for information on how to upgrade to Second Generation instances.
        /// Specific to read instances, indicates
        /// when crash-safe replication flags are enabled.
        /// </summary>
        public readonly bool? CrashSafeReplication;
        public readonly ImmutableArray<Outputs.DatabaseInstanceSettingsDatabaseFlag> DatabaseFlags;
        /// <summary>
        /// Configuration to increase storage size automatically.  Note that future `pulumi apply` calls will attempt to resize the disk to the value specified in `disk_size` - if this is set, do not set `disk_size`.
        /// </summary>
        public readonly bool? DiskAutoresize;
        public readonly int? DiskAutoresizeLimit;
        /// <summary>
        /// The size of data disk, in GB. Size of a running instance cannot be reduced but can be increased.
        /// </summary>
        public readonly int? DiskSize;
        /// <summary>
        /// The type of data disk: PD_SSD or PD_HDD.
        /// </summary>
        public readonly string? DiskType;
        public readonly Outputs.DatabaseInstanceSettingsInsightsConfig? InsightsConfig;
        public readonly Outputs.DatabaseInstanceSettingsIpConfiguration? IpConfiguration;
        public readonly Outputs.DatabaseInstanceSettingsLocationPreference? LocationPreference;
        public readonly Outputs.DatabaseInstanceSettingsMaintenanceWindow? MaintenanceWindow;
        /// <summary>
        /// Pricing plan for this instance, can only be `PER_USE`.
        /// </summary>
        public readonly string? PricingPlan;
        /// <summary>
        /// This property is only applicable to First Generation instances.
        /// First Generation instances are now deprecated, see [here](https://cloud.google.com/sql/docs/mysql/upgrade-2nd-gen)
        /// for information on how to upgrade to Second Generation instances.
        /// Replication type for this instance, can be one of `ASYNCHRONOUS` or `SYNCHRONOUS`.
        /// </summary>
        public readonly string? ReplicationType;
        /// <summary>
        /// The machine type to use. See [tiers](https://cloud.google.com/sql/docs/admin-api/v1beta4/tiers)
        /// for more details and supported versions. Postgres supports only shared-core machine types,
        /// and custom machine types such as `db-custom-2-13312`. See the [Custom Machine Type Documentation](https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#create) to learn about specifying custom machine types.
        /// </summary>
        public readonly string Tier;
        /// <summary>
        /// A set of key/value user label pairs to assign to the instance.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? UserLabels;
        public readonly int? Version;

        [OutputConstructor]
        private DatabaseInstanceSettings(
            string? activationPolicy,

            ImmutableArray<string> authorizedGaeApplications,

            string? availabilityType,

            Outputs.DatabaseInstanceSettingsBackupConfiguration? backupConfiguration,

            string? collation,

            bool? crashSafeReplication,

            ImmutableArray<Outputs.DatabaseInstanceSettingsDatabaseFlag> databaseFlags,

            bool? diskAutoresize,

            int? diskAutoresizeLimit,

            int? diskSize,

            string? diskType,

            Outputs.DatabaseInstanceSettingsInsightsConfig? insightsConfig,

            Outputs.DatabaseInstanceSettingsIpConfiguration? ipConfiguration,

            Outputs.DatabaseInstanceSettingsLocationPreference? locationPreference,

            Outputs.DatabaseInstanceSettingsMaintenanceWindow? maintenanceWindow,

            string? pricingPlan,

            string? replicationType,

            string tier,

            ImmutableDictionary<string, string>? userLabels,

            int? version)
        {
            ActivationPolicy = activationPolicy;
            AuthorizedGaeApplications = authorizedGaeApplications;
            AvailabilityType = availabilityType;
            BackupConfiguration = backupConfiguration;
            Collation = collation;
            CrashSafeReplication = crashSafeReplication;
            DatabaseFlags = databaseFlags;
            DiskAutoresize = diskAutoresize;
            DiskAutoresizeLimit = diskAutoresizeLimit;
            DiskSize = diskSize;
            DiskType = diskType;
            InsightsConfig = insightsConfig;
            IpConfiguration = ipConfiguration;
            LocationPreference = locationPreference;
            MaintenanceWindow = maintenanceWindow;
            PricingPlan = pricingPlan;
            ReplicationType = replicationType;
            Tier = tier;
            UserLabels = userLabels;
            Version = version;
        }
    }
}
