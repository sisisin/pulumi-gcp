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
    public sealed class GetHealthCheckTcpHealthCheckResult
    {
        public readonly int Port;
        public readonly string PortName;
        public readonly string PortSpecification;
        public readonly string ProxyHeader;
        public readonly string Request;
        public readonly string Response;

        [OutputConstructor]
        private GetHealthCheckTcpHealthCheckResult(
            int port,

            string portName,

            string portSpecification,

            string proxyHeader,

            string request,

            string response)
        {
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            Request = request;
            Response = response;
        }
    }
}
