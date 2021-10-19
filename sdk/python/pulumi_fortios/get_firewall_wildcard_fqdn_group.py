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
    'GetFirewallWildcardFqdnGroupResult',
    'AwaitableGetFirewallWildcardFqdnGroupResult',
    'get_firewall_wildcard_fqdn_group',
]

@pulumi.output_type
class GetFirewallWildcardFqdnGroupResult:
    """
    A collection of values returned by GetFirewallWildcardFqdnGroup.
    """
    def __init__(__self__, color=None, comment=None, id=None, members=None, name=None, uuid=None, vdomparam=None, visibility=None):
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        GUI icon color.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def members(self) -> Sequence['outputs.GetFirewallWildcardFqdnGroupMemberResult']:
        """
        Address group members. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Address name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        Enable/disable address visibility.
        """
        return pulumi.get(self, "visibility")


class AwaitableGetFirewallWildcardFqdnGroupResult(GetFirewallWildcardFqdnGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallWildcardFqdnGroupResult(
            color=self.color,
            comment=self.comment,
            id=self.id,
            members=self.members,
            name=self.name,
            uuid=self.uuid,
            vdomparam=self.vdomparam,
            visibility=self.visibility)


def get_firewall_wildcard_fqdn_group(name: Optional[str] = None,
                                     vdomparam: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallWildcardFqdnGroupResult:
    """
    Use this data source to get information on an fortios firewallwildcardfqdn group


    :param str name: Specify the name of the desired firewallwildcardfqdn group.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallWildcardFqdnGroup:GetFirewallWildcardFqdnGroup', __args__, opts=opts, typ=GetFirewallWildcardFqdnGroupResult).value

    return AwaitableGetFirewallWildcardFqdnGroupResult(
        color=__ret__.color,
        comment=__ret__.comment,
        id=__ret__.id,
        members=__ret__.members,
        name=__ret__.name,
        uuid=__ret__.uuid,
        vdomparam=__ret__.vdomparam,
        visibility=__ret__.visibility)
