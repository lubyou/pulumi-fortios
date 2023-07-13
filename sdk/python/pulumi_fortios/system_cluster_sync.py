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

__all__ = ['SystemClusterSyncArgs', 'SystemClusterSync']

@pulumi.input_type
class SystemClusterSyncArgs:
    def __init__(__self__, *,
                 down_intfs_before_sess_syncs: Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 hb_interval: Optional[pulumi.Input[int]] = None,
                 hb_lost_threshold: Optional[pulumi.Input[int]] = None,
                 ike_heartbeat_interval: Optional[pulumi.Input[int]] = None,
                 ike_monitor: Optional[pulumi.Input[str]] = None,
                 ike_monitor_interval: Optional[pulumi.Input[int]] = None,
                 ipsec_tunnel_sync: Optional[pulumi.Input[str]] = None,
                 peerip: Optional[pulumi.Input[str]] = None,
                 peervd: Optional[pulumi.Input[str]] = None,
                 secondary_add_ipsec_routes: Optional[pulumi.Input[str]] = None,
                 session_sync_filter: Optional[pulumi.Input['SystemClusterSyncSessionSyncFilterArgs']] = None,
                 slave_add_ike_routes: Optional[pulumi.Input[str]] = None,
                 sync_id: Optional[pulumi.Input[int]] = None,
                 syncvds: Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncSyncvdArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemClusterSync resource.
        """
        if down_intfs_before_sess_syncs is not None:
            pulumi.set(__self__, "down_intfs_before_sess_syncs", down_intfs_before_sess_syncs)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if hb_interval is not None:
            pulumi.set(__self__, "hb_interval", hb_interval)
        if hb_lost_threshold is not None:
            pulumi.set(__self__, "hb_lost_threshold", hb_lost_threshold)
        if ike_heartbeat_interval is not None:
            pulumi.set(__self__, "ike_heartbeat_interval", ike_heartbeat_interval)
        if ike_monitor is not None:
            pulumi.set(__self__, "ike_monitor", ike_monitor)
        if ike_monitor_interval is not None:
            pulumi.set(__self__, "ike_monitor_interval", ike_monitor_interval)
        if ipsec_tunnel_sync is not None:
            pulumi.set(__self__, "ipsec_tunnel_sync", ipsec_tunnel_sync)
        if peerip is not None:
            pulumi.set(__self__, "peerip", peerip)
        if peervd is not None:
            pulumi.set(__self__, "peervd", peervd)
        if secondary_add_ipsec_routes is not None:
            pulumi.set(__self__, "secondary_add_ipsec_routes", secondary_add_ipsec_routes)
        if session_sync_filter is not None:
            pulumi.set(__self__, "session_sync_filter", session_sync_filter)
        if slave_add_ike_routes is not None:
            pulumi.set(__self__, "slave_add_ike_routes", slave_add_ike_routes)
        if sync_id is not None:
            pulumi.set(__self__, "sync_id", sync_id)
        if syncvds is not None:
            pulumi.set(__self__, "syncvds", syncvds)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="downIntfsBeforeSessSyncs")
    def down_intfs_before_sess_syncs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]]:
        return pulumi.get(self, "down_intfs_before_sess_syncs")

    @down_intfs_before_sess_syncs.setter
    def down_intfs_before_sess_syncs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]]):
        pulumi.set(self, "down_intfs_before_sess_syncs", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="hbInterval")
    def hb_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "hb_interval")

    @hb_interval.setter
    def hb_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hb_interval", value)

    @property
    @pulumi.getter(name="hbLostThreshold")
    def hb_lost_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "hb_lost_threshold")

    @hb_lost_threshold.setter
    def hb_lost_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hb_lost_threshold", value)

    @property
    @pulumi.getter(name="ikeHeartbeatInterval")
    def ike_heartbeat_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ike_heartbeat_interval")

    @ike_heartbeat_interval.setter
    def ike_heartbeat_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ike_heartbeat_interval", value)

    @property
    @pulumi.getter(name="ikeMonitor")
    def ike_monitor(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ike_monitor")

    @ike_monitor.setter
    def ike_monitor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ike_monitor", value)

    @property
    @pulumi.getter(name="ikeMonitorInterval")
    def ike_monitor_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ike_monitor_interval")

    @ike_monitor_interval.setter
    def ike_monitor_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ike_monitor_interval", value)

    @property
    @pulumi.getter(name="ipsecTunnelSync")
    def ipsec_tunnel_sync(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipsec_tunnel_sync")

    @ipsec_tunnel_sync.setter
    def ipsec_tunnel_sync(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec_tunnel_sync", value)

    @property
    @pulumi.getter
    def peerip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peerip")

    @peerip.setter
    def peerip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peerip", value)

    @property
    @pulumi.getter
    def peervd(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peervd")

    @peervd.setter
    def peervd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peervd", value)

    @property
    @pulumi.getter(name="secondaryAddIpsecRoutes")
    def secondary_add_ipsec_routes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secondary_add_ipsec_routes")

    @secondary_add_ipsec_routes.setter
    def secondary_add_ipsec_routes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_add_ipsec_routes", value)

    @property
    @pulumi.getter(name="sessionSyncFilter")
    def session_sync_filter(self) -> Optional[pulumi.Input['SystemClusterSyncSessionSyncFilterArgs']]:
        return pulumi.get(self, "session_sync_filter")

    @session_sync_filter.setter
    def session_sync_filter(self, value: Optional[pulumi.Input['SystemClusterSyncSessionSyncFilterArgs']]):
        pulumi.set(self, "session_sync_filter", value)

    @property
    @pulumi.getter(name="slaveAddIkeRoutes")
    def slave_add_ike_routes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slave_add_ike_routes")

    @slave_add_ike_routes.setter
    def slave_add_ike_routes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slave_add_ike_routes", value)

    @property
    @pulumi.getter(name="syncId")
    def sync_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "sync_id")

    @sync_id.setter
    def sync_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sync_id", value)

    @property
    @pulumi.getter
    def syncvds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncSyncvdArgs']]]]:
        return pulumi.get(self, "syncvds")

    @syncvds.setter
    def syncvds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncSyncvdArgs']]]]):
        pulumi.set(self, "syncvds", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemClusterSyncState:
    def __init__(__self__, *,
                 down_intfs_before_sess_syncs: Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 hb_interval: Optional[pulumi.Input[int]] = None,
                 hb_lost_threshold: Optional[pulumi.Input[int]] = None,
                 ike_heartbeat_interval: Optional[pulumi.Input[int]] = None,
                 ike_monitor: Optional[pulumi.Input[str]] = None,
                 ike_monitor_interval: Optional[pulumi.Input[int]] = None,
                 ipsec_tunnel_sync: Optional[pulumi.Input[str]] = None,
                 peerip: Optional[pulumi.Input[str]] = None,
                 peervd: Optional[pulumi.Input[str]] = None,
                 secondary_add_ipsec_routes: Optional[pulumi.Input[str]] = None,
                 session_sync_filter: Optional[pulumi.Input['SystemClusterSyncSessionSyncFilterArgs']] = None,
                 slave_add_ike_routes: Optional[pulumi.Input[str]] = None,
                 sync_id: Optional[pulumi.Input[int]] = None,
                 syncvds: Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncSyncvdArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemClusterSync resources.
        """
        if down_intfs_before_sess_syncs is not None:
            pulumi.set(__self__, "down_intfs_before_sess_syncs", down_intfs_before_sess_syncs)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if hb_interval is not None:
            pulumi.set(__self__, "hb_interval", hb_interval)
        if hb_lost_threshold is not None:
            pulumi.set(__self__, "hb_lost_threshold", hb_lost_threshold)
        if ike_heartbeat_interval is not None:
            pulumi.set(__self__, "ike_heartbeat_interval", ike_heartbeat_interval)
        if ike_monitor is not None:
            pulumi.set(__self__, "ike_monitor", ike_monitor)
        if ike_monitor_interval is not None:
            pulumi.set(__self__, "ike_monitor_interval", ike_monitor_interval)
        if ipsec_tunnel_sync is not None:
            pulumi.set(__self__, "ipsec_tunnel_sync", ipsec_tunnel_sync)
        if peerip is not None:
            pulumi.set(__self__, "peerip", peerip)
        if peervd is not None:
            pulumi.set(__self__, "peervd", peervd)
        if secondary_add_ipsec_routes is not None:
            pulumi.set(__self__, "secondary_add_ipsec_routes", secondary_add_ipsec_routes)
        if session_sync_filter is not None:
            pulumi.set(__self__, "session_sync_filter", session_sync_filter)
        if slave_add_ike_routes is not None:
            pulumi.set(__self__, "slave_add_ike_routes", slave_add_ike_routes)
        if sync_id is not None:
            pulumi.set(__self__, "sync_id", sync_id)
        if syncvds is not None:
            pulumi.set(__self__, "syncvds", syncvds)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="downIntfsBeforeSessSyncs")
    def down_intfs_before_sess_syncs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]]:
        return pulumi.get(self, "down_intfs_before_sess_syncs")

    @down_intfs_before_sess_syncs.setter
    def down_intfs_before_sess_syncs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]]):
        pulumi.set(self, "down_intfs_before_sess_syncs", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="hbInterval")
    def hb_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "hb_interval")

    @hb_interval.setter
    def hb_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hb_interval", value)

    @property
    @pulumi.getter(name="hbLostThreshold")
    def hb_lost_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "hb_lost_threshold")

    @hb_lost_threshold.setter
    def hb_lost_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hb_lost_threshold", value)

    @property
    @pulumi.getter(name="ikeHeartbeatInterval")
    def ike_heartbeat_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ike_heartbeat_interval")

    @ike_heartbeat_interval.setter
    def ike_heartbeat_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ike_heartbeat_interval", value)

    @property
    @pulumi.getter(name="ikeMonitor")
    def ike_monitor(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ike_monitor")

    @ike_monitor.setter
    def ike_monitor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ike_monitor", value)

    @property
    @pulumi.getter(name="ikeMonitorInterval")
    def ike_monitor_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ike_monitor_interval")

    @ike_monitor_interval.setter
    def ike_monitor_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ike_monitor_interval", value)

    @property
    @pulumi.getter(name="ipsecTunnelSync")
    def ipsec_tunnel_sync(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipsec_tunnel_sync")

    @ipsec_tunnel_sync.setter
    def ipsec_tunnel_sync(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec_tunnel_sync", value)

    @property
    @pulumi.getter
    def peerip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peerip")

    @peerip.setter
    def peerip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peerip", value)

    @property
    @pulumi.getter
    def peervd(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peervd")

    @peervd.setter
    def peervd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peervd", value)

    @property
    @pulumi.getter(name="secondaryAddIpsecRoutes")
    def secondary_add_ipsec_routes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secondary_add_ipsec_routes")

    @secondary_add_ipsec_routes.setter
    def secondary_add_ipsec_routes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_add_ipsec_routes", value)

    @property
    @pulumi.getter(name="sessionSyncFilter")
    def session_sync_filter(self) -> Optional[pulumi.Input['SystemClusterSyncSessionSyncFilterArgs']]:
        return pulumi.get(self, "session_sync_filter")

    @session_sync_filter.setter
    def session_sync_filter(self, value: Optional[pulumi.Input['SystemClusterSyncSessionSyncFilterArgs']]):
        pulumi.set(self, "session_sync_filter", value)

    @property
    @pulumi.getter(name="slaveAddIkeRoutes")
    def slave_add_ike_routes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slave_add_ike_routes")

    @slave_add_ike_routes.setter
    def slave_add_ike_routes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slave_add_ike_routes", value)

    @property
    @pulumi.getter(name="syncId")
    def sync_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "sync_id")

    @sync_id.setter
    def sync_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sync_id", value)

    @property
    @pulumi.getter
    def syncvds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncSyncvdArgs']]]]:
        return pulumi.get(self, "syncvds")

    @syncvds.setter
    def syncvds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemClusterSyncSyncvdArgs']]]]):
        pulumi.set(self, "syncvds", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemClusterSync(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 down_intfs_before_sess_syncs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 hb_interval: Optional[pulumi.Input[int]] = None,
                 hb_lost_threshold: Optional[pulumi.Input[int]] = None,
                 ike_heartbeat_interval: Optional[pulumi.Input[int]] = None,
                 ike_monitor: Optional[pulumi.Input[str]] = None,
                 ike_monitor_interval: Optional[pulumi.Input[int]] = None,
                 ipsec_tunnel_sync: Optional[pulumi.Input[str]] = None,
                 peerip: Optional[pulumi.Input[str]] = None,
                 peervd: Optional[pulumi.Input[str]] = None,
                 secondary_add_ipsec_routes: Optional[pulumi.Input[str]] = None,
                 session_sync_filter: Optional[pulumi.Input[pulumi.InputType['SystemClusterSyncSessionSyncFilterArgs']]] = None,
                 slave_add_ike_routes: Optional[pulumi.Input[str]] = None,
                 sync_id: Optional[pulumi.Input[int]] = None,
                 syncvds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemClusterSyncSyncvdArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemClusterSync resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemClusterSyncArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemClusterSync resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemClusterSyncArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemClusterSyncArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 down_intfs_before_sess_syncs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 hb_interval: Optional[pulumi.Input[int]] = None,
                 hb_lost_threshold: Optional[pulumi.Input[int]] = None,
                 ike_heartbeat_interval: Optional[pulumi.Input[int]] = None,
                 ike_monitor: Optional[pulumi.Input[str]] = None,
                 ike_monitor_interval: Optional[pulumi.Input[int]] = None,
                 ipsec_tunnel_sync: Optional[pulumi.Input[str]] = None,
                 peerip: Optional[pulumi.Input[str]] = None,
                 peervd: Optional[pulumi.Input[str]] = None,
                 secondary_add_ipsec_routes: Optional[pulumi.Input[str]] = None,
                 session_sync_filter: Optional[pulumi.Input[pulumi.InputType['SystemClusterSyncSessionSyncFilterArgs']]] = None,
                 slave_add_ike_routes: Optional[pulumi.Input[str]] = None,
                 sync_id: Optional[pulumi.Input[int]] = None,
                 syncvds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemClusterSyncSyncvdArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemClusterSyncArgs.__new__(SystemClusterSyncArgs)

            __props__.__dict__["down_intfs_before_sess_syncs"] = down_intfs_before_sess_syncs
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["hb_interval"] = hb_interval
            __props__.__dict__["hb_lost_threshold"] = hb_lost_threshold
            __props__.__dict__["ike_heartbeat_interval"] = ike_heartbeat_interval
            __props__.__dict__["ike_monitor"] = ike_monitor
            __props__.__dict__["ike_monitor_interval"] = ike_monitor_interval
            __props__.__dict__["ipsec_tunnel_sync"] = ipsec_tunnel_sync
            __props__.__dict__["peerip"] = peerip
            __props__.__dict__["peervd"] = peervd
            __props__.__dict__["secondary_add_ipsec_routes"] = secondary_add_ipsec_routes
            __props__.__dict__["session_sync_filter"] = session_sync_filter
            __props__.__dict__["slave_add_ike_routes"] = slave_add_ike_routes
            __props__.__dict__["sync_id"] = sync_id
            __props__.__dict__["syncvds"] = syncvds
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemClusterSync, __self__).__init__(
            'fortios:index/systemClusterSync:SystemClusterSync',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            down_intfs_before_sess_syncs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemClusterSyncDownIntfsBeforeSessSyncArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            hb_interval: Optional[pulumi.Input[int]] = None,
            hb_lost_threshold: Optional[pulumi.Input[int]] = None,
            ike_heartbeat_interval: Optional[pulumi.Input[int]] = None,
            ike_monitor: Optional[pulumi.Input[str]] = None,
            ike_monitor_interval: Optional[pulumi.Input[int]] = None,
            ipsec_tunnel_sync: Optional[pulumi.Input[str]] = None,
            peerip: Optional[pulumi.Input[str]] = None,
            peervd: Optional[pulumi.Input[str]] = None,
            secondary_add_ipsec_routes: Optional[pulumi.Input[str]] = None,
            session_sync_filter: Optional[pulumi.Input[pulumi.InputType['SystemClusterSyncSessionSyncFilterArgs']]] = None,
            slave_add_ike_routes: Optional[pulumi.Input[str]] = None,
            sync_id: Optional[pulumi.Input[int]] = None,
            syncvds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemClusterSyncSyncvdArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemClusterSync':
        """
        Get an existing SystemClusterSync resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemClusterSyncState.__new__(_SystemClusterSyncState)

        __props__.__dict__["down_intfs_before_sess_syncs"] = down_intfs_before_sess_syncs
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["hb_interval"] = hb_interval
        __props__.__dict__["hb_lost_threshold"] = hb_lost_threshold
        __props__.__dict__["ike_heartbeat_interval"] = ike_heartbeat_interval
        __props__.__dict__["ike_monitor"] = ike_monitor
        __props__.__dict__["ike_monitor_interval"] = ike_monitor_interval
        __props__.__dict__["ipsec_tunnel_sync"] = ipsec_tunnel_sync
        __props__.__dict__["peerip"] = peerip
        __props__.__dict__["peervd"] = peervd
        __props__.__dict__["secondary_add_ipsec_routes"] = secondary_add_ipsec_routes
        __props__.__dict__["session_sync_filter"] = session_sync_filter
        __props__.__dict__["slave_add_ike_routes"] = slave_add_ike_routes
        __props__.__dict__["sync_id"] = sync_id
        __props__.__dict__["syncvds"] = syncvds
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemClusterSync(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="downIntfsBeforeSessSyncs")
    def down_intfs_before_sess_syncs(self) -> pulumi.Output[Optional[Sequence['outputs.SystemClusterSyncDownIntfsBeforeSessSync']]]:
        return pulumi.get(self, "down_intfs_before_sess_syncs")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="hbInterval")
    def hb_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "hb_interval")

    @property
    @pulumi.getter(name="hbLostThreshold")
    def hb_lost_threshold(self) -> pulumi.Output[int]:
        return pulumi.get(self, "hb_lost_threshold")

    @property
    @pulumi.getter(name="ikeHeartbeatInterval")
    def ike_heartbeat_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "ike_heartbeat_interval")

    @property
    @pulumi.getter(name="ikeMonitor")
    def ike_monitor(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ike_monitor")

    @property
    @pulumi.getter(name="ikeMonitorInterval")
    def ike_monitor_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "ike_monitor_interval")

    @property
    @pulumi.getter(name="ipsecTunnelSync")
    def ipsec_tunnel_sync(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ipsec_tunnel_sync")

    @property
    @pulumi.getter
    def peerip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "peerip")

    @property
    @pulumi.getter
    def peervd(self) -> pulumi.Output[str]:
        return pulumi.get(self, "peervd")

    @property
    @pulumi.getter(name="secondaryAddIpsecRoutes")
    def secondary_add_ipsec_routes(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secondary_add_ipsec_routes")

    @property
    @pulumi.getter(name="sessionSyncFilter")
    def session_sync_filter(self) -> pulumi.Output['outputs.SystemClusterSyncSessionSyncFilter']:
        return pulumi.get(self, "session_sync_filter")

    @property
    @pulumi.getter(name="slaveAddIkeRoutes")
    def slave_add_ike_routes(self) -> pulumi.Output[str]:
        return pulumi.get(self, "slave_add_ike_routes")

    @property
    @pulumi.getter(name="syncId")
    def sync_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "sync_id")

    @property
    @pulumi.getter
    def syncvds(self) -> pulumi.Output[Optional[Sequence['outputs.SystemClusterSyncSyncvd']]]:
        return pulumi.get(self, "syncvds")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

