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
    'GetSystemIpv6TunnelResult',
    'AwaitableGetSystemIpv6TunnelResult',
    'get_system_ipv6_tunnel',
    'get_system_ipv6_tunnel_output',
]

@pulumi.output_type
class GetSystemIpv6TunnelResult:
    """
    A collection of values returned by GetSystemIpv6Tunnel.
    """
    def __init__(__self__, auto_asic_offload=None, destination=None, id=None, interface=None, name=None, source=None, use_sdwan=None, vdomparam=None):
        if auto_asic_offload and not isinstance(auto_asic_offload, str):
            raise TypeError("Expected argument 'auto_asic_offload' to be a str")
        pulumi.set(__self__, "auto_asic_offload", auto_asic_offload)
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
        if use_sdwan and not isinstance(use_sdwan, str):
            raise TypeError("Expected argument 'use_sdwan' to be a str")
        pulumi.set(__self__, "use_sdwan", use_sdwan)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoAsicOffload")
    def auto_asic_offload(self) -> str:
        return pulumi.get(self, "auto_asic_offload")

    @property
    @pulumi.getter
    def destination(self) -> str:
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
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> str:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="useSdwan")
    def use_sdwan(self) -> str:
        return pulumi.get(self, "use_sdwan")

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
            auto_asic_offload=self.auto_asic_offload,
            destination=self.destination,
            id=self.id,
            interface=self.interface,
            name=self.name,
            source=self.source,
            use_sdwan=self.use_sdwan,
            vdomparam=self.vdomparam)


def get_system_ipv6_tunnel(name: Optional[str] = None,
                           vdomparam: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemIpv6TunnelResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemIpv6Tunnel:GetSystemIpv6Tunnel', __args__, opts=opts, typ=GetSystemIpv6TunnelResult).value

    return AwaitableGetSystemIpv6TunnelResult(
        auto_asic_offload=__ret__.auto_asic_offload,
        destination=__ret__.destination,
        id=__ret__.id,
        interface=__ret__.interface,
        name=__ret__.name,
        source=__ret__.source,
        use_sdwan=__ret__.use_sdwan,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_ipv6_tunnel)
def get_system_ipv6_tunnel_output(name: Optional[pulumi.Input[str]] = None,
                                  vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemIpv6TunnelResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
