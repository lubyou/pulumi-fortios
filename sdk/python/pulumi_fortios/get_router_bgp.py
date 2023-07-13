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
    'GetRouterBgpResult',
    'AwaitableGetRouterBgpResult',
    'get_router_bgp',
    'get_router_bgp_output',
]

@pulumi.output_type
class GetRouterBgpResult:
    """
    A collection of values returned by GetRouterBgp.
    """
    def __init__(__self__, additional_path=None, additional_path6=None, additional_path_select=None, additional_path_select6=None, additional_path_select_vpnv4=None, additional_path_vpnv4=None, admin_distances=None, aggregate_address6s=None, aggregate_addresses=None, always_compare_med=None, as_=None, as_string=None, bestpath_as_path_ignore=None, bestpath_cmp_confed_aspath=None, bestpath_cmp_routerid=None, bestpath_med_confed=None, bestpath_med_missing_as_worst=None, client_to_client_reflection=None, cluster_id=None, confederation_identifier=None, confederation_peers=None, cross_family_conditional_adv=None, dampening=None, dampening_max_suppress_time=None, dampening_reachability_half_life=None, dampening_reuse=None, dampening_route_map=None, dampening_suppress=None, dampening_unreachability_half_life=None, default_local_preference=None, deterministic_med=None, distance_external=None, distance_internal=None, distance_local=None, ebgp_multipath=None, enforce_first_as=None, fast_external_failover=None, graceful_end_on_timer=None, graceful_restart=None, graceful_restart_time=None, graceful_stalepath_time=None, graceful_update_delay=None, holdtime_timer=None, ibgp_multipath=None, id=None, ignore_optional_capability=None, keepalive_timer=None, log_neighbour_changes=None, multipath_recursive_distance=None, neighbor_groups=None, neighbor_range6s=None, neighbor_ranges=None, neighbors=None, network6s=None, network_import_check=None, networks=None, recursive_inherit_priority=None, recursive_next_hop=None, redistribute6s=None, redistributes=None, router_id=None, scan_time=None, synchronization=None, tag_resolve_mode=None, vdomparam=None, vrf6s=None, vrf_leak6s=None, vrf_leaks=None, vrves=None):
        if additional_path and not isinstance(additional_path, str):
            raise TypeError("Expected argument 'additional_path' to be a str")
        pulumi.set(__self__, "additional_path", additional_path)
        if additional_path6 and not isinstance(additional_path6, str):
            raise TypeError("Expected argument 'additional_path6' to be a str")
        pulumi.set(__self__, "additional_path6", additional_path6)
        if additional_path_select and not isinstance(additional_path_select, int):
            raise TypeError("Expected argument 'additional_path_select' to be a int")
        pulumi.set(__self__, "additional_path_select", additional_path_select)
        if additional_path_select6 and not isinstance(additional_path_select6, int):
            raise TypeError("Expected argument 'additional_path_select6' to be a int")
        pulumi.set(__self__, "additional_path_select6", additional_path_select6)
        if additional_path_select_vpnv4 and not isinstance(additional_path_select_vpnv4, int):
            raise TypeError("Expected argument 'additional_path_select_vpnv4' to be a int")
        pulumi.set(__self__, "additional_path_select_vpnv4", additional_path_select_vpnv4)
        if additional_path_vpnv4 and not isinstance(additional_path_vpnv4, str):
            raise TypeError("Expected argument 'additional_path_vpnv4' to be a str")
        pulumi.set(__self__, "additional_path_vpnv4", additional_path_vpnv4)
        if admin_distances and not isinstance(admin_distances, list):
            raise TypeError("Expected argument 'admin_distances' to be a list")
        pulumi.set(__self__, "admin_distances", admin_distances)
        if aggregate_address6s and not isinstance(aggregate_address6s, list):
            raise TypeError("Expected argument 'aggregate_address6s' to be a list")
        pulumi.set(__self__, "aggregate_address6s", aggregate_address6s)
        if aggregate_addresses and not isinstance(aggregate_addresses, list):
            raise TypeError("Expected argument 'aggregate_addresses' to be a list")
        pulumi.set(__self__, "aggregate_addresses", aggregate_addresses)
        if always_compare_med and not isinstance(always_compare_med, str):
            raise TypeError("Expected argument 'always_compare_med' to be a str")
        pulumi.set(__self__, "always_compare_med", always_compare_med)
        if as_ and not isinstance(as_, int):
            raise TypeError("Expected argument 'as_' to be a int")
        pulumi.set(__self__, "as_", as_)
        if as_string and not isinstance(as_string, str):
            raise TypeError("Expected argument 'as_string' to be a str")
        pulumi.set(__self__, "as_string", as_string)
        if bestpath_as_path_ignore and not isinstance(bestpath_as_path_ignore, str):
            raise TypeError("Expected argument 'bestpath_as_path_ignore' to be a str")
        pulumi.set(__self__, "bestpath_as_path_ignore", bestpath_as_path_ignore)
        if bestpath_cmp_confed_aspath and not isinstance(bestpath_cmp_confed_aspath, str):
            raise TypeError("Expected argument 'bestpath_cmp_confed_aspath' to be a str")
        pulumi.set(__self__, "bestpath_cmp_confed_aspath", bestpath_cmp_confed_aspath)
        if bestpath_cmp_routerid and not isinstance(bestpath_cmp_routerid, str):
            raise TypeError("Expected argument 'bestpath_cmp_routerid' to be a str")
        pulumi.set(__self__, "bestpath_cmp_routerid", bestpath_cmp_routerid)
        if bestpath_med_confed and not isinstance(bestpath_med_confed, str):
            raise TypeError("Expected argument 'bestpath_med_confed' to be a str")
        pulumi.set(__self__, "bestpath_med_confed", bestpath_med_confed)
        if bestpath_med_missing_as_worst and not isinstance(bestpath_med_missing_as_worst, str):
            raise TypeError("Expected argument 'bestpath_med_missing_as_worst' to be a str")
        pulumi.set(__self__, "bestpath_med_missing_as_worst", bestpath_med_missing_as_worst)
        if client_to_client_reflection and not isinstance(client_to_client_reflection, str):
            raise TypeError("Expected argument 'client_to_client_reflection' to be a str")
        pulumi.set(__self__, "client_to_client_reflection", client_to_client_reflection)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if confederation_identifier and not isinstance(confederation_identifier, int):
            raise TypeError("Expected argument 'confederation_identifier' to be a int")
        pulumi.set(__self__, "confederation_identifier", confederation_identifier)
        if confederation_peers and not isinstance(confederation_peers, list):
            raise TypeError("Expected argument 'confederation_peers' to be a list")
        pulumi.set(__self__, "confederation_peers", confederation_peers)
        if cross_family_conditional_adv and not isinstance(cross_family_conditional_adv, str):
            raise TypeError("Expected argument 'cross_family_conditional_adv' to be a str")
        pulumi.set(__self__, "cross_family_conditional_adv", cross_family_conditional_adv)
        if dampening and not isinstance(dampening, str):
            raise TypeError("Expected argument 'dampening' to be a str")
        pulumi.set(__self__, "dampening", dampening)
        if dampening_max_suppress_time and not isinstance(dampening_max_suppress_time, int):
            raise TypeError("Expected argument 'dampening_max_suppress_time' to be a int")
        pulumi.set(__self__, "dampening_max_suppress_time", dampening_max_suppress_time)
        if dampening_reachability_half_life and not isinstance(dampening_reachability_half_life, int):
            raise TypeError("Expected argument 'dampening_reachability_half_life' to be a int")
        pulumi.set(__self__, "dampening_reachability_half_life", dampening_reachability_half_life)
        if dampening_reuse and not isinstance(dampening_reuse, int):
            raise TypeError("Expected argument 'dampening_reuse' to be a int")
        pulumi.set(__self__, "dampening_reuse", dampening_reuse)
        if dampening_route_map and not isinstance(dampening_route_map, str):
            raise TypeError("Expected argument 'dampening_route_map' to be a str")
        pulumi.set(__self__, "dampening_route_map", dampening_route_map)
        if dampening_suppress and not isinstance(dampening_suppress, int):
            raise TypeError("Expected argument 'dampening_suppress' to be a int")
        pulumi.set(__self__, "dampening_suppress", dampening_suppress)
        if dampening_unreachability_half_life and not isinstance(dampening_unreachability_half_life, int):
            raise TypeError("Expected argument 'dampening_unreachability_half_life' to be a int")
        pulumi.set(__self__, "dampening_unreachability_half_life", dampening_unreachability_half_life)
        if default_local_preference and not isinstance(default_local_preference, int):
            raise TypeError("Expected argument 'default_local_preference' to be a int")
        pulumi.set(__self__, "default_local_preference", default_local_preference)
        if deterministic_med and not isinstance(deterministic_med, str):
            raise TypeError("Expected argument 'deterministic_med' to be a str")
        pulumi.set(__self__, "deterministic_med", deterministic_med)
        if distance_external and not isinstance(distance_external, int):
            raise TypeError("Expected argument 'distance_external' to be a int")
        pulumi.set(__self__, "distance_external", distance_external)
        if distance_internal and not isinstance(distance_internal, int):
            raise TypeError("Expected argument 'distance_internal' to be a int")
        pulumi.set(__self__, "distance_internal", distance_internal)
        if distance_local and not isinstance(distance_local, int):
            raise TypeError("Expected argument 'distance_local' to be a int")
        pulumi.set(__self__, "distance_local", distance_local)
        if ebgp_multipath and not isinstance(ebgp_multipath, str):
            raise TypeError("Expected argument 'ebgp_multipath' to be a str")
        pulumi.set(__self__, "ebgp_multipath", ebgp_multipath)
        if enforce_first_as and not isinstance(enforce_first_as, str):
            raise TypeError("Expected argument 'enforce_first_as' to be a str")
        pulumi.set(__self__, "enforce_first_as", enforce_first_as)
        if fast_external_failover and not isinstance(fast_external_failover, str):
            raise TypeError("Expected argument 'fast_external_failover' to be a str")
        pulumi.set(__self__, "fast_external_failover", fast_external_failover)
        if graceful_end_on_timer and not isinstance(graceful_end_on_timer, str):
            raise TypeError("Expected argument 'graceful_end_on_timer' to be a str")
        pulumi.set(__self__, "graceful_end_on_timer", graceful_end_on_timer)
        if graceful_restart and not isinstance(graceful_restart, str):
            raise TypeError("Expected argument 'graceful_restart' to be a str")
        pulumi.set(__self__, "graceful_restart", graceful_restart)
        if graceful_restart_time and not isinstance(graceful_restart_time, int):
            raise TypeError("Expected argument 'graceful_restart_time' to be a int")
        pulumi.set(__self__, "graceful_restart_time", graceful_restart_time)
        if graceful_stalepath_time and not isinstance(graceful_stalepath_time, int):
            raise TypeError("Expected argument 'graceful_stalepath_time' to be a int")
        pulumi.set(__self__, "graceful_stalepath_time", graceful_stalepath_time)
        if graceful_update_delay and not isinstance(graceful_update_delay, int):
            raise TypeError("Expected argument 'graceful_update_delay' to be a int")
        pulumi.set(__self__, "graceful_update_delay", graceful_update_delay)
        if holdtime_timer and not isinstance(holdtime_timer, int):
            raise TypeError("Expected argument 'holdtime_timer' to be a int")
        pulumi.set(__self__, "holdtime_timer", holdtime_timer)
        if ibgp_multipath and not isinstance(ibgp_multipath, str):
            raise TypeError("Expected argument 'ibgp_multipath' to be a str")
        pulumi.set(__self__, "ibgp_multipath", ibgp_multipath)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_optional_capability and not isinstance(ignore_optional_capability, str):
            raise TypeError("Expected argument 'ignore_optional_capability' to be a str")
        pulumi.set(__self__, "ignore_optional_capability", ignore_optional_capability)
        if keepalive_timer and not isinstance(keepalive_timer, int):
            raise TypeError("Expected argument 'keepalive_timer' to be a int")
        pulumi.set(__self__, "keepalive_timer", keepalive_timer)
        if log_neighbour_changes and not isinstance(log_neighbour_changes, str):
            raise TypeError("Expected argument 'log_neighbour_changes' to be a str")
        pulumi.set(__self__, "log_neighbour_changes", log_neighbour_changes)
        if multipath_recursive_distance and not isinstance(multipath_recursive_distance, str):
            raise TypeError("Expected argument 'multipath_recursive_distance' to be a str")
        pulumi.set(__self__, "multipath_recursive_distance", multipath_recursive_distance)
        if neighbor_groups and not isinstance(neighbor_groups, list):
            raise TypeError("Expected argument 'neighbor_groups' to be a list")
        pulumi.set(__self__, "neighbor_groups", neighbor_groups)
        if neighbor_range6s and not isinstance(neighbor_range6s, list):
            raise TypeError("Expected argument 'neighbor_range6s' to be a list")
        pulumi.set(__self__, "neighbor_range6s", neighbor_range6s)
        if neighbor_ranges and not isinstance(neighbor_ranges, list):
            raise TypeError("Expected argument 'neighbor_ranges' to be a list")
        pulumi.set(__self__, "neighbor_ranges", neighbor_ranges)
        if neighbors and not isinstance(neighbors, list):
            raise TypeError("Expected argument 'neighbors' to be a list")
        pulumi.set(__self__, "neighbors", neighbors)
        if network6s and not isinstance(network6s, list):
            raise TypeError("Expected argument 'network6s' to be a list")
        pulumi.set(__self__, "network6s", network6s)
        if network_import_check and not isinstance(network_import_check, str):
            raise TypeError("Expected argument 'network_import_check' to be a str")
        pulumi.set(__self__, "network_import_check", network_import_check)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if recursive_inherit_priority and not isinstance(recursive_inherit_priority, str):
            raise TypeError("Expected argument 'recursive_inherit_priority' to be a str")
        pulumi.set(__self__, "recursive_inherit_priority", recursive_inherit_priority)
        if recursive_next_hop and not isinstance(recursive_next_hop, str):
            raise TypeError("Expected argument 'recursive_next_hop' to be a str")
        pulumi.set(__self__, "recursive_next_hop", recursive_next_hop)
        if redistribute6s and not isinstance(redistribute6s, list):
            raise TypeError("Expected argument 'redistribute6s' to be a list")
        pulumi.set(__self__, "redistribute6s", redistribute6s)
        if redistributes and not isinstance(redistributes, list):
            raise TypeError("Expected argument 'redistributes' to be a list")
        pulumi.set(__self__, "redistributes", redistributes)
        if router_id and not isinstance(router_id, str):
            raise TypeError("Expected argument 'router_id' to be a str")
        pulumi.set(__self__, "router_id", router_id)
        if scan_time and not isinstance(scan_time, int):
            raise TypeError("Expected argument 'scan_time' to be a int")
        pulumi.set(__self__, "scan_time", scan_time)
        if synchronization and not isinstance(synchronization, str):
            raise TypeError("Expected argument 'synchronization' to be a str")
        pulumi.set(__self__, "synchronization", synchronization)
        if tag_resolve_mode and not isinstance(tag_resolve_mode, str):
            raise TypeError("Expected argument 'tag_resolve_mode' to be a str")
        pulumi.set(__self__, "tag_resolve_mode", tag_resolve_mode)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if vrf6s and not isinstance(vrf6s, list):
            raise TypeError("Expected argument 'vrf6s' to be a list")
        pulumi.set(__self__, "vrf6s", vrf6s)
        if vrf_leak6s and not isinstance(vrf_leak6s, list):
            raise TypeError("Expected argument 'vrf_leak6s' to be a list")
        pulumi.set(__self__, "vrf_leak6s", vrf_leak6s)
        if vrf_leaks and not isinstance(vrf_leaks, list):
            raise TypeError("Expected argument 'vrf_leaks' to be a list")
        pulumi.set(__self__, "vrf_leaks", vrf_leaks)
        if vrves and not isinstance(vrves, list):
            raise TypeError("Expected argument 'vrves' to be a list")
        pulumi.set(__self__, "vrves", vrves)

    @property
    @pulumi.getter(name="additionalPath")
    def additional_path(self) -> str:
        return pulumi.get(self, "additional_path")

    @property
    @pulumi.getter(name="additionalPath6")
    def additional_path6(self) -> str:
        return pulumi.get(self, "additional_path6")

    @property
    @pulumi.getter(name="additionalPathSelect")
    def additional_path_select(self) -> int:
        return pulumi.get(self, "additional_path_select")

    @property
    @pulumi.getter(name="additionalPathSelect6")
    def additional_path_select6(self) -> int:
        return pulumi.get(self, "additional_path_select6")

    @property
    @pulumi.getter(name="additionalPathSelectVpnv4")
    def additional_path_select_vpnv4(self) -> int:
        return pulumi.get(self, "additional_path_select_vpnv4")

    @property
    @pulumi.getter(name="additionalPathVpnv4")
    def additional_path_vpnv4(self) -> str:
        return pulumi.get(self, "additional_path_vpnv4")

    @property
    @pulumi.getter(name="adminDistances")
    def admin_distances(self) -> Sequence['outputs.GetRouterBgpAdminDistanceResult']:
        return pulumi.get(self, "admin_distances")

    @property
    @pulumi.getter(name="aggregateAddress6s")
    def aggregate_address6s(self) -> Sequence['outputs.GetRouterBgpAggregateAddress6Result']:
        return pulumi.get(self, "aggregate_address6s")

    @property
    @pulumi.getter(name="aggregateAddresses")
    def aggregate_addresses(self) -> Sequence['outputs.GetRouterBgpAggregateAddressResult']:
        return pulumi.get(self, "aggregate_addresses")

    @property
    @pulumi.getter(name="alwaysCompareMed")
    def always_compare_med(self) -> str:
        return pulumi.get(self, "always_compare_med")

    @property
    @pulumi.getter(name="as")
    def as_(self) -> int:
        return pulumi.get(self, "as_")

    @property
    @pulumi.getter(name="asString")
    def as_string(self) -> str:
        return pulumi.get(self, "as_string")

    @property
    @pulumi.getter(name="bestpathAsPathIgnore")
    def bestpath_as_path_ignore(self) -> str:
        return pulumi.get(self, "bestpath_as_path_ignore")

    @property
    @pulumi.getter(name="bestpathCmpConfedAspath")
    def bestpath_cmp_confed_aspath(self) -> str:
        return pulumi.get(self, "bestpath_cmp_confed_aspath")

    @property
    @pulumi.getter(name="bestpathCmpRouterid")
    def bestpath_cmp_routerid(self) -> str:
        return pulumi.get(self, "bestpath_cmp_routerid")

    @property
    @pulumi.getter(name="bestpathMedConfed")
    def bestpath_med_confed(self) -> str:
        return pulumi.get(self, "bestpath_med_confed")

    @property
    @pulumi.getter(name="bestpathMedMissingAsWorst")
    def bestpath_med_missing_as_worst(self) -> str:
        return pulumi.get(self, "bestpath_med_missing_as_worst")

    @property
    @pulumi.getter(name="clientToClientReflection")
    def client_to_client_reflection(self) -> str:
        return pulumi.get(self, "client_to_client_reflection")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="confederationIdentifier")
    def confederation_identifier(self) -> int:
        return pulumi.get(self, "confederation_identifier")

    @property
    @pulumi.getter(name="confederationPeers")
    def confederation_peers(self) -> Sequence['outputs.GetRouterBgpConfederationPeerResult']:
        return pulumi.get(self, "confederation_peers")

    @property
    @pulumi.getter(name="crossFamilyConditionalAdv")
    def cross_family_conditional_adv(self) -> str:
        return pulumi.get(self, "cross_family_conditional_adv")

    @property
    @pulumi.getter
    def dampening(self) -> str:
        return pulumi.get(self, "dampening")

    @property
    @pulumi.getter(name="dampeningMaxSuppressTime")
    def dampening_max_suppress_time(self) -> int:
        return pulumi.get(self, "dampening_max_suppress_time")

    @property
    @pulumi.getter(name="dampeningReachabilityHalfLife")
    def dampening_reachability_half_life(self) -> int:
        return pulumi.get(self, "dampening_reachability_half_life")

    @property
    @pulumi.getter(name="dampeningReuse")
    def dampening_reuse(self) -> int:
        return pulumi.get(self, "dampening_reuse")

    @property
    @pulumi.getter(name="dampeningRouteMap")
    def dampening_route_map(self) -> str:
        return pulumi.get(self, "dampening_route_map")

    @property
    @pulumi.getter(name="dampeningSuppress")
    def dampening_suppress(self) -> int:
        return pulumi.get(self, "dampening_suppress")

    @property
    @pulumi.getter(name="dampeningUnreachabilityHalfLife")
    def dampening_unreachability_half_life(self) -> int:
        return pulumi.get(self, "dampening_unreachability_half_life")

    @property
    @pulumi.getter(name="defaultLocalPreference")
    def default_local_preference(self) -> int:
        return pulumi.get(self, "default_local_preference")

    @property
    @pulumi.getter(name="deterministicMed")
    def deterministic_med(self) -> str:
        return pulumi.get(self, "deterministic_med")

    @property
    @pulumi.getter(name="distanceExternal")
    def distance_external(self) -> int:
        return pulumi.get(self, "distance_external")

    @property
    @pulumi.getter(name="distanceInternal")
    def distance_internal(self) -> int:
        return pulumi.get(self, "distance_internal")

    @property
    @pulumi.getter(name="distanceLocal")
    def distance_local(self) -> int:
        return pulumi.get(self, "distance_local")

    @property
    @pulumi.getter(name="ebgpMultipath")
    def ebgp_multipath(self) -> str:
        return pulumi.get(self, "ebgp_multipath")

    @property
    @pulumi.getter(name="enforceFirstAs")
    def enforce_first_as(self) -> str:
        return pulumi.get(self, "enforce_first_as")

    @property
    @pulumi.getter(name="fastExternalFailover")
    def fast_external_failover(self) -> str:
        return pulumi.get(self, "fast_external_failover")

    @property
    @pulumi.getter(name="gracefulEndOnTimer")
    def graceful_end_on_timer(self) -> str:
        return pulumi.get(self, "graceful_end_on_timer")

    @property
    @pulumi.getter(name="gracefulRestart")
    def graceful_restart(self) -> str:
        return pulumi.get(self, "graceful_restart")

    @property
    @pulumi.getter(name="gracefulRestartTime")
    def graceful_restart_time(self) -> int:
        return pulumi.get(self, "graceful_restart_time")

    @property
    @pulumi.getter(name="gracefulStalepathTime")
    def graceful_stalepath_time(self) -> int:
        return pulumi.get(self, "graceful_stalepath_time")

    @property
    @pulumi.getter(name="gracefulUpdateDelay")
    def graceful_update_delay(self) -> int:
        return pulumi.get(self, "graceful_update_delay")

    @property
    @pulumi.getter(name="holdtimeTimer")
    def holdtime_timer(self) -> int:
        return pulumi.get(self, "holdtime_timer")

    @property
    @pulumi.getter(name="ibgpMultipath")
    def ibgp_multipath(self) -> str:
        return pulumi.get(self, "ibgp_multipath")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreOptionalCapability")
    def ignore_optional_capability(self) -> str:
        return pulumi.get(self, "ignore_optional_capability")

    @property
    @pulumi.getter(name="keepaliveTimer")
    def keepalive_timer(self) -> int:
        return pulumi.get(self, "keepalive_timer")

    @property
    @pulumi.getter(name="logNeighbourChanges")
    def log_neighbour_changes(self) -> str:
        return pulumi.get(self, "log_neighbour_changes")

    @property
    @pulumi.getter(name="multipathRecursiveDistance")
    def multipath_recursive_distance(self) -> str:
        return pulumi.get(self, "multipath_recursive_distance")

    @property
    @pulumi.getter(name="neighborGroups")
    def neighbor_groups(self) -> Sequence['outputs.GetRouterBgpNeighborGroupResult']:
        return pulumi.get(self, "neighbor_groups")

    @property
    @pulumi.getter(name="neighborRange6s")
    def neighbor_range6s(self) -> Sequence['outputs.GetRouterBgpNeighborRange6Result']:
        return pulumi.get(self, "neighbor_range6s")

    @property
    @pulumi.getter(name="neighborRanges")
    def neighbor_ranges(self) -> Sequence['outputs.GetRouterBgpNeighborRangeResult']:
        return pulumi.get(self, "neighbor_ranges")

    @property
    @pulumi.getter
    def neighbors(self) -> Sequence['outputs.GetRouterBgpNeighborResult']:
        return pulumi.get(self, "neighbors")

    @property
    @pulumi.getter
    def network6s(self) -> Sequence['outputs.GetRouterBgpNetwork6Result']:
        return pulumi.get(self, "network6s")

    @property
    @pulumi.getter(name="networkImportCheck")
    def network_import_check(self) -> str:
        return pulumi.get(self, "network_import_check")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetRouterBgpNetworkResult']:
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="recursiveInheritPriority")
    def recursive_inherit_priority(self) -> str:
        return pulumi.get(self, "recursive_inherit_priority")

    @property
    @pulumi.getter(name="recursiveNextHop")
    def recursive_next_hop(self) -> str:
        return pulumi.get(self, "recursive_next_hop")

    @property
    @pulumi.getter
    def redistribute6s(self) -> Sequence['outputs.GetRouterBgpRedistribute6Result']:
        return pulumi.get(self, "redistribute6s")

    @property
    @pulumi.getter
    def redistributes(self) -> Sequence['outputs.GetRouterBgpRedistributeResult']:
        return pulumi.get(self, "redistributes")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> str:
        return pulumi.get(self, "router_id")

    @property
    @pulumi.getter(name="scanTime")
    def scan_time(self) -> int:
        return pulumi.get(self, "scan_time")

    @property
    @pulumi.getter
    def synchronization(self) -> str:
        return pulumi.get(self, "synchronization")

    @property
    @pulumi.getter(name="tagResolveMode")
    def tag_resolve_mode(self) -> str:
        return pulumi.get(self, "tag_resolve_mode")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vrf6s(self) -> Sequence['outputs.GetRouterBgpVrf6Result']:
        return pulumi.get(self, "vrf6s")

    @property
    @pulumi.getter(name="vrfLeak6s")
    def vrf_leak6s(self) -> Sequence['outputs.GetRouterBgpVrfLeak6Result']:
        return pulumi.get(self, "vrf_leak6s")

    @property
    @pulumi.getter(name="vrfLeaks")
    def vrf_leaks(self) -> Sequence['outputs.GetRouterBgpVrfLeakResult']:
        return pulumi.get(self, "vrf_leaks")

    @property
    @pulumi.getter
    def vrves(self) -> Sequence['outputs.GetRouterBgpVrfResult']:
        return pulumi.get(self, "vrves")


