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
    'GetRouterKeyChainResult',
    'AwaitableGetRouterKeyChainResult',
    'get_router_key_chain',
]

@pulumi.output_type
class GetRouterKeyChainResult:
    """
    A collection of values returned by GetRouterKeyChain.
    """
    def __init__(__self__, id=None, keys=None, name=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
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
    def keys(self) -> Sequence['outputs.GetRouterKeyChainKeyResult']:
        """
        Configuration method to edit key settings. The structure of `key` block is documented below.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Key-chain name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterKeyChainResult(GetRouterKeyChainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterKeyChainResult(
            id=self.id,
            keys=self.keys,
            name=self.name,
            vdomparam=self.vdomparam)


def get_router_key_chain(name: Optional[str] = None,
                         vdomparam: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterKeyChainResult:
    """
    Use this data source to get information on an fortios router keychain


    :param str name: Specify the name of the desired router keychain.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterKeyChain:GetRouterKeyChain', __args__, opts=opts, typ=GetRouterKeyChainResult).value

    return AwaitableGetRouterKeyChainResult(
        id=__ret__.id,
        keys=__ret__.keys,
        name=__ret__.name,
        vdomparam=__ret__.vdomparam)