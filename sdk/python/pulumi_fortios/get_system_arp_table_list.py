# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemArpTableListResult',
    'AwaitableGetSystemArpTableListResult',
    'get_system_arp_table_list',
    'get_system_arp_table_list_output',
]

@pulumi.output_type
class GetSystemArpTableListResult:
    """
    A collection of values returned by GetSystemArpTableList.
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
        """
        A list of the `SystemArpTable`.
        """
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


class AwaitableGetSystemArpTableListResult(GetSystemArpTableListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemArpTableListResult(
            filter=self.filter,
            fosidlists=self.fosidlists,
            id=self.id,
            vdomparam=self.vdomparam)


def get_system_arp_table_list(filter: Optional[str] = None,
                              vdomparam: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemArpTableListResult:
    """
    Provides a list of `SystemArpTable`.


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemArpTableList:GetSystemArpTableList', __args__, opts=opts, typ=GetSystemArpTableListResult).value

    return AwaitableGetSystemArpTableListResult(
        filter=__ret__.filter,
        fosidlists=__ret__.fosidlists,
        id=__ret__.id,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_arp_table_list)
def get_system_arp_table_list_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                                     vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemArpTableListResult]:
    """
    Provides a list of `SystemArpTable`.


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
