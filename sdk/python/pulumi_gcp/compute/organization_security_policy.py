# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class OrganizationSecurityPolicy(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A textual description for the organization security policy.
    """
    display_name: pulumi.Output[str]
    """
    A textual name of the security policy.
    """
    fingerprint: pulumi.Output[str]
    """
    Fingerprint of this resource. This field is used internally during updates of this resource.
    """
    parent: pulumi.Output[str]
    """
    The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
    Format: organizations/{organization_id} or folders/{folder_id}
    """
    policy_id: pulumi.Output[str]
    """
    The unique identifier for the resource. This identifier is defined by the server.
    """
    type: pulumi.Output[str]
    """
    The type indicates the intended use of the security policy.
    For organization security policies, the only supported type
    is "FIREWALL".
    Default value is `FIREWALL`.
    Possible values are `FIREWALL`.
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, parent=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a OrganizationSecurityPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A textual description for the organization security policy.
        :param pulumi.Input[str] display_name: A textual name of the security policy.
        :param pulumi.Input[str] parent: The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
               Format: organizations/{organization_id} or folders/{folder_id}
        :param pulumi.Input[str] type: The type indicates the intended use of the security policy.
               For organization security policies, the only supported type
               is "FIREWALL".
               Default value is `FIREWALL`.
               Possible values are `FIREWALL`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['type'] = type
            __props__['fingerprint'] = None
            __props__['policy_id'] = None
        super(OrganizationSecurityPolicy, __self__).__init__(
            'gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, display_name=None, fingerprint=None, parent=None, policy_id=None, type=None):
        """
        Get an existing OrganizationSecurityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A textual description for the organization security policy.
        :param pulumi.Input[str] display_name: A textual name of the security policy.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. This field is used internally during updates of this resource.
        :param pulumi.Input[str] parent: The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
               Format: organizations/{organization_id} or folders/{folder_id}
        :param pulumi.Input[str] policy_id: The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] type: The type indicates the intended use of the security policy.
               For organization security policies, the only supported type
               is "FIREWALL".
               Default value is `FIREWALL`.
               Possible values are `FIREWALL`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["fingerprint"] = fingerprint
        __props__["parent"] = parent
        __props__["policy_id"] = policy_id
        __props__["type"] = type
        return OrganizationSecurityPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
