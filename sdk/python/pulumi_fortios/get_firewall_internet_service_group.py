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
    'GetFirewallInternetServiceGroupResult',
    'AwaitableGetFirewallInternetServiceGroupResult',
    'get_firewall_internet_service_group',
    'get_firewall_internet_service_group_output',
]

@pulumi.output_type
class GetFirewallInternetServiceGroupResult:
    """
    A collection of values returned by GetFirewallInternetServiceGroup.
    """
    def __init__(__self__, comment=None, direction=None, id=None, members=None, name=None, vdomparam=None):
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def direction(self) -> str:
        """
        How this service may be used (source, destination or both).
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def members(self) -> Sequence['outputs.GetFirewallInternetServiceGroupMemberResult']:
        """
        Internet Service group member. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Internet Service name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallInternetServiceGroupResult(GetFirewallInternetServiceGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallInternetServiceGroupResult(
            comment=self.comment,
            direction=self.direction,
            id=self.id,
            members=self.members,
            name=self.name,
            vdomparam=self.vdomparam)


def get_firewall_internet_service_group(name: Optional[str] = None,
                                        vdomparam: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallInternetServiceGroupResult:
    """
    Use this data source to get information on an fortios firewall internetservicegroup


    :param str name: Specify the name of the desired firewall internetservicegroup.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallInternetServiceGroup:GetFirewallInternetServiceGroup', __args__, opts=opts, typ=GetFirewallInternetServiceGroupResult).value

    return AwaitableGetFirewallInternetServiceGroupResult(
        comment=__ret__.comment,
        direction=__ret__.direction,
        id=__ret__.id,
        members=__ret__.members,
        name=__ret__.name,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_firewall_internet_service_group)
def get_firewall_internet_service_group_output(name: Optional[pulumi.Input[str]] = None,
                                               vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallInternetServiceGroupResult]:
    """
    Use this data source to get information on an fortios firewall internetservicegroup


    :param str name: Specify the name of the desired firewall internetservicegroup.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
