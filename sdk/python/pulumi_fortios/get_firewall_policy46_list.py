# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFirewallPolicy46ListResult',
    'AwaitableGetFirewallPolicy46ListResult',
    'get_firewall_policy46_list',
]

@pulumi.output_type
class GetFirewallPolicy46ListResult:
    """
    A collection of values returned by GetFirewallPolicy46List.
    """
    def __init__(__self__, filter=None, id=None, policyidlists=None, vdomparam=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policyidlists and not isinstance(policyidlists, list):
            raise TypeError("Expected argument 'policyidlists' to be a list")
        pulumi.set(__self__, "policyidlists", policyidlists)
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
    def policyidlists(self) -> Sequence[int]:
        """
        A list of the `FirewallPolicy46`.
        """
        return pulumi.get(self, "policyidlists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallPolicy46ListResult(GetFirewallPolicy46ListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallPolicy46ListResult(
            filter=self.filter,
            id=self.id,
            policyidlists=self.policyidlists,
            vdomparam=self.vdomparam)


def get_firewall_policy46_list(filter: Optional[str] = None,
                               vdomparam: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallPolicy46ListResult:
    """
    Provides a list of `FirewallPolicy46`.


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallPolicy46List:GetFirewallPolicy46List', __args__, opts=opts, typ=GetFirewallPolicy46ListResult).value

    return AwaitableGetFirewallPolicy46ListResult(
        filter=__ret__.filter,
        id=__ret__.id,
        policyidlists=__ret__.policyidlists,
        vdomparam=__ret__.vdomparam)
