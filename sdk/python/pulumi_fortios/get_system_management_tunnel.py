# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemManagementTunnelResult',
    'AwaitableGetSystemManagementTunnelResult',
    'get_system_management_tunnel',
]

@pulumi.output_type
class GetSystemManagementTunnelResult:
    """
    A collection of values returned by GetSystemManagementTunnel.
    """
    def __init__(__self__, allow_collect_statistics=None, allow_config_restore=None, allow_push_configuration=None, allow_push_firmware=None, authorized_manager_only=None, id=None, serial_number=None, status=None, vdomparam=None):
        if allow_collect_statistics and not isinstance(allow_collect_statistics, str):
            raise TypeError("Expected argument 'allow_collect_statistics' to be a str")
        pulumi.set(__self__, "allow_collect_statistics", allow_collect_statistics)
        if allow_config_restore and not isinstance(allow_config_restore, str):
            raise TypeError("Expected argument 'allow_config_restore' to be a str")
        pulumi.set(__self__, "allow_config_restore", allow_config_restore)
        if allow_push_configuration and not isinstance(allow_push_configuration, str):
            raise TypeError("Expected argument 'allow_push_configuration' to be a str")
        pulumi.set(__self__, "allow_push_configuration", allow_push_configuration)
        if allow_push_firmware and not isinstance(allow_push_firmware, str):
            raise TypeError("Expected argument 'allow_push_firmware' to be a str")
        pulumi.set(__self__, "allow_push_firmware", allow_push_firmware)
        if authorized_manager_only and not isinstance(authorized_manager_only, str):
            raise TypeError("Expected argument 'authorized_manager_only' to be a str")
        pulumi.set(__self__, "authorized_manager_only", authorized_manager_only)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if serial_number and not isinstance(serial_number, str):
            raise TypeError("Expected argument 'serial_number' to be a str")
        pulumi.set(__self__, "serial_number", serial_number)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="allowCollectStatistics")
    def allow_collect_statistics(self) -> str:
        """
        Enable/disable collection of run time statistics.
        """
        return pulumi.get(self, "allow_collect_statistics")

    @property
    @pulumi.getter(name="allowConfigRestore")
    def allow_config_restore(self) -> str:
        """
        Enable/disable allow config restore.
        """
        return pulumi.get(self, "allow_config_restore")

    @property
    @pulumi.getter(name="allowPushConfiguration")
    def allow_push_configuration(self) -> str:
        """
        Enable/disable push configuration.
        """
        return pulumi.get(self, "allow_push_configuration")

    @property
    @pulumi.getter(name="allowPushFirmware")
    def allow_push_firmware(self) -> str:
        """
        Enable/disable push firmware.
        """
        return pulumi.get(self, "allow_push_firmware")

    @property
    @pulumi.getter(name="authorizedManagerOnly")
    def authorized_manager_only(self) -> str:
        """
        Enable/disable restriction of authorized manager only.
        """
        return pulumi.get(self, "authorized_manager_only")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> str:
        """
        Serial number.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable FGFM tunnel.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemManagementTunnelResult(GetSystemManagementTunnelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemManagementTunnelResult(
            allow_collect_statistics=self.allow_collect_statistics,
            allow_config_restore=self.allow_config_restore,
            allow_push_configuration=self.allow_push_configuration,
            allow_push_firmware=self.allow_push_firmware,
            authorized_manager_only=self.authorized_manager_only,
            id=self.id,
            serial_number=self.serial_number,
            status=self.status,
            vdomparam=self.vdomparam)


def get_system_management_tunnel(vdomparam: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemManagementTunnelResult:
    """
    Use this data source to get information on fortios system managementtunnel


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemManagementTunnel:GetSystemManagementTunnel', __args__, opts=opts, typ=GetSystemManagementTunnelResult).value

    return AwaitableGetSystemManagementTunnelResult(
        allow_collect_statistics=__ret__.allow_collect_statistics,
        allow_config_restore=__ret__.allow_config_restore,
        allow_push_configuration=__ret__.allow_push_configuration,
        allow_push_firmware=__ret__.allow_push_firmware,
        authorized_manager_only=__ret__.authorized_manager_only,
        id=__ret__.id,
        serial_number=__ret__.serial_number,
        status=__ret__.status,
        vdomparam=__ret__.vdomparam)
