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
    'GetSystemDnsServerResult',
    'AwaitableGetSystemDnsServerResult',
    'get_system_dns_server',
    'get_system_dns_server_output',
]

@pulumi.output_type
class GetSystemDnsServerResult:
    """
    A collection of values returned by GetSystemDnsServer.
    """
    def __init__(__self__, dnsfilter_profile=None, doh=None, id=None, mode=None, name=None, vdomparam=None):
        if dnsfilter_profile and not isinstance(dnsfilter_profile, str):
            raise TypeError("Expected argument 'dnsfilter_profile' to be a str")
        pulumi.set(__self__, "dnsfilter_profile", dnsfilter_profile)
        if doh and not isinstance(doh, str):
            raise TypeError("Expected argument 'doh' to be a str")
        pulumi.set(__self__, "doh", doh)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dnsfilterProfile")
    def dnsfilter_profile(self) -> str:
        return pulumi.get(self, "dnsfilter_profile")

    @property
    @pulumi.getter
    def doh(self) -> str:
        return pulumi.get(self, "doh")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def mode(self) -> str:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemDnsServerResult(GetSystemDnsServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemDnsServerResult(
            dnsfilter_profile=self.dnsfilter_profile,
            doh=self.doh,
            id=self.id,
            mode=self.mode,
            name=self.name,
            vdomparam=self.vdomparam)


def get_system_dns_server(name: Optional[str] = None,
                          vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemDnsServerResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemDnsServer:GetSystemDnsServer', __args__, opts=opts, typ=GetSystemDnsServerResult).value

    return AwaitableGetSystemDnsServerResult(
        dnsfilter_profile=pulumi.get(__ret__, 'dnsfilter_profile'),
        doh=pulumi.get(__ret__, 'doh'),
        id=pulumi.get(__ret__, 'id'),
        mode=pulumi.get(__ret__, 'mode'),
        name=pulumi.get(__ret__, 'name'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_system_dns_server)
def get_system_dns_server_output(name: Optional[pulumi.Input[str]] = None,
                                 vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemDnsServerResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
