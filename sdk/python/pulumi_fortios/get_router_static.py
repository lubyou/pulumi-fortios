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

__all__ = [
    'GetRouterStaticResult',
    'AwaitableGetRouterStaticResult',
    'get_router_static',
    'get_router_static_output',
]

@pulumi.output_type
class GetRouterStaticResult:
    """
    A collection of values returned by GetRouterStatic.
    """
    def __init__(__self__, bfd=None, blackhole=None, comment=None, device=None, distance=None, dst=None, dstaddr=None, dynamic_gateway=None, gateway=None, id=None, internet_service=None, internet_service_custom=None, link_monitor_exempt=None, priority=None, sdwan=None, sdwan_zones=None, seq_num=None, src=None, status=None, vdomparam=None, virtual_wan_link=None, vrf=None, weight=None):
        if bfd and not isinstance(bfd, str):
            raise TypeError("Expected argument 'bfd' to be a str")
        pulumi.set(__self__, "bfd", bfd)
        if blackhole and not isinstance(blackhole, str):
            raise TypeError("Expected argument 'blackhole' to be a str")
        pulumi.set(__self__, "blackhole", blackhole)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if distance and not isinstance(distance, int):
            raise TypeError("Expected argument 'distance' to be a int")
        pulumi.set(__self__, "distance", distance)
        if dst and not isinstance(dst, str):
            raise TypeError("Expected argument 'dst' to be a str")
        pulumi.set(__self__, "dst", dst)
        if dstaddr and not isinstance(dstaddr, str):
            raise TypeError("Expected argument 'dstaddr' to be a str")
        pulumi.set(__self__, "dstaddr", dstaddr)
        if dynamic_gateway and not isinstance(dynamic_gateway, str):
            raise TypeError("Expected argument 'dynamic_gateway' to be a str")
        pulumi.set(__self__, "dynamic_gateway", dynamic_gateway)
        if gateway and not isinstance(gateway, str):
            raise TypeError("Expected argument 'gateway' to be a str")
        pulumi.set(__self__, "gateway", gateway)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internet_service and not isinstance(internet_service, int):
            raise TypeError("Expected argument 'internet_service' to be a int")
        pulumi.set(__self__, "internet_service", internet_service)
        if internet_service_custom and not isinstance(internet_service_custom, str):
            raise TypeError("Expected argument 'internet_service_custom' to be a str")
        pulumi.set(__self__, "internet_service_custom", internet_service_custom)
        if link_monitor_exempt and not isinstance(link_monitor_exempt, str):
            raise TypeError("Expected argument 'link_monitor_exempt' to be a str")
        pulumi.set(__self__, "link_monitor_exempt", link_monitor_exempt)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if sdwan and not isinstance(sdwan, str):
            raise TypeError("Expected argument 'sdwan' to be a str")
        pulumi.set(__self__, "sdwan", sdwan)
        if sdwan_zones and not isinstance(sdwan_zones, list):
            raise TypeError("Expected argument 'sdwan_zones' to be a list")
        pulumi.set(__self__, "sdwan_zones", sdwan_zones)
        if seq_num and not isinstance(seq_num, int):
            raise TypeError("Expected argument 'seq_num' to be a int")
        pulumi.set(__self__, "seq_num", seq_num)
        if src and not isinstance(src, str):
            raise TypeError("Expected argument 'src' to be a str")
        pulumi.set(__self__, "src", src)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if virtual_wan_link and not isinstance(virtual_wan_link, str):
            raise TypeError("Expected argument 'virtual_wan_link' to be a str")
        pulumi.set(__self__, "virtual_wan_link", virtual_wan_link)
        if vrf and not isinstance(vrf, int):
            raise TypeError("Expected argument 'vrf' to be a int")
        pulumi.set(__self__, "vrf", vrf)
        if weight and not isinstance(weight, int):
            raise TypeError("Expected argument 'weight' to be a int")
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def bfd(self) -> str:
        return pulumi.get(self, "bfd")

    @property
    @pulumi.getter
    def blackhole(self) -> str:
        return pulumi.get(self, "blackhole")

    @property
    @pulumi.getter
    def comment(self) -> str:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def device(self) -> str:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def distance(self) -> int:
        return pulumi.get(self, "distance")

    @property
    @pulumi.getter
    def dst(self) -> str:
        return pulumi.get(self, "dst")

    @property
    @pulumi.getter
    def dstaddr(self) -> str:
        return pulumi.get(self, "dstaddr")

    @property
    @pulumi.getter(name="dynamicGateway")
    def dynamic_gateway(self) -> str:
        return pulumi.get(self, "dynamic_gateway")

    @property
    @pulumi.getter
    def gateway(self) -> str:
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="internetService")
    def internet_service(self) -> int:
        return pulumi.get(self, "internet_service")

    @property
    @pulumi.getter(name="internetServiceCustom")
    def internet_service_custom(self) -> str:
        return pulumi.get(self, "internet_service_custom")

    @property
    @pulumi.getter(name="linkMonitorExempt")
    def link_monitor_exempt(self) -> str:
        return pulumi.get(self, "link_monitor_exempt")

    @property
    @pulumi.getter
    def priority(self) -> int:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def sdwan(self) -> str:
        return pulumi.get(self, "sdwan")

    @property
    @pulumi.getter(name="sdwanZones")
    def sdwan_zones(self) -> Sequence['outputs.GetRouterStaticSdwanZoneResult']:
        return pulumi.get(self, "sdwan_zones")

    @property
    @pulumi.getter(name="seqNum")
    def seq_num(self) -> int:
        return pulumi.get(self, "seq_num")

    @property
    @pulumi.getter
    def src(self) -> str:
        return pulumi.get(self, "src")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="virtualWanLink")
    def virtual_wan_link(self) -> str:
        return pulumi.get(self, "virtual_wan_link")

    @property
    @pulumi.getter
    def vrf(self) -> int:
        return pulumi.get(self, "vrf")

    @property
    @pulumi.getter
    def weight(self) -> int:
        return pulumi.get(self, "weight")


