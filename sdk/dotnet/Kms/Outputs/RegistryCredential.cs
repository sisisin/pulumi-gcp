// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms.Outputs
{

    [OutputType]
    public sealed class RegistryCredential
    {
        public readonly ImmutableDictionary<string, object> PublicKeyCertificate;

        [OutputConstructor]
        private RegistryCredential(ImmutableDictionary<string, object> publicKeyCertificate)
        {
            PublicKeyCertificate = publicKeyCertificate;
        }
    }
}
