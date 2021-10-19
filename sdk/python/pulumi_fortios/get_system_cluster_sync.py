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
    'GetSystemClusterSyncResult',
    'AwaitableGetSystemClusterSyncResult',
    'get_system_cluster_sync',
]

@pulumi.output_type
class GetSystemClusterSyncResult:
    """
    A collection of values returned by GetSystemClusterSync.
    """
    def __init__(__self__, down_intfs_before_sess_syncs=None, hb_interval=None, hb_lost_threshold=None, id=None, ipsec_tunnel_sync=None, peerip=None, peervd=None, session_sync_filter=None, slave_add_ike_routes=None, sync_id=None, syncvds=None, vdomparam=None):
        if down_intfs_before_sess_syncs and not isinstance(down_intfs_before_sess_syncs, list):
            raise TypeError("Expected argument 'down_intfs_before_sess_syncs' to be a list")
        pulumi.set(__self__, "down_intfs_before_sess_syncs", down_intfs_before_sess_syncs)
        if hb_interval and not isinstance(hb_interval, int):
            raise TypeError("Expected argument 'hb_interval' to be a int")
        pulumi.set(__self__, "hb_interval", hb_interval)
        if hb_lost_threshold and not isinstance(hb_lost_threshold, int):
            raise TypeError("Expected argument 'hb_lost_threshold' to be a int")
        pulumi.set(__self__, "hb_lost_threshold", hb_lost_threshold)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipsec_tunnel_sync and not isinstance(ipsec_tunnel_sync, str):
            raise TypeError("Expected argument 'ipsec_tunnel_sync' to be a str")
        pulumi.set(__self__, "ipsec_tunnel_sync", ipsec_tunnel_sync)
        if peerip and not isinstance(peerip, str):
            raise TypeError("Expected argument 'peerip' to be a str")
        pulumi.set(__self__, "peerip", peerip)
        if peervd and not isinstance(peervd, str):
            raise TypeError("Expected argument 'peervd' to be a str")
        pulumi.set(__self__, "peervd", peervd)
        if session_sync_filter and not isinstance(session_sync_filter, dict):
            raise TypeError("Expected argument 'session_sync_filter' to be a dict")
        pulumi.set(__self__, "session_sync_filter", session_sync_filter)
        if slave_add_ike_routes and not isinstance(slave_add_ike_routes, str):
            raise TypeError("Expected argument 'slave_add_ike_routes' to be a str")
        pulumi.set(__self__, "slave_add_ike_routes", slave_add_ike_routes)
        if sync_id and not isinstance(sync_id, int):
            raise TypeError("Expected argument 'sync_id' to be a int")
        pulumi.set(__self__, "sync_id", sync_id)
        if syncvds and not isinstance(syncvds, list):
            raise TypeError("Expected argument 'syncvds' to be a list")
        pulumi.set(__self__, "syncvds", syncvds)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="downIntfsBeforeSessSyncs")
    def down_intfs_before_sess_syncs(self) -> Sequence['outputs.GetSystemClusterSyncDownIntfsBeforeSessSyncResult']:
        """
        List of interfaces to be turned down before session synchronization is complete. The structure of `down_intfs_before_sess_sync` block is documented below.
        """
        return pulumi.get(self, "down_intfs_before_sess_syncs")

    @property
    @pulumi.getter(name="hbInterval")
    def hb_interval(self) -> int:
        """
        Heartbeat interval (1 - 10 sec).
        """
        return pulumi.get(self, "hb_interval")

    @property
    @pulumi.getter(name="hbLostThreshold")
    def hb_lost_threshold(self) -> int:
        """
        Lost heartbeat threshold (1 - 10).
        """
        return pulumi.get(self, "hb_lost_threshold")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipsecTunnelSync")
    def ipsec_tunnel_sync(self) -> str:
        """
        Enable/disable IPsec tunnel synchronization.
        """
        return pulumi.get(self, "ipsec_tunnel_sync")

    @property
    @pulumi.getter
    def peerip(self) -> str:
        """
        IP address of the interface on the peer unit that is used for the session synchronization link.
        """
        return pulumi.get(self, "peerip")

    @property
    @pulumi.getter
    def peervd(self) -> str:
        """
        VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
        """
        return pulumi.get(self, "peervd")

    @property
    @pulumi.getter(name="sessionSyncFilter")
    def session_sync_filter(self) -> 'outputs.GetSystemClusterSyncSessionSyncFilterResult':
        """
        Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `session_sync_filter` block is documented below.
        """
        return pulumi.get(self, "session_sync_filter")

    @property
    @pulumi.getter(name="slaveAddIkeRoutes")
    def slave_add_ike_routes(self) -> str:
        """
        Enable/disable IKE route announcement on the backup unit.
        """
        return pulumi.get(self, "slave_add_ike_routes")

    @property
    @pulumi.getter(name="syncId")
    def sync_id(self) -> int:
        """
        Sync ID.
        """
        return pulumi.get(self, "sync_id")

    @property
    @pulumi.getter
    def syncvds(self) -> Sequence['outputs.GetSystemClusterSyncSyncvdResult']:
        """
        Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
        """
        return pulumi.get(self, "syncvds")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemClusterSyncResult(GetSystemClusterSyncResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemClusterSyncResult(
            down_intfs_before_sess_syncs=self.down_intfs_before_sess_syncs,
            hb_interval=self.hb_interval,
            hb_lost_threshold=self.hb_lost_threshold,
            id=self.id,
            ipsec_tunnel_sync=self.ipsec_tunnel_sync,
            peerip=self.peerip,
            peervd=self.peervd,
            session_sync_filter=self.session_sync_filter,
            slave_add_ike_routes=self.slave_add_ike_routes,
            sync_id=self.sync_id,
            syncvds=self.syncvds,
            vdomparam=self.vdomparam)


def get_system_cluster_sync(sync_id: Optional[int] = None,
                            vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemClusterSyncResult:
    """
    Use this data source to get information on an fortios system clustersync


    :param int sync_id: Specify the sync_id of the desired system clustersync.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['syncId'] = sync_id
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemClusterSync:GetSystemClusterSync', __args__, opts=opts, typ=GetSystemClusterSyncResult).value

    return AwaitableGetSystemClusterSyncResult(
        down_intfs_before_sess_syncs=__ret__.down_intfs_before_sess_syncs,
        hb_interval=__ret__.hb_interval,
        hb_lost_threshold=__ret__.hb_lost_threshold,
        id=__ret__.id,
        ipsec_tunnel_sync=__ret__.ipsec_tunnel_sync,
        peerip=__ret__.peerip,
        peervd=__ret__.peervd,
        session_sync_filter=__ret__.session_sync_filter,
        slave_add_ike_routes=__ret__.slave_add_ike_routes,
        sync_id=__ret__.sync_id,
        syncvds=__ret__.syncvds,
        vdomparam=__ret__.vdomparam)
