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
    'GetFirewallProxyAddrgrpResult',
    'AwaitableGetFirewallProxyAddrgrpResult',
    'get_firewall_proxy_addrgrp',
]

@pulumi.output_type
class GetFirewallProxyAddrgrpResult:
    """
    A collection of values returned by GetFirewallProxyAddrgrp.
    """
    def __init__(__self__, color=None, comment=None, id=None, members=None, name=None, taggings=None, type=None, uuid=None, vdomparam=None, visibility=None):
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
        if taggings and not isinstance(taggings, list):
            raise TypeError("Expected argument 'taggings' to be a list")
        pulumi.set(__self__, "taggings", taggings)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
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
        Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Optional comments.
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
    def members(self) -> Sequence['outputs.GetFirewallProxyAddrgrpMemberResult']:
        """
        Members of address group. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Tag name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def taggings(self) -> Sequence['outputs.GetFirewallProxyAddrgrpTaggingResult']:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Source or destination address group type.
        """
        return pulumi.get(self, "type")

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
        Enable/disable visibility of the object in the GUI.
        """
        return pulumi.get(self, "visibility")


class AwaitableGetFirewallProxyAddrgrpResult(GetFirewallProxyAddrgrpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallProxyAddrgrpResult(
            color=self.color,
            comment=self.comment,
            id=self.id,
            members=self.members,
            name=self.name,
            taggings=self.taggings,
            type=self.type,
            uuid=self.uuid,
            vdomparam=self.vdomparam,
            visibility=self.visibility)


def get_firewall_proxy_addrgrp(name: Optional[str] = None,
                               vdomparam: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallProxyAddrgrpResult:
    """
    Use this data source to get information on an fortios firewall proxyaddrgrp


    :param str name: Specify the name of the desired firewall proxyaddrgrp.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallProxyAddrgrp:GetFirewallProxyAddrgrp', __args__, opts=opts, typ=GetFirewallProxyAddrgrpResult).value

    return AwaitableGetFirewallProxyAddrgrpResult(
        color=__ret__.color,
        comment=__ret__.comment,
        id=__ret__.id,
        members=__ret__.members,
        name=__ret__.name,
        taggings=__ret__.taggings,
        type=__ret__.type,
        uuid=__ret__.uuid,
        vdomparam=__ret__.vdomparam,
        visibility=__ret__.visibility)
