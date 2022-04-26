# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemMacAddressTableArgs', 'SystemMacAddressTable']

@pulumi.input_type
class SystemMacAddressTableArgs:
    def __init__(__self__, *,
                 interface: pulumi.Input[str],
                 mac: pulumi.Input[str],
                 reply_substitute: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemMacAddressTable resource.
        :param pulumi.Input[str] interface: Interface name.
        :param pulumi.Input[str] mac: MAC address.
        :param pulumi.Input[str] reply_substitute: New MAC for reply traffic.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "mac", mac)
        if reply_substitute is not None:
            pulumi.set(__self__, "reply_substitute", reply_substitute)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        """
        Interface name.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Input[str]:
        """
        MAC address.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: pulumi.Input[str]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter(name="replySubstitute")
    def reply_substitute(self) -> Optional[pulumi.Input[str]]:
        """
        New MAC for reply traffic.
        """
        return pulumi.get(self, "reply_substitute")

    @reply_substitute.setter
    def reply_substitute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reply_substitute", value)

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
class _SystemMacAddressTableState:
    def __init__(__self__, *,
                 interface: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 reply_substitute: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemMacAddressTable resources.
        :param pulumi.Input[str] interface: Interface name.
        :param pulumi.Input[str] mac: MAC address.
        :param pulumi.Input[str] reply_substitute: New MAC for reply traffic.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if reply_substitute is not None:
            pulumi.set(__self__, "reply_substitute", reply_substitute)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter(name="replySubstitute")
    def reply_substitute(self) -> Optional[pulumi.Input[str]]:
        """
        New MAC for reply traffic.
        """
        return pulumi.get(self, "reply_substitute")

    @reply_substitute.setter
    def reply_substitute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reply_substitute", value)

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


class SystemMacAddressTable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 reply_substitute: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure MAC address tables.

        ## Import

        System MacAddressTable can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemMacAddressTable:SystemMacAddressTable labelname {{mac}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemMacAddressTable:SystemMacAddressTable labelname {{mac}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] interface: Interface name.
        :param pulumi.Input[str] mac: MAC address.
        :param pulumi.Input[str] reply_substitute: New MAC for reply traffic.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemMacAddressTableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure MAC address tables.

        ## Import

        System MacAddressTable can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemMacAddressTable:SystemMacAddressTable labelname {{mac}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemMacAddressTable:SystemMacAddressTable labelname {{mac}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemMacAddressTableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemMacAddressTableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 reply_substitute: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemMacAddressTableArgs.__new__(SystemMacAddressTableArgs)

            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            if mac is None and not opts.urn:
                raise TypeError("Missing required property 'mac'")
            __props__.__dict__["mac"] = mac
            __props__.__dict__["reply_substitute"] = reply_substitute
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemMacAddressTable, __self__).__init__(
            'fortios:index/systemMacAddressTable:SystemMacAddressTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            interface: Optional[pulumi.Input[str]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            reply_substitute: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemMacAddressTable':
        """
        Get an existing SystemMacAddressTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] interface: Interface name.
        :param pulumi.Input[str] mac: MAC address.
        :param pulumi.Input[str] reply_substitute: New MAC for reply traffic.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemMacAddressTableState.__new__(_SystemMacAddressTableState)

        __props__.__dict__["interface"] = interface
        __props__.__dict__["mac"] = mac
        __props__.__dict__["reply_substitute"] = reply_substitute
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemMacAddressTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Interface name.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        MAC address.
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter(name="replySubstitute")
    def reply_substitute(self) -> pulumi.Output[str]:
        """
        New MAC for reply traffic.
        """
        return pulumi.get(self, "reply_substitute")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

