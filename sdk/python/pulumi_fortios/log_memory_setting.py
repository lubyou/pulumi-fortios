# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LogMemorySettingArgs', 'LogMemorySetting']

@pulumi.input_type
class LogMemorySettingArgs:
    def __init__(__self__, *,
                 diskfull: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogMemorySetting resource.
        :param pulumi.Input[str] diskfull: Action to take when memory is full. Valid values: `overwrite`.
        :param pulumi.Input[str] status: Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if diskfull is not None:
            pulumi.set(__self__, "diskfull", diskfull)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def diskfull(self) -> Optional[pulumi.Input[str]]:
        """
        Action to take when memory is full. Valid values: `overwrite`.
        """
        return pulumi.get(self, "diskfull")

    @diskfull.setter
    def diskfull(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diskfull", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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
class _LogMemorySettingState:
    def __init__(__self__, *,
                 diskfull: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogMemorySetting resources.
        :param pulumi.Input[str] diskfull: Action to take when memory is full. Valid values: `overwrite`.
        :param pulumi.Input[str] status: Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if diskfull is not None:
            pulumi.set(__self__, "diskfull", diskfull)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def diskfull(self) -> Optional[pulumi.Input[str]]:
        """
        Action to take when memory is full. Valid values: `overwrite`.
        """
        return pulumi.get(self, "diskfull")

    @diskfull.setter
    def diskfull(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diskfull", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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


class LogMemorySetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 diskfull: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Settings for memory buffer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.LogMemorySetting("trname",
            diskfull="overwrite",
            status="disable")
        ```

        ## Import

        LogMemory Setting can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/logMemorySetting:LogMemorySetting labelname LogMemorySetting
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/logMemorySetting:LogMemorySetting labelname LogMemorySetting
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] diskfull: Action to take when memory is full. Valid values: `overwrite`.
        :param pulumi.Input[str] status: Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LogMemorySettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Settings for memory buffer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.LogMemorySetting("trname",
            diskfull="overwrite",
            status="disable")
        ```

        ## Import

        LogMemory Setting can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/logMemorySetting:LogMemorySetting labelname LogMemorySetting
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/logMemorySetting:LogMemorySetting labelname LogMemorySetting
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param LogMemorySettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogMemorySettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 diskfull: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
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
            __props__ = LogMemorySettingArgs.__new__(LogMemorySettingArgs)

            __props__.__dict__["diskfull"] = diskfull
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(LogMemorySetting, __self__).__init__(
            'fortios:index/logMemorySetting:LogMemorySetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            diskfull: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'LogMemorySetting':
        """
        Get an existing LogMemorySetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] diskfull: Action to take when memory is full. Valid values: `overwrite`.
        :param pulumi.Input[str] status: Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogMemorySettingState.__new__(_LogMemorySettingState)

        __props__.__dict__["diskfull"] = diskfull
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return LogMemorySetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def diskfull(self) -> pulumi.Output[str]:
        """
        Action to take when memory is full. Valid values: `overwrite`.
        """
        return pulumi.get(self, "diskfull")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

