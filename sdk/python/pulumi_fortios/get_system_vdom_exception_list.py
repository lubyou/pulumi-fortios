# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemVdomExceptionListResult',
    'AwaitableGetSystemVdomExceptionListResult',
    'get_system_vdom_exception_list',
]

@pulumi.output_type
class GetSystemVdomExceptionListResult:
    """
    A collection of values returned by GetSystemVdomExceptionList.
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
        A list of the `SystemVdomException`.
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


class AwaitableGetSystemVdomExceptionListResult(GetSystemVdomExceptionListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemVdomExceptionListResult(
            filter=self.filter,
            fosidlists=self.fosidlists,
            id=self.id,
            vdomparam=self.vdomparam)


def get_system_vdom_exception_list(filter: Optional[str] = None,
                                   vdomparam: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemVdomExceptionListResult:
    """
    Provides a list of `SystemVdomException`.


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemVdomExceptionList:GetSystemVdomExceptionList', __args__, opts=opts, typ=GetSystemVdomExceptionListResult).value

    return AwaitableGetSystemVdomExceptionListResult(
        filter=__ret__.filter,
        fosidlists=__ret__.fosidlists,
        id=__ret__.id,
        vdomparam=__ret__.vdomparam)