class AwaitableGetRouterBgpResult(GetRouterBgpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterBgpResult(
            additional_path=self.additional_path,
            additional_path6=self.additional_path6,
            additional_path_select=self.additional_path_select,
            additional_path_select6=self.additional_path_select6,
            additional_path_select_vpnv4=self.additional_path_select_vpnv4,
            additional_path_vpnv4=self.additional_path_vpnv4,
            admin_distances=self.admin_distances,
            aggregate_address6s=self.aggregate_address6s,
            aggregate_addresses=self.aggregate_addresses,
            always_compare_med=self.always_compare_med,
            as_=self.as_,
            as_string=self.as_string,
            bestpath_as_path_ignore=self.bestpath_as_path_ignore,
            bestpath_cmp_confed_aspath=self.bestpath_cmp_confed_aspath,
            bestpath_cmp_routerid=self.bestpath_cmp_routerid,
            bestpath_med_confed=self.bestpath_med_confed,
            bestpath_med_missing_as_worst=self.bestpath_med_missing_as_worst,
            client_to_client_reflection=self.client_to_client_reflection,
            cluster_id=self.cluster_id,
            confederation_identifier=self.confederation_identifier,
            confederation_peers=self.confederation_peers,
            cross_family_conditional_adv=self.cross_family_conditional_adv,
            dampening=self.dampening,
            dampening_max_suppress_time=self.dampening_max_suppress_time,
            dampening_reachability_half_life=self.dampening_reachability_half_life,
            dampening_reuse=self.dampening_reuse,
            dampening_route_map=self.dampening_route_map,
            dampening_suppress=self.dampening_suppress,
            dampening_unreachability_half_life=self.dampening_unreachability_half_life,
            default_local_preference=self.default_local_preference,
            deterministic_med=self.deterministic_med,
            distance_external=self.distance_external,
            distance_internal=self.distance_internal,
            distance_local=self.distance_local,
            ebgp_multipath=self.ebgp_multipath,
            enforce_first_as=self.enforce_first_as,
            fast_external_failover=self.fast_external_failover,
            graceful_end_on_timer=self.graceful_end_on_timer,
            graceful_restart=self.graceful_restart,
            graceful_restart_time=self.graceful_restart_time,
            graceful_stalepath_time=self.graceful_stalepath_time,
            graceful_update_delay=self.graceful_update_delay,
            holdtime_timer=self.holdtime_timer,
            ibgp_multipath=self.ibgp_multipath,
            id=self.id,
            ignore_optional_capability=self.ignore_optional_capability,
            keepalive_timer=self.keepalive_timer,
            log_neighbour_changes=self.log_neighbour_changes,
            multipath_recursive_distance=self.multipath_recursive_distance,
            neighbor_groups=self.neighbor_groups,
            neighbor_range6s=self.neighbor_range6s,
            neighbor_ranges=self.neighbor_ranges,
            neighbors=self.neighbors,
            network6s=self.network6s,
            network_import_check=self.network_import_check,
            networks=self.networks,
            recursive_inherit_priority=self.recursive_inherit_priority,
            recursive_next_hop=self.recursive_next_hop,
            redistribute6s=self.redistribute6s,
            redistributes=self.redistributes,
            router_id=self.router_id,
            scan_time=self.scan_time,
            synchronization=self.synchronization,
            tag_resolve_mode=self.tag_resolve_mode,
            vdomparam=self.vdomparam,
            vrf6s=self.vrf6s,
            vrf_leak6s=self.vrf_leak6s,
            vrf_leaks=self.vrf_leaks,
            vrves=self.vrves)


def get_router_bgp(vdomparam: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterBgpResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterBgp:GetRouterBgp', __args__, opts=opts, typ=GetRouterBgpResult).value

    return AwaitableGetRouterBgpResult(
        additional_path=pulumi.get(__ret__, 'additional_path'),
        additional_path6=pulumi.get(__ret__, 'additional_path6'),
        additional_path_select=pulumi.get(__ret__, 'additional_path_select'),
        additional_path_select6=pulumi.get(__ret__, 'additional_path_select6'),
        additional_path_select_vpnv4=pulumi.get(__ret__, 'additional_path_select_vpnv4'),
        additional_path_vpnv4=pulumi.get(__ret__, 'additional_path_vpnv4'),
        admin_distances=pulumi.get(__ret__, 'admin_distances'),
        aggregate_address6s=pulumi.get(__ret__, 'aggregate_address6s'),
        aggregate_addresses=pulumi.get(__ret__, 'aggregate_addresses'),
        always_compare_med=pulumi.get(__ret__, 'always_compare_med'),
        as_=pulumi.get(__ret__, 'as_'),
        as_string=pulumi.get(__ret__, 'as_string'),
        bestpath_as_path_ignore=pulumi.get(__ret__, 'bestpath_as_path_ignore'),
        bestpath_cmp_confed_aspath=pulumi.get(__ret__, 'bestpath_cmp_confed_aspath'),
        bestpath_cmp_routerid=pulumi.get(__ret__, 'bestpath_cmp_routerid'),
        bestpath_med_confed=pulumi.get(__ret__, 'bestpath_med_confed'),
        bestpath_med_missing_as_worst=pulumi.get(__ret__, 'bestpath_med_missing_as_worst'),
        client_to_client_reflection=pulumi.get(__ret__, 'client_to_client_reflection'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        confederation_identifier=pulumi.get(__ret__, 'confederation_identifier'),
        confederation_peers=pulumi.get(__ret__, 'confederation_peers'),
        cross_family_conditional_adv=pulumi.get(__ret__, 'cross_family_conditional_adv'),
        dampening=pulumi.get(__ret__, 'dampening'),
        dampening_max_suppress_time=pulumi.get(__ret__, 'dampening_max_suppress_time'),
        dampening_reachability_half_life=pulumi.get(__ret__, 'dampening_reachability_half_life'),
        dampening_reuse=pulumi.get(__ret__, 'dampening_reuse'),
        dampening_route_map=pulumi.get(__ret__, 'dampening_route_map'),
        dampening_suppress=pulumi.get(__ret__, 'dampening_suppress'),
        dampening_unreachability_half_life=pulumi.get(__ret__, 'dampening_unreachability_half_life'),
        default_local_preference=pulumi.get(__ret__, 'default_local_preference'),
        deterministic_med=pulumi.get(__ret__, 'deterministic_med'),
        distance_external=pulumi.get(__ret__, 'distance_external'),
        distance_internal=pulumi.get(__ret__, 'distance_internal'),
        distance_local=pulumi.get(__ret__, 'distance_local'),
        ebgp_multipath=pulumi.get(__ret__, 'ebgp_multipath'),
        enforce_first_as=pulumi.get(__ret__, 'enforce_first_as'),
        fast_external_failover=pulumi.get(__ret__, 'fast_external_failover'),
        graceful_end_on_timer=pulumi.get(__ret__, 'graceful_end_on_timer'),
        graceful_restart=pulumi.get(__ret__, 'graceful_restart'),
        graceful_restart_time=pulumi.get(__ret__, 'graceful_restart_time'),
        graceful_stalepath_time=pulumi.get(__ret__, 'graceful_stalepath_time'),
        graceful_update_delay=pulumi.get(__ret__, 'graceful_update_delay'),
        holdtime_timer=pulumi.get(__ret__, 'holdtime_timer'),
        ibgp_multipath=pulumi.get(__ret__, 'ibgp_multipath'),
        id=pulumi.get(__ret__, 'id'),
        ignore_optional_capability=pulumi.get(__ret__, 'ignore_optional_capability'),
        keepalive_timer=pulumi.get(__ret__, 'keepalive_timer'),
        log_neighbour_changes=pulumi.get(__ret__, 'log_neighbour_changes'),
        multipath_recursive_distance=pulumi.get(__ret__, 'multipath_recursive_distance'),
        neighbor_groups=pulumi.get(__ret__, 'neighbor_groups'),
        neighbor_range6s=pulumi.get(__ret__, 'neighbor_range6s'),
        neighbor_ranges=pulumi.get(__ret__, 'neighbor_ranges'),
        neighbors=pulumi.get(__ret__, 'neighbors'),
        network6s=pulumi.get(__ret__, 'network6s'),
        network_import_check=pulumi.get(__ret__, 'network_import_check'),
        networks=pulumi.get(__ret__, 'networks'),
        recursive_inherit_priority=pulumi.get(__ret__, 'recursive_inherit_priority'),
        recursive_next_hop=pulumi.get(__ret__, 'recursive_next_hop'),
        redistribute6s=pulumi.get(__ret__, 'redistribute6s'),
        redistributes=pulumi.get(__ret__, 'redistributes'),
        router_id=pulumi.get(__ret__, 'router_id'),
        scan_time=pulumi.get(__ret__, 'scan_time'),
        synchronization=pulumi.get(__ret__, 'synchronization'),
        tag_resolve_mode=pulumi.get(__ret__, 'tag_resolve_mode'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'),
        vrf6s=pulumi.get(__ret__, 'vrf6s'),
        vrf_leak6s=pulumi.get(__ret__, 'vrf_leak6s'),
        vrf_leaks=pulumi.get(__ret__, 'vrf_leaks'),
        vrves=pulumi.get(__ret__, 'vrves'))


@_utilities.lift_output_func(get_router_bgp)
def get_router_bgp_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterBgpResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
