# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['AccessApprovalSettings']


class AccessApprovalSettings(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enrolled_services: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AccessApprovalSettingsEnrolledServiceArgs']]]]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 notification_emails: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a AccessApprovalSettings resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AccessApprovalSettingsEnrolledServiceArgs']]]] enrolled_services: A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
               resource given by name against any of these services contained here will be required to have explicit approval.
               Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
               expanded as the set of supported services is expanded.
        :param pulumi.Input[str] folder_id: ID of the folder of the access approval settings.
        :param pulumi.Input[List[pulumi.Input[str]]] notification_emails: A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
               a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
               addresses are allowed.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if enrolled_services is None:
                raise TypeError("Missing required property 'enrolled_services'")
            __props__['enrolled_services'] = enrolled_services
            if folder_id is None:
                raise TypeError("Missing required property 'folder_id'")
            __props__['folder_id'] = folder_id
            __props__['notification_emails'] = notification_emails
            __props__['enrolled_ancestor'] = None
            __props__['name'] = None
        super(AccessApprovalSettings, __self__).__init__(
            'gcp:folder/accessApprovalSettings:AccessApprovalSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enrolled_ancestor: Optional[pulumi.Input[bool]] = None,
            enrolled_services: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AccessApprovalSettingsEnrolledServiceArgs']]]]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_emails: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'AccessApprovalSettings':
        """
        Get an existing AccessApprovalSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enrolled_ancestor: If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
               of the Folder.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AccessApprovalSettingsEnrolledServiceArgs']]]] enrolled_services: A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
               resource given by name against any of these services contained here will be required to have explicit approval.
               Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
               expanded as the set of supported services is expanded.
        :param pulumi.Input[str] folder_id: ID of the folder of the access approval settings.
        :param pulumi.Input[str] name: The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
        :param pulumi.Input[List[pulumi.Input[str]]] notification_emails: A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
               a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
               addresses are allowed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["enrolled_ancestor"] = enrolled_ancestor
        __props__["enrolled_services"] = enrolled_services
        __props__["folder_id"] = folder_id
        __props__["name"] = name
        __props__["notification_emails"] = notification_emails
        return AccessApprovalSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enrolledAncestor")
    def enrolled_ancestor(self) -> pulumi.Output[bool]:
        """
        If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
        of the Folder.
        """
        return pulumi.get(self, "enrolled_ancestor")

    @property
    @pulumi.getter(name="enrolledServices")
    def enrolled_services(self) -> pulumi.Output[List['outputs.AccessApprovalSettingsEnrolledService']]:
        """
        A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
        resource given by name against any of these services contained here will be required to have explicit approval.
        Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
        expanded as the set of supported services is expanded.
        """
        return pulumi.get(self, "enrolled_services")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        ID of the folder of the access approval settings.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationEmails")
    def notification_emails(self) -> pulumi.Output[List[str]]:
        """
        A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
        a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
        addresses are allowed.
        """
        return pulumi.get(self, "notification_emails")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
