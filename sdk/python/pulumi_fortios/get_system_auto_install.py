# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemAutoInstallResult',
    'AwaitableGetSystemAutoInstallResult',
    'get_system_auto_install',
    'get_system_auto_install_output',
]

@pulumi.output_type
class GetSystemAutoInstallResult:
    """
    A collection of values returned by GetSystemAutoInstall.
    """
    def __init__(__self__, auto_install_config=None, auto_install_image=None, default_config_file=None, default_image_file=None, id=None, vdomparam=None):
        if auto_install_config and not isinstance(auto_install_config, str):
            raise TypeError("Expected argument 'auto_install_config' to be a str")
        pulumi.set(__self__, "auto_install_config", auto_install_config)
        if auto_install_image and not isinstance(auto_install_image, str):
            raise TypeError("Expected argument 'auto_install_image' to be a str")
        pulumi.set(__self__, "auto_install_image", auto_install_image)
        if default_config_file and not isinstance(default_config_file, str):
            raise TypeError("Expected argument 'default_config_file' to be a str")
        pulumi.set(__self__, "default_config_file", default_config_file)
        if default_image_file and not isinstance(default_image_file, str):
            raise TypeError("Expected argument 'default_image_file' to be a str")
        pulumi.set(__self__, "default_image_file", default_image_file)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoInstallConfig")
    def auto_install_config(self) -> str:
        return pulumi.get(self, "auto_install_config")

    @property
    @pulumi.getter(name="autoInstallImage")
    def auto_install_image(self) -> str:
        return pulumi.get(self, "auto_install_image")

    @property
    @pulumi.getter(name="defaultConfigFile")
    def default_config_file(self) -> str:
        return pulumi.get(self, "default_config_file")

    @property
    @pulumi.getter(name="defaultImageFile")
    def default_image_file(self) -> str:
        return pulumi.get(self, "default_image_file")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemAutoInstallResult(GetSystemAutoInstallResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAutoInstallResult(
            auto_install_config=self.auto_install_config,
            auto_install_image=self.auto_install_image,
            default_config_file=self.default_config_file,
            default_image_file=self.default_image_file,
            id=self.id,
            vdomparam=self.vdomparam)


def get_system_auto_install(vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAutoInstallResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemAutoInstall:GetSystemAutoInstall', __args__, opts=opts, typ=GetSystemAutoInstallResult).value

    return AwaitableGetSystemAutoInstallResult(
        auto_install_config=pulumi.get(__ret__, 'auto_install_config'),
        auto_install_image=pulumi.get(__ret__, 'auto_install_image'),
        default_config_file=pulumi.get(__ret__, 'default_config_file'),
        default_image_file=pulumi.get(__ret__, 'default_image_file'),
        id=pulumi.get(__ret__, 'id'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_system_auto_install)
def get_system_auto_install_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemAutoInstallResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
