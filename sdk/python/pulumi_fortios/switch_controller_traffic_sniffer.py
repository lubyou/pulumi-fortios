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

__all__ = ['SwitchControllerTrafficSnifferArgs', 'SwitchControllerTrafficSniffer']

@pulumi.input_type
class SwitchControllerTrafficSnifferArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 erspan_ip: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 target_ips: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetIpArgs']]]] = None,
                 target_macs: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetMacArgs']]]] = None,
                 target_ports: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetPortArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerTrafficSniffer resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if erspan_ip is not None:
            pulumi.set(__self__, "erspan_ip", erspan_ip)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if target_ips is not None:
            pulumi.set(__self__, "target_ips", target_ips)
        if target_macs is not None:
            pulumi.set(__self__, "target_macs", target_macs)
        if target_ports is not None:
            pulumi.set(__self__, "target_ports", target_ports)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="erspanIp")
    def erspan_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "erspan_ip")

    @erspan_ip.setter
    def erspan_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "erspan_ip", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="targetIps")
    def target_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetIpArgs']]]]:
        return pulumi.get(self, "target_ips")

    @target_ips.setter
    def target_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetIpArgs']]]]):
        pulumi.set(self, "target_ips", value)

    @property
    @pulumi.getter(name="targetMacs")
    def target_macs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetMacArgs']]]]:
        return pulumi.get(self, "target_macs")

    @target_macs.setter
    def target_macs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetMacArgs']]]]):
        pulumi.set(self, "target_macs", value)

    @property
    @pulumi.getter(name="targetPorts")
    def target_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetPortArgs']]]]:
        return pulumi.get(self, "target_ports")

    @target_ports.setter
    def target_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetPortArgs']]]]):
        pulumi.set(self, "target_ports", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SwitchControllerTrafficSnifferState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 erspan_ip: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 target_ips: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetIpArgs']]]] = None,
                 target_macs: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetMacArgs']]]] = None,
                 target_ports: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetPortArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerTrafficSniffer resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if erspan_ip is not None:
            pulumi.set(__self__, "erspan_ip", erspan_ip)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if target_ips is not None:
            pulumi.set(__self__, "target_ips", target_ips)
        if target_macs is not None:
            pulumi.set(__self__, "target_macs", target_macs)
        if target_ports is not None:
            pulumi.set(__self__, "target_ports", target_ports)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="erspanIp")
    def erspan_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "erspan_ip")

    @erspan_ip.setter
    def erspan_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "erspan_ip", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="targetIps")
    def target_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetIpArgs']]]]:
        return pulumi.get(self, "target_ips")

    @target_ips.setter
    def target_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetIpArgs']]]]):
        pulumi.set(self, "target_ips", value)

    @property
    @pulumi.getter(name="targetMacs")
    def target_macs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetMacArgs']]]]:
        return pulumi.get(self, "target_macs")

    @target_macs.setter
    def target_macs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetMacArgs']]]]):
        pulumi.set(self, "target_macs", value)

    @property
    @pulumi.getter(name="targetPorts")
    def target_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetPortArgs']]]]:
        return pulumi.get(self, "target_ports")

    @target_ports.setter
    def target_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerTrafficSnifferTargetPortArgs']]]]):
        pulumi.set(self, "target_ports", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SwitchControllerTrafficSniffer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 erspan_ip: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 target_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetIpArgs']]]]] = None,
                 target_macs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetMacArgs']]]]] = None,
                 target_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetPortArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SwitchControllerTrafficSniffer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerTrafficSnifferArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SwitchControllerTrafficSniffer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SwitchControllerTrafficSnifferArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerTrafficSnifferArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 erspan_ip: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 target_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetIpArgs']]]]] = None,
                 target_macs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetMacArgs']]]]] = None,
                 target_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetPortArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchControllerTrafficSnifferArgs.__new__(SwitchControllerTrafficSnifferArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["erspan_ip"] = erspan_ip
            __props__.__dict__["mode"] = mode
            __props__.__dict__["target_ips"] = target_ips
            __props__.__dict__["target_macs"] = target_macs
            __props__.__dict__["target_ports"] = target_ports
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerTrafficSniffer, __self__).__init__(
            'fortios:index/switchControllerTrafficSniffer:SwitchControllerTrafficSniffer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            erspan_ip: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            target_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetIpArgs']]]]] = None,
            target_macs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetMacArgs']]]]] = None,
            target_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerTrafficSnifferTargetPortArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerTrafficSniffer':
        """
        Get an existing SwitchControllerTrafficSniffer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerTrafficSnifferState.__new__(_SwitchControllerTrafficSnifferState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["erspan_ip"] = erspan_ip
        __props__.__dict__["mode"] = mode
        __props__.__dict__["target_ips"] = target_ips
        __props__.__dict__["target_macs"] = target_macs
        __props__.__dict__["target_ports"] = target_ports
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerTrafficSniffer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="erspanIp")
    def erspan_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "erspan_ip")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="targetIps")
    def target_ips(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerTrafficSnifferTargetIp']]]:
        return pulumi.get(self, "target_ips")

    @property
    @pulumi.getter(name="targetMacs")
    def target_macs(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerTrafficSnifferTargetMac']]]:
        return pulumi.get(self, "target_macs")

    @property
    @pulumi.getter(name="targetPorts")
    def target_ports(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerTrafficSnifferTargetPort']]]:
        return pulumi.get(self, "target_ports")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

