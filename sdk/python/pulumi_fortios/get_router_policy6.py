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
    'GetRouterPolicy6Result',
    'AwaitableGetRouterPolicy6Result',
    'get_router_policy6',
    'get_router_policy6_output',
]

@pulumi.output_type
class GetRouterPolicy6Result:
    """
    A collection of values returned by GetRouterPolicy6.
    """
    def __init__(__self__, action=None, comments=None, dst=None, dst_negate=None, dstaddrs=None, end_port=None, gateway=None, id=None, input_device=None, input_device_negate=None, internet_service_customs=None, internet_service_ids=None, output_device=None, protocol=None, seq_num=None, src=None, src_negate=None, srcaddrs=None, start_port=None, status=None, tos=None, tos_mask=None, vdomparam=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if dst and not isinstance(dst, str):
            raise TypeError("Expected argument 'dst' to be a str")
        pulumi.set(__self__, "dst", dst)
        if dst_negate and not isinstance(dst_negate, str):
            raise TypeError("Expected argument 'dst_negate' to be a str")
        pulumi.set(__self__, "dst_negate", dst_negate)
        if dstaddrs and not isinstance(dstaddrs, list):
            raise TypeError("Expected argument 'dstaddrs' to be a list")
        pulumi.set(__self__, "dstaddrs", dstaddrs)
        if end_port and not isinstance(end_port, int):
            raise TypeError("Expected argument 'end_port' to be a int")
        pulumi.set(__self__, "end_port", end_port)
        if gateway and not isinstance(gateway, str):
            raise TypeError("Expected argument 'gateway' to be a str")
        pulumi.set(__self__, "gateway", gateway)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if input_device and not isinstance(input_device, str):
            raise TypeError("Expected argument 'input_device' to be a str")
        pulumi.set(__self__, "input_device", input_device)
        if input_device_negate and not isinstance(input_device_negate, str):
            raise TypeError("Expected argument 'input_device_negate' to be a str")
        pulumi.set(__self__, "input_device_negate", input_device_negate)
        if internet_service_customs and not isinstance(internet_service_customs, list):
            raise TypeError("Expected argument 'internet_service_customs' to be a list")
        pulumi.set(__self__, "internet_service_customs", internet_service_customs)
        if internet_service_ids and not isinstance(internet_service_ids, list):
            raise TypeError("Expected argument 'internet_service_ids' to be a list")
        pulumi.set(__self__, "internet_service_ids", internet_service_ids)
        if output_device and not isinstance(output_device, str):
            raise TypeError("Expected argument 'output_device' to be a str")
        pulumi.set(__self__, "output_device", output_device)
        if protocol and not isinstance(protocol, int):
            raise TypeError("Expected argument 'protocol' to be a int")
        pulumi.set(__self__, "protocol", protocol)
        if seq_num and not isinstance(seq_num, int):
            raise TypeError("Expected argument 'seq_num' to be a int")
        pulumi.set(__self__, "seq_num", seq_num)
        if src and not isinstance(src, str):
            raise TypeError("Expected argument 'src' to be a str")
        pulumi.set(__self__, "src", src)
        if src_negate and not isinstance(src_negate, str):
            raise TypeError("Expected argument 'src_negate' to be a str")
        pulumi.set(__self__, "src_negate", src_negate)
        if srcaddrs and not isinstance(srcaddrs, list):
            raise TypeError("Expected argument 'srcaddrs' to be a list")
        pulumi.set(__self__, "srcaddrs", srcaddrs)
        if start_port and not isinstance(start_port, int):
            raise TypeError("Expected argument 'start_port' to be a int")
        pulumi.set(__self__, "start_port", start_port)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tos and not isinstance(tos, str):
            raise TypeError("Expected argument 'tos' to be a str")
        pulumi.set(__self__, "tos", tos)
        if tos_mask and not isinstance(tos_mask, str):
            raise TypeError("Expected argument 'tos_mask' to be a str")
        pulumi.set(__self__, "tos_mask", tos_mask)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def comments(self) -> str:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def dst(self) -> str:
        return pulumi.get(self, "dst")

    @property
    @pulumi.getter(name="dstNegate")
    def dst_negate(self) -> str:
        return pulumi.get(self, "dst_negate")

    @property
    @pulumi.getter
    def dstaddrs(self) -> Sequence['outputs.GetRouterPolicy6DstaddrResult']:
        return pulumi.get(self, "dstaddrs")

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> int:
        return pulumi.get(self, "end_port")

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
    @pulumi.getter(name="inputDevice")
    def input_device(self) -> str:
        return pulumi.get(self, "input_device")

    @property
    @pulumi.getter(name="inputDeviceNegate")
    def input_device_negate(self) -> str:
        return pulumi.get(self, "input_device_negate")

    @property
    @pulumi.getter(name="internetServiceCustoms")
    def internet_service_customs(self) -> Sequence['outputs.GetRouterPolicy6InternetServiceCustomResult']:
        return pulumi.get(self, "internet_service_customs")

    @property
    @pulumi.getter(name="internetServiceIds")
    def internet_service_ids(self) -> Sequence['outputs.GetRouterPolicy6InternetServiceIdResult']:
        return pulumi.get(self, "internet_service_ids")

    @property
    @pulumi.getter(name="outputDevice")
    def output_device(self) -> str:
        return pulumi.get(self, "output_device")

    @property
    @pulumi.getter
    def protocol(self) -> int:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="seqNum")
    def seq_num(self) -> int:
        return pulumi.get(self, "seq_num")

    @property
    @pulumi.getter
    def src(self) -> str:
        return pulumi.get(self, "src")

    @property
    @pulumi.getter(name="srcNegate")
    def src_negate(self) -> str:
        return pulumi.get(self, "src_negate")

    @property
    @pulumi.getter
    def srcaddrs(self) -> Sequence['outputs.GetRouterPolicy6SrcaddrResult']:
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter(name="startPort")
    def start_port(self) -> int:
        return pulumi.get(self, "start_port")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tos(self) -> str:
        return pulumi.get(self, "tos")

    @property
    @pulumi.getter(name="tosMask")
    def tos_mask(self) -> str:
        return pulumi.get(self, "tos_mask")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterPolicy6Result(GetRouterPolicy6Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterPolicy6Result(
            action=self.action,
            comments=self.comments,
            dst=self.dst,
            dst_negate=self.dst_negate,
            dstaddrs=self.dstaddrs,
            end_port=self.end_port,
            gateway=self.gateway,
            id=self.id,
            input_device=self.input_device,
            input_device_negate=self.input_device_negate,
            internet_service_customs=self.internet_service_customs,
            internet_service_ids=self.internet_service_ids,
            output_device=self.output_device,
            protocol=self.protocol,
            seq_num=self.seq_num,
            src=self.src,
            src_negate=self.src_negate,
            srcaddrs=self.srcaddrs,
            start_port=self.start_port,
            status=self.status,
            tos=self.tos,
            tos_mask=self.tos_mask,
            vdomparam=self.vdomparam)


def get_router_policy6(seq_num: Optional[int] = None,
                       vdomparam: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterPolicy6Result:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['seqNum'] = seq_num
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterPolicy6:GetRouterPolicy6', __args__, opts=opts, typ=GetRouterPolicy6Result).value

    return AwaitableGetRouterPolicy6Result(
        action=__ret__.action,
        comments=__ret__.comments,
        dst=__ret__.dst,
        dst_negate=__ret__.dst_negate,
        dstaddrs=__ret__.dstaddrs,
        end_port=__ret__.end_port,
        gateway=__ret__.gateway,
        id=__ret__.id,
        input_device=__ret__.input_device,
        input_device_negate=__ret__.input_device_negate,
        internet_service_customs=__ret__.internet_service_customs,
        internet_service_ids=__ret__.internet_service_ids,
        output_device=__ret__.output_device,
        protocol=__ret__.protocol,
        seq_num=__ret__.seq_num,
        src=__ret__.src,
        src_negate=__ret__.src_negate,
        srcaddrs=__ret__.srcaddrs,
        start_port=__ret__.start_port,
        status=__ret__.status,
        tos=__ret__.tos,
        tos_mask=__ret__.tos_mask,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_router_policy6)
def get_router_policy6_output(seq_num: Optional[pulumi.Input[int]] = None,
                              vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterPolicy6Result]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
