# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemFortimanagerResult',
    'AwaitableGetSystemFortimanagerResult',
    'get_system_fortimanager',
    'get_system_fortimanager_output',
]

@pulumi.output_type
class GetSystemFortimanagerResult:
    """
    A collection of values returned by GetSystemFortimanager.
    """
    def __init__(__self__, central_management=None, central_mgmt_auto_backup=None, central_mgmt_schedule_config_restore=None, central_mgmt_schedule_script_restore=None, id=None, ip=None, ipsec=None, vdom=None, vdomparam=None):
        if central_management and not isinstance(central_management, str):
            raise TypeError("Expected argument 'central_management' to be a str")
        pulumi.set(__self__, "central_management", central_management)
        if central_mgmt_auto_backup and not isinstance(central_mgmt_auto_backup, str):
            raise TypeError("Expected argument 'central_mgmt_auto_backup' to be a str")
        pulumi.set(__self__, "central_mgmt_auto_backup", central_mgmt_auto_backup)
        if central_mgmt_schedule_config_restore and not isinstance(central_mgmt_schedule_config_restore, str):
            raise TypeError("Expected argument 'central_mgmt_schedule_config_restore' to be a str")
        pulumi.set(__self__, "central_mgmt_schedule_config_restore", central_mgmt_schedule_config_restore)
        if central_mgmt_schedule_script_restore and not isinstance(central_mgmt_schedule_script_restore, str):
            raise TypeError("Expected argument 'central_mgmt_schedule_script_restore' to be a str")
        pulumi.set(__self__, "central_mgmt_schedule_script_restore", central_mgmt_schedule_script_restore)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if ipsec and not isinstance(ipsec, str):
            raise TypeError("Expected argument 'ipsec' to be a str")
        pulumi.set(__self__, "ipsec", ipsec)
        if vdom and not isinstance(vdom, str):
            raise TypeError("Expected argument 'vdom' to be a str")
        pulumi.set(__self__, "vdom", vdom)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="centralManagement")
    def central_management(self) -> str:
        return pulumi.get(self, "central_management")

    @property
    @pulumi.getter(name="centralMgmtAutoBackup")
    def central_mgmt_auto_backup(self) -> str:
        return pulumi.get(self, "central_mgmt_auto_backup")

    @property
    @pulumi.getter(name="centralMgmtScheduleConfigRestore")
    def central_mgmt_schedule_config_restore(self) -> str:
        return pulumi.get(self, "central_mgmt_schedule_config_restore")

    @property
    @pulumi.getter(name="centralMgmtScheduleScriptRestore")
    def central_mgmt_schedule_script_restore(self) -> str:
        return pulumi.get(self, "central_mgmt_schedule_script_restore")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def ipsec(self) -> str:
        return pulumi.get(self, "ipsec")

    @property
    @pulumi.getter
    def vdom(self) -> str:
        return pulumi.get(self, "vdom")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemFortimanagerResult(GetSystemFortimanagerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemFortimanagerResult(
            central_management=self.central_management,
            central_mgmt_auto_backup=self.central_mgmt_auto_backup,
            central_mgmt_schedule_config_restore=self.central_mgmt_schedule_config_restore,
            central_mgmt_schedule_script_restore=self.central_mgmt_schedule_script_restore,
            id=self.id,
            ip=self.ip,
            ipsec=self.ipsec,
            vdom=self.vdom,
            vdomparam=self.vdomparam)


def get_system_fortimanager(vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemFortimanagerResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemFortimanager:GetSystemFortimanager', __args__, opts=opts, typ=GetSystemFortimanagerResult).value

    return AwaitableGetSystemFortimanagerResult(
        central_management=__ret__.central_management,
        central_mgmt_auto_backup=__ret__.central_mgmt_auto_backup,
        central_mgmt_schedule_config_restore=__ret__.central_mgmt_schedule_config_restore,
        central_mgmt_schedule_script_restore=__ret__.central_mgmt_schedule_script_restore,
        id=__ret__.id,
        ip=__ret__.ip,
        ipsec=__ret__.ipsec,
        vdom=__ret__.vdom,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_fortimanager)
def get_system_fortimanager_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemFortimanagerResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
