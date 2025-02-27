# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DicomStoreArgs', 'DicomStore']

@pulumi.input_type
class DicomStoreArgs:
    def __init__(__self__, *,
                 dataset: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input['DicomStoreNotificationConfigArgs']] = None,
                 stream_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DicomStoreStreamConfigArgs']]]] = None):
        """
        The set of arguments for constructing a DicomStore resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize DICOM stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the DicomStore.
               ** Changing this property may recreate the Dicom store (removing all data) **
        :param pulumi.Input['DicomStoreNotificationConfigArgs'] notification_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['DicomStoreStreamConfigArgs']]] stream_configs: To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store. streamConfigs is an array, so
               you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery
               tables in a BigQuery dataset.
        """
        pulumi.set(__self__, "dataset", dataset)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_config is not None:
            pulumi.set(__self__, "notification_config", notification_config)
        if stream_configs is not None:
            pulumi.set(__self__, "stream_configs", stream_configs)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Input[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-supplied key-value pairs used to organize DICOM stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the DicomStore.
        ** Changing this property may recreate the Dicom store (removing all data) **
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> Optional[pulumi.Input['DicomStoreNotificationConfigArgs']]:
        """
        A nested object resource
        Structure is documented below.
        """
        return pulumi.get(self, "notification_config")

    @notification_config.setter
    def notification_config(self, value: Optional[pulumi.Input['DicomStoreNotificationConfigArgs']]):
        pulumi.set(self, "notification_config", value)

    @property
    @pulumi.getter(name="streamConfigs")
    def stream_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DicomStoreStreamConfigArgs']]]]:
        """
        To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store. streamConfigs is an array, so
        you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery
        tables in a BigQuery dataset.
        """
        return pulumi.get(self, "stream_configs")

    @stream_configs.setter
    def stream_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DicomStoreStreamConfigArgs']]]]):
        pulumi.set(self, "stream_configs", value)


@pulumi.input_type
class _DicomStoreState:
    def __init__(__self__, *,
                 dataset: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input['DicomStoreNotificationConfigArgs']] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 stream_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DicomStoreStreamConfigArgs']]]] = None):
        """
        Input properties used for looking up and filtering DicomStore resources.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize DICOM stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the DicomStore.
               ** Changing this property may recreate the Dicom store (removing all data) **
        :param pulumi.Input['DicomStoreNotificationConfigArgs'] notification_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset
        :param pulumi.Input[Sequence[pulumi.Input['DicomStoreStreamConfigArgs']]] stream_configs: To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store. streamConfigs is an array, so
               you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery
               tables in a BigQuery dataset.
        """
        if dataset is not None:
            pulumi.set(__self__, "dataset", dataset)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_config is not None:
            pulumi.set(__self__, "notification_config", notification_config)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if stream_configs is not None:
            pulumi.set(__self__, "stream_configs", stream_configs)

    @property
    @pulumi.getter
    def dataset(self) -> Optional[pulumi.Input[str]]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-supplied key-value pairs used to organize DICOM stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the DicomStore.
        ** Changing this property may recreate the Dicom store (removing all data) **
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> Optional[pulumi.Input['DicomStoreNotificationConfigArgs']]:
        """
        A nested object resource
        Structure is documented below.
        """
        return pulumi.get(self, "notification_config")

    @notification_config.setter
    def notification_config(self, value: Optional[pulumi.Input['DicomStoreNotificationConfigArgs']]):
        pulumi.set(self, "notification_config", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified name of this dataset
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="streamConfigs")
    def stream_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DicomStoreStreamConfigArgs']]]]:
        """
        To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store. streamConfigs is an array, so
        you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery
        tables in a BigQuery dataset.
        """
        return pulumi.get(self, "stream_configs")

    @stream_configs.setter
    def stream_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DicomStoreStreamConfigArgs']]]]):
        pulumi.set(self, "stream_configs", value)


class DicomStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']]] = None,
                 stream_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DicomStoreStreamConfigArgs']]]]] = None,
                 __props__=None):
        """
        A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
        (https://www.dicomstandard.org/about/) standard for Healthcare information exchange

        To get more information about DicomStore, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.dicomStores)
        * How-to Guides
            * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)

        ## Example Usage
        ### Healthcare Dicom Store Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        topic = gcp.pubsub.Topic("topic")
        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        default = gcp.healthcare.DicomStore("default",
            dataset=dataset.id,
            notification_config=gcp.healthcare.DicomStoreNotificationConfigArgs(
                pubsub_topic=topic.id,
            ),
            labels={
                "label1": "labelvalue1",
            })
        ```
        ### Healthcare Dicom Store Bq Stream

        ```python
        import pulumi
        import pulumi_gcp as gcp

        topic = gcp.pubsub.Topic("topic", opts=pulumi.ResourceOptions(provider=google_beta))
        dataset = gcp.healthcare.Dataset("dataset", location="us-central1",
        opts=pulumi.ResourceOptions(provider=google_beta))
        bq_dataset = gcp.bigquery.Dataset("bqDataset",
            dataset_id="dicom_bq_ds",
            friendly_name="test",
            description="This is a test description",
            location="US",
            delete_contents_on_destroy=True,
            opts=pulumi.ResourceOptions(provider=google_beta))
        bq_table = gcp.bigquery.Table("bqTable",
            deletion_protection=False,
            dataset_id=bq_dataset.dataset_id,
            table_id="dicom_bq_tb",
            opts=pulumi.ResourceOptions(provider=google_beta))
        default = gcp.healthcare.DicomStore("default",
            dataset=dataset.id,
            notification_config=gcp.healthcare.DicomStoreNotificationConfigArgs(
                pubsub_topic=topic.id,
            ),
            labels={
                "label1": "labelvalue1",
            },
            stream_configs=[gcp.healthcare.DicomStoreStreamConfigArgs(
                bigquery_destination=gcp.healthcare.DicomStoreStreamConfigBigqueryDestinationArgs(
                    table_uri=pulumi.Output.all(bq_dataset.project, bq_dataset.dataset_id, bq_table.table_id).apply(lambda project, dataset_id, table_id: f"bq://{project}.{dataset_id}.{table_id}"),
                ),
            )],
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        DicomStore can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:healthcare/dicomStore:DicomStore default {{dataset}}/dicomStores/{{name}}
        ```

        ```sh
         $ pulumi import gcp:healthcare/dicomStore:DicomStore default {{dataset}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize DICOM stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the DicomStore.
               ** Changing this property may recreate the Dicom store (removing all data) **
        :param pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']] notification_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DicomStoreStreamConfigArgs']]]] stream_configs: To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store. streamConfigs is an array, so
               you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery
               tables in a BigQuery dataset.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DicomStoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
        (https://www.dicomstandard.org/about/) standard for Healthcare information exchange

        To get more information about DicomStore, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.dicomStores)
        * How-to Guides
            * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)

        ## Example Usage
        ### Healthcare Dicom Store Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        topic = gcp.pubsub.Topic("topic")
        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        default = gcp.healthcare.DicomStore("default",
            dataset=dataset.id,
            notification_config=gcp.healthcare.DicomStoreNotificationConfigArgs(
                pubsub_topic=topic.id,
            ),
            labels={
                "label1": "labelvalue1",
            })
        ```
        ### Healthcare Dicom Store Bq Stream

        ```python
        import pulumi
        import pulumi_gcp as gcp

        topic = gcp.pubsub.Topic("topic", opts=pulumi.ResourceOptions(provider=google_beta))
        dataset = gcp.healthcare.Dataset("dataset", location="us-central1",
        opts=pulumi.ResourceOptions(provider=google_beta))
        bq_dataset = gcp.bigquery.Dataset("bqDataset",
            dataset_id="dicom_bq_ds",
            friendly_name="test",
            description="This is a test description",
            location="US",
            delete_contents_on_destroy=True,
            opts=pulumi.ResourceOptions(provider=google_beta))
        bq_table = gcp.bigquery.Table("bqTable",
            deletion_protection=False,
            dataset_id=bq_dataset.dataset_id,
            table_id="dicom_bq_tb",
            opts=pulumi.ResourceOptions(provider=google_beta))
        default = gcp.healthcare.DicomStore("default",
            dataset=dataset.id,
            notification_config=gcp.healthcare.DicomStoreNotificationConfigArgs(
                pubsub_topic=topic.id,
            ),
            labels={
                "label1": "labelvalue1",
            },
            stream_configs=[gcp.healthcare.DicomStoreStreamConfigArgs(
                bigquery_destination=gcp.healthcare.DicomStoreStreamConfigBigqueryDestinationArgs(
                    table_uri=pulumi.Output.all(bq_dataset.project, bq_dataset.dataset_id, bq_table.table_id).apply(lambda project, dataset_id, table_id: f"bq://{project}.{dataset_id}.{table_id}"),
                ),
            )],
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        DicomStore can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:healthcare/dicomStore:DicomStore default {{dataset}}/dicomStores/{{name}}
        ```

        ```sh
         $ pulumi import gcp:healthcare/dicomStore:DicomStore default {{dataset}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param DicomStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DicomStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']]] = None,
                 stream_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DicomStoreStreamConfigArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DicomStoreArgs.__new__(DicomStoreArgs)

            if dataset is None and not opts.urn:
                raise TypeError("Missing required property 'dataset'")
            __props__.__dict__["dataset"] = dataset
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_config"] = notification_config
            __props__.__dict__["stream_configs"] = stream_configs
            __props__.__dict__["self_link"] = None
        super(DicomStore, __self__).__init__(
            'gcp:healthcare/dicomStore:DicomStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dataset: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_config: Optional[pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            stream_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DicomStoreStreamConfigArgs']]]]] = None) -> 'DicomStore':
        """
        Get an existing DicomStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize DICOM stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the DicomStore.
               ** Changing this property may recreate the Dicom store (removing all data) **
        :param pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']] notification_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DicomStoreStreamConfigArgs']]]] stream_configs: To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store. streamConfigs is an array, so
               you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery
               tables in a BigQuery dataset.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DicomStoreState.__new__(_DicomStoreState)

        __props__.__dict__["dataset"] = dataset
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["notification_config"] = notification_config
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["stream_configs"] = stream_configs
        return DicomStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Output[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        """
        return pulumi.get(self, "dataset")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        User-supplied key-value pairs used to organize DICOM stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the DicomStore.
        ** Changing this property may recreate the Dicom store (removing all data) **
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> pulumi.Output[Optional['outputs.DicomStoreNotificationConfig']]:
        """
        A nested object resource
        Structure is documented below.
        """
        return pulumi.get(self, "notification_config")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The fully qualified name of this dataset
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="streamConfigs")
    def stream_configs(self) -> pulumi.Output[Optional[Sequence['outputs.DicomStoreStreamConfig']]]:
        """
        To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store. streamConfigs is an array, so
        you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery
        tables in a BigQuery dataset.
        """
        return pulumi.get(self, "stream_configs")

