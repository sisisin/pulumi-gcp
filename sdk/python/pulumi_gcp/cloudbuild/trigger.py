# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Trigger(pulumi.CustomResource):
    """
    Creates a new build trigger within GCR. For more information, see
    [the official documentation](https://cloud.google.com/container-builder/docs/running-builds/automate-builds)
    and
    [API](https://godoc.org/google.golang.org/api/cloudbuild/v1#BuildTrigger).
    """
    def __init__(__self__, __name__, __opts__=None, build=None, description=None, filename=None, project=None, substitutions=None, trigger_template=None):
        """Create a Trigger resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['build'] = build

        __props__['description'] = description

        __props__['filename'] = filename

        __props__['project'] = project

        __props__['substitutions'] = substitutions

        __props__['triggerTemplate'] = trigger_template

        super(Trigger, __self__).__init__(
            'gcp:cloudbuild/trigger:Trigger',
            __name__,
            __props__,
            __opts__)
