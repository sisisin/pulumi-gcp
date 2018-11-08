# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class InstanceFromTemplate(pulumi.CustomResource):
    """
    Manages a VM instance resource within GCE. For more information see
    [the official documentation](https://cloud.google.com/compute/docs/instances)
    and
    [API](https://cloud.google.com/compute/docs/reference/latest/instances).
    
    This resource is specifically to create a compute instance from a given
    `source_instance_template`. To create an instance without a template, use the
    `google_compute_instance` resource.
    
    """
    def __init__(__self__, __name__, __opts__=None, allow_stopping_for_update=None, attached_disks=None, boot_disk=None, can_ip_forward=None, deletion_protection=None, description=None, guest_accelerators=None, labels=None, machine_type=None, metadata=None, metadata_startup_script=None, min_cpu_platform=None, name=None, network_interfaces=None, project=None, scheduling=None, scratch_disks=None, service_account=None, source_instance_template=None, tags=None, zone=None):
        """Create a InstanceFromTemplate resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allowStoppingForUpdate'] = allow_stopping_for_update

        __props__['attachedDisks'] = attached_disks

        __props__['bootDisk'] = boot_disk

        __props__['canIpForward'] = can_ip_forward

        __props__['deletionProtection'] = deletion_protection

        __props__['description'] = description

        __props__['guestAccelerators'] = guest_accelerators

        __props__['labels'] = labels

        __props__['machineType'] = machine_type

        __props__['metadata'] = metadata

        __props__['metadataStartupScript'] = metadata_startup_script

        __props__['minCpuPlatform'] = min_cpu_platform

        __props__['name'] = name

        __props__['networkInterfaces'] = network_interfaces

        __props__['project'] = project

        __props__['scheduling'] = scheduling

        __props__['scratchDisks'] = scratch_disks

        __props__['serviceAccount'] = service_account

        if not source_instance_template:
            raise TypeError('Missing required property source_instance_template')
        __props__['sourceInstanceTemplate'] = source_instance_template

        __props__['tags'] = tags

        __props__['zone'] = zone

        __props__['cpu_platform'] = None
        __props__['instance_id'] = None
        __props__['label_fingerprint'] = None
        __props__['metadata_fingerprint'] = None
        __props__['self_link'] = None
        __props__['tags_fingerprint'] = None

        super(InstanceFromTemplate, __self__).__init__(
            'gcp:compute/instanceFromTemplate:InstanceFromTemplate',
            __name__,
            __props__,
            __opts__)
