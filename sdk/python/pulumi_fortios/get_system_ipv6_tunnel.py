# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemIpv6TunnelResult',
    'AwaitableGetSystemIpv6TunnelResult',
    'get_system_ipv6_tunnel',
]

@pulumi.output_type
class GetSystemIpv6TunnelResult:
    """
    A collection of values returned by GetSystemIpv6Tunnel.
    """
    def __init__(__self__, destination=None, id=None, interface=None, name=None, source=None, vdomparam=None):
        if destination and not isinstance(destination, str):
            raise TypeError("Expected argument 'destination' to be a str")
        pulumi.set(__self__, "destination", destination)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def destination(self) -> str:
        """
        Remote IPv6 address of the tunnel.
        """
        return pulumi.get(self, "destination")

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
        Interface name.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        IPv6 tunnel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> str:
        """
        Local IPv6 address of the tunnel.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemIpv6TunnelResult(GetSystemIpv6TunnelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemIpv6TunnelResult(
            destination=self.destination,
            id=self.id,
            interface=self.interface,
            name=self.name,
            source=self.source,
            vdomparam=self.vdomparam)


def get_system_ipv6_tunnel(name: Optional[str] = None,
                           vdomparam: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemIpv6TunnelResult:
    """
    Use this data source to get information on an fortios system ipv6tunnel


    :param str name: Specify the name of the desired system ipv6tunnel.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemIpv6Tunnel:GetSystemIpv6Tunnel', __args__, opts=opts, typ=GetSystemIpv6TunnelResult).value

    return AwaitableGetSystemIpv6TunnelResult(
        destination=__ret__.destination,
        id=__ret__.id,
        interface=__ret__.interface,
        name=__ret__.name,
        source=__ret__.source,
        vdomparam=__ret__.vdomparam)
