# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemReplacemsgSslvpnArgs', 'SystemReplacemsgSslvpn']

@pulumi.input_type
class SystemReplacemsgSslvpnArgs:
    def __init__(__self__, *,
                 msg_type: pulumi.Input[str],
                 buffer: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 header: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemReplacemsgSslvpn resource.
        :param pulumi.Input[str] msg_type: Message type.
        :param pulumi.Input[str] buffer: Message string.
        :param pulumi.Input[str] format: Format flag.
        :param pulumi.Input[str] header: Header flag. Valid values: `none`, `http`, `8bit`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "msg_type", msg_type)
        if buffer is not None:
            pulumi.set(__self__, "buffer", buffer)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if header is not None:
            pulumi.set(__self__, "header", header)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="msgType")
    def msg_type(self) -> pulumi.Input[str]:
        """
        Message type.
        """
        return pulumi.get(self, "msg_type")

    @msg_type.setter
    def msg_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "msg_type", value)

    @property
    @pulumi.getter
    def buffer(self) -> Optional[pulumi.Input[str]]:
        """
        Message string.
        """
        return pulumi.get(self, "buffer")

    @buffer.setter
    def buffer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "buffer", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Format flag.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def header(self) -> Optional[pulumi.Input[str]]:
        """
        Header flag. Valid values: `none`, `http`, `8bit`.
        """
        return pulumi.get(self, "header")

    @header.setter
    def header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header", value)

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
class _SystemReplacemsgSslvpnState:
    def __init__(__self__, *,
                 buffer: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 header: Optional[pulumi.Input[str]] = None,
                 msg_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemReplacemsgSslvpn resources.
        :param pulumi.Input[str] buffer: Message string.
        :param pulumi.Input[str] format: Format flag.
        :param pulumi.Input[str] header: Header flag. Valid values: `none`, `http`, `8bit`.
        :param pulumi.Input[str] msg_type: Message type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if buffer is not None:
            pulumi.set(__self__, "buffer", buffer)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if header is not None:
            pulumi.set(__self__, "header", header)
        if msg_type is not None:
            pulumi.set(__self__, "msg_type", msg_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def buffer(self) -> Optional[pulumi.Input[str]]:
        """
        Message string.
        """
        return pulumi.get(self, "buffer")

    @buffer.setter
    def buffer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "buffer", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Format flag.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def header(self) -> Optional[pulumi.Input[str]]:
        """
        Header flag. Valid values: `none`, `http`, `8bit`.
        """
        return pulumi.get(self, "header")

    @header.setter
    def header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header", value)

    @property
    @pulumi.getter(name="msgType")
    def msg_type(self) -> Optional[pulumi.Input[str]]:
        """
        Message type.
        """
        return pulumi.get(self, "msg_type")

    @msg_type.setter
    def msg_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "msg_type", value)

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


class SystemReplacemsgSslvpn(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 buffer: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 header: Optional[pulumi.Input[str]] = None,
                 msg_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Replacement messages.

        ## Import

        SystemReplacemsg Sslvpn can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemReplacemsgSslvpn:SystemReplacemsgSslvpn labelname {{msg_type}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] buffer: Message string.
        :param pulumi.Input[str] format: Format flag.
        :param pulumi.Input[str] header: Header flag. Valid values: `none`, `http`, `8bit`.
        :param pulumi.Input[str] msg_type: Message type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemReplacemsgSslvpnArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Replacement messages.

        ## Import

        SystemReplacemsg Sslvpn can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemReplacemsgSslvpn:SystemReplacemsgSslvpn labelname {{msg_type}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemReplacemsgSslvpnArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemReplacemsgSslvpnArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 buffer: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 header: Optional[pulumi.Input[str]] = None,
                 msg_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemReplacemsgSslvpnArgs.__new__(SystemReplacemsgSslvpnArgs)

            __props__.__dict__["buffer"] = buffer
            __props__.__dict__["format"] = format
            __props__.__dict__["header"] = header
            if msg_type is None and not opts.urn:
                raise TypeError("Missing required property 'msg_type'")
            __props__.__dict__["msg_type"] = msg_type
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemReplacemsgSslvpn, __self__).__init__(
            'fortios:index/systemReplacemsgSslvpn:SystemReplacemsgSslvpn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            buffer: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            header: Optional[pulumi.Input[str]] = None,
            msg_type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemReplacemsgSslvpn':
        """
        Get an existing SystemReplacemsgSslvpn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] buffer: Message string.
        :param pulumi.Input[str] format: Format flag.
        :param pulumi.Input[str] header: Header flag. Valid values: `none`, `http`, `8bit`.
        :param pulumi.Input[str] msg_type: Message type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemReplacemsgSslvpnState.__new__(_SystemReplacemsgSslvpnState)

        __props__.__dict__["buffer"] = buffer
        __props__.__dict__["format"] = format
        __props__.__dict__["header"] = header
        __props__.__dict__["msg_type"] = msg_type
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemReplacemsgSslvpn(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def buffer(self) -> pulumi.Output[Optional[str]]:
        """
        Message string.
        """
        return pulumi.get(self, "buffer")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[str]:
        """
        Format flag.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter
    def header(self) -> pulumi.Output[str]:
        """
        Header flag. Valid values: `none`, `http`, `8bit`.
        """
        return pulumi.get(self, "header")

    @property
    @pulumi.getter(name="msgType")
    def msg_type(self) -> pulumi.Output[str]:
        """
        Message type.
        """
        return pulumi.get(self, "msg_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

