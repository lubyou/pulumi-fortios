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

__all__ = ['FirewallMulticastPolicyArgs', 'FirewallMulticastPolicy']

@pulumi.input_type
class FirewallMulticastPolicyArgs:
    def __init__(__self__, *,
                 dstaddrs: pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicyDstaddrArgs']]],
                 dstintf: pulumi.Input[str],
                 srcaddrs: pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicySrcaddrArgs']]],
                 srcintf: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 auto_asic_offload: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dnat: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 end_port: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 logtraffic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[int]] = None,
                 snat: Optional[pulumi.Input[str]] = None,
                 snat_ip: Optional[pulumi.Input[str]] = None,
                 start_port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 traffic_shaper: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallMulticastPolicy resource.
        """
        pulumi.set(__self__, "dstaddrs", dstaddrs)
        pulumi.set(__self__, "dstintf", dstintf)
        pulumi.set(__self__, "srcaddrs", srcaddrs)
        pulumi.set(__self__, "srcintf", srcintf)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if auto_asic_offload is not None:
            pulumi.set(__self__, "auto_asic_offload", auto_asic_offload)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dnat is not None:
            pulumi.set(__self__, "dnat", dnat)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if end_port is not None:
            pulumi.set(__self__, "end_port", end_port)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if logtraffic is not None:
            pulumi.set(__self__, "logtraffic", logtraffic)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if snat is not None:
            pulumi.set(__self__, "snat", snat)
        if snat_ip is not None:
            pulumi.set(__self__, "snat_ip", snat_ip)
        if start_port is not None:
            pulumi.set(__self__, "start_port", start_port)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if traffic_shaper is not None:
            pulumi.set(__self__, "traffic_shaper", traffic_shaper)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def dstaddrs(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicyDstaddrArgs']]]:
        return pulumi.get(self, "dstaddrs")

    @dstaddrs.setter
    def dstaddrs(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicyDstaddrArgs']]]):
        pulumi.set(self, "dstaddrs", value)

    @property
    @pulumi.getter
    def dstintf(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dstintf")

    @dstintf.setter
    def dstintf(self, value: pulumi.Input[str]):
        pulumi.set(self, "dstintf", value)

    @property
    @pulumi.getter
    def srcaddrs(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicySrcaddrArgs']]]:
        return pulumi.get(self, "srcaddrs")

    @srcaddrs.setter
    def srcaddrs(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicySrcaddrArgs']]]):
        pulumi.set(self, "srcaddrs", value)

    @property
    @pulumi.getter
    def srcintf(self) -> pulumi.Input[str]:
        return pulumi.get(self, "srcintf")

    @srcintf.setter
    def srcintf(self, value: pulumi.Input[str]):
        pulumi.set(self, "srcintf", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="autoAsicOffload")
    def auto_asic_offload(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_asic_offload")

    @auto_asic_offload.setter
    def auto_asic_offload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_asic_offload", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def dnat(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dnat")

    @dnat.setter
    def dnat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dnat", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end_port")

    @end_port.setter
    def end_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_port", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def logtraffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "logtraffic")

    @logtraffic.setter
    def logtraffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logtraffic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def snat(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "snat")

    @snat.setter
    def snat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat", value)

    @property
    @pulumi.getter(name="snatIp")
    def snat_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "snat_ip")

    @snat_ip.setter
    def snat_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_ip", value)

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start_port")

    @start_port.setter
    def start_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_port", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="trafficShaper")
    def traffic_shaper(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "traffic_shaper")

    @traffic_shaper.setter
    def traffic_shaper(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_shaper", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallMulticastPolicyState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 auto_asic_offload: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dnat: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicyDstaddrArgs']]]] = None,
                 dstintf: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 end_port: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 logtraffic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[int]] = None,
                 snat: Optional[pulumi.Input[str]] = None,
                 snat_ip: Optional[pulumi.Input[str]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicySrcaddrArgs']]]] = None,
                 srcintf: Optional[pulumi.Input[str]] = None,
                 start_port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 traffic_shaper: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallMulticastPolicy resources.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if auto_asic_offload is not None:
            pulumi.set(__self__, "auto_asic_offload", auto_asic_offload)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dnat is not None:
            pulumi.set(__self__, "dnat", dnat)
        if dstaddrs is not None:
            pulumi.set(__self__, "dstaddrs", dstaddrs)
        if dstintf is not None:
            pulumi.set(__self__, "dstintf", dstintf)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if end_port is not None:
            pulumi.set(__self__, "end_port", end_port)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if logtraffic is not None:
            pulumi.set(__self__, "logtraffic", logtraffic)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if snat is not None:
            pulumi.set(__self__, "snat", snat)
        if snat_ip is not None:
            pulumi.set(__self__, "snat_ip", snat_ip)
        if srcaddrs is not None:
            pulumi.set(__self__, "srcaddrs", srcaddrs)
        if srcintf is not None:
            pulumi.set(__self__, "srcintf", srcintf)
        if start_port is not None:
            pulumi.set(__self__, "start_port", start_port)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if traffic_shaper is not None:
            pulumi.set(__self__, "traffic_shaper", traffic_shaper)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="autoAsicOffload")
    def auto_asic_offload(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_asic_offload")

    @auto_asic_offload.setter
    def auto_asic_offload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_asic_offload", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def dnat(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dnat")

    @dnat.setter
    def dnat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dnat", value)

    @property
    @pulumi.getter
    def dstaddrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicyDstaddrArgs']]]]:
        return pulumi.get(self, "dstaddrs")

    @dstaddrs.setter
    def dstaddrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicyDstaddrArgs']]]]):
        pulumi.set(self, "dstaddrs", value)

    @property
    @pulumi.getter
    def dstintf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dstintf")

    @dstintf.setter
    def dstintf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dstintf", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end_port")

    @end_port.setter
    def end_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_port", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def logtraffic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "logtraffic")

    @logtraffic.setter
    def logtraffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logtraffic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def snat(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "snat")

    @snat.setter
    def snat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat", value)

    @property
    @pulumi.getter(name="snatIp")
    def snat_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "snat_ip")

    @snat_ip.setter
    def snat_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_ip", value)

    @property
    @pulumi.getter
    def srcaddrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicySrcaddrArgs']]]]:
        return pulumi.get(self, "srcaddrs")

    @srcaddrs.setter
    def srcaddrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastPolicySrcaddrArgs']]]]):
        pulumi.set(self, "srcaddrs", value)

    @property
    @pulumi.getter
    def srcintf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "srcintf")

    @srcintf.setter
    def srcintf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "srcintf", value)

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start_port")

    @start_port.setter
    def start_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_port", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="trafficShaper")
    def traffic_shaper(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "traffic_shaper")

    @traffic_shaper.setter
    def traffic_shaper(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_shaper", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallMulticastPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 auto_asic_offload: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dnat: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastPolicyDstaddrArgs']]]]] = None,
                 dstintf: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 end_port: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 logtraffic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[int]] = None,
                 snat: Optional[pulumi.Input[str]] = None,
                 snat_ip: Optional[pulumi.Input[str]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastPolicySrcaddrArgs']]]]] = None,
                 srcintf: Optional[pulumi.Input[str]] = None,
                 start_port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 traffic_shaper: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallMulticastPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallMulticastPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallMulticastPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallMulticastPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallMulticastPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 auto_asic_offload: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dnat: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastPolicyDstaddrArgs']]]]] = None,
                 dstintf: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 end_port: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 logtraffic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[int]] = None,
                 snat: Optional[pulumi.Input[str]] = None,
                 snat_ip: Optional[pulumi.Input[str]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastPolicySrcaddrArgs']]]]] = None,
                 srcintf: Optional[pulumi.Input[str]] = None,
                 start_port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 traffic_shaper: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallMulticastPolicyArgs.__new__(FirewallMulticastPolicyArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["auto_asic_offload"] = auto_asic_offload
            __props__.__dict__["comments"] = comments
            __props__.__dict__["dnat"] = dnat
            if dstaddrs is None and not opts.urn:
                raise TypeError("Missing required property 'dstaddrs'")
            __props__.__dict__["dstaddrs"] = dstaddrs
            if dstintf is None and not opts.urn:
                raise TypeError("Missing required property 'dstintf'")
            __props__.__dict__["dstintf"] = dstintf
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["end_port"] = end_port
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["logtraffic"] = logtraffic
            __props__.__dict__["name"] = name
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["snat"] = snat
            __props__.__dict__["snat_ip"] = snat_ip
            if srcaddrs is None and not opts.urn:
                raise TypeError("Missing required property 'srcaddrs'")
            __props__.__dict__["srcaddrs"] = srcaddrs
            if srcintf is None and not opts.urn:
                raise TypeError("Missing required property 'srcintf'")
            __props__.__dict__["srcintf"] = srcintf
            __props__.__dict__["start_port"] = start_port
            __props__.__dict__["status"] = status
            __props__.__dict__["traffic_shaper"] = traffic_shaper
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallMulticastPolicy, __self__).__init__(
            'fortios:index/firewallMulticastPolicy:FirewallMulticastPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            auto_asic_offload: Optional[pulumi.Input[str]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            dnat: Optional[pulumi.Input[str]] = None,
            dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastPolicyDstaddrArgs']]]]] = None,
            dstintf: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            end_port: Optional[pulumi.Input[int]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            logtraffic: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[int]] = None,
            snat: Optional[pulumi.Input[str]] = None,
            snat_ip: Optional[pulumi.Input[str]] = None,
            srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastPolicySrcaddrArgs']]]]] = None,
            srcintf: Optional[pulumi.Input[str]] = None,
            start_port: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            traffic_shaper: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallMulticastPolicy':
        """
        Get an existing FirewallMulticastPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallMulticastPolicyState.__new__(_FirewallMulticastPolicyState)

        __props__.__dict__["action"] = action
        __props__.__dict__["auto_asic_offload"] = auto_asic_offload
        __props__.__dict__["comments"] = comments
        __props__.__dict__["dnat"] = dnat
        __props__.__dict__["dstaddrs"] = dstaddrs
        __props__.__dict__["dstintf"] = dstintf
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["end_port"] = end_port
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["logtraffic"] = logtraffic
        __props__.__dict__["name"] = name
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["snat"] = snat
        __props__.__dict__["snat_ip"] = snat_ip
        __props__.__dict__["srcaddrs"] = srcaddrs
        __props__.__dict__["srcintf"] = srcintf
        __props__.__dict__["start_port"] = start_port
        __props__.__dict__["status"] = status
        __props__.__dict__["traffic_shaper"] = traffic_shaper
        __props__.__dict__["uuid"] = uuid
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallMulticastPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="autoAsicOffload")
    def auto_asic_offload(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auto_asic_offload")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def dnat(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dnat")

    @property
    @pulumi.getter
    def dstaddrs(self) -> pulumi.Output[Sequence['outputs.FirewallMulticastPolicyDstaddr']]:
        return pulumi.get(self, "dstaddrs")

    @property
    @pulumi.getter
    def dstintf(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dstintf")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "end_port")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def logtraffic(self) -> pulumi.Output[str]:
        return pulumi.get(self, "logtraffic")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[int]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def snat(self) -> pulumi.Output[str]:
        return pulumi.get(self, "snat")

    @property
    @pulumi.getter(name="snatIp")
    def snat_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "snat_ip")

    @property
    @pulumi.getter
    def srcaddrs(self) -> pulumi.Output[Sequence['outputs.FirewallMulticastPolicySrcaddr']]:
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter
    def srcintf(self) -> pulumi.Output[str]:
        return pulumi.get(self, "srcintf")

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "start_port")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trafficShaper")
    def traffic_shaper(self) -> pulumi.Output[str]:
        return pulumi.get(self, "traffic_shaper")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