class AwaitableGetRouterStaticResult(GetRouterStaticResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterStaticResult(
            bfd=self.bfd,
            blackhole=self.blackhole,
            comment=self.comment,
            device=self.device,
            distance=self.distance,
            dst=self.dst,
            dstaddr=self.dstaddr,
            dynamic_gateway=self.dynamic_gateway,
            gateway=self.gateway,
            id=self.id,
            internet_service=self.internet_service,
            internet_service_custom=self.internet_service_custom,
            link_monitor_exempt=self.link_monitor_exempt,
            priority=self.priority,
            sdwan=self.sdwan,
            sdwan_zones=self.sdwan_zones,
            seq_num=self.seq_num,
            src=self.src,
            status=self.status,
            vdomparam=self.vdomparam,
            virtual_wan_link=self.virtual_wan_link,
            vrf=self.vrf,
            weight=self.weight)


def get_router_static(seq_num: Optional[int] = None,
                      vdomparam: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterStaticResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['seqNum'] = seq_num
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterStatic:GetRouterStatic', __args__, opts=opts, typ=GetRouterStaticResult).value

    return AwaitableGetRouterStaticResult(
        bfd=__ret__.bfd,
        blackhole=__ret__.blackhole,
        comment=__ret__.comment,
        device=__ret__.device,
        distance=__ret__.distance,
        dst=__ret__.dst,
        dstaddr=__ret__.dstaddr,
        dynamic_gateway=__ret__.dynamic_gateway,
        gateway=__ret__.gateway,
        id=__ret__.id,
        internet_service=__ret__.internet_service,
        internet_service_custom=__ret__.internet_service_custom,
        link_monitor_exempt=__ret__.link_monitor_exempt,
        priority=__ret__.priority,
        sdwan=__ret__.sdwan,
        sdwan_zones=__ret__.sdwan_zones,
        seq_num=__ret__.seq_num,
        src=__ret__.src,
        status=__ret__.status,
        vdomparam=__ret__.vdomparam,
        virtual_wan_link=__ret__.virtual_wan_link,
        vrf=__ret__.vrf,
        weight=__ret__.weight)


@_utilities.lift_output_func(get_router_static)
def get_router_static_output(seq_num: Optional[pulumi.Input[int]] = None,
                             vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterStaticResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
