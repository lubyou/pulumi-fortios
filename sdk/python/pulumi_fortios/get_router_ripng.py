# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetRouterRipngResult',
    'AwaitableGetRouterRipngResult',
    'get_router_ripng',
    'get_router_ripng_output',
]

@pulumi.output_type
class GetRouterRipngResult:
    """
    A collection of values returned by GetRouterRipng.
    """
    def __init__(__self__, aggregate_addresses=None, default_information_originate=None, default_metric=None, distances=None, distribute_lists=None, garbage_timer=None, id=None, interfaces=None, max_out_metric=None, neighbors=None, networks=None, offset_lists=None, passive_interfaces=None, redistributes=None, timeout_timer=None, update_timer=None, vdomparam=None):
        if aggregate_addresses and not isinstance(aggregate_addresses, list):
            raise TypeError("Expected argument 'aggregate_addresses' to be a list")
        pulumi.set(__self__, "aggregate_addresses", aggregate_addresses)
        if default_information_originate and not isinstance(default_information_originate, str):
            raise TypeError("Expected argument 'default_information_originate' to be a str")
        pulumi.set(__self__, "default_information_originate", default_information_originate)
        if default_metric and not isinstance(default_metric, int):
            raise TypeError("Expected argument 'default_metric' to be a int")
        pulumi.set(__self__, "default_metric", default_metric)
        if distances and not isinstance(distances, list):
            raise TypeError("Expected argument 'distances' to be a list")
        pulumi.set(__self__, "distances", distances)
        if distribute_lists and not isinstance(distribute_lists, list):
            raise TypeError("Expected argument 'distribute_lists' to be a list")
        pulumi.set(__self__, "distribute_lists", distribute_lists)
        if garbage_timer and not isinstance(garbage_timer, int):
            raise TypeError("Expected argument 'garbage_timer' to be a int")
        pulumi.set(__self__, "garbage_timer", garbage_timer)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if max_out_metric and not isinstance(max_out_metric, int):
            raise TypeError("Expected argument 'max_out_metric' to be a int")
        pulumi.set(__self__, "max_out_metric", max_out_metric)
        if neighbors and not isinstance(neighbors, list):
            raise TypeError("Expected argument 'neighbors' to be a list")
        pulumi.set(__self__, "neighbors", neighbors)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if offset_lists and not isinstance(offset_lists, list):
            raise TypeError("Expected argument 'offset_lists' to be a list")
        pulumi.set(__self__, "offset_lists", offset_lists)
        if passive_interfaces and not isinstance(passive_interfaces, list):
            raise TypeError("Expected argument 'passive_interfaces' to be a list")
        pulumi.set(__self__, "passive_interfaces", passive_interfaces)
        if redistributes and not isinstance(redistributes, list):
            raise TypeError("Expected argument 'redistributes' to be a list")
        pulumi.set(__self__, "redistributes", redistributes)
        if timeout_timer and not isinstance(timeout_timer, int):
            raise TypeError("Expected argument 'timeout_timer' to be a int")
        pulumi.set(__self__, "timeout_timer", timeout_timer)
        if update_timer and not isinstance(update_timer, int):
            raise TypeError("Expected argument 'update_timer' to be a int")
        pulumi.set(__self__, "update_timer", update_timer)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="aggregateAddresses")
    def aggregate_addresses(self) -> Sequence['outputs.GetRouterRipngAggregateAddressResult']:
        """
        Aggregate address. The structure of `aggregate_address` block is documented below.
        """
        return pulumi.get(self, "aggregate_addresses")

    @property
    @pulumi.getter(name="defaultInformationOriginate")
    def default_information_originate(self) -> str:
        """
        Enable/disable generation of default route.
        """
        return pulumi.get(self, "default_information_originate")

    @property
    @pulumi.getter(name="defaultMetric")
    def default_metric(self) -> int:
        """
        Default metric.
        """
        return pulumi.get(self, "default_metric")

    @property
    @pulumi.getter
    def distances(self) -> Sequence['outputs.GetRouterRipngDistanceResult']:
        """
        Distance (1 - 255).
        """
        return pulumi.get(self, "distances")

    @property
    @pulumi.getter(name="distributeLists")
    def distribute_lists(self) -> Sequence['outputs.GetRouterRipngDistributeListResult']:
        """
        Distribute list. The structure of `distribute_list` block is documented below.
        """
        return pulumi.get(self, "distribute_lists")

    @property
    @pulumi.getter(name="garbageTimer")
    def garbage_timer(self) -> int:
        """
        Garbage timer.
        """
        return pulumi.get(self, "garbage_timer")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetRouterRipngInterfaceResult']:
        """
        Interface name.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="maxOutMetric")
    def max_out_metric(self) -> int:
        """
        Maximum metric allowed to output(0 means 'not set').
        """
        return pulumi.get(self, "max_out_metric")

    @property
    @pulumi.getter
    def neighbors(self) -> Sequence['outputs.GetRouterRipngNeighborResult']:
        """
        neighbor The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetRouterRipngNetworkResult']:
        """
        Network. The structure of `network` block is documented below.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="offsetLists")
    def offset_lists(self) -> Sequence['outputs.GetRouterRipngOffsetListResult']:
        """
        Offset list. The structure of `offset_list` block is documented below.
        """
        return pulumi.get(self, "offset_lists")

    @property
    @pulumi.getter(name="passiveInterfaces")
    def passive_interfaces(self) -> Sequence['outputs.GetRouterRipngPassiveInterfaceResult']:
        """
        Passive interface configuration. The structure of `passive_interface` block is documented below.
        """
        return pulumi.get(self, "passive_interfaces")

    @property
    @pulumi.getter
    def redistributes(self) -> Sequence['outputs.GetRouterRipngRedistributeResult']:
        """
        Redistribute configuration. The structure of `redistribute` block is documented below.
        """
        return pulumi.get(self, "redistributes")

    @property
    @pulumi.getter(name="timeoutTimer")
    def timeout_timer(self) -> int:
        """
        Timeout timer.
        """
        return pulumi.get(self, "timeout_timer")

    @property
    @pulumi.getter(name="updateTimer")
    def update_timer(self) -> int:
        """
        Update timer.
        """
        return pulumi.get(self, "update_timer")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterRipngResult(GetRouterRipngResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterRipngResult(
            aggregate_addresses=self.aggregate_addresses,
            default_information_originate=self.default_information_originate,
            default_metric=self.default_metric,
            distances=self.distances,
            distribute_lists=self.distribute_lists,
            garbage_timer=self.garbage_timer,
            id=self.id,
            interfaces=self.interfaces,
            max_out_metric=self.max_out_metric,
            neighbors=self.neighbors,
            networks=self.networks,
            offset_lists=self.offset_lists,
            passive_interfaces=self.passive_interfaces,
            redistributes=self.redistributes,
            timeout_timer=self.timeout_timer,
            update_timer=self.update_timer,
            vdomparam=self.vdomparam)


def get_router_ripng(vdomparam: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterRipngResult:
    """
    Use this data source to get information on fortios router ripng


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterRipng:GetRouterRipng', __args__, opts=opts, typ=GetRouterRipngResult).value

    return AwaitableGetRouterRipngResult(
        aggregate_addresses=__ret__.aggregate_addresses,
        default_information_originate=__ret__.default_information_originate,
        default_metric=__ret__.default_metric,
        distances=__ret__.distances,
        distribute_lists=__ret__.distribute_lists,
        garbage_timer=__ret__.garbage_timer,
        id=__ret__.id,
        interfaces=__ret__.interfaces,
        max_out_metric=__ret__.max_out_metric,
        neighbors=__ret__.neighbors,
        networks=__ret__.networks,
        offset_lists=__ret__.offset_lists,
        passive_interfaces=__ret__.passive_interfaces,
        redistributes=__ret__.redistributes,
        timeout_timer=__ret__.timeout_timer,
        update_timer=__ret__.update_timer,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_router_ripng)
def get_router_ripng_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterRipngResult]:
    """
    Use this data source to get information on fortios router ripng


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
