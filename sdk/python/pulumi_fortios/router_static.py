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

__all__ = ['RouterStaticArgs', 'RouterStatic']

@pulumi.input_type
class RouterStaticArgs:
    def __init__(__self__, *,
                 bfd: Optional[pulumi.Input[str]] = None,
                 blackhole: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 dst: Optional[pulumi.Input[str]] = None,
                 dstaddr: Optional[pulumi.Input[str]] = None,
                 dynamic_gateway: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 internet_service: Optional[pulumi.Input[int]] = None,
                 internet_service_custom: Optional[pulumi.Input[str]] = None,
                 link_monitor_exempt: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 sdwan_zones: Optional[pulumi.Input[Sequence[pulumi.Input['RouterStaticSdwanZoneArgs']]]] = None,
                 seq_num: Optional[pulumi.Input[int]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 virtual_wan_link: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[int]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RouterStatic resource.
        """
        if bfd is not None:
            pulumi.set(__self__, "bfd", bfd)
        if blackhole is not None:
            pulumi.set(__self__, "blackhole", blackhole)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if dst is not None:
            pulumi.set(__self__, "dst", dst)
        if dstaddr is not None:
            pulumi.set(__self__, "dstaddr", dstaddr)
        if dynamic_gateway is not None:
            pulumi.set(__self__, "dynamic_gateway", dynamic_gateway)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if internet_service is not None:
            pulumi.set(__self__, "internet_service", internet_service)
        if internet_service_custom is not None:
            pulumi.set(__self__, "internet_service_custom", internet_service_custom)
        if link_monitor_exempt is not None:
            pulumi.set(__self__, "link_monitor_exempt", link_monitor_exempt)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if sdwan is not None:
            pulumi.set(__self__, "sdwan", sdwan)
        if sdwan_zones is not None:
            pulumi.set(__self__, "sdwan_zones", sdwan_zones)
        if seq_num is not None:
            pulumi.set(__self__, "seq_num", seq_num)
        if src is not None:
            pulumi.set(__self__, "src", src)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if virtual_wan_link is not None:
            pulumi.set(__self__, "virtual_wan_link", virtual_wan_link)
        if vrf is not None:
            pulumi.set(__self__, "vrf", vrf)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def bfd(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bfd")

    @bfd.setter
    def bfd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bfd", value)

    @property
    @pulumi.getter
    def blackhole(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "blackhole")

    @blackhole.setter
    def blackhole(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blackhole", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def dst(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dst")

    @dst.setter
    def dst(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst", value)

    @property
    @pulumi.getter
    def dstaddr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dstaddr")

    @dstaddr.setter
    def dstaddr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dstaddr", value)

    @property
    @pulumi.getter(name="dynamicGateway")
    def dynamic_gateway(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_gateway")

    @dynamic_gateway.setter
    def dynamic_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_gateway", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="internetService")
    def internet_service(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "internet_service")

    @internet_service.setter
    def internet_service(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_service", value)

    @property
    @pulumi.getter(name="internetServiceCustom")
    def internet_service_custom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "internet_service_custom")

    @internet_service_custom.setter
    def internet_service_custom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_service_custom", value)

    @property
    @pulumi.getter(name="linkMonitorExempt")
    def link_monitor_exempt(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "link_monitor_exempt")

    @link_monitor_exempt.setter
    def link_monitor_exempt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_monitor_exempt", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def sdwan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sdwan")

    @sdwan.setter
    def sdwan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdwan", value)

    @property
    @pulumi.getter(name="sdwanZones")
    def sdwan_zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterStaticSdwanZoneArgs']]]]:
        return pulumi.get(self, "sdwan_zones")

    @sdwan_zones.setter
    def sdwan_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterStaticSdwanZoneArgs']]]]):
        pulumi.set(self, "sdwan_zones", value)

    @property
    @pulumi.getter(name="seqNum")
    def seq_num(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "seq_num")

    @seq_num.setter
    def seq_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "seq_num", value)

    @property
    @pulumi.getter
    def src(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "src")

    @src.setter
    def src(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="virtualWanLink")
    def virtual_wan_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "virtual_wan_link")

    @virtual_wan_link.setter
    def virtual_wan_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_wan_link", value)

    @property
    @pulumi.getter
    def vrf(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vrf")

    @vrf.setter
    def vrf(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vrf", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _RouterStaticState:
    def __init__(__self__, *,
                 bfd: Optional[pulumi.Input[str]] = None,
                 blackhole: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 dst: Optional[pulumi.Input[str]] = None,
                 dstaddr: Optional[pulumi.Input[str]] = None,
                 dynamic_gateway: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 internet_service: Optional[pulumi.Input[int]] = None,
                 internet_service_custom: Optional[pulumi.Input[str]] = None,
                 link_monitor_exempt: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 sdwan_zones: Optional[pulumi.Input[Sequence[pulumi.Input['RouterStaticSdwanZoneArgs']]]] = None,
                 seq_num: Optional[pulumi.Input[int]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 virtual_wan_link: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[int]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RouterStatic resources.
        """
        if bfd is not None:
            pulumi.set(__self__, "bfd", bfd)
        if blackhole is not None:
            pulumi.set(__self__, "blackhole", blackhole)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if dst is not None:
            pulumi.set(__self__, "dst", dst)
        if dstaddr is not None:
            pulumi.set(__self__, "dstaddr", dstaddr)
        if dynamic_gateway is not None:
            pulumi.set(__self__, "dynamic_gateway", dynamic_gateway)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if internet_service is not None:
            pulumi.set(__self__, "internet_service", internet_service)
        if internet_service_custom is not None:
            pulumi.set(__self__, "internet_service_custom", internet_service_custom)
        if link_monitor_exempt is not None:
            pulumi.set(__self__, "link_monitor_exempt", link_monitor_exempt)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if sdwan is not None:
            pulumi.set(__self__, "sdwan", sdwan)
        if sdwan_zones is not None:
            pulumi.set(__self__, "sdwan_zones", sdwan_zones)
        if seq_num is not None:
            pulumi.set(__self__, "seq_num", seq_num)
        if src is not None:
            pulumi.set(__self__, "src", src)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if virtual_wan_link is not None:
            pulumi.set(__self__, "virtual_wan_link", virtual_wan_link)
        if vrf is not None:
            pulumi.set(__self__, "vrf", vrf)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def bfd(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bfd")

    @bfd.setter
    def bfd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bfd", value)

    @property
    @pulumi.getter
    def blackhole(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "blackhole")

    @blackhole.setter
    def blackhole(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blackhole", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def dst(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dst")

    @dst.setter
    def dst(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst", value)

    @property
    @pulumi.getter
    def dstaddr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dstaddr")

    @dstaddr.setter
    def dstaddr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dstaddr", value)

    @property
    @pulumi.getter(name="dynamicGateway")
    def dynamic_gateway(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_gateway")

    @dynamic_gateway.setter
    def dynamic_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_gateway", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="internetService")
    def internet_service(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "internet_service")

    @internet_service.setter
    def internet_service(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_service", value)

    @property
    @pulumi.getter(name="internetServiceCustom")
    def internet_service_custom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "internet_service_custom")

    @internet_service_custom.setter
    def internet_service_custom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_service_custom", value)

    @property
    @pulumi.getter(name="linkMonitorExempt")
    def link_monitor_exempt(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "link_monitor_exempt")

    @link_monitor_exempt.setter
    def link_monitor_exempt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_monitor_exempt", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def sdwan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sdwan")

    @sdwan.setter
    def sdwan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdwan", value)

    @property
    @pulumi.getter(name="sdwanZones")
    def sdwan_zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterStaticSdwanZoneArgs']]]]:
        return pulumi.get(self, "sdwan_zones")

    @sdwan_zones.setter
    def sdwan_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterStaticSdwanZoneArgs']]]]):
        pulumi.set(self, "sdwan_zones", value)

    @property
    @pulumi.getter(name="seqNum")
    def seq_num(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "seq_num")

    @seq_num.setter
    def seq_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "seq_num", value)

    @property
    @pulumi.getter
    def src(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "src")

    @src.setter
    def src(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="virtualWanLink")
    def virtual_wan_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "virtual_wan_link")

    @virtual_wan_link.setter
    def virtual_wan_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_wan_link", value)

    @property
    @pulumi.getter
    def vrf(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vrf")

    @vrf.setter
    def vrf(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vrf", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


class RouterStatic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bfd: Optional[pulumi.Input[str]] = None,
                 blackhole: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 dst: Optional[pulumi.Input[str]] = None,
                 dstaddr: Optional[pulumi.Input[str]] = None,
                 dynamic_gateway: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 internet_service: Optional[pulumi.Input[int]] = None,
                 internet_service_custom: Optional[pulumi.Input[str]] = None,
                 link_monitor_exempt: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 sdwan_zones: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterStaticSdwanZoneArgs']]]]] = None,
                 seq_num: Optional[pulumi.Input[int]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 virtual_wan_link: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[int]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a RouterStatic resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RouterStaticArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RouterStatic resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RouterStaticArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterStaticArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bfd: Optional[pulumi.Input[str]] = None,
                 blackhole: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 dst: Optional[pulumi.Input[str]] = None,
                 dstaddr: Optional[pulumi.Input[str]] = None,
                 dynamic_gateway: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 internet_service: Optional[pulumi.Input[int]] = None,
                 internet_service_custom: Optional[pulumi.Input[str]] = None,
                 link_monitor_exempt: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 sdwan_zones: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterStaticSdwanZoneArgs']]]]] = None,
                 seq_num: Optional[pulumi.Input[int]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 virtual_wan_link: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[int]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouterStaticArgs.__new__(RouterStaticArgs)

            __props__.__dict__["bfd"] = bfd
            __props__.__dict__["blackhole"] = blackhole
            __props__.__dict__["comment"] = comment
            __props__.__dict__["device"] = device
            __props__.__dict__["distance"] = distance
            __props__.__dict__["dst"] = dst
            __props__.__dict__["dstaddr"] = dstaddr
            __props__.__dict__["dynamic_gateway"] = dynamic_gateway
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["gateway"] = gateway
            __props__.__dict__["internet_service"] = internet_service
            __props__.__dict__["internet_service_custom"] = internet_service_custom
            __props__.__dict__["link_monitor_exempt"] = link_monitor_exempt
            __props__.__dict__["priority"] = priority
            __props__.__dict__["sdwan"] = sdwan
            __props__.__dict__["sdwan_zones"] = sdwan_zones
            __props__.__dict__["seq_num"] = seq_num
            __props__.__dict__["src"] = src
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["virtual_wan_link"] = virtual_wan_link
            __props__.__dict__["vrf"] = vrf
            __props__.__dict__["weight"] = weight
        super(RouterStatic, __self__).__init__(
            'fortios:index/routerStatic:RouterStatic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bfd: Optional[pulumi.Input[str]] = None,
            blackhole: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            distance: Optional[pulumi.Input[int]] = None,
            dst: Optional[pulumi.Input[str]] = None,
            dstaddr: Optional[pulumi.Input[str]] = None,
            dynamic_gateway: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            internet_service: Optional[pulumi.Input[int]] = None,
            internet_service_custom: Optional[pulumi.Input[str]] = None,
            link_monitor_exempt: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            sdwan: Optional[pulumi.Input[str]] = None,
            sdwan_zones: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterStaticSdwanZoneArgs']]]]] = None,
            seq_num: Optional[pulumi.Input[int]] = None,
            src: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            virtual_wan_link: Optional[pulumi.Input[str]] = None,
            vrf: Optional[pulumi.Input[int]] = None,
            weight: Optional[pulumi.Input[int]] = None) -> 'RouterStatic':
        """
        Get an existing RouterStatic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouterStaticState.__new__(_RouterStaticState)

        __props__.__dict__["bfd"] = bfd
        __props__.__dict__["blackhole"] = blackhole
        __props__.__dict__["comment"] = comment
        __props__.__dict__["device"] = device
        __props__.__dict__["distance"] = distance
        __props__.__dict__["dst"] = dst
        __props__.__dict__["dstaddr"] = dstaddr
        __props__.__dict__["dynamic_gateway"] = dynamic_gateway
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["internet_service"] = internet_service
        __props__.__dict__["internet_service_custom"] = internet_service_custom
        __props__.__dict__["link_monitor_exempt"] = link_monitor_exempt
        __props__.__dict__["priority"] = priority
        __props__.__dict__["sdwan"] = sdwan
        __props__.__dict__["sdwan_zones"] = sdwan_zones
        __props__.__dict__["seq_num"] = seq_num
        __props__.__dict__["src"] = src
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["virtual_wan_link"] = virtual_wan_link
        __props__.__dict__["vrf"] = vrf
        __props__.__dict__["weight"] = weight
        return RouterStatic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bfd(self) -> pulumi.Output[str]:
        return pulumi.get(self, "bfd")

    @property
    @pulumi.getter
    def blackhole(self) -> pulumi.Output[str]:
        return pulumi.get(self, "blackhole")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def distance(self) -> pulumi.Output[int]:
        return pulumi.get(self, "distance")

    @property
    @pulumi.getter
    def dst(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dst")

    @property
    @pulumi.getter
    def dstaddr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dstaddr")

    @property
    @pulumi.getter(name="dynamicGateway")
    def dynamic_gateway(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dynamic_gateway")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="internetService")
    def internet_service(self) -> pulumi.Output[int]:
        return pulumi.get(self, "internet_service")

    @property
    @pulumi.getter(name="internetServiceCustom")
    def internet_service_custom(self) -> pulumi.Output[str]:
        return pulumi.get(self, "internet_service_custom")

    @property
    @pulumi.getter(name="linkMonitorExempt")
    def link_monitor_exempt(self) -> pulumi.Output[str]:
        return pulumi.get(self, "link_monitor_exempt")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def sdwan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sdwan")

    @property
    @pulumi.getter(name="sdwanZones")
    def sdwan_zones(self) -> pulumi.Output[Optional[Sequence['outputs.RouterStaticSdwanZone']]]:
        return pulumi.get(self, "sdwan_zones")

    @property
    @pulumi.getter(name="seqNum")
    def seq_num(self) -> pulumi.Output[int]:
        return pulumi.get(self, "seq_num")

    @property
    @pulumi.getter
    def src(self) -> pulumi.Output[str]:
        return pulumi.get(self, "src")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="virtualWanLink")
    def virtual_wan_link(self) -> pulumi.Output[str]:
        return pulumi.get(self, "virtual_wan_link")

    @property
    @pulumi.getter
    def vrf(self) -> pulumi.Output[int]:
        return pulumi.get(self, "vrf")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[int]:
        return pulumi.get(self, "weight")

