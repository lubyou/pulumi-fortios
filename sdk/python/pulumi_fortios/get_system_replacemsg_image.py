# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemReplacemsgImageResult',
    'AwaitableGetSystemReplacemsgImageResult',
    'get_system_replacemsg_image',
]

@pulumi.output_type
class GetSystemReplacemsgImageResult:
    """
    A collection of values returned by GetSystemReplacemsgImage.
    """
    def __init__(__self__, id=None, image_base64=None, image_type=None, name=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_base64 and not isinstance(image_base64, str):
            raise TypeError("Expected argument 'image_base64' to be a str")
        pulumi.set(__self__, "image_base64", image_base64)
        if image_type and not isinstance(image_type, str):
            raise TypeError("Expected argument 'image_type' to be a str")
        pulumi.set(__self__, "image_type", image_type)
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
    @pulumi.getter(name="imageBase64")
    def image_base64(self) -> str:
        """
        Image data.
        """
        return pulumi.get(self, "image_base64")

    @property
    @pulumi.getter(name="imageType")
    def image_type(self) -> str:
        """
        Image type.
        """
        return pulumi.get(self, "image_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Image name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemReplacemsgImageResult(GetSystemReplacemsgImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemReplacemsgImageResult(
            id=self.id,
            image_base64=self.image_base64,
            image_type=self.image_type,
            name=self.name,
            vdomparam=self.vdomparam)


def get_system_replacemsg_image(name: Optional[str] = None,
                                vdomparam: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemReplacemsgImageResult:
    """
    Use this data source to get information on an fortios system replacemsgimage


    :param str name: Specify the name of the desired system replacemsgimage.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemReplacemsgImage:GetSystemReplacemsgImage', __args__, opts=opts, typ=GetSystemReplacemsgImageResult).value

    return AwaitableGetSystemReplacemsgImageResult(
        id=__ret__.id,
        image_base64=__ret__.image_base64,
        image_type=__ret__.image_type,
        name=__ret__.name,
        vdomparam=__ret__.vdomparam)
