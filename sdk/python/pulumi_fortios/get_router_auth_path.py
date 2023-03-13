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
    'GetRouterAuthPathResult',
    'AwaitableGetRouterAuthPathResult',
    'get_router_auth_path',
    'get_router_auth_path_output',
]

@pulumi.output_type
class GetRouterAuthPathResult:
    """
    A collection of values returned by GetRouterAuthPath.
    """
    def __init__(__self__, device=None, gateway=None, id=None, name=None, vdomparam=None):
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if gateway and not isinstance(gateway, str):
            raise TypeError("Expected argument 'gateway' to be a str")
        pulumi.set(__self__, "gateway", gateway)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def device(self) -> str:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def gateway(self) -> str:
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterAuthPathResult(GetRouterAuthPathResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterAuthPathResult(
            device=self.device,
            gateway=self.gateway,
            id=self.id,
            name=self.name,
            vdomparam=self.vdomparam)


def get_router_auth_path(name: Optional[str] = None,
                         vdomparam: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterAuthPathResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterAuthPath:GetRouterAuthPath', __args__, opts=opts, typ=GetRouterAuthPathResult).value

    return AwaitableGetRouterAuthPathResult(
        device=__ret__.device,
        gateway=__ret__.gateway,
        id=__ret__.id,
        name=__ret__.name,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_router_auth_path)
def get_router_auth_path_output(name: Optional[pulumi.Input[str]] = None,
                                vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterAuthPathResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
