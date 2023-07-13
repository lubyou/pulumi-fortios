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

__all__ = ['VpnOcvpnArgs', 'VpnOcvpn']

@pulumi.input_type
class VpnOcvpnArgs:
    def __init__(__self__, *,
                 auto_discovery: Optional[pulumi.Input[str]] = None,
                 auto_discovery_shortcut_mode: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 eap: Optional[pulumi.Input[str]] = None,
                 eap_users: Optional[pulumi.Input[str]] = None,
                 forticlient_access: Optional[pulumi.Input['VpnOcvpnForticlientAccessArgs']] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ip_allocation_block: Optional[pulumi.Input[str]] = None,
                 multipath: Optional[pulumi.Input[str]] = None,
                 nat: Optional[pulumi.Input[str]] = None,
                 overlays: Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnOverlayArgs']]]] = None,
                 poll_interval: Optional[pulumi.Input[int]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 sdwan_zone: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wan_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnWanInterfaceArgs']]]] = None):
        """
        The set of arguments for constructing a VpnOcvpn resource.
        """
        if auto_discovery is not None:
            pulumi.set(__self__, "auto_discovery", auto_discovery)
        if auto_discovery_shortcut_mode is not None:
            pulumi.set(__self__, "auto_discovery_shortcut_mode", auto_discovery_shortcut_mode)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if eap is not None:
            pulumi.set(__self__, "eap", eap)
        if eap_users is not None:
            pulumi.set(__self__, "eap_users", eap_users)
        if forticlient_access is not None:
            pulumi.set(__self__, "forticlient_access", forticlient_access)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ip_allocation_block is not None:
            pulumi.set(__self__, "ip_allocation_block", ip_allocation_block)
        if multipath is not None:
            pulumi.set(__self__, "multipath", multipath)
        if nat is not None:
            pulumi.set(__self__, "nat", nat)
        if overlays is not None:
            pulumi.set(__self__, "overlays", overlays)
        if poll_interval is not None:
            pulumi.set(__self__, "poll_interval", poll_interval)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if sdwan is not None:
            pulumi.set(__self__, "sdwan", sdwan)
        if sdwan_zone is not None:
            pulumi.set(__self__, "sdwan_zone", sdwan_zone)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wan_interfaces is not None:
            pulumi.set(__self__, "wan_interfaces", wan_interfaces)

    @property
    @pulumi.getter(name="autoDiscovery")
    def auto_discovery(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_discovery")

    @auto_discovery.setter
    def auto_discovery(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_discovery", value)

    @property
    @pulumi.getter(name="autoDiscoveryShortcutMode")
    def auto_discovery_shortcut_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_discovery_shortcut_mode")

    @auto_discovery_shortcut_mode.setter
    def auto_discovery_shortcut_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_discovery_shortcut_mode", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def eap(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eap")

    @eap.setter
    def eap(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eap", value)

    @property
    @pulumi.getter(name="eapUsers")
    def eap_users(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eap_users")

    @eap_users.setter
    def eap_users(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eap_users", value)

    @property
    @pulumi.getter(name="forticlientAccess")
    def forticlient_access(self) -> Optional[pulumi.Input['VpnOcvpnForticlientAccessArgs']]:
        return pulumi.get(self, "forticlient_access")

    @forticlient_access.setter
    def forticlient_access(self, value: Optional[pulumi.Input['VpnOcvpnForticlientAccessArgs']]):
        pulumi.set(self, "forticlient_access", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="ipAllocationBlock")
    def ip_allocation_block(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_allocation_block")

    @ip_allocation_block.setter
    def ip_allocation_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_allocation_block", value)

    @property
    @pulumi.getter
    def multipath(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "multipath")

    @multipath.setter
    def multipath(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multipath", value)

    @property
    @pulumi.getter
    def nat(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nat")

    @nat.setter
    def nat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat", value)

    @property
    @pulumi.getter
    def overlays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnOverlayArgs']]]]:
        return pulumi.get(self, "overlays")

    @overlays.setter
    def overlays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnOverlayArgs']]]]):
        pulumi.set(self, "overlays", value)

    @property
    @pulumi.getter(name="pollInterval")
    def poll_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "poll_interval")

    @poll_interval.setter
    def poll_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "poll_interval", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def sdwan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sdwan")

    @sdwan.setter
    def sdwan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdwan", value)

    @property
    @pulumi.getter(name="sdwanZone")
    def sdwan_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sdwan_zone")

    @sdwan_zone.setter
    def sdwan_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdwan_zone", value)

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
    @pulumi.getter(name="wanInterfaces")
    def wan_interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnWanInterfaceArgs']]]]:
        return pulumi.get(self, "wan_interfaces")

    @wan_interfaces.setter
    def wan_interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnWanInterfaceArgs']]]]):
        pulumi.set(self, "wan_interfaces", value)


@pulumi.input_type
class _VpnOcvpnState:
    def __init__(__self__, *,
                 auto_discovery: Optional[pulumi.Input[str]] = None,
                 auto_discovery_shortcut_mode: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 eap: Optional[pulumi.Input[str]] = None,
                 eap_users: Optional[pulumi.Input[str]] = None,
                 forticlient_access: Optional[pulumi.Input['VpnOcvpnForticlientAccessArgs']] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ip_allocation_block: Optional[pulumi.Input[str]] = None,
                 multipath: Optional[pulumi.Input[str]] = None,
                 nat: Optional[pulumi.Input[str]] = None,
                 overlays: Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnOverlayArgs']]]] = None,
                 poll_interval: Optional[pulumi.Input[int]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 sdwan_zone: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wan_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnWanInterfaceArgs']]]] = None):
        """
        Input properties used for looking up and filtering VpnOcvpn resources.
        """
        if auto_discovery is not None:
            pulumi.set(__self__, "auto_discovery", auto_discovery)
        if auto_discovery_shortcut_mode is not None:
            pulumi.set(__self__, "auto_discovery_shortcut_mode", auto_discovery_shortcut_mode)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if eap is not None:
            pulumi.set(__self__, "eap", eap)
        if eap_users is not None:
            pulumi.set(__self__, "eap_users", eap_users)
        if forticlient_access is not None:
            pulumi.set(__self__, "forticlient_access", forticlient_access)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ip_allocation_block is not None:
            pulumi.set(__self__, "ip_allocation_block", ip_allocation_block)
        if multipath is not None:
            pulumi.set(__self__, "multipath", multipath)
        if nat is not None:
            pulumi.set(__self__, "nat", nat)
        if overlays is not None:
            pulumi.set(__self__, "overlays", overlays)
        if poll_interval is not None:
            pulumi.set(__self__, "poll_interval", poll_interval)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if sdwan is not None:
            pulumi.set(__self__, "sdwan", sdwan)
        if sdwan_zone is not None:
            pulumi.set(__self__, "sdwan_zone", sdwan_zone)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wan_interfaces is not None:
            pulumi.set(__self__, "wan_interfaces", wan_interfaces)

    @property
    @pulumi.getter(name="autoDiscovery")
    def auto_discovery(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_discovery")

    @auto_discovery.setter
    def auto_discovery(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_discovery", value)

    @property
    @pulumi.getter(name="autoDiscoveryShortcutMode")
    def auto_discovery_shortcut_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_discovery_shortcut_mode")

    @auto_discovery_shortcut_mode.setter
    def auto_discovery_shortcut_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_discovery_shortcut_mode", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def eap(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eap")

    @eap.setter
    def eap(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eap", value)

    @property
    @pulumi.getter(name="eapUsers")
    def eap_users(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "eap_users")

    @eap_users.setter
    def eap_users(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eap_users", value)

    @property
    @pulumi.getter(name="forticlientAccess")
    def forticlient_access(self) -> Optional[pulumi.Input['VpnOcvpnForticlientAccessArgs']]:
        return pulumi.get(self, "forticlient_access")

    @forticlient_access.setter
    def forticlient_access(self, value: Optional[pulumi.Input['VpnOcvpnForticlientAccessArgs']]):
        pulumi.set(self, "forticlient_access", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="ipAllocationBlock")
    def ip_allocation_block(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_allocation_block")

    @ip_allocation_block.setter
    def ip_allocation_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_allocation_block", value)

    @property
    @pulumi.getter
    def multipath(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "multipath")

    @multipath.setter
    def multipath(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multipath", value)

    @property
    @pulumi.getter
    def nat(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nat")

    @nat.setter
    def nat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat", value)

    @property
    @pulumi.getter
    def overlays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnOverlayArgs']]]]:
        return pulumi.get(self, "overlays")

    @overlays.setter
    def overlays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnOverlayArgs']]]]):
        pulumi.set(self, "overlays", value)

    @property
    @pulumi.getter(name="pollInterval")
    def poll_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "poll_interval")

    @poll_interval.setter
    def poll_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "poll_interval", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def sdwan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sdwan")

    @sdwan.setter
    def sdwan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdwan", value)

    @property
    @pulumi.getter(name="sdwanZone")
    def sdwan_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sdwan_zone")

    @sdwan_zone.setter
    def sdwan_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdwan_zone", value)

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
    @pulumi.getter(name="wanInterfaces")
    def wan_interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnWanInterfaceArgs']]]]:
        return pulumi.get(self, "wan_interfaces")

    @wan_interfaces.setter
    def wan_interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnOcvpnWanInterfaceArgs']]]]):
        pulumi.set(self, "wan_interfaces", value)


class VpnOcvpn(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_discovery: Optional[pulumi.Input[str]] = None,
                 auto_discovery_shortcut_mode: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 eap: Optional[pulumi.Input[str]] = None,
                 eap_users: Optional[pulumi.Input[str]] = None,
                 forticlient_access: Optional[pulumi.Input[pulumi.InputType['VpnOcvpnForticlientAccessArgs']]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ip_allocation_block: Optional[pulumi.Input[str]] = None,
                 multipath: Optional[pulumi.Input[str]] = None,
                 nat: Optional[pulumi.Input[str]] = None,
                 overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnOcvpnOverlayArgs']]]]] = None,
                 poll_interval: Optional[pulumi.Input[int]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 sdwan_zone: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wan_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnOcvpnWanInterfaceArgs']]]]] = None,
                 __props__=None):
        """
        Create a VpnOcvpn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VpnOcvpnArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpnOcvpn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpnOcvpnArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpnOcvpnArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_discovery: Optional[pulumi.Input[str]] = None,
                 auto_discovery_shortcut_mode: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 eap: Optional[pulumi.Input[str]] = None,
                 eap_users: Optional[pulumi.Input[str]] = None,
                 forticlient_access: Optional[pulumi.Input[pulumi.InputType['VpnOcvpnForticlientAccessArgs']]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ip_allocation_block: Optional[pulumi.Input[str]] = None,
                 multipath: Optional[pulumi.Input[str]] = None,
                 nat: Optional[pulumi.Input[str]] = None,
                 overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnOcvpnOverlayArgs']]]]] = None,
                 poll_interval: Optional[pulumi.Input[int]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 sdwan_zone: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wan_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnOcvpnWanInterfaceArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpnOcvpnArgs.__new__(VpnOcvpnArgs)

            __props__.__dict__["auto_discovery"] = auto_discovery
            __props__.__dict__["auto_discovery_shortcut_mode"] = auto_discovery_shortcut_mode
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["eap"] = eap
            __props__.__dict__["eap_users"] = eap_users
            __props__.__dict__["forticlient_access"] = forticlient_access
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["ip_allocation_block"] = ip_allocation_block
            __props__.__dict__["multipath"] = multipath
            __props__.__dict__["nat"] = nat
            __props__.__dict__["overlays"] = overlays
            __props__.__dict__["poll_interval"] = poll_interval
            __props__.__dict__["role"] = role
            __props__.__dict__["sdwan"] = sdwan
            __props__.__dict__["sdwan_zone"] = sdwan_zone
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["wan_interfaces"] = wan_interfaces
        super(VpnOcvpn, __self__).__init__(
            'fortios:index/vpnOcvpn:VpnOcvpn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_discovery: Optional[pulumi.Input[str]] = None,
            auto_discovery_shortcut_mode: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            eap: Optional[pulumi.Input[str]] = None,
            eap_users: Optional[pulumi.Input[str]] = None,
            forticlient_access: Optional[pulumi.Input[pulumi.InputType['VpnOcvpnForticlientAccessArgs']]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            ip_allocation_block: Optional[pulumi.Input[str]] = None,
            multipath: Optional[pulumi.Input[str]] = None,
            nat: Optional[pulumi.Input[str]] = None,
            overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnOcvpnOverlayArgs']]]]] = None,
            poll_interval: Optional[pulumi.Input[int]] = None,
            role: Optional[pulumi.Input[str]] = None,
            sdwan: Optional[pulumi.Input[str]] = None,
            sdwan_zone: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            wan_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnOcvpnWanInterfaceArgs']]]]] = None) -> 'VpnOcvpn':
        """
        Get an existing VpnOcvpn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpnOcvpnState.__new__(_VpnOcvpnState)

        __props__.__dict__["auto_discovery"] = auto_discovery
        __props__.__dict__["auto_discovery_shortcut_mode"] = auto_discovery_shortcut_mode
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["eap"] = eap
        __props__.__dict__["eap_users"] = eap_users
        __props__.__dict__["forticlient_access"] = forticlient_access
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["ip_allocation_block"] = ip_allocation_block
        __props__.__dict__["multipath"] = multipath
        __props__.__dict__["nat"] = nat
        __props__.__dict__["overlays"] = overlays
        __props__.__dict__["poll_interval"] = poll_interval
        __props__.__dict__["role"] = role
        __props__.__dict__["sdwan"] = sdwan
        __props__.__dict__["sdwan_zone"] = sdwan_zone
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["wan_interfaces"] = wan_interfaces
        return VpnOcvpn(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDiscovery")
    def auto_discovery(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auto_discovery")

    @property
    @pulumi.getter(name="autoDiscoveryShortcutMode")
    def auto_discovery_shortcut_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auto_discovery_shortcut_mode")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def eap(self) -> pulumi.Output[str]:
        return pulumi.get(self, "eap")

    @property
    @pulumi.getter(name="eapUsers")
    def eap_users(self) -> pulumi.Output[str]:
        return pulumi.get(self, "eap_users")

    @property
    @pulumi.getter(name="forticlientAccess")
    def forticlient_access(self) -> pulumi.Output['outputs.VpnOcvpnForticlientAccess']:
        return pulumi.get(self, "forticlient_access")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="ipAllocationBlock")
    def ip_allocation_block(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ip_allocation_block")

    @property
    @pulumi.getter
    def multipath(self) -> pulumi.Output[str]:
        return pulumi.get(self, "multipath")

    @property
    @pulumi.getter
    def nat(self) -> pulumi.Output[str]:
        return pulumi.get(self, "nat")

    @property
    @pulumi.getter
    def overlays(self) -> pulumi.Output[Optional[Sequence['outputs.VpnOcvpnOverlay']]]:
        return pulumi.get(self, "overlays")

    @property
    @pulumi.getter(name="pollInterval")
    def poll_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "poll_interval")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def sdwan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sdwan")

    @property
    @pulumi.getter(name="sdwanZone")
    def sdwan_zone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sdwan_zone")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="wanInterfaces")
    def wan_interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.VpnOcvpnWanInterface']]]:
        return pulumi.get(self, "wan_interfaces")

