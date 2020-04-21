// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.IdentityPlatform.Inputs
{

    public sealed class InboundSamlConfigSpConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("callbackUri")]
        public Input<string>? CallbackUri { get; set; }

        [Input("spCertificates")]
        private InputList<Inputs.InboundSamlConfigSpConfigSpCertificateGetArgs>? _spCertificates;
        public InputList<Inputs.InboundSamlConfigSpConfigSpCertificateGetArgs> SpCertificates
        {
            get => _spCertificates ?? (_spCertificates = new InputList<Inputs.InboundSamlConfigSpConfigSpCertificateGetArgs>());
            set => _spCertificates = value;
        }

        [Input("spEntityId")]
        public Input<string>? SpEntityId { get; set; }

        public InboundSamlConfigSpConfigGetArgs()
        {
        }
    }
}