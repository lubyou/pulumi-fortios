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
    'GetRouterAccessListResult',
    'AwaitableGetRouterAccessListResult',
    'get_router_access_list',
    'get_router_access_list_output',
]

@pulumi.output_type
class GetRouterAccessListResult:
    """
    A collection of values returned by GetRouterAccessList.
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
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetRouterAccessListRuleResult']:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterAccessListResult(GetRouterAccessListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterAccessListResult(
            comments=self.comments,
            id=self.id,
            name=self.name,
            rules=self.rules,
            vdomparam=self.vdomparam)


def get_router_access_list(name: Optional[str] = None,
                           vdomparam: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterAccessListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterAccessList:GetRouterAccessList', __args__, opts=opts, typ=GetRouterAccessListResult).value

    return AwaitableGetRouterAccessListResult(
        comments=pulumi.get(__ret__, 'comments'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        rules=pulumi.get(__ret__, 'rules'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_router_access_list)
def get_router_access_list_output(name: Optional[pulumi.Input[str]] = None,
                                  vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterAccessListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
