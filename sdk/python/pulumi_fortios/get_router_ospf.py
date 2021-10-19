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
    'GetRouterOspfResult',
    'AwaitableGetRouterOspfResult',
    'get_router_ospf',
]

@pulumi.output_type
class GetRouterOspfResult:
    """
    A collection of values returned by GetRouterOspf.
    """
    def __init__(__self__, abr_type=None, areas=None, auto_cost_ref_bandwidth=None, bfd=None, database_overflow=None, database_overflow_max_lsas=None, database_overflow_time_to_recover=None, default_information_metric=None, default_information_metric_type=None, default_information_originate=None, default_information_route_map=None, default_metric=None, distance=None, distance_external=None, distance_inter_area=None, distance_intra_area=None, distribute_list_in=None, distribute_lists=None, distribute_route_map_in=None, id=None, log_neighbour_changes=None, neighbors=None, networks=None, ospf_interfaces=None, passive_interfaces=None, redistributes=None, restart_mode=None, restart_period=None, rfc1583_compatible=None, router_id=None, spf_timers=None, summary_addresses=None, vdomparam=None):
        if abr_type and not isinstance(abr_type, str):
            raise TypeError("Expected argument 'abr_type' to be a str")
        pulumi.set(__self__, "abr_type", abr_type)
        if areas and not isinstance(areas, list):
            raise TypeError("Expected argument 'areas' to be a list")
        pulumi.set(__self__, "areas", areas)
        if auto_cost_ref_bandwidth and not isinstance(auto_cost_ref_bandwidth, int):
            raise TypeError("Expected argument 'auto_cost_ref_bandwidth' to be a int")
        pulumi.set(__self__, "auto_cost_ref_bandwidth", auto_cost_ref_bandwidth)
        if bfd and not isinstance(bfd, str):
            raise TypeError("Expected argument 'bfd' to be a str")
        pulumi.set(__self__, "bfd", bfd)
        if database_overflow and not isinstance(database_overflow, str):
            raise TypeError("Expected argument 'database_overflow' to be a str")
        pulumi.set(__self__, "database_overflow", database_overflow)
        if database_overflow_max_lsas and not isinstance(database_overflow_max_lsas, int):
            raise TypeError("Expected argument 'database_overflow_max_lsas' to be a int")
        pulumi.set(__self__, "database_overflow_max_lsas", database_overflow_max_lsas)
        if database_overflow_time_to_recover and not isinstance(database_overflow_time_to_recover, int):
            raise TypeError("Expected argument 'database_overflow_time_to_recover' to be a int")
        pulumi.set(__self__, "database_overflow_time_to_recover", database_overflow_time_to_recover)
        if default_information_metric and not isinstance(default_information_metric, int):
            raise TypeError("Expected argument 'default_information_metric' to be a int")
        pulumi.set(__self__, "default_information_metric", default_information_metric)
        if default_information_metric_type and not isinstance(default_information_metric_type, str):
            raise TypeError("Expected argument 'default_information_metric_type' to be a str")
        pulumi.set(__self__, "default_information_metric_type", default_information_metric_type)
        if default_information_originate and not isinstance(default_information_originate, str):
            raise TypeError("Expected argument 'default_information_originate' to be a str")
        pulumi.set(__self__, "default_information_originate", default_information_originate)
        if default_information_route_map and not isinstance(default_information_route_map, str):
            raise TypeError("Expected argument 'default_information_route_map' to be a str")
        pulumi.set(__self__, "default_information_route_map", default_information_route_map)
        if default_metric and not isinstance(default_metric, int):
            raise TypeError("Expected argument 'default_metric' to be a int")
        pulumi.set(__self__, "default_metric", default_metric)
        if distance and not isinstance(distance, int):
            raise TypeError("Expected argument 'distance' to be a int")
        pulumi.set(__self__, "distance", distance)
        if distance_external and not isinstance(distance_external, int):
            raise TypeError("Expected argument 'distance_external' to be a int")
        pulumi.set(__self__, "distance_external", distance_external)
        if distance_inter_area and not isinstance(distance_inter_area, int):
            raise TypeError("Expected argument 'distance_inter_area' to be a int")
        pulumi.set(__self__, "distance_inter_area", distance_inter_area)
        if distance_intra_area and not isinstance(distance_intra_area, int):
            raise TypeError("Expected argument 'distance_intra_area' to be a int")
        pulumi.set(__self__, "distance_intra_area", distance_intra_area)
        if distribute_list_in and not isinstance(distribute_list_in, str):
            raise TypeError("Expected argument 'distribute_list_in' to be a str")
        pulumi.set(__self__, "distribute_list_in", distribute_list_in)
        if distribute_lists and not isinstance(distribute_lists, list):
            raise TypeError("Expected argument 'distribute_lists' to be a list")
        pulumi.set(__self__, "distribute_lists", distribute_lists)
        if distribute_route_map_in and not isinstance(distribute_route_map_in, str):
            raise TypeError("Expected argument 'distribute_route_map_in' to be a str")
        pulumi.set(__self__, "distribute_route_map_in", distribute_route_map_in)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if log_neighbour_changes and not isinstance(log_neighbour_changes, str):
            raise TypeError("Expected argument 'log_neighbour_changes' to be a str")
        pulumi.set(__self__, "log_neighbour_changes", log_neighbour_changes)
        if neighbors and not isinstance(neighbors, list):
            raise TypeError("Expected argument 'neighbors' to be a list")
        pulumi.set(__self__, "neighbors", neighbors)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if ospf_interfaces and not isinstance(ospf_interfaces, list):
            raise TypeError("Expected argument 'ospf_interfaces' to be a list")
        pulumi.set(__self__, "ospf_interfaces", ospf_interfaces)
        if passive_interfaces and not isinstance(passive_interfaces, list):
            raise TypeError("Expected argument 'passive_interfaces' to be a list")
        pulumi.set(__self__, "passive_interfaces", passive_interfaces)
        if redistributes and not isinstance(redistributes, list):
            raise TypeError("Expected argument 'redistributes' to be a list")
        pulumi.set(__self__, "redistributes", redistributes)
        if restart_mode and not isinstance(restart_mode, str):
            raise TypeError("Expected argument 'restart_mode' to be a str")
        pulumi.set(__self__, "restart_mode", restart_mode)
        if restart_period and not isinstance(restart_period, int):
            raise TypeError("Expected argument 'restart_period' to be a int")
        pulumi.set(__self__, "restart_period", restart_period)
        if rfc1583_compatible and not isinstance(rfc1583_compatible, str):
            raise TypeError("Expected argument 'rfc1583_compatible' to be a str")
        pulumi.set(__self__, "rfc1583_compatible", rfc1583_compatible)
        if router_id and not isinstance(router_id, str):
            raise TypeError("Expected argument 'router_id' to be a str")
        pulumi.set(__self__, "router_id", router_id)
        if spf_timers and not isinstance(spf_timers, str):
            raise TypeError("Expected argument 'spf_timers' to be a str")
        pulumi.set(__self__, "spf_timers", spf_timers)
        if summary_addresses and not isinstance(summary_addresses, list):
            raise TypeError("Expected argument 'summary_addresses' to be a list")
        pulumi.set(__self__, "summary_addresses", summary_addresses)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="abrType")
    def abr_type(self) -> str:
        """
        Area border router type.
        """
        return pulumi.get(self, "abr_type")

    @property
    @pulumi.getter
    def areas(self) -> Sequence['outputs.GetRouterOspfAreaResult']:
        """
        Attach the network to area.
        """
        return pulumi.get(self, "areas")

    @property
    @pulumi.getter(name="autoCostRefBandwidth")
    def auto_cost_ref_bandwidth(self) -> int:
        """
        Reference bandwidth in terms of megabits per second.
        """
        return pulumi.get(self, "auto_cost_ref_bandwidth")

    @property
    @pulumi.getter
    def bfd(self) -> str:
        """
        Bidirectional Forwarding Detection (BFD).
        """
        return pulumi.get(self, "bfd")

    @property
    @pulumi.getter(name="databaseOverflow")
    def database_overflow(self) -> str:
        """
        Enable/disable database overflow.
        """
        return pulumi.get(self, "database_overflow")

    @property
    @pulumi.getter(name="databaseOverflowMaxLsas")
    def database_overflow_max_lsas(self) -> int:
        """
        Database overflow maximum LSAs.
        """
        return pulumi.get(self, "database_overflow_max_lsas")

    @property
    @pulumi.getter(name="databaseOverflowTimeToRecover")
    def database_overflow_time_to_recover(self) -> int:
        """
        Database overflow time to recover (sec).
        """
        return pulumi.get(self, "database_overflow_time_to_recover")

    @property
    @pulumi.getter(name="defaultInformationMetric")
    def default_information_metric(self) -> int:
        """
        Default information metric.
        """
        return pulumi.get(self, "default_information_metric")

    @property
    @pulumi.getter(name="defaultInformationMetricType")
    def default_information_metric_type(self) -> str:
        """
        Default information metric type.
        """
        return pulumi.get(self, "default_information_metric_type")

    @property
    @pulumi.getter(name="defaultInformationOriginate")
    def default_information_originate(self) -> str:
        """
        Enable/disable generation of default route.
        """
        return pulumi.get(self, "default_information_originate")

    @property
    @pulumi.getter(name="defaultInformationRouteMap")
    def default_information_route_map(self) -> str:
        """
        Default information route map.
        """
        return pulumi.get(self, "default_information_route_map")

    @property
    @pulumi.getter(name="defaultMetric")
    def default_metric(self) -> int:
        """
        Default metric of redistribute routes.
        """
        return pulumi.get(self, "default_metric")

    @property
    @pulumi.getter
    def distance(self) -> int:
        """
        Distance of the route.
        """
        return pulumi.get(self, "distance")

    @property
    @pulumi.getter(name="distanceExternal")
    def distance_external(self) -> int:
        """
        Administrative external distance.
        """
        return pulumi.get(self, "distance_external")

    @property
    @pulumi.getter(name="distanceInterArea")
    def distance_inter_area(self) -> int:
        """
        Administrative inter-area distance.
        """
        return pulumi.get(self, "distance_inter_area")

    @property
    @pulumi.getter(name="distanceIntraArea")
    def distance_intra_area(self) -> int:
        """
        Administrative intra-area distance.
        """
        return pulumi.get(self, "distance_intra_area")

    @property
    @pulumi.getter(name="distributeListIn")
    def distribute_list_in(self) -> str:
        """
        Filter incoming routes.
        """
        return pulumi.get(self, "distribute_list_in")

    @property
    @pulumi.getter(name="distributeLists")
    def distribute_lists(self) -> Sequence['outputs.GetRouterOspfDistributeListResult']:
        """
        Distribute list configuration. The structure of `distribute_list` block is documented below.
        """
        return pulumi.get(self, "distribute_lists")

    @property
    @pulumi.getter(name="distributeRouteMapIn")
    def distribute_route_map_in(self) -> str:
        """
        Filter incoming external routes by route-map.
        """
        return pulumi.get(self, "distribute_route_map_in")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logNeighbourChanges")
    def log_neighbour_changes(self) -> str:
        """
        Enable logging of OSPF neighbour's changes
        """
        return pulumi.get(self, "log_neighbour_changes")

    @property
    @pulumi.getter
    def neighbors(self) -> Sequence['outputs.GetRouterOspfNeighborResult']:
        """
        OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetRouterOspfNetworkResult']:
        """
        OSPF network configuration. The structure of `network` block is documented below.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="ospfInterfaces")
    def ospf_interfaces(self) -> Sequence['outputs.GetRouterOspfOspfInterfaceResult']:
        """
        OSPF interface configuration. The structure of `ospf_interface` block is documented below.
        """
        return pulumi.get(self, "ospf_interfaces")

    @property
    @pulumi.getter(name="passiveInterfaces")
    def passive_interfaces(self) -> Sequence['outputs.GetRouterOspfPassiveInterfaceResult']:
        """
        Passive interface configuration. The structure of `passive_interface` block is documented below.
        """
        return pulumi.get(self, "passive_interfaces")

    @property
    @pulumi.getter
    def redistributes(self) -> Sequence['outputs.GetRouterOspfRedistributeResult']:
        """
        Redistribute configuration. The structure of `redistribute` block is documented below.
        """
        return pulumi.get(self, "redistributes")

    @property
    @pulumi.getter(name="restartMode")
    def restart_mode(self) -> str:
        """
        OSPF restart mode (graceful or LLS).
        """
        return pulumi.get(self, "restart_mode")

    @property
    @pulumi.getter(name="restartPeriod")
    def restart_period(self) -> int:
        """
        Graceful restart period.
        """
        return pulumi.get(self, "restart_period")

    @property
    @pulumi.getter(name="rfc1583Compatible")
    def rfc1583_compatible(self) -> str:
        """
        Enable/disable RFC1583 compatibility.
        """
        return pulumi.get(self, "rfc1583_compatible")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> str:
        """
        Router ID.
        """
        return pulumi.get(self, "router_id")

    @property
    @pulumi.getter(name="spfTimers")
    def spf_timers(self) -> str:
        """
        SPF calculation frequency.
        """
        return pulumi.get(self, "spf_timers")

    @property
    @pulumi.getter(name="summaryAddresses")
    def summary_addresses(self) -> Sequence['outputs.GetRouterOspfSummaryAddressResult']:
        """
        IP address summary configuration. The structure of `summary_address` block is documented below.
        """
        return pulumi.get(self, "summary_addresses")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterOspfResult(GetRouterOspfResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterOspfResult(
            abr_type=self.abr_type,
            areas=self.areas,
            auto_cost_ref_bandwidth=self.auto_cost_ref_bandwidth,
            bfd=self.bfd,
            database_overflow=self.database_overflow,
            database_overflow_max_lsas=self.database_overflow_max_lsas,
            database_overflow_time_to_recover=self.database_overflow_time_to_recover,
            default_information_metric=self.default_information_metric,
            default_information_metric_type=self.default_information_metric_type,
            default_information_originate=self.default_information_originate,
            default_information_route_map=self.default_information_route_map,
            default_metric=self.default_metric,
            distance=self.distance,
            distance_external=self.distance_external,
            distance_inter_area=self.distance_inter_area,
            distance_intra_area=self.distance_intra_area,
            distribute_list_in=self.distribute_list_in,
            distribute_lists=self.distribute_lists,
            distribute_route_map_in=self.distribute_route_map_in,
            id=self.id,
            log_neighbour_changes=self.log_neighbour_changes,
            neighbors=self.neighbors,
            networks=self.networks,
            ospf_interfaces=self.ospf_interfaces,
            passive_interfaces=self.passive_interfaces,
            redistributes=self.redistributes,
            restart_mode=self.restart_mode,
            restart_period=self.restart_period,
            rfc1583_compatible=self.rfc1583_compatible,
            router_id=self.router_id,
            spf_timers=self.spf_timers,
            summary_addresses=self.summary_addresses,
            vdomparam=self.vdomparam)


def get_router_ospf(vdomparam: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterOspfResult:
    """
    Use this data source to get information on fortios router ospf


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterOspf:GetRouterOspf', __args__, opts=opts, typ=GetRouterOspfResult).value

    return AwaitableGetRouterOspfResult(
        abr_type=__ret__.abr_type,
        areas=__ret__.areas,
        auto_cost_ref_bandwidth=__ret__.auto_cost_ref_bandwidth,
        bfd=__ret__.bfd,
        database_overflow=__ret__.database_overflow,
        database_overflow_max_lsas=__ret__.database_overflow_max_lsas,
        database_overflow_time_to_recover=__ret__.database_overflow_time_to_recover,
        default_information_metric=__ret__.default_information_metric,
        default_information_metric_type=__ret__.default_information_metric_type,
        default_information_originate=__ret__.default_information_originate,
        default_information_route_map=__ret__.default_information_route_map,
        default_metric=__ret__.default_metric,
        distance=__ret__.distance,
        distance_external=__ret__.distance_external,
        distance_inter_area=__ret__.distance_inter_area,
        distance_intra_area=__ret__.distance_intra_area,
        distribute_list_in=__ret__.distribute_list_in,
        distribute_lists=__ret__.distribute_lists,
        distribute_route_map_in=__ret__.distribute_route_map_in,
        id=__ret__.id,
        log_neighbour_changes=__ret__.log_neighbour_changes,
        neighbors=__ret__.neighbors,
        networks=__ret__.networks,
        ospf_interfaces=__ret__.ospf_interfaces,
        passive_interfaces=__ret__.passive_interfaces,
        redistributes=__ret__.redistributes,
        restart_mode=__ret__.restart_mode,
        restart_period=__ret__.restart_period,
        rfc1583_compatible=__ret__.rfc1583_compatible,
        router_id=__ret__.router_id,
        spf_timers=__ret__.spf_timers,
        summary_addresses=__ret__.summary_addresses,
        vdomparam=__ret__.vdomparam)
