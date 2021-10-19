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
    'GetRouterPrefixListResult',
    'AwaitableGetRouterPrefixListResult',
    'get_router_prefix_list',
]

@pulumi.output_type
class GetRouterPrefixListResult:
    """
    A collection of values returned by GetRouterPrefixList.
    """
    def __init__(__self__, comments=None, id=None, name=None, rules=None, vdomparam=None):
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comments(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetRouterPrefixListRuleResult']:
        """
        IPv4 prefix list rule. The structure of `rule` block is documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterPrefixListResult(GetRouterPrefixListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterPrefixListResult(
            comments=self.comments,
            id=self.id,
            name=self.name,
            rules=self.rules,
            vdomparam=self.vdomparam)


def get_router_prefix_list(name: Optional[str] = None,
                           vdomparam: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterPrefixListResult:
    """
    Use this data source to get information on an fortios router prefixlist


    :param str name: Specify the name of the desired router prefixlist.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterPrefixList:GetRouterPrefixList', __args__, opts=opts, typ=GetRouterPrefixListResult).value

    return AwaitableGetRouterPrefixListResult(
        comments=__ret__.comments,
        id=__ret__.id,
        name=__ret__.name,
        rules=__ret__.rules,
        vdomparam=__ret__.vdomparam)
