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
    'GetSystemExternalResourceListResult',
    'AwaitableGetSystemExternalResourceListResult',
    'get_system_external_resource_list',
    'get_system_external_resource_list_output',
]

@pulumi.output_type
class GetSystemExternalResourceListResult:
    """
    A collection of values returned by GetSystemExternalResourceList.
    """
    def __init__(__self__, filter=None, id=None, namelists=None, vdomparam=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namelists and not isinstance(namelists, list):
            raise TypeError("Expected argument 'namelists' to be a list")
        pulumi.set(__self__, "namelists", namelists)
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
    @pulumi.getter
    def namelists(self) -> Sequence[str]:
        return pulumi.get(self, "namelists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemExternalResourceListResult(GetSystemExternalResourceListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemExternalResourceListResult(
            filter=self.filter,
            id=self.id,
            namelists=self.namelists,
            vdomparam=self.vdomparam)


def get_system_external_resource_list(filter: Optional[str] = None,
                                      vdomparam: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemExternalResourceListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemExternalResourceList:GetSystemExternalResourceList', __args__, opts=opts, typ=GetSystemExternalResourceListResult).value

    return AwaitableGetSystemExternalResourceListResult(
        filter=pulumi.get(__ret__, 'filter'),
        id=pulumi.get(__ret__, 'id'),
        namelists=pulumi.get(__ret__, 'namelists'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_system_external_resource_list)
def get_system_external_resource_list_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                                             vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemExternalResourceListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
