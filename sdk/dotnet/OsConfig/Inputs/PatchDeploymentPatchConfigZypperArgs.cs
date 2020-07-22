// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Inputs
{

    public sealed class PatchDeploymentPatchConfigZypperArgs : Pulumi.ResourceArgs
    {
        [Input("categories")]
        private InputList<string>? _categories;

        /// <summary>
        /// Install only patches with these categories. Common categories include security, recommended, and feature.
        /// </summary>
        public InputList<string> Categories
        {
            get => _categories ?? (_categories = new InputList<string>());
            set => _categories = value;
        }

        [Input("excludes")]
        private InputList<string>? _excludes;

        /// <summary>
        /// List of KBs to exclude from update.
        /// </summary>
        public InputList<string> Excludes
        {
            get => _excludes ?? (_excludes = new InputList<string>());
            set => _excludes = value;
        }

        [Input("exclusivePatches")]
        private InputList<string>? _exclusivePatches;

        /// <summary>
        /// An exclusive list of kbs to be updated. These are the only patches that will be updated.
        /// This field must not be used with other patch configurations.
        /// </summary>
        public InputList<string> ExclusivePatches
        {
            get => _exclusivePatches ?? (_exclusivePatches = new InputList<string>());
            set => _exclusivePatches = value;
        }

        [Input("severities")]
        private InputList<string>? _severities;

        /// <summary>
        /// Install only patches with these severities. Common severities include critical, important, moderate, and low.
        /// </summary>
        public InputList<string> Severities
        {
            get => _severities ?? (_severities = new InputList<string>());
            set => _severities = value;
        }

        /// <summary>
        /// Adds the --with-optional flag to zypper patch.
        /// </summary>
        [Input("withOptional")]
        public Input<bool>? WithOptional { get; set; }

        /// <summary>
        /// Adds the --with-update flag, to zypper patch.
        /// </summary>
        [Input("withUpdate")]
        public Input<bool>? WithUpdate { get; set; }

        public PatchDeploymentPatchConfigZypperArgs()
        {
        }
    }
}