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
    'GetSystemSessionHelperListResult',
    'AwaitableGetSystemSessionHelperListResult',
    'get_system_session_helper_list',
    'get_system_session_helper_list_output',
]

@pulumi.output_type
class GetSystemSessionHelperListResult:
    """
    A collection of values returned by GetSystemSessionHelperList.
    """
    def __init__(__self__, filter=None, fosidlists=None, id=None, vdomparam=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if fosidlists and not isinstance(fosidlists, list):
            raise TypeError("Expected argument 'fosidlists' to be a list")
        pulumi.set(__self__, "fosidlists", fosidlists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def fosidlists(self) -> Sequence[int]:
        return pulumi.get(self, "fosidlists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemSessionHelperListResult(GetSystemSessionHelperListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemSessionHelperListResult(
            filter=self.filter,
            fosidlists=self.fosidlists,
            id=self.id,
            vdomparam=self.vdomparam)


def get_system_session_helper_list(filter: Optional[str] = None,
                                   vdomparam: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemSessionHelperListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemSessionHelperList:GetSystemSessionHelperList', __args__, opts=opts, typ=GetSystemSessionHelperListResult).value

    return AwaitableGetSystemSessionHelperListResult(
        filter=__ret__.filter,
        fosidlists=__ret__.fosidlists,
        id=__ret__.id,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_session_helper_list)
def get_system_session_helper_list_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                                          vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemSessionHelperListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
