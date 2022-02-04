# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemAutoInstallArgs', 'SystemAutoInstall']

@pulumi.input_type
class SystemAutoInstallArgs:
    def __init__(__self__, *,
                 auto_install_config: Optional[pulumi.Input[str]] = None,
                 auto_install_image: Optional[pulumi.Input[str]] = None,
                 default_config_file: Optional[pulumi.Input[str]] = None,
                 default_image_file: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemAutoInstall resource.
        :param pulumi.Input[str] auto_install_config: Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] auto_install_image: Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] default_config_file: Default config file name in USB disk.
        :param pulumi.Input[str] default_image_file: Default image file name in USB disk.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auto_install_config is not None:
            pulumi.set(__self__, "auto_install_config", auto_install_config)
        if auto_install_image is not None:
            pulumi.set(__self__, "auto_install_image", auto_install_image)
        if default_config_file is not None:
            pulumi.set(__self__, "default_config_file", default_config_file)
        if default_image_file is not None:
            pulumi.set(__self__, "default_image_file", default_image_file)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoInstallConfig")
    def auto_install_config(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_install_config")

    @auto_install_config.setter
    def auto_install_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_install_config", value)

    @property
    @pulumi.getter(name="autoInstallImage")
    def auto_install_image(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_install_image")

    @auto_install_image.setter
    def auto_install_image(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_install_image", value)

    @property
    @pulumi.getter(name="defaultConfigFile")
    def default_config_file(self) -> Optional[pulumi.Input[str]]:
        """
        Default config file name in USB disk.
        """
        return pulumi.get(self, "default_config_file")

    @default_config_file.setter
    def default_config_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_config_file", value)

    @property
    @pulumi.getter(name="defaultImageFile")
    def default_image_file(self) -> Optional[pulumi.Input[str]]:
        """
        Default image file name in USB disk.
        """
        return pulumi.get(self, "default_image_file")

    @default_image_file.setter
    def default_image_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_image_file", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemAutoInstallState:
    def __init__(__self__, *,
                 auto_install_config: Optional[pulumi.Input[str]] = None,
                 auto_install_image: Optional[pulumi.Input[str]] = None,
                 default_config_file: Optional[pulumi.Input[str]] = None,
                 default_image_file: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemAutoInstall resources.
        :param pulumi.Input[str] auto_install_config: Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] auto_install_image: Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] default_config_file: Default config file name in USB disk.
        :param pulumi.Input[str] default_image_file: Default image file name in USB disk.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auto_install_config is not None:
            pulumi.set(__self__, "auto_install_config", auto_install_config)
        if auto_install_image is not None:
            pulumi.set(__self__, "auto_install_image", auto_install_image)
        if default_config_file is not None:
            pulumi.set(__self__, "default_config_file", default_config_file)
        if default_image_file is not None:
            pulumi.set(__self__, "default_image_file", default_image_file)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoInstallConfig")
    def auto_install_config(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_install_config")

    @auto_install_config.setter
    def auto_install_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_install_config", value)

    @property
    @pulumi.getter(name="autoInstallImage")
    def auto_install_image(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_install_image")

    @auto_install_image.setter
    def auto_install_image(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_install_image", value)

    @property
    @pulumi.getter(name="defaultConfigFile")
    def default_config_file(self) -> Optional[pulumi.Input[str]]:
        """
        Default config file name in USB disk.
        """
        return pulumi.get(self, "default_config_file")

    @default_config_file.setter
    def default_config_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_config_file", value)

    @property
    @pulumi.getter(name="defaultImageFile")
    def default_image_file(self) -> Optional[pulumi.Input[str]]:
        """
        Default image file name in USB disk.
        """
        return pulumi.get(self, "default_image_file")

    @default_image_file.setter
    def default_image_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_image_file", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemAutoInstall(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_install_config: Optional[pulumi.Input[str]] = None,
                 auto_install_image: Optional[pulumi.Input[str]] = None,
                 default_config_file: Optional[pulumi.Input[str]] = None,
                 default_image_file: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure USB auto installation.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemAutoInstall("trname",
            auto_install_config="enable",
            auto_install_image="enable",
            default_config_file="fgt_system.conf",
            default_image_file="image.out")
        ```

        ## Import

        System AutoInstall can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemAutoInstall:SystemAutoInstall labelname SystemAutoInstall
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_install_config: Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] auto_install_image: Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] default_config_file: Default config file name in USB disk.
        :param pulumi.Input[str] default_image_file: Default image file name in USB disk.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemAutoInstallArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure USB auto installation.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemAutoInstall("trname",
            auto_install_config="enable",
            auto_install_image="enable",
            default_config_file="fgt_system.conf",
            default_image_file="image.out")
        ```

        ## Import

        System AutoInstall can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemAutoInstall:SystemAutoInstall labelname SystemAutoInstall
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemAutoInstallArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemAutoInstallArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_install_config: Optional[pulumi.Input[str]] = None,
                 auto_install_image: Optional[pulumi.Input[str]] = None,
                 default_config_file: Optional[pulumi.Input[str]] = None,
                 default_image_file: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemAutoInstallArgs.__new__(SystemAutoInstallArgs)

            __props__.__dict__["auto_install_config"] = auto_install_config
            __props__.__dict__["auto_install_image"] = auto_install_image
            __props__.__dict__["default_config_file"] = default_config_file
            __props__.__dict__["default_image_file"] = default_image_file
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemAutoInstall, __self__).__init__(
            'fortios:index/systemAutoInstall:SystemAutoInstall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_install_config: Optional[pulumi.Input[str]] = None,
            auto_install_image: Optional[pulumi.Input[str]] = None,
            default_config_file: Optional[pulumi.Input[str]] = None,
            default_image_file: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemAutoInstall':
        """
        Get an existing SystemAutoInstall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_install_config: Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] auto_install_image: Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] default_config_file: Default config file name in USB disk.
        :param pulumi.Input[str] default_image_file: Default image file name in USB disk.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemAutoInstallState.__new__(_SystemAutoInstallState)

        __props__.__dict__["auto_install_config"] = auto_install_config
        __props__.__dict__["auto_install_image"] = auto_install_image
        __props__.__dict__["default_config_file"] = default_config_file
        __props__.__dict__["default_image_file"] = default_image_file
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemAutoInstall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoInstallConfig")
    def auto_install_config(self) -> pulumi.Output[str]:
        """
        Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_install_config")

    @property
    @pulumi.getter(name="autoInstallImage")
    def auto_install_image(self) -> pulumi.Output[str]:
        """
        Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_install_image")

    @property
    @pulumi.getter(name="defaultConfigFile")
    def default_config_file(self) -> pulumi.Output[str]:
        """
        Default config file name in USB disk.
        """
        return pulumi.get(self, "default_config_file")

    @property
    @pulumi.getter(name="defaultImageFile")
    def default_image_file(self) -> pulumi.Output[str]:
        """
        Default image file name in USB disk.
        """
        return pulumi.get(self, "default_image_file")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

