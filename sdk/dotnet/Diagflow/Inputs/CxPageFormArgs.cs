// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Diagflow.Inputs
{

    public sealed class CxPageFormArgs : Pulumi.ResourceArgs
    {
        [Input("parameters")]
        private InputList<Inputs.CxPageFormParameterArgs>? _parameters;

        /// <summary>
        /// Parameters to collect from the user.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.CxPageFormParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.CxPageFormParameterArgs>());
            set => _parameters = value;
        }

        public CxPageFormArgs()
        {
        }
    }
}