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
    'GetRouterBfd6Result',
    'AwaitableGetRouterBfd6Result',
    'get_router_bfd6',
]

@pulumi.output_type
class GetRouterBfd6Result:
    """
    A collection of values returned by GetRouterBfd6.
    """
    def __init__(__self__, id=None, neighbors=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if neighbors and not isinstance(neighbors, list):
            raise TypeError("Expected argument 'neighbors' to be a list")
        pulumi.set(__self__, "neighbors", neighbors)
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
    def neighbors(self) -> Sequence['outputs.GetRouterBfd6NeighborResult']:
        """
        Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterBfd6Result(GetRouterBfd6Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterBfd6Result(
            id=self.id,
            neighbors=self.neighbors,
            vdomparam=self.vdomparam)


def get_router_bfd6(vdomparam: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterBfd6Result:
    """
    Use this data source to get information on fortios router bfd6


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterBfd6:GetRouterBfd6', __args__, opts=opts, typ=GetRouterBfd6Result).value

    return AwaitableGetRouterBfd6Result(
        id=__ret__.id,
        neighbors=__ret__.neighbors,
        vdomparam=__ret__.vdomparam)
