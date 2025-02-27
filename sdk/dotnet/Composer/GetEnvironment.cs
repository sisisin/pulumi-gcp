// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer
{
    public static class GetEnvironment
    {
        /// <summary>
        /// Provides access to Cloud Composer environment configuration in a region for a given project.
        /// </summary>
        public static Task<GetEnvironmentResult> InvokeAsync(GetEnvironmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentResult>("gcp:composer/getEnvironment:getEnvironment", args ?? new GetEnvironmentArgs(), options.WithVersion());
    }


    public sealed class GetEnvironmentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the environment.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The location or Compute Engine region of the environment.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetEnvironmentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnvironmentResult
    {
        /// <summary>
        /// Configuration parameters for the environment.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEnvironmentConfigResult> Configs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly string? Project;
        public readonly string? Region;

        [OutputConstructor]
        private GetEnvironmentResult(
            ImmutableArray<Outputs.GetEnvironmentConfigResult> configs,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            string? project,

            string? region)
        {
            Configs = configs;
            Id = id;
            Labels = labels;
            Name = name;
            Project = project;
            Region = region;
        }
    }
}
