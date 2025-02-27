// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Outputs
{

    [OutputType]
    public sealed class EnvironmentConfigEncryptionConfig
    {
        /// <summary>
        /// Customer-managed Encryption Key available through Google's Key Management Service. It must
        /// be the fully qualified resource name,
        /// i.e. projects/project-id/locations/location/keyRings/keyring/cryptoKeys/key. Cannot be updated.
        /// </summary>
        public readonly string KmsKeyName;

        [OutputConstructor]
        private EnvironmentConfigEncryptionConfig(string kmsKeyName)
        {
            KmsKeyName = kmsKeyName;
        }
    }
}
