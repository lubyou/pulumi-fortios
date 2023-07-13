# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SwitchControllerFlowTrackingArgs', 'SwitchControllerFlowTracking']

@pulumi.input_type
class SwitchControllerFlowTrackingArgs:
    def __init__(__self__, *,
                 aggregates: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingAggregateArgs']]]] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 collectors: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingCollectorArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 max_export_pkt_size: Optional[pulumi.Input[int]] = None,
                 sample_mode: Optional[pulumi.Input[str]] = None,
                 sample_rate: Optional[pulumi.Input[int]] = None,
                 template_export_period: Optional[pulumi.Input[int]] = None,
                 timeout_general: Optional[pulumi.Input[int]] = None,
                 timeout_icmp: Optional[pulumi.Input[int]] = None,
                 timeout_max: Optional[pulumi.Input[int]] = None,
                 timeout_tcp: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_fin: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_rst: Optional[pulumi.Input[int]] = None,
                 timeout_udp: Optional[pulumi.Input[int]] = None,
                 transport: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerFlowTracking resource.
        """
        if aggregates is not None:
            pulumi.set(__self__, "aggregates", aggregates)
        if collector_ip is not None:
            pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port is not None:
            pulumi.set(__self__, "collector_port", collector_port)
        if collectors is not None:
            pulumi.set(__self__, "collectors", collectors)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if level is not None:
            pulumi.set(__self__, "level", level)
        if max_export_pkt_size is not None:
            pulumi.set(__self__, "max_export_pkt_size", max_export_pkt_size)
        if sample_mode is not None:
            pulumi.set(__self__, "sample_mode", sample_mode)
        if sample_rate is not None:
            pulumi.set(__self__, "sample_rate", sample_rate)
        if template_export_period is not None:
            pulumi.set(__self__, "template_export_period", template_export_period)
        if timeout_general is not None:
            pulumi.set(__self__, "timeout_general", timeout_general)
        if timeout_icmp is not None:
            pulumi.set(__self__, "timeout_icmp", timeout_icmp)
        if timeout_max is not None:
            pulumi.set(__self__, "timeout_max", timeout_max)
        if timeout_tcp is not None:
            pulumi.set(__self__, "timeout_tcp", timeout_tcp)
        if timeout_tcp_fin is not None:
            pulumi.set(__self__, "timeout_tcp_fin", timeout_tcp_fin)
        if timeout_tcp_rst is not None:
            pulumi.set(__self__, "timeout_tcp_rst", timeout_tcp_rst)
        if timeout_udp is not None:
            pulumi.set(__self__, "timeout_udp", timeout_udp)
        if transport is not None:
            pulumi.set(__self__, "transport", transport)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def aggregates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingAggregateArgs']]]]:
        return pulumi.get(self, "aggregates")

    @aggregates.setter
    def aggregates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingAggregateArgs']]]]):
        pulumi.set(self, "aggregates", value)

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "collector_ip")

    @collector_ip.setter
    def collector_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collector_ip", value)

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "collector_port")

    @collector_port.setter
    def collector_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "collector_port", value)

    @property
    @pulumi.getter
    def collectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingCollectorArgs']]]]:
        return pulumi.get(self, "collectors")

    @collectors.setter
    def collectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingCollectorArgs']]]]):
        pulumi.set(self, "collectors", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "level", value)

    @property
    @pulumi.getter(name="maxExportPktSize")
    def max_export_pkt_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_export_pkt_size")

    @max_export_pkt_size.setter
    def max_export_pkt_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_export_pkt_size", value)

    @property
    @pulumi.getter(name="sampleMode")
    def sample_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sample_mode")

    @sample_mode.setter
    def sample_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sample_mode", value)

    @property
    @pulumi.getter(name="sampleRate")
    def sample_rate(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "sample_rate")

    @sample_rate.setter
    def sample_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sample_rate", value)

    @property
    @pulumi.getter(name="templateExportPeriod")
    def template_export_period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "template_export_period")

    @template_export_period.setter
    def template_export_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_export_period", value)

    @property
    @pulumi.getter(name="timeoutGeneral")
    def timeout_general(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_general")

    @timeout_general.setter
    def timeout_general(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_general", value)

    @property
    @pulumi.getter(name="timeoutIcmp")
    def timeout_icmp(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_icmp")

    @timeout_icmp.setter
    def timeout_icmp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_icmp", value)

    @property
    @pulumi.getter(name="timeoutMax")
    def timeout_max(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_max")

    @timeout_max.setter
    def timeout_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_max", value)

    @property
    @pulumi.getter(name="timeoutTcp")
    def timeout_tcp(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_tcp")

    @timeout_tcp.setter
    def timeout_tcp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_tcp", value)

    @property
    @pulumi.getter(name="timeoutTcpFin")
    def timeout_tcp_fin(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_tcp_fin")

    @timeout_tcp_fin.setter
    def timeout_tcp_fin(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_tcp_fin", value)

    @property
    @pulumi.getter(name="timeoutTcpRst")
    def timeout_tcp_rst(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_tcp_rst")

    @timeout_tcp_rst.setter
    def timeout_tcp_rst(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_tcp_rst", value)

    @property
    @pulumi.getter(name="timeoutUdp")
    def timeout_udp(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_udp")

    @timeout_udp.setter
    def timeout_udp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_udp", value)

    @property
    @pulumi.getter
    def transport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "transport")

    @transport.setter
    def transport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transport", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SwitchControllerFlowTrackingState:
    def __init__(__self__, *,
                 aggregates: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingAggregateArgs']]]] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 collectors: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingCollectorArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 max_export_pkt_size: Optional[pulumi.Input[int]] = None,
                 sample_mode: Optional[pulumi.Input[str]] = None,
                 sample_rate: Optional[pulumi.Input[int]] = None,
                 template_export_period: Optional[pulumi.Input[int]] = None,
                 timeout_general: Optional[pulumi.Input[int]] = None,
                 timeout_icmp: Optional[pulumi.Input[int]] = None,
                 timeout_max: Optional[pulumi.Input[int]] = None,
                 timeout_tcp: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_fin: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_rst: Optional[pulumi.Input[int]] = None,
                 timeout_udp: Optional[pulumi.Input[int]] = None,
                 transport: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerFlowTracking resources.
        """
        if aggregates is not None:
            pulumi.set(__self__, "aggregates", aggregates)
        if collector_ip is not None:
            pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port is not None:
            pulumi.set(__self__, "collector_port", collector_port)
        if collectors is not None:
            pulumi.set(__self__, "collectors", collectors)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if level is not None:
            pulumi.set(__self__, "level", level)
        if max_export_pkt_size is not None:
            pulumi.set(__self__, "max_export_pkt_size", max_export_pkt_size)
        if sample_mode is not None:
            pulumi.set(__self__, "sample_mode", sample_mode)
        if sample_rate is not None:
            pulumi.set(__self__, "sample_rate", sample_rate)
        if template_export_period is not None:
            pulumi.set(__self__, "template_export_period", template_export_period)
        if timeout_general is not None:
            pulumi.set(__self__, "timeout_general", timeout_general)
        if timeout_icmp is not None:
            pulumi.set(__self__, "timeout_icmp", timeout_icmp)
        if timeout_max is not None:
            pulumi.set(__self__, "timeout_max", timeout_max)
        if timeout_tcp is not None:
            pulumi.set(__self__, "timeout_tcp", timeout_tcp)
        if timeout_tcp_fin is not None:
            pulumi.set(__self__, "timeout_tcp_fin", timeout_tcp_fin)
        if timeout_tcp_rst is not None:
            pulumi.set(__self__, "timeout_tcp_rst", timeout_tcp_rst)
        if timeout_udp is not None:
            pulumi.set(__self__, "timeout_udp", timeout_udp)
        if transport is not None:
            pulumi.set(__self__, "transport", transport)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def aggregates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingAggregateArgs']]]]:
        return pulumi.get(self, "aggregates")

    @aggregates.setter
    def aggregates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingAggregateArgs']]]]):
        pulumi.set(self, "aggregates", value)

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "collector_ip")

    @collector_ip.setter
    def collector_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collector_ip", value)

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "collector_port")

    @collector_port.setter
    def collector_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "collector_port", value)

    @property
    @pulumi.getter
    def collectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingCollectorArgs']]]]:
        return pulumi.get(self, "collectors")

    @collectors.setter
    def collectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerFlowTrackingCollectorArgs']]]]):
        pulumi.set(self, "collectors", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "level", value)

    @property
    @pulumi.getter(name="maxExportPktSize")
    def max_export_pkt_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_export_pkt_size")

    @max_export_pkt_size.setter
    def max_export_pkt_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_export_pkt_size", value)

    @property
    @pulumi.getter(name="sampleMode")
    def sample_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sample_mode")

    @sample_mode.setter
    def sample_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sample_mode", value)

    @property
    @pulumi.getter(name="sampleRate")
    def sample_rate(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "sample_rate")

    @sample_rate.setter
    def sample_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sample_rate", value)

    @property
    @pulumi.getter(name="templateExportPeriod")
    def template_export_period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "template_export_period")

    @template_export_period.setter
    def template_export_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_export_period", value)

    @property
    @pulumi.getter(name="timeoutGeneral")
    def timeout_general(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_general")

    @timeout_general.setter
    def timeout_general(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_general", value)

    @property
    @pulumi.getter(name="timeoutIcmp")
    def timeout_icmp(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_icmp")

    @timeout_icmp.setter
    def timeout_icmp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_icmp", value)

    @property
    @pulumi.getter(name="timeoutMax")
    def timeout_max(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_max")

    @timeout_max.setter
    def timeout_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_max", value)

    @property
    @pulumi.getter(name="timeoutTcp")
    def timeout_tcp(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_tcp")

    @timeout_tcp.setter
    def timeout_tcp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_tcp", value)

    @property
    @pulumi.getter(name="timeoutTcpFin")
    def timeout_tcp_fin(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_tcp_fin")

    @timeout_tcp_fin.setter
    def timeout_tcp_fin(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_tcp_fin", value)

    @property
    @pulumi.getter(name="timeoutTcpRst")
    def timeout_tcp_rst(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_tcp_rst")

    @timeout_tcp_rst.setter
    def timeout_tcp_rst(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_tcp_rst", value)

    @property
    @pulumi.getter(name="timeoutUdp")
    def timeout_udp(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_udp")

    @timeout_udp.setter
    def timeout_udp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_udp", value)

    @property
    @pulumi.getter
    def transport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "transport")

    @transport.setter
    def transport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transport", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SwitchControllerFlowTracking(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerFlowTrackingAggregateArgs']]]]] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 collectors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerFlowTrackingCollectorArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 max_export_pkt_size: Optional[pulumi.Input[int]] = None,
                 sample_mode: Optional[pulumi.Input[str]] = None,
                 sample_rate: Optional[pulumi.Input[int]] = None,
                 template_export_period: Optional[pulumi.Input[int]] = None,
                 timeout_general: Optional[pulumi.Input[int]] = None,
                 timeout_icmp: Optional[pulumi.Input[int]] = None,
                 timeout_max: Optional[pulumi.Input[int]] = None,
                 timeout_tcp: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_fin: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_rst: Optional[pulumi.Input[int]] = None,
                 timeout_udp: Optional[pulumi.Input[int]] = None,
                 transport: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SwitchControllerFlowTracking resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerFlowTrackingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SwitchControllerFlowTracking resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SwitchControllerFlowTrackingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerFlowTrackingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerFlowTrackingAggregateArgs']]]]] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 collectors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerFlowTrackingCollectorArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 max_export_pkt_size: Optional[pulumi.Input[int]] = None,
                 sample_mode: Optional[pulumi.Input[str]] = None,
                 sample_rate: Optional[pulumi.Input[int]] = None,
                 template_export_period: Optional[pulumi.Input[int]] = None,
                 timeout_general: Optional[pulumi.Input[int]] = None,
                 timeout_icmp: Optional[pulumi.Input[int]] = None,
                 timeout_max: Optional[pulumi.Input[int]] = None,
                 timeout_tcp: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_fin: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_rst: Optional[pulumi.Input[int]] = None,
                 timeout_udp: Optional[pulumi.Input[int]] = None,
                 transport: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchControllerFlowTrackingArgs.__new__(SwitchControllerFlowTrackingArgs)

            __props__.__dict__["aggregates"] = aggregates
            __props__.__dict__["collector_ip"] = collector_ip
            __props__.__dict__["collector_port"] = collector_port
            __props__.__dict__["collectors"] = collectors
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["format"] = format
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["level"] = level
            __props__.__dict__["max_export_pkt_size"] = max_export_pkt_size
            __props__.__dict__["sample_mode"] = sample_mode
            __props__.__dict__["sample_rate"] = sample_rate
            __props__.__dict__["template_export_period"] = template_export_period
            __props__.__dict__["timeout_general"] = timeout_general
            __props__.__dict__["timeout_icmp"] = timeout_icmp
            __props__.__dict__["timeout_max"] = timeout_max
            __props__.__dict__["timeout_tcp"] = timeout_tcp
            __props__.__dict__["timeout_tcp_fin"] = timeout_tcp_fin
            __props__.__dict__["timeout_tcp_rst"] = timeout_tcp_rst
            __props__.__dict__["timeout_udp"] = timeout_udp
            __props__.__dict__["transport"] = transport
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerFlowTracking, __self__).__init__(
            'fortios:index/switchControllerFlowTracking:SwitchControllerFlowTracking',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aggregates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerFlowTrackingAggregateArgs']]]]] = None,
            collector_ip: Optional[pulumi.Input[str]] = None,
            collector_port: Optional[pulumi.Input[int]] = None,
            collectors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerFlowTrackingCollectorArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            level: Optional[pulumi.Input[str]] = None,
            max_export_pkt_size: Optional[pulumi.Input[int]] = None,
            sample_mode: Optional[pulumi.Input[str]] = None,
            sample_rate: Optional[pulumi.Input[int]] = None,
            template_export_period: Optional[pulumi.Input[int]] = None,
            timeout_general: Optional[pulumi.Input[int]] = None,
            timeout_icmp: Optional[pulumi.Input[int]] = None,
            timeout_max: Optional[pulumi.Input[int]] = None,
            timeout_tcp: Optional[pulumi.Input[int]] = None,
            timeout_tcp_fin: Optional[pulumi.Input[int]] = None,
            timeout_tcp_rst: Optional[pulumi.Input[int]] = None,
            timeout_udp: Optional[pulumi.Input[int]] = None,
            transport: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerFlowTracking':
        """
        Get an existing SwitchControllerFlowTracking resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerFlowTrackingState.__new__(_SwitchControllerFlowTrackingState)

        __props__.__dict__["aggregates"] = aggregates
        __props__.__dict__["collector_ip"] = collector_ip
        __props__.__dict__["collector_port"] = collector_port
        __props__.__dict__["collectors"] = collectors
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["format"] = format
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["level"] = level
        __props__.__dict__["max_export_pkt_size"] = max_export_pkt_size
        __props__.__dict__["sample_mode"] = sample_mode
        __props__.__dict__["sample_rate"] = sample_rate
        __props__.__dict__["template_export_period"] = template_export_period
        __props__.__dict__["timeout_general"] = timeout_general
        __props__.__dict__["timeout_icmp"] = timeout_icmp
        __props__.__dict__["timeout_max"] = timeout_max
        __props__.__dict__["timeout_tcp"] = timeout_tcp
        __props__.__dict__["timeout_tcp_fin"] = timeout_tcp_fin
        __props__.__dict__["timeout_tcp_rst"] = timeout_tcp_rst
        __props__.__dict__["timeout_udp"] = timeout_udp
        __props__.__dict__["transport"] = transport
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerFlowTracking(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def aggregates(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerFlowTrackingAggregate']]]:
        return pulumi.get(self, "aggregates")

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "collector_ip")

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "collector_port")

    @property
    @pulumi.getter
    def collectors(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerFlowTrackingCollector']]]:
        return pulumi.get(self, "collectors")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[str]:
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def level(self) -> pulumi.Output[str]:
        return pulumi.get(self, "level")

    @property
    @pulumi.getter(name="maxExportPktSize")
    def max_export_pkt_size(self) -> pulumi.Output[int]:
        return pulumi.get(self, "max_export_pkt_size")

    @property
    @pulumi.getter(name="sampleMode")
    def sample_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sample_mode")

    @property
    @pulumi.getter(name="sampleRate")
    def sample_rate(self) -> pulumi.Output[int]:
        return pulumi.get(self, "sample_rate")

    @property
    @pulumi.getter(name="templateExportPeriod")
    def template_export_period(self) -> pulumi.Output[int]:
        return pulumi.get(self, "template_export_period")

    @property
    @pulumi.getter(name="timeoutGeneral")
    def timeout_general(self) -> pulumi.Output[int]:
        return pulumi.get(self, "timeout_general")

    @property
    @pulumi.getter(name="timeoutIcmp")
    def timeout_icmp(self) -> pulumi.Output[int]:
        return pulumi.get(self, "timeout_icmp")

    @property
    @pulumi.getter(name="timeoutMax")
    def timeout_max(self) -> pulumi.Output[int]:
        return pulumi.get(self, "timeout_max")

    @property
    @pulumi.getter(name="timeoutTcp")
    def timeout_tcp(self) -> pulumi.Output[int]:
        return pulumi.get(self, "timeout_tcp")

    @property
    @pulumi.getter(name="timeoutTcpFin")
    def timeout_tcp_fin(self) -> pulumi.Output[int]:
        return pulumi.get(self, "timeout_tcp_fin")

    @property
    @pulumi.getter(name="timeoutTcpRst")
    def timeout_tcp_rst(self) -> pulumi.Output[int]:
        return pulumi.get(self, "timeout_tcp_rst")

    @property
    @pulumi.getter(name="timeoutUdp")
    def timeout_udp(self) -> pulumi.Output[int]:
        return pulumi.get(self, "timeout_udp")

    @property
    @pulumi.getter
    def transport(self) -> pulumi.Output[str]:
        return pulumi.get(self, "transport")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

