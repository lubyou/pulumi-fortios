# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemProxyArpResult',
    'AwaitableGetSystemProxyArpResult',
    'get_system_proxy_arp',
]

@pulumi.output_type
class GetSystemProxyArpResult:
    """
    A collection of values returned by GetSystemProxyArp.
    """
    def __init__(__self__, end_ip=None, fosid=None, id=None, interface=None, ip=None, vdomparam=None):
        if end_ip and not isinstance(end_ip, str):
            raise TypeError("Expected argument 'end_ip' to be a str")
        pulumi.set(__self__, "end_ip", end_ip)
        if fosid and not isinstance(fosid, int):
            raise TypeError("Expected argument 'fosid' to be a int")
        pulumi.set(__self__, "fosid", fosid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> str:
        """
        End IP of IP range to be proxied.
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def fosid(self) -> int:
        """
        Unique integer ID of the entry.
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
    def interface(self) -> str:
        """
        Interface acting proxy-ARP.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        IP address or start IP to be proxied.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemProxyArpResult(GetSystemProxyArpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemProxyArpResult(
            end_ip=self.end_ip,
            fosid=self.fosid,
            id=self.id,
            interface=self.interface,
            ip=self.ip,
            vdomparam=self.vdomparam)


def get_system_proxy_arp(fosid: Optional[int] = None,
                         vdomparam: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemProxyArpResult:
    """
    Use this data source to get information on an fortios system proxyarp


    :param int fosid: Specify the fosid of the desired system proxyarp.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['fosid'] = fosid
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemProxyArp:GetSystemProxyArp', __args__, opts=opts, typ=GetSystemProxyArpResult).value

    return AwaitableGetSystemProxyArpResult(
        end_ip=__ret__.end_ip,
        fosid=__ret__.fosid,
        id=__ret__.id,
        interface=__ret__.interface,
        ip=__ret__.ip,
        vdomparam=__ret__.vdomparam)