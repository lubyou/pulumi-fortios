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
    'GetSystemVirtualWanLinkResult',
    'AwaitableGetSystemVirtualWanLinkResult',
    'get_system_virtual_wan_link',
]

@pulumi.output_type
class GetSystemVirtualWanLinkResult:
    """
    A collection of values returned by GetSystemVirtualWanLink.
    """
    def __init__(__self__, fail_alert_interfaces=None, fail_detect=None, health_checks=None, id=None, load_balance_mode=None, members=None, neighbor_hold_boot_time=None, neighbor_hold_down=None, neighbor_hold_down_time=None, neighbors=None, services=None, status=None, vdomparam=None, zones=None):
        if fail_alert_interfaces and not isinstance(fail_alert_interfaces, list):
            raise TypeError("Expected argument 'fail_alert_interfaces' to be a list")
        pulumi.set(__self__, "fail_alert_interfaces", fail_alert_interfaces)
        if fail_detect and not isinstance(fail_detect, str):
            raise TypeError("Expected argument 'fail_detect' to be a str")
        pulumi.set(__self__, "fail_detect", fail_detect)
        if health_checks and not isinstance(health_checks, list):
            raise TypeError("Expected argument 'health_checks' to be a list")
        pulumi.set(__self__, "health_checks", health_checks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if load_balance_mode and not isinstance(load_balance_mode, str):
            raise TypeError("Expected argument 'load_balance_mode' to be a str")
        pulumi.set(__self__, "load_balance_mode", load_balance_mode)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if neighbor_hold_boot_time and not isinstance(neighbor_hold_boot_time, int):
            raise TypeError("Expected argument 'neighbor_hold_boot_time' to be a int")
        pulumi.set(__self__, "neighbor_hold_boot_time", neighbor_hold_boot_time)
        if neighbor_hold_down and not isinstance(neighbor_hold_down, str):
            raise TypeError("Expected argument 'neighbor_hold_down' to be a str")
        pulumi.set(__self__, "neighbor_hold_down", neighbor_hold_down)
        if neighbor_hold_down_time and not isinstance(neighbor_hold_down_time, int):
            raise TypeError("Expected argument 'neighbor_hold_down_time' to be a int")
        pulumi.set(__self__, "neighbor_hold_down_time", neighbor_hold_down_time)
        if neighbors and not isinstance(neighbors, list):
            raise TypeError("Expected argument 'neighbors' to be a list")
        pulumi.set(__self__, "neighbors", neighbors)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="failAlertInterfaces")
    def fail_alert_interfaces(self) -> Sequence['outputs.GetSystemVirtualWanLinkFailAlertInterfaceResult']:
        """
        Physical interfaces that will be alerted. The structure of `fail_alert_interfaces` block is documented below.
        """
        return pulumi.get(self, "fail_alert_interfaces")

    @property
    @pulumi.getter(name="failDetect")
    def fail_detect(self) -> str:
        """
        Enable/disable SD-WAN Internet connection status checking (failure detection).
        """
        return pulumi.get(self, "fail_detect")

    @property
    @pulumi.getter(name="healthChecks")
    def health_checks(self) -> Sequence['outputs.GetSystemVirtualWanLinkHealthCheckResult']:
        """
        Virtual WAN Link health-check.
        """
        return pulumi.get(self, "health_checks")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loadBalanceMode")
    def load_balance_mode(self) -> str:
        """
        Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.
        """
        return pulumi.get(self, "load_balance_mode")

    @property
    @pulumi.getter
    def members(self) -> Sequence['outputs.GetSystemVirtualWanLinkMemberResult']:
        """
        Member sequence number list. The structure of `members` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter(name="neighborHoldBootTime")
    def neighbor_hold_boot_time(self) -> int:
        """
        Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
        """
        return pulumi.get(self, "neighbor_hold_boot_time")

    @property
    @pulumi.getter(name="neighborHoldDown")
    def neighbor_hold_down(self) -> str:
        """
        Enable/disable hold switching from the secondary neighbor to the primary neighbor.
        """
        return pulumi.get(self, "neighbor_hold_down")

    @property
    @pulumi.getter(name="neighborHoldDownTime")
    def neighbor_hold_down_time(self) -> int:
        """
        Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
        """
        return pulumi.get(self, "neighbor_hold_down_time")

    @property
    @pulumi.getter
    def neighbors(self) -> Sequence['outputs.GetSystemVirtualWanLinkNeighborResult']:
        """
        Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetSystemVirtualWanLinkServiceResult']:
        """
        Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable SD-WAN service.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def zones(self) -> Sequence['outputs.GetSystemVirtualWanLinkZoneResult']:
        """
        Configure SD-WAN zones. The structure of `zone` block is documented below.
        """
        return pulumi.get(self, "zones")


class AwaitableGetSystemVirtualWanLinkResult(GetSystemVirtualWanLinkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemVirtualWanLinkResult(
            fail_alert_interfaces=self.fail_alert_interfaces,
            fail_detect=self.fail_detect,
            health_checks=self.health_checks,
            id=self.id,
            load_balance_mode=self.load_balance_mode,
            members=self.members,
            neighbor_hold_boot_time=self.neighbor_hold_boot_time,
            neighbor_hold_down=self.neighbor_hold_down,
            neighbor_hold_down_time=self.neighbor_hold_down_time,
            neighbors=self.neighbors,
            services=self.services,
            status=self.status,
            vdomparam=self.vdomparam,
            zones=self.zones)


def get_system_virtual_wan_link(vdomparam: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemVirtualWanLinkResult:
    """
    Use this data source to get information on fortios system virtualwanlink


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemVirtualWanLink:GetSystemVirtualWanLink', __args__, opts=opts, typ=GetSystemVirtualWanLinkResult).value

    return AwaitableGetSystemVirtualWanLinkResult(
        fail_alert_interfaces=__ret__.fail_alert_interfaces,
        fail_detect=__ret__.fail_detect,
        health_checks=__ret__.health_checks,
        id=__ret__.id,
        load_balance_mode=__ret__.load_balance_mode,
        members=__ret__.members,
        neighbor_hold_boot_time=__ret__.neighbor_hold_boot_time,
        neighbor_hold_down=__ret__.neighbor_hold_down,
        neighbor_hold_down_time=__ret__.neighbor_hold_down_time,
        neighbors=__ret__.neighbors,
        services=__ret__.services,
        status=__ret__.status,
        vdomparam=__ret__.vdomparam,
        zones=__ret__.zones)