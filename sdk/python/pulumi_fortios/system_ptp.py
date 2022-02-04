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

__all__ = ['SystemPtpArgs', 'SystemPtp']

@pulumi.input_type
class SystemPtpArgs:
    def __init__(__self__, *,
                 interface: pulumi.Input[str],
                 delay_mechanism: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 request_interval: Optional[pulumi.Input[int]] = None,
                 server_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['SystemPtpServerInterfaceArgs']]]] = None,
                 server_mode: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemPtp resource.
        :param pulumi.Input[str] interface: PTP slave will reply through this interface.
        :param pulumi.Input[str] delay_mechanism: End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] mode: Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        :param pulumi.Input[int] request_interval: The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        :param pulumi.Input[Sequence[pulumi.Input['SystemPtpServerInterfaceArgs']]] server_interfaces: FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        :param pulumi.Input[str] server_mode: Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "interface", interface)
        if delay_mechanism is not None:
            pulumi.set(__self__, "delay_mechanism", delay_mechanism)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if request_interval is not None:
            pulumi.set(__self__, "request_interval", request_interval)
        if server_interfaces is not None:
            pulumi.set(__self__, "server_interfaces", server_interfaces)
        if server_mode is not None:
            pulumi.set(__self__, "server_mode", server_mode)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        """
        PTP slave will reply through this interface.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="delayMechanism")
    def delay_mechanism(self) -> Optional[pulumi.Input[str]]:
        """
        End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        """
        return pulumi.get(self, "delay_mechanism")

    @delay_mechanism.setter
    def delay_mechanism(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delay_mechanism", value)

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
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="requestInterval")
    def request_interval(self) -> Optional[pulumi.Input[int]]:
        """
        The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        """
        return pulumi.get(self, "request_interval")

    @request_interval.setter
    def request_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "request_interval", value)

    @property
    @pulumi.getter(name="serverInterfaces")
    def server_interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemPtpServerInterfaceArgs']]]]:
        """
        FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        """
        return pulumi.get(self, "server_interfaces")

    @server_interfaces.setter
    def server_interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemPtpServerInterfaceArgs']]]]):
        pulumi.set(self, "server_interfaces", value)

    @property
    @pulumi.getter(name="serverMode")
    def server_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "server_mode")

    @server_mode.setter
    def server_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_mode", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
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
class _SystemPtpState:
    def __init__(__self__, *,
                 delay_mechanism: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 request_interval: Optional[pulumi.Input[int]] = None,
                 server_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['SystemPtpServerInterfaceArgs']]]] = None,
                 server_mode: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemPtp resources.
        :param pulumi.Input[str] delay_mechanism: End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] interface: PTP slave will reply through this interface.
        :param pulumi.Input[str] mode: Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        :param pulumi.Input[int] request_interval: The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        :param pulumi.Input[Sequence[pulumi.Input['SystemPtpServerInterfaceArgs']]] server_interfaces: FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        :param pulumi.Input[str] server_mode: Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if delay_mechanism is not None:
            pulumi.set(__self__, "delay_mechanism", delay_mechanism)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if request_interval is not None:
            pulumi.set(__self__, "request_interval", request_interval)
        if server_interfaces is not None:
            pulumi.set(__self__, "server_interfaces", server_interfaces)
        if server_mode is not None:
            pulumi.set(__self__, "server_mode", server_mode)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="delayMechanism")
    def delay_mechanism(self) -> Optional[pulumi.Input[str]]:
        """
        End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        """
        return pulumi.get(self, "delay_mechanism")

    @delay_mechanism.setter
    def delay_mechanism(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delay_mechanism", value)

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
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        PTP slave will reply through this interface.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="requestInterval")
    def request_interval(self) -> Optional[pulumi.Input[int]]:
        """
        The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        """
        return pulumi.get(self, "request_interval")

    @request_interval.setter
    def request_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "request_interval", value)

    @property
    @pulumi.getter(name="serverInterfaces")
    def server_interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemPtpServerInterfaceArgs']]]]:
        """
        FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        """
        return pulumi.get(self, "server_interfaces")

    @server_interfaces.setter
    def server_interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemPtpServerInterfaceArgs']]]]):
        pulumi.set(self, "server_interfaces", value)

    @property
    @pulumi.getter(name="serverMode")
    def server_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "server_mode")

    @server_mode.setter
    def server_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_mode", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
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


class SystemPtp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delay_mechanism: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 request_interval: Optional[pulumi.Input[int]] = None,
                 server_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemPtpServerInterfaceArgs']]]]] = None,
                 server_mode: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure system PTP information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemPtp("trname",
            delay_mechanism="E2E",
            interface="port3",
            mode="multicast",
            request_interval=1,
            status="enable")
        ```

        ## Import

        System Ptp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemPtp:SystemPtp labelname SystemPtp
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] delay_mechanism: End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] interface: PTP slave will reply through this interface.
        :param pulumi.Input[str] mode: Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        :param pulumi.Input[int] request_interval: The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemPtpServerInterfaceArgs']]]] server_interfaces: FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        :param pulumi.Input[str] server_mode: Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemPtpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure system PTP information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemPtp("trname",
            delay_mechanism="E2E",
            interface="port3",
            mode="multicast",
            request_interval=1,
            status="enable")
        ```

        ## Import

        System Ptp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemPtp:SystemPtp labelname SystemPtp
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemPtpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemPtpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delay_mechanism: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 request_interval: Optional[pulumi.Input[int]] = None,
                 server_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemPtpServerInterfaceArgs']]]]] = None,
                 server_mode: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemPtpArgs.__new__(SystemPtpArgs)

            __props__.__dict__["delay_mechanism"] = delay_mechanism
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            __props__.__dict__["mode"] = mode
            __props__.__dict__["request_interval"] = request_interval
            __props__.__dict__["server_interfaces"] = server_interfaces
            __props__.__dict__["server_mode"] = server_mode
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemPtp, __self__).__init__(
            'fortios:index/systemPtp:SystemPtp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delay_mechanism: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            request_interval: Optional[pulumi.Input[int]] = None,
            server_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemPtpServerInterfaceArgs']]]]] = None,
            server_mode: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemPtp':
        """
        Get an existing SystemPtp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] delay_mechanism: End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] interface: PTP slave will reply through this interface.
        :param pulumi.Input[str] mode: Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        :param pulumi.Input[int] request_interval: The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemPtpServerInterfaceArgs']]]] server_interfaces: FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        :param pulumi.Input[str] server_mode: Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemPtpState.__new__(_SystemPtpState)

        __props__.__dict__["delay_mechanism"] = delay_mechanism
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["interface"] = interface
        __props__.__dict__["mode"] = mode
        __props__.__dict__["request_interval"] = request_interval
        __props__.__dict__["server_interfaces"] = server_interfaces
        __props__.__dict__["server_mode"] = server_mode
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemPtp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="delayMechanism")
    def delay_mechanism(self) -> pulumi.Output[str]:
        """
        End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        """
        return pulumi.get(self, "delay_mechanism")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        PTP slave will reply through this interface.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="requestInterval")
    def request_interval(self) -> pulumi.Output[int]:
        """
        The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        """
        return pulumi.get(self, "request_interval")

    @property
    @pulumi.getter(name="serverInterfaces")
    def server_interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.SystemPtpServerInterface']]]:
        """
        FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        """
        return pulumi.get(self, "server_interfaces")

    @property
    @pulumi.getter(name="serverMode")
    def server_mode(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "server_mode")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

