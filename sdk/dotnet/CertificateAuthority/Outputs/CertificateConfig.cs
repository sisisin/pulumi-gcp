// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority.Outputs
{

    [OutputType]
    public sealed class CertificateConfig
    {
        /// <summary>
        /// A PublicKey describes a public key.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.CertificateConfigPublicKey PublicKey;
        /// <summary>
        /// A resource path to a ReusableConfig in the format
        /// `projects/*/locations/*/reusableConfigs/*`.
        /// </summary>
        public readonly Outputs.CertificateConfigReusableConfig ReusableConfig;
        /// <summary>
        /// Specifies some of the values in a certificate that are related to the subject.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.CertificateConfigSubjectConfig SubjectConfig;

        [OutputConstructor]
        private CertificateConfig(
            Outputs.CertificateConfigPublicKey publicKey,

            Outputs.CertificateConfigReusableConfig reusableConfig,

            Outputs.CertificateConfigSubjectConfig subjectConfig)
        {
            PublicKey = publicKey;
            ReusableConfig = reusableConfig;
            SubjectConfig = subjectConfig;
        }
    }
}