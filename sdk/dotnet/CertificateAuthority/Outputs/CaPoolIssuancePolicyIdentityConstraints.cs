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
    public sealed class CaPoolIssuancePolicyIdentityConstraints
    {
        /// <summary>
        /// If this is set, the SubjectAltNames extension may be copied from a certificate request into the signed certificate.
        /// Otherwise, the requested SubjectAltNames will be discarded.
        /// </summary>
        public readonly bool AllowSubjectAltNamesPassthrough;
        /// <summary>
        /// If this is set, the Subject field may be copied from a certificate request into the signed certificate.
        /// Otherwise, the requested Subject will be discarded.
        /// </summary>
        public readonly bool AllowSubjectPassthrough;
        /// <summary>
        /// A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a
        /// certificate is signed. To see the full allowed syntax and some examples,
        /// see https://cloud.google.com/certificate-authority-service/docs/cel-guide
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.CaPoolIssuancePolicyIdentityConstraintsCelExpression? CelExpression;

        [OutputConstructor]
        private CaPoolIssuancePolicyIdentityConstraints(
            bool allowSubjectAltNamesPassthrough,

            bool allowSubjectPassthrough,

            Outputs.CaPoolIssuancePolicyIdentityConstraintsCelExpression? celExpression)
        {
            AllowSubjectAltNamesPassthrough = allowSubjectAltNamesPassthrough;
            AllowSubjectPassthrough = allowSubjectPassthrough;
            CelExpression = celExpression;
        }
    }
}