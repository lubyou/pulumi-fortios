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
    'GetRouterMulticast6Result',
    'AwaitableGetRouterMulticast6Result',
    'get_router_multicast6',
]

@pulumi.output_type
class GetRouterMulticast6Result:
    """
    A collection of values returned by GetRouterMulticast6.
    """
    def __init__(__self__, id=None, interfaces=None, multicast_pmtu=None, multicast_routing=None, pim_sm_global=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if multicast_pmtu and not isinstance(multicast_pmtu, str):
            raise TypeError("Expected argument 'multicast_pmtu' to be a str")
        pulumi.set(__self__, "multicast_pmtu", multicast_pmtu)
        if multicast_routing and not isinstance(multicast_routing, str):
            raise TypeError("Expected argument 'multicast_routing' to be a str")
        pulumi.set(__self__, "multicast_routing", multicast_routing)
        if pim_sm_global and not isinstance(pim_sm_global, dict):
            raise TypeError("Expected argument 'pim_sm_global' to be a dict")
        pulumi.set(__self__, "pim_sm_global", pim_sm_global)
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
    def interfaces(self) -> Sequence['outputs.GetRouterMulticast6InterfaceResult']:
        """
        Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="multicastPmtu")
    def multicast_pmtu(self) -> str:
        """
        Enable/disable PMTU for IPv6 multicast.
        """
        return pulumi.get(self, "multicast_pmtu")

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> str:
        """
        Enable/disable IPv6 multicast routing.
        """
        return pulumi.get(self, "multicast_routing")

    @property
    @pulumi.getter(name="pimSmGlobal")
    def pim_sm_global(self) -> 'outputs.GetRouterMulticast6PimSmGlobalResult':
        """
        PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        """
        return pulumi.get(self, "pim_sm_global")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterMulticast6Result(GetRouterMulticast6Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterMulticast6Result(
            id=self.id,
            interfaces=self.interfaces,
            multicast_pmtu=self.multicast_pmtu,
            multicast_routing=self.multicast_routing,
            pim_sm_global=self.pim_sm_global,
            vdomparam=self.vdomparam)


def get_router_multicast6(vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterMulticast6Result:
    """
    Use this data source to get information on fortios router multicast6


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterMulticast6:GetRouterMulticast6', __args__, opts=opts, typ=GetRouterMulticast6Result).value

    return AwaitableGetRouterMulticast6Result(
        id=__ret__.id,
        interfaces=__ret__.interfaces,
        multicast_pmtu=__ret__.multicast_pmtu,
        multicast_routing=__ret__.multicast_routing,
        pim_sm_global=__ret__.pim_sm_global,
        vdomparam=__ret__.vdomparam)