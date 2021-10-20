# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetRouterStatic6ListResult',
    'AwaitableGetRouterStatic6ListResult',
    'get_router_static6_list',
]

@pulumi.output_type
class GetRouterStatic6ListResult:
    """
    A collection of values returned by GetRouterStatic6List.
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
        """
        A list of the `RouterStatic6`.
        """
        return pulumi.get(self, "seq_numlists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterStatic6ListResult(GetRouterStatic6ListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterStatic6ListResult(
            filter=self.filter,
            id=self.id,
            seq_numlists=self.seq_numlists,
            vdomparam=self.vdomparam)


def get_router_static6_list(filter: Optional[str] = None,
                            vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterStatic6ListResult:
    """
    Provides a list of `RouterStatic6`.


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterStatic6List:GetRouterStatic6List', __args__, opts=opts, typ=GetRouterStatic6ListResult).value

    return AwaitableGetRouterStatic6ListResult(
        filter=__ret__.filter,
        id=__ret__.id,
        seq_numlists=__ret__.seq_numlists,
        vdomparam=__ret__.vdomparam)