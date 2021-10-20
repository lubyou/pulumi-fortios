# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['WirelessControllerApcfgProfileArgs', 'WirelessControllerApcfgProfile']

@pulumi.input_type
class WirelessControllerApcfgProfileArgs:
    def __init__(__self__, *,
                 ac_ip: Optional[pulumi.Input[str]] = None,
                 ac_port: Optional[pulumi.Input[int]] = None,
                 ac_timer: Optional[pulumi.Input[int]] = None,
                 ac_type: Optional[pulumi.Input[str]] = None,
                 command_lists: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerApcfgProfileCommandListArgs']]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerApcfgProfile resource.
        :param pulumi.Input[str] ac_ip: IP address of the validation controller that AP must be able to join after applying AP local configuration.
        :param pulumi.Input[int] ac_port: Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
        :param pulumi.Input[int] ac_timer: Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
        :param pulumi.Input[str] ac_type: Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerApcfgProfileCommandListArgs']]] command_lists: AP local configuration command list. The structure of `command_list` block is documented below.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: AP local configuration command name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ac_ip is not None:
            pulumi.set(__self__, "ac_ip", ac_ip)
        if ac_port is not None:
            pulumi.set(__self__, "ac_port", ac_port)
        if ac_timer is not None:
            pulumi.set(__self__, "ac_timer", ac_timer)
        if ac_type is not None:
            pulumi.set(__self__, "ac_type", ac_type)
        if command_lists is not None:
            pulumi.set(__self__, "command_lists", command_lists)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="acIp")
    def ac_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the validation controller that AP must be able to join after applying AP local configuration.
        """
        return pulumi.get(self, "ac_ip")

    @ac_ip.setter
    def ac_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ac_ip", value)

    @property
    @pulumi.getter(name="acPort")
    def ac_port(self) -> Optional[pulumi.Input[int]]:
        """
        Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
        """
        return pulumi.get(self, "ac_port")

    @ac_port.setter
    def ac_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ac_port", value)

    @property
    @pulumi.getter(name="acTimer")
    def ac_timer(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
        """
        return pulumi.get(self, "ac_timer")

    @ac_timer.setter
    def ac_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ac_timer", value)

    @property
    @pulumi.getter(name="acType")
    def ac_type(self) -> Optional[pulumi.Input[str]]:
        """
        Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
        """
        return pulumi.get(self, "ac_type")

    @ac_type.setter
    def ac_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ac_type", value)

    @property
    @pulumi.getter(name="commandLists")
    def command_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerApcfgProfileCommandListArgs']]]]:
        """
        AP local configuration command list. The structure of `command_list` block is documented below.
        """
        return pulumi.get(self, "command_lists")

    @command_lists.setter
    def command_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerApcfgProfileCommandListArgs']]]]):
        pulumi.set(self, "command_lists", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        AP local configuration command name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _WirelessControllerApcfgProfileState:
    def __init__(__self__, *,
                 ac_ip: Optional[pulumi.Input[str]] = None,
                 ac_port: Optional[pulumi.Input[int]] = None,
                 ac_timer: Optional[pulumi.Input[int]] = None,
                 ac_type: Optional[pulumi.Input[str]] = None,
                 command_lists: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerApcfgProfileCommandListArgs']]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerApcfgProfile resources.
        :param pulumi.Input[str] ac_ip: IP address of the validation controller that AP must be able to join after applying AP local configuration.
        :param pulumi.Input[int] ac_port: Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
        :param pulumi.Input[int] ac_timer: Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
        :param pulumi.Input[str] ac_type: Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerApcfgProfileCommandListArgs']]] command_lists: AP local configuration command list. The structure of `command_list` block is documented below.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: AP local configuration command name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ac_ip is not None:
            pulumi.set(__self__, "ac_ip", ac_ip)
        if ac_port is not None:
            pulumi.set(__self__, "ac_port", ac_port)
        if ac_timer is not None:
            pulumi.set(__self__, "ac_timer", ac_timer)
        if ac_type is not None:
            pulumi.set(__self__, "ac_type", ac_type)
        if command_lists is not None:
            pulumi.set(__self__, "command_lists", command_lists)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="acIp")
    def ac_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the validation controller that AP must be able to join after applying AP local configuration.
        """
        return pulumi.get(self, "ac_ip")

    @ac_ip.setter
    def ac_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ac_ip", value)

    @property
    @pulumi.getter(name="acPort")
    def ac_port(self) -> Optional[pulumi.Input[int]]:
        """
        Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
        """
        return pulumi.get(self, "ac_port")

    @ac_port.setter
    def ac_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ac_port", value)

    @property
    @pulumi.getter(name="acTimer")
    def ac_timer(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
        """
        return pulumi.get(self, "ac_timer")

    @ac_timer.setter
    def ac_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ac_timer", value)

    @property
    @pulumi.getter(name="acType")
    def ac_type(self) -> Optional[pulumi.Input[str]]:
        """
        Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
        """
        return pulumi.get(self, "ac_type")

    @ac_type.setter
    def ac_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ac_type", value)

    @property
    @pulumi.getter(name="commandLists")
    def command_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerApcfgProfileCommandListArgs']]]]:
        """
        AP local configuration command list. The structure of `command_list` block is documented below.
        """
        return pulumi.get(self, "command_lists")

    @command_lists.setter
    def command_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerApcfgProfileCommandListArgs']]]]):
        pulumi.set(self, "command_lists", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        AP local configuration command name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class WirelessControllerApcfgProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ac_ip: Optional[pulumi.Input[str]] = None,
                 ac_port: Optional[pulumi.Input[int]] = None,
                 ac_timer: Optional[pulumi.Input[int]] = None,
                 ac_type: Optional[pulumi.Input[str]] = None,
                 command_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerApcfgProfileCommandListArgs']]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure AP local configuration profiles. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        WirelessController ApcfgProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerApcfgProfile:WirelessControllerApcfgProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ac_ip: IP address of the validation controller that AP must be able to join after applying AP local configuration.
        :param pulumi.Input[int] ac_port: Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
        :param pulumi.Input[int] ac_timer: Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
        :param pulumi.Input[str] ac_type: Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerApcfgProfileCommandListArgs']]]] command_lists: AP local configuration command list. The structure of `command_list` block is documented below.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: AP local configuration command name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerApcfgProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure AP local configuration profiles. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        WirelessController ApcfgProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerApcfgProfile:WirelessControllerApcfgProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelessControllerApcfgProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerApcfgProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ac_ip: Optional[pulumi.Input[str]] = None,
                 ac_port: Optional[pulumi.Input[int]] = None,
                 ac_timer: Optional[pulumi.Input[int]] = None,
                 ac_type: Optional[pulumi.Input[str]] = None,
                 command_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerApcfgProfileCommandListArgs']]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelessControllerApcfgProfileArgs.__new__(WirelessControllerApcfgProfileArgs)

            __props__.__dict__["ac_ip"] = ac_ip
            __props__.__dict__["ac_port"] = ac_port
            __props__.__dict__["ac_timer"] = ac_timer
            __props__.__dict__["ac_type"] = ac_type
            __props__.__dict__["command_lists"] = command_lists
            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerApcfgProfile, __self__).__init__(
            'fortios:index/wirelessControllerApcfgProfile:WirelessControllerApcfgProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ac_ip: Optional[pulumi.Input[str]] = None,
            ac_port: Optional[pulumi.Input[int]] = None,
            ac_timer: Optional[pulumi.Input[int]] = None,
            ac_type: Optional[pulumi.Input[str]] = None,
            command_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerApcfgProfileCommandListArgs']]]]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerApcfgProfile':
        """
        Get an existing WirelessControllerApcfgProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ac_ip: IP address of the validation controller that AP must be able to join after applying AP local configuration.
        :param pulumi.Input[int] ac_port: Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
        :param pulumi.Input[int] ac_timer: Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
        :param pulumi.Input[str] ac_type: Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerApcfgProfileCommandListArgs']]]] command_lists: AP local configuration command list. The structure of `command_list` block is documented below.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: AP local configuration command name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerApcfgProfileState.__new__(_WirelessControllerApcfgProfileState)

        __props__.__dict__["ac_ip"] = ac_ip
        __props__.__dict__["ac_port"] = ac_port
        __props__.__dict__["ac_timer"] = ac_timer
        __props__.__dict__["ac_type"] = ac_type
        __props__.__dict__["command_lists"] = command_lists
        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerApcfgProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acIp")
    def ac_ip(self) -> pulumi.Output[str]:
        """
        IP address of the validation controller that AP must be able to join after applying AP local configuration.
        """
        return pulumi.get(self, "ac_ip")

    @property
    @pulumi.getter(name="acPort")
    def ac_port(self) -> pulumi.Output[int]:
        """
        Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
        """
        return pulumi.get(self, "ac_port")

    @property
    @pulumi.getter(name="acTimer")
    def ac_timer(self) -> pulumi.Output[int]:
        """
        Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
        """
        return pulumi.get(self, "ac_timer")

    @property
    @pulumi.getter(name="acType")
    def ac_type(self) -> pulumi.Output[str]:
        """
        Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
        """
        return pulumi.get(self, "ac_type")

    @property
    @pulumi.getter(name="commandLists")
    def command_lists(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerApcfgProfileCommandList']]]:
        """
        AP local configuration command list. The structure of `command_list` block is documented below.
        """
        return pulumi.get(self, "command_lists")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        AP local configuration command name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
