# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'AiDatasetEncryptionSpec',
    'AiFeatureStoreEntityTypeMonitoringConfig',
    'AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis',
    'AiFeatureStoreOnlineServingConfig',
]

@pulumi.output_type
class AiDatasetEncryptionSpec(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "kmsKeyName":
            suggest = "kms_key_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AiDatasetEncryptionSpec. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AiDatasetEncryptionSpec.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AiDatasetEncryptionSpec.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kms_key_name: Optional[str] = None):
        """
        :param str kms_key_name: Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
               Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
        """
        if kms_key_name is not None:
            pulumi.set(__self__, "kms_key_name", kms_key_name)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> Optional[str]:
        """
        Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
        Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
        """
        return pulumi.get(self, "kms_key_name")


@pulumi.output_type
class AiFeatureStoreEntityTypeMonitoringConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "snapshotAnalysis":
            suggest = "snapshot_analysis"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AiFeatureStoreEntityTypeMonitoringConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AiFeatureStoreEntityTypeMonitoringConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AiFeatureStoreEntityTypeMonitoringConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 snapshot_analysis: Optional['outputs.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis'] = None):
        """
        :param 'AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs' snapshot_analysis: Configuration of how features in Featurestore are monitored.
               Structure is documented below.
        """
        if snapshot_analysis is not None:
            pulumi.set(__self__, "snapshot_analysis", snapshot_analysis)

    @property
    @pulumi.getter(name="snapshotAnalysis")
    def snapshot_analysis(self) -> Optional['outputs.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis']:
        """
        Configuration of how features in Featurestore are monitored.
        Structure is documented below.
        """
        return pulumi.get(self, "snapshot_analysis")


@pulumi.output_type
class AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "monitoringInterval":
            suggest = "monitoring_interval"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 disabled: Optional[bool] = None,
                 monitoring_interval: Optional[str] = None):
        """
        :param bool disabled: The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
        :param str monitoring_interval: Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if monitoring_interval is not None:
            pulumi.set(__self__, "monitoring_interval", monitoring_interval)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[bool]:
        """
        The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="monitoringInterval")
    def monitoring_interval(self) -> Optional[str]:
        """
        Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "monitoring_interval")


@pulumi.output_type
class AiFeatureStoreOnlineServingConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fixedNodeCount":
            suggest = "fixed_node_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AiFeatureStoreOnlineServingConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AiFeatureStoreOnlineServingConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AiFeatureStoreOnlineServingConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 fixed_node_count: int):
        """
        :param int fixed_node_count: The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.
        """
        pulumi.set(__self__, "fixed_node_count", fixed_node_count)

    @property
    @pulumi.getter(name="fixedNodeCount")
    def fixed_node_count(self) -> int:
        """
        The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.
        """
        return pulumi.get(self, "fixed_node_count")

