// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Endpoints.Outputs
{

    [OutputType]
    public sealed class ServiceApiMethod
    {
        public readonly string? Name;
        public readonly string? RequestType;
        public readonly string? ResponseType;
        public readonly string? Syntax;

        [OutputConstructor]
        private ServiceApiMethod(
            string? name,

            string? requestType,

            string? responseType,

            string? syntax)
        {
            Name = name;
            RequestType = requestType;
            ResponseType = responseType;
            Syntax = syntax;
        }
    }
}