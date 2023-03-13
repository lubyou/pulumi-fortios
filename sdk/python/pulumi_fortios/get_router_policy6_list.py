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
    'GetRouterPolicy6ListResult',
    'AwaitableGetRouterPolicy6ListResult',
    'get_router_policy6_list',
    'get_router_policy6_list_output',
]

@pulumi.output_type
class GetRouterPolicy6ListResult:
    """
    A collection of values returned by GetRouterPolicy6List.
    """
    def __init__(__self__, filter=None, id=None, seq_numlists=None, vdomparam=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if seq_numlists and not isinstance(seq_numlists, list):
            raise TypeError("Expected argument 'seq_numlists' to be a list")
        pulumi.set(__self__, "seq_numlists", seq_numlists)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="seqNumlists")
    def seq_numlists(self) -> Sequence[int]:
        return pulumi.get(self, "seq_numlists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterPolicy6ListResult(GetRouterPolicy6ListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterPolicy6ListResult(
            filter=self.filter,
            id=self.id,
            seq_numlists=self.seq_numlists,
            vdomparam=self.vdomparam)


def get_router_policy6_list(filter: Optional[str] = None,
                            vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterPolicy6ListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterPolicy6List:GetRouterPolicy6List', __args__, opts=opts, typ=GetRouterPolicy6ListResult).value

    return AwaitableGetRouterPolicy6ListResult(
        filter=__ret__.filter,
        id=__ret__.id,
        seq_numlists=__ret__.seq_numlists,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_router_policy6_list)
def get_router_policy6_list_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                                   vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterPolicy6ListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
