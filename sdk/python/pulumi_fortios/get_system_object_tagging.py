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
    'GetSystemObjectTaggingResult',
    'AwaitableGetSystemObjectTaggingResult',
    'get_system_object_tagging',
    'get_system_object_tagging_output',
]

@pulumi.output_type
class GetSystemObjectTaggingResult:
    """
    A collection of values returned by GetSystemObjectTagging.
    """
    def __init__(__self__, address=None, category=None, color=None, device=None, id=None, interface=None, multiple=None, tags=None, vdomparam=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if category and not isinstance(category, str):
            raise TypeError("Expected argument 'category' to be a str")
        pulumi.set(__self__, "category", category)
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if multiple and not isinstance(multiple, str):
            raise TypeError("Expected argument 'multiple' to be a str")
        pulumi.set(__self__, "multiple", multiple)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        Address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def category(self) -> str:
        """
        Tag Category.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def device(self) -> str:
        """
        Device.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        Interface.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def multiple(self) -> str:
        """
        Allow multiple tag selection.
        """
        return pulumi.get(self, "multiple")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetSystemObjectTaggingTagResult']:
        """
        Tags. The structure of `tags` block is documented below.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemObjectTaggingResult(GetSystemObjectTaggingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemObjectTaggingResult(
            address=self.address,
            category=self.category,
            color=self.color,
            device=self.device,
            id=self.id,
            interface=self.interface,
            multiple=self.multiple,
            tags=self.tags,
            vdomparam=self.vdomparam)


def get_system_object_tagging(category: Optional[str] = None,
                              vdomparam: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemObjectTaggingResult:
    """
    Use this data source to get information on an fortios system objecttagging


    :param str category: Specify the category of the desired system objecttagging.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['category'] = category
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemObjectTagging:GetSystemObjectTagging', __args__, opts=opts, typ=GetSystemObjectTaggingResult).value

    return AwaitableGetSystemObjectTaggingResult(
        address=__ret__.address,
        category=__ret__.category,
        color=__ret__.color,
        device=__ret__.device,
        id=__ret__.id,
        interface=__ret__.interface,
        multiple=__ret__.multiple,
        tags=__ret__.tags,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_object_tagging)
def get_system_object_tagging_output(category: Optional[pulumi.Input[str]] = None,
                                     vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemObjectTaggingResult]:
    """
    Use this data source to get information on an fortios system objecttagging


    :param str category: Specify the category of the desired system objecttagging.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
