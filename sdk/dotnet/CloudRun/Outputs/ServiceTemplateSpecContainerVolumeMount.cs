// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Outputs
{

    [OutputType]
    public sealed class ServiceTemplateSpecContainerVolumeMount
    {
        /// <summary>
        /// Path within the container at which the volume should be mounted.  Must
        /// not contain ':'.
        /// </summary>
        public readonly string MountPath;
        /// <summary>
        /// Volume's name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ServiceTemplateSpecContainerVolumeMount(
            string mountPath,

            string name)
        {
            MountPath = mountPath;
            Name = name;
        }
    }
}
