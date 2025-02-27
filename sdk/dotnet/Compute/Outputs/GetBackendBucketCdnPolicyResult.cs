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
    public sealed class GetBackendBucketCdnPolicyResult
    {
        public readonly string CacheMode;
        public readonly int ClientTtl;
        public readonly int DefaultTtl;
        public readonly int MaxTtl;
        public readonly bool NegativeCaching;
        public readonly ImmutableArray<Outputs.GetBackendBucketCdnPolicyNegativeCachingPolicyResult> NegativeCachingPolicies;
        public readonly int ServeWhileStale;
        public readonly int SignedUrlCacheMaxAgeSec;

        [OutputConstructor]
        private GetBackendBucketCdnPolicyResult(
            string cacheMode,

            int clientTtl,

            int defaultTtl,

            int maxTtl,

            bool negativeCaching,

            ImmutableArray<Outputs.GetBackendBucketCdnPolicyNegativeCachingPolicyResult> negativeCachingPolicies,

            int serveWhileStale,

            int signedUrlCacheMaxAgeSec)
        {
            CacheMode = cacheMode;
            ClientTtl = clientTtl;
            DefaultTtl = defaultTtl;
            MaxTtl = maxTtl;
            NegativeCaching = negativeCaching;
            NegativeCachingPolicies = negativeCachingPolicies;
            ServeWhileStale = serveWhileStale;
            SignedUrlCacheMaxAgeSec = signedUrlCacheMaxAgeSec;
        }
    }
}
