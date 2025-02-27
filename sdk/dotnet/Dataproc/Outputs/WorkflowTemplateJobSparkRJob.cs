// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class WorkflowTemplateJobSparkRJob
    {
        /// <summary>
        /// Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
        /// </summary>
        public readonly ImmutableArray<string> ArchiveUris;
        /// <summary>
        /// Optional. The arguments to pass to the driver. Do not include arguments, such as `--conf`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.
        /// </summary>
        public readonly ImmutableArray<string> FileUris;
        /// <summary>
        /// Optional. The runtime log config for job execution.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobSparkRJobLoggingConfig? LoggingConfig;
        /// <summary>
        /// Required. The HCFS URI of the main R file to use as the driver. Must be a .R file.
        /// </summary>
        public readonly string MainRFileUri;
        /// <summary>
        /// Optional. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `core:hadoop.tmp.dir`. The following are supported prefixes and their mappings: * capacity-scheduler: `capacity-scheduler.xml` * core: `core-site.xml` * distcp: `distcp-default.xml` * hdfs: `hdfs-site.xml` * hive: `hive-site.xml` * mapred: `mapred-site.xml` * pig: `pig.properties` * spark: `spark-defaults.conf` * yarn: `yarn-site.xml` For more information, see (https://cloud.google.com/dataproc/docs/concepts/cluster-properties).
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Properties;

        [OutputConstructor]
        private WorkflowTemplateJobSparkRJob(
            ImmutableArray<string> archiveUris,

            ImmutableArray<string> args,

            ImmutableArray<string> fileUris,

            Outputs.WorkflowTemplateJobSparkRJobLoggingConfig? loggingConfig,

            string mainRFileUri,

            ImmutableDictionary<string, string>? properties)
        {
            ArchiveUris = archiveUris;
            Args = args;
            FileUris = fileUris;
            LoggingConfig = loggingConfig;
            MainRFileUri = mainRFileUri;
            Properties = properties;
        }
    }
}
