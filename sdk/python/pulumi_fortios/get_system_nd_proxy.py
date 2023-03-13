# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetSystemNdProxyResult',
    'AwaitableGetSystemNdProxyResult',
    'get_system_nd_proxy',
    'get_system_nd_proxy_output',
]

@pulumi.output_type
class GetSystemNdProxyResult:
    """
    A collection of values returned by GetSystemNdProxy.
    """
    def __init__(__self__, id=None, members=None, status=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
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
    def members(self) -> Sequence['outputs.GetSystemNdProxyMemberResult']:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemNdProxyResult(GetSystemNdProxyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemNdProxyResult(
            id=self.id,
            members=self.members,
            status=self.status,
            vdomparam=self.vdomparam)


def get_system_nd_proxy(vdomparam: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemNdProxyResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemNdProxy:GetSystemNdProxy', __args__, opts=opts, typ=GetSystemNdProxyResult).value

    return AwaitableGetSystemNdProxyResult(
        id=__ret__.id,
        members=__ret__.members,
        status=__ret__.status,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_nd_proxy)
def get_system_nd_proxy_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemNdProxyResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
