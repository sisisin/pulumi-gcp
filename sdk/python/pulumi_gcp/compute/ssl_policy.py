# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class SSLPolicy(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, custom_features=None, description=None, min_tls_version=None, name=None, profile=None, project=None):
        """Create a SSLPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['customFeatures'] = custom_features

        __props__['description'] = description

        __props__['minTlsVersion'] = min_tls_version

        __props__['name'] = name

        __props__['profile'] = profile

        __props__['project'] = project

        __props__['creation_timestamp'] = None
        __props__['enabled_features'] = None
        __props__['fingerprint'] = None
        __props__['self_link'] = None

        super(SSLPolicy, __self__).__init__(
            'gcp:compute/sSLPolicy:SSLPolicy',
            __name__,
            __props__,
            __opts__)
