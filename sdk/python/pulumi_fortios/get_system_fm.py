# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemFmResult',
    'AwaitableGetSystemFmResult',
    'get_system_fm',
]

@pulumi.output_type
class GetSystemFmResult:
    """
    A collection of values returned by GetSystemFm.
    """
    def __init__(__self__, auto_backup=None, fosid=None, id=None, ip=None, ipsec=None, scheduled_config_restore=None, status=None, vdom=None, vdomparam=None):
        if auto_backup and not isinstance(auto_backup, str):
            raise TypeError("Expected argument 'auto_backup' to be a str")
        pulumi.set(__self__, "auto_backup", auto_backup)
        if fosid and not isinstance(fosid, str):
            raise TypeError("Expected argument 'fosid' to be a str")
        pulumi.set(__self__, "fosid", fosid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if ipsec and not isinstance(ipsec, str):
            raise TypeError("Expected argument 'ipsec' to be a str")
        pulumi.set(__self__, "ipsec", ipsec)
        if scheduled_config_restore and not isinstance(scheduled_config_restore, str):
            raise TypeError("Expected argument 'scheduled_config_restore' to be a str")
        pulumi.set(__self__, "scheduled_config_restore", scheduled_config_restore)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdom and not isinstance(vdom, str):
            raise TypeError("Expected argument 'vdom' to be a str")
        pulumi.set(__self__, "vdom", vdom)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoBackup")
    def auto_backup(self) -> str:
        """
        Enable/disable automatic backup.
        """
        return pulumi.get(self, "auto_backup")

    @property
    @pulumi.getter
    def fosid(self) -> str:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

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
        """
        IP address.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def ipsec(self) -> str:
        """
        Enable/disable IPsec.
        """
        return pulumi.get(self, "ipsec")

    @property
    @pulumi.getter(name="scheduledConfigRestore")
    def scheduled_config_restore(self) -> str:
        """
        Enable/disable scheduled configuration restore.
        """
        return pulumi.get(self, "scheduled_config_restore")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable FM.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdom(self) -> str:
        """
        VDOM.
        """
        return pulumi.get(self, "vdom")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemFmResult(GetSystemFmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemFmResult(
            auto_backup=self.auto_backup,
            fosid=self.fosid,
            id=self.id,
            ip=self.ip,
            ipsec=self.ipsec,
            scheduled_config_restore=self.scheduled_config_restore,
            status=self.status,
            vdom=self.vdom,
            vdomparam=self.vdomparam)


def get_system_fm(vdomparam: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemFmResult:
    """
    Use this data source to get information on fortios system fm


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemFm:GetSystemFm', __args__, opts=opts, typ=GetSystemFmResult).value

    return AwaitableGetSystemFmResult(
        auto_backup=__ret__.auto_backup,
        fosid=__ret__.fosid,
        id=__ret__.id,
        ip=__ret__.ip,
        ipsec=__ret__.ipsec,
        scheduled_config_restore=__ret__.scheduled_config_restore,
        status=__ret__.status,
        vdom=__ret__.vdom,
        vdomparam=__ret__.vdomparam)
