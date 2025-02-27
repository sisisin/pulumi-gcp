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
    public sealed class WorkflowTemplateJob
    {
        /// <summary>
        /// Optional. Job is a Hadoop job.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobHadoopJob? HadoopJob;
        /// <summary>
        /// Optional. Job is a Hive job.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobHiveJob? HiveJob;
        /// <summary>
        /// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// Optional. Job is a Pig job.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobPigJob? PigJob;
        /// <summary>
        /// Optional. The optional list of prerequisite job step_ids. If not specified, the job will start at the beginning of workflow.
        /// </summary>
        public readonly ImmutableArray<string> PrerequisiteStepIds;
        /// <summary>
        /// Optional. Job is a Presto job.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobPrestoJob? PrestoJob;
        /// <summary>
        /// Optional. Job is a PySpark job.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobPysparkJob? PysparkJob;
        /// <summary>
        /// Optional. Job scheduling configuration.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobScheduling? Scheduling;
        /// <summary>
        /// Optional. Job is a Spark job.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobSparkJob? SparkJob;
        /// <summary>
        /// Optional. Job is a SparkR job.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobSparkRJob? SparkRJob;
        /// <summary>
        /// Optional. Job is a SparkSql job.
        /// </summary>
        public readonly Outputs.WorkflowTemplateJobSparkSqlJob? SparkSqlJob;
        /// <summary>
        /// Required. The step id. The id must be unique among all jobs within the template. The step id is used as prefix for job id, as job `goog-dataproc-workflow-step-id` label, and in field from other steps. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
        /// </summary>
        public readonly string StepId;

        [OutputConstructor]
        private WorkflowTemplateJob(
            Outputs.WorkflowTemplateJobHadoopJob? hadoopJob,

            Outputs.WorkflowTemplateJobHiveJob? hiveJob,

            ImmutableDictionary<string, string>? labels,

            Outputs.WorkflowTemplateJobPigJob? pigJob,

            ImmutableArray<string> prerequisiteStepIds,

            Outputs.WorkflowTemplateJobPrestoJob? prestoJob,

            Outputs.WorkflowTemplateJobPysparkJob? pysparkJob,

            Outputs.WorkflowTemplateJobScheduling? scheduling,

            Outputs.WorkflowTemplateJobSparkJob? sparkJob,

            Outputs.WorkflowTemplateJobSparkRJob? sparkRJob,

            Outputs.WorkflowTemplateJobSparkSqlJob? sparkSqlJob,

            string stepId)
        {
            HadoopJob = hadoopJob;
            HiveJob = hiveJob;
            Labels = labels;
            PigJob = pigJob;
            PrerequisiteStepIds = prerequisiteStepIds;
            PrestoJob = prestoJob;
            PysparkJob = pysparkJob;
            Scheduling = scheduling;
            SparkJob = sparkJob;
            SparkRJob = sparkRJob;
            SparkSqlJob = sparkSqlJob;
            StepId = stepId;
        }
    }
}
