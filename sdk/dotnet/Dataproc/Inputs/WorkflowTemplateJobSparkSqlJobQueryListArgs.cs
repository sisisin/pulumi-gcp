// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class WorkflowTemplateJobSparkSqlJobQueryListArgs : Pulumi.ResourceArgs
    {
        [Input("queries", required: true)]
        private InputList<string>? _queries;

        /// <summary>
        /// Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: "hiveJob": { "queryList": { "queries": } }
        /// </summary>
        public InputList<string> Queries
        {
            get => _queries ?? (_queries = new InputList<string>());
            set => _queries = value;
        }

        public WorkflowTemplateJobSparkSqlJobQueryListArgs()
        {
        }
    }
}