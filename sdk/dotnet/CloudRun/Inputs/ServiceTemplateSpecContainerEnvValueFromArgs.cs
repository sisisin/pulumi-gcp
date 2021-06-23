// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceTemplateSpecContainerEnvValueFromArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Selects a key (version) of a secret in Secret Manager.
        /// Structure is documented below.
        /// </summary>
        [Input("secretKeyRef", required: true)]
        public Input<Inputs.ServiceTemplateSpecContainerEnvValueFromSecretKeyRefArgs> SecretKeyRef { get; set; } = null!;

        public ServiceTemplateSpecContainerEnvValueFromArgs()
        {
        }
    }
}