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
    'GetFirewallInternetServiceCustomResult',
    'AwaitableGetFirewallInternetServiceCustomResult',
    'get_firewall_internet_service_custom',
    'get_firewall_internet_service_custom_output',
]

@pulumi.output_type
class GetFirewallInternetServiceCustomResult:
    """
    A collection of values returned by GetFirewallInternetServiceCustom.
    """
    def __init__(__self__, comment=None, entries=None, id=None, name=None, reputation=None, vdomparam=None):
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if entries and not isinstance(entries, list):
            raise TypeError("Expected argument 'entries' to be a list")
        pulumi.set(__self__, "entries", entries)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if reputation and not isinstance(reputation, int):
            raise TypeError("Expected argument 'reputation' to be a int")
        pulumi.set(__self__, "reputation", reputation)
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
    def entries(self) -> Sequence['outputs.GetFirewallInternetServiceCustomEntryResult']:
        """
        Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
        """
        return pulumi.get(self, "entries")

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
        Select the destination address or address group object from available options.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def reputation(self) -> int:
        """
        Reputation level of the custom Internet Service.
        """
        return pulumi.get(self, "reputation")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallInternetServiceCustomResult(GetFirewallInternetServiceCustomResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallInternetServiceCustomResult(
            comment=self.comment,
            entries=self.entries,
            id=self.id,
            name=self.name,
            reputation=self.reputation,
            vdomparam=self.vdomparam)


def get_firewall_internet_service_custom(name: Optional[str] = None,
                                         vdomparam: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallInternetServiceCustomResult:
    """
    Use this data source to get information on an fortios firewall internetservicecustom


    :param str name: Specify the name of the desired firewall internetservicecustom.
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
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallInternetServiceCustom:GetFirewallInternetServiceCustom', __args__, opts=opts, typ=GetFirewallInternetServiceCustomResult).value

    return AwaitableGetFirewallInternetServiceCustomResult(
        comment=__ret__.comment,
        entries=__ret__.entries,
        id=__ret__.id,
        name=__ret__.name,
        reputation=__ret__.reputation,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_firewall_internet_service_custom)
def get_firewall_internet_service_custom_output(name: Optional[pulumi.Input[str]] = None,
                                                vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallInternetServiceCustomResult]:
    """
    Use this data source to get information on an fortios firewall internetservicecustom


    :param str name: Specify the name of the desired firewall internetservicecustom.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
