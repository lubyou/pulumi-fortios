# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetRouterMulticastResult',
    'AwaitableGetRouterMulticastResult',
    'get_router_multicast',
]

@pulumi.output_type
class GetRouterMulticastResult:
    """
    A collection of values returned by GetRouterMulticast.
    """
    def __init__(__self__, id=None, interfaces=None, multicast_routing=None, pim_sm_global=None, route_limit=None, route_threshold=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if multicast_routing and not isinstance(multicast_routing, str):
            raise TypeError("Expected argument 'multicast_routing' to be a str")
        pulumi.set(__self__, "multicast_routing", multicast_routing)
        if pim_sm_global and not isinstance(pim_sm_global, dict):
            raise TypeError("Expected argument 'pim_sm_global' to be a dict")
        pulumi.set(__self__, "pim_sm_global", pim_sm_global)
        if route_limit and not isinstance(route_limit, int):
            raise TypeError("Expected argument 'route_limit' to be a int")
        pulumi.set(__self__, "route_limit", route_limit)
        if route_threshold and not isinstance(route_threshold, int):
            raise TypeError("Expected argument 'route_threshold' to be a int")
        pulumi.set(__self__, "route_threshold", route_threshold)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetRouterMulticastInterfaceResult']:
        """
        PIM interfaces. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> str:
        """
        Enable/disable IP multicast routing.
        """
        return pulumi.get(self, "multicast_routing")

    @property
    @pulumi.getter(name="pimSmGlobal")
    def pim_sm_global(self) -> 'outputs.GetRouterMulticastPimSmGlobalResult':
        """
        PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        """
        return pulumi.get(self, "pim_sm_global")

    @property
    @pulumi.getter(name="routeLimit")
    def route_limit(self) -> int:
        """
        Maximum number of multicast routes.
        """
        return pulumi.get(self, "route_limit")

    @property
    @pulumi.getter(name="routeThreshold")
    def route_threshold(self) -> int:
        """
        Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
        """
        return pulumi.get(self, "route_threshold")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterMulticastResult(GetRouterMulticastResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterMulticastResult(
            id=self.id,
            interfaces=self.interfaces,
            multicast_routing=self.multicast_routing,
            pim_sm_global=self.pim_sm_global,
            route_limit=self.route_limit,
            route_threshold=self.route_threshold,
            vdomparam=self.vdomparam)


def get_router_multicast(vdomparam: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterMulticastResult:
    """
    Use this data source to get information on fortios router multicast


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterMulticast:GetRouterMulticast', __args__, opts=opts, typ=GetRouterMulticastResult).value

    return AwaitableGetRouterMulticastResult(
        id=__ret__.id,
        interfaces=__ret__.interfaces,
        multicast_routing=__ret__.multicast_routing,
        pim_sm_global=__ret__.pim_sm_global,
        route_limit=__ret__.route_limit,
        route_threshold=__ret__.route_threshold,
        vdomparam=__ret__.vdomparam)
