// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceTemplateSpecContainerArgs : Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// Arguments to the entrypoint.
        /// The docker image's CMD is used if this is not provided.
        /// Variable references $(VAR_NAME) are expanded using the container's
        /// environment. If a variable cannot be resolved, the reference in the input
        /// string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
        /// double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
        /// regardless of whether the variable exists or not.
        /// More info:
        /// https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// Entrypoint array. Not executed within a shell.
        /// The docker image's ENTRYPOINT is used if this is not provided.
        /// Variable references $(VAR_NAME) are expanded using the container's
        /// environment. If a variable cannot be resolved, the reference in the input
        /// string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
        /// double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
        /// regardless of whether the variable exists or not.
        /// More info:
        /// https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("envFroms")]
        private InputList<Inputs.ServiceTemplateSpecContainerEnvFromArgs>? _envFroms;

        /// <summary>
        /// -
        /// (Optional, Deprecated)
        /// List of sources to populate environment variables in the container.
        /// All invalid keys will be reported as an event when the container is starting.
        /// When a key exists in multiple sources, the value associated with the last source will
        /// take precedence. Values defined by an Env with a duplicate key will take
        /// precedence.
        /// Structure is documented below.
        /// </summary>
        [Obsolete(@"Not supported by Cloud Run fully managed")]
        public InputList<Inputs.ServiceTemplateSpecContainerEnvFromArgs> EnvFroms
        {
            get => _envFroms ?? (_envFroms = new InputList<Inputs.ServiceTemplateSpecContainerEnvFromArgs>());
            set => _envFroms = value;
        }

        [Input("envs")]
        private InputList<Inputs.ServiceTemplateSpecContainerEnvArgs>? _envs;

        /// <summary>
        /// List of environment variables to set in the container.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceTemplateSpecContainerEnvArgs> Envs
        {
            get => _envs ?? (_envs = new InputList<Inputs.ServiceTemplateSpecContainerEnvArgs>());
            set => _envs = value;
        }

        /// <summary>
        /// Docker image name. This is most often a reference to a container located
        /// in the container registry, such as gcr.io/cloudrun/hello
        /// More info: https://kubernetes.io/docs/concepts/containers/images
        /// </summary>
        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        [Input("ports")]
        private InputList<Inputs.ServiceTemplateSpecContainerPortArgs>? _ports;

        /// <summary>
        /// List of open ports in the container.
        /// More Info:
        /// https://cloud.google.com/run/docs/reference/rest/v1/RevisionSpec#ContainerPort
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceTemplateSpecContainerPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ServiceTemplateSpecContainerPortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Compute Resources required by this container. Used to set values such as max memory
        /// More info:
        /// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#requests-and-limits
        /// Structure is documented below.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.ServiceTemplateSpecContainerResourcesArgs>? Resources { get; set; }

        [Input("volumeMounts")]
        private InputList<Inputs.ServiceTemplateSpecContainerVolumeMountArgs>? _volumeMounts;
        public InputList<Inputs.ServiceTemplateSpecContainerVolumeMountArgs> VolumeMounts
        {
            get => _volumeMounts ?? (_volumeMounts = new InputList<Inputs.ServiceTemplateSpecContainerVolumeMountArgs>());
            set => _volumeMounts = value;
        }

        /// <summary>
        /// -
        /// (Optional, Deprecated)
        /// Container's working directory.
        /// If not specified, the container runtime's default will be used, which
        /// might be configured in the container image.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public ServiceTemplateSpecContainerArgs()
        {
        }
    }
}
